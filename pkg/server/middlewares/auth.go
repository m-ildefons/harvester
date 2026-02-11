package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	ctlcorev1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/core/v1"

	"github.com/harvester/harvester/pkg/util"
	"github.com/rancher/apiserver/pkg/apierror"
	"github.com/rancher/rancher/pkg/auth/tokens/hashers"
	"github.com/rancher/wrangler/v3/pkg/schemas/validation"
	"github.com/sirupsen/logrus"
)

const (
	localRKEStateSecretName = "local-rke-state"
	serverTokenKey          = "serverToken"
)

type AuthMiddleware struct {
	secretCache ctlcorev1.SecretCache

	tokenOnce sync.Once
	tokenHash string
	tokenErr  error
}

func NewAuthMiddleware(secretCache ctlcorev1.SecretCache) *AuthMiddleware {
	amw := &AuthMiddleware{
		secretCache: secretCache,
		tokenErr:    nil,
	}
	return amw
}

func (amw *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		amw.loadToken()
		if amw.tokenErr != nil {
			http.Error(w, "failed to load token", http.StatusInternalServerError)
		}

		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "bearer token required", http.StatusForbidden)
		} else {
			token := strings.TrimPrefix(header, "Bearer ")
			err := amw.validateToken(token)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
				}).Warning("failed authentication attempt")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}

func (amw *AuthMiddleware) loadToken() {
	secret, err := amw.secretCache.Get(util.FleetLocalNamespaceName, localRKEStateSecretName)
	if err != nil {
		amw.tokenErr = err
		return
	}
	bytes, ok := secret.Data[serverTokenKey]
	if !ok {
		amw.tokenErr = fmt.Errorf("RKE state secret does not contain server token")
		return
	}
	token := string(bytes)
	hasher := hashers.GetHasher()
	amw.tokenHash, err = hasher.CreateHash(token)
	if err != nil {
		amw.tokenErr = err
	}
}

func (amw *AuthMiddleware) validateToken(token string) error {
	hasher, err := hashers.GetHasherForHash(amw.tokenHash)
	if err != nil {
		return fmt.Errorf("failed to get hasher: %s", err.Error())
	}
	err = hasher.VerifyHash(amw.tokenHash, token)
	if err == nil {
		return nil
	}
	return apierror.NewAPIError(validation.PermissionDenied, "Invalid token")
}
