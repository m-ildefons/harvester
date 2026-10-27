module github.com/harvester/harvester/pkg/builder

go 1.26.0

replace (
	github.com/cilium/ebpf => github.com/cilium/ebpf v0.7.0
	github.com/containerd/containerd => github.com/containerd/containerd v1.7.35
	github.com/cyphar/filepath-securejoin => github.com/cyphar/filepath-securejoin v0.3.6
	github.com/docker/distribution => github.com/docker/distribution v2.8.2+incompatible // oras dep requires a replace is set
	github.com/docker/docker => github.com/docker/docker v25.0.6+incompatible // oras dep requires a replace is set
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.9.1
	github.com/google/gnostic-models => github.com/google/gnostic-models v0.7.1

	github.com/harvester/harvester => github.com/m-ildefons/harvester v0.0.0-20261027125242-5d0d4408b0c5
	github.com/harvester/harvester-network-controller => github.com/harvester/harvester-network-controller v1.6.0-rc3
	github.com/nsf/termbox-go => github.com/Harvester/termbox-go v1.1.1-0.20210318083914-8ab92204a400
	github.com/opencontainers/runc => github.com/opencontainers/runc v1.1.14
	github.com/openshift/api => github.com/openshift/api v0.0.0-20191219222812-2987a591a72c
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20200521150516-05eb9880269c
	github.com/operator-framework/operator-lifecycle-manager => github.com/operator-framework/operator-lifecycle-manager v0.0.0-20190128024246-5eb7ae5bdb7a
	github.com/rancher/lasso => github.com/rancher/lasso v0.0.0-20241202185148-04649f379358
	github.com/rancher/rancher => github.com/rancher/rancher v0.0.0-20250828140533-07a90f09a491
	github.com/rancher/rancher/pkg/apis => github.com/rancher/rancher/pkg/apis v0.0.0-20250828140533-07a90f09a491
	github.com/rancher/rancher/pkg/client => github.com/rancher/rancher/pkg/client v0.0.0-20250828140533-07a90f09a491
	// handle rancher dependenices
	go.qase.io/client => github.com/rancher/qase-go/client v0.0.0-20250627195016-142ff3dfec16

	helm.sh/helm/v3 => github.com/rancher/helm/v3 v3.15.1-rancher2
	k8s.io/api => k8s.io/api v0.34.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.34.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.34.2
	k8s.io/apiserver => k8s.io/apiserver v0.34.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.34.2
	k8s.io/client-go => k8s.io/client-go v0.34.2
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.34.2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.34.2
	k8s.io/code-generator => k8s.io/code-generator v0.34.2
	k8s.io/component-base => k8s.io/component-base v0.34.2
	k8s.io/component-helpers => k8s.io/component-helpers v0.34.2
	k8s.io/controller-manager => k8s.io/controller-manager v0.34.2
	k8s.io/cri-api => k8s.io/cri-api v0.34.2
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.34.2
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.34.2
	k8s.io/endpointslice => k8s.io/endpointslice v0.34.2
	k8s.io/gengo => k8s.io/gengo v0.0.0-20240826214909-a7b603a56eb7
	k8s.io/kms => k8s.io/kms v0.34.2
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.34.2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.34.2
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20250710124328-f3f2b991d03b
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.34.2
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.34.2
	k8s.io/kubectl => k8s.io/kubectl v0.34.2
	k8s.io/kubelet => k8s.io/kubelet v0.34.2
	k8s.io/kubernetes => k8s.io/kubernetes v1.34.2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.34.2
	k8s.io/metrics => k8s.io/metrics v0.34.2
	k8s.io/mount-utils => k8s.io/mount-utils v0.34.2
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.34.2
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.34.2
)

require (
	github.com/harvester/harvester v1.5.0
	github.com/sirupsen/logrus v1.9.4
	github.com/stretchr/testify v1.11.1
	k8s.io/api v0.34.5
	k8s.io/apimachinery v0.34.5
	k8s.io/utils v0.0.0-20260319190234-28399d86e0b5
	kubevirt.io/api v1.9.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/c9s/goprocinfo v0.0.0-20210130143923-c95fcf8c64a8 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cockroachdb/errors v1.12.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/containernetworking/cni v1.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.6.0 // indirect
	github.com/cyphar/filepath-securejoin v0.6.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.13.0 // indirect
	github.com/evanphx/json-patch v5.9.11+incompatible // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/fxamacker/cbor/v2 v2.9.0 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.22.4 // indirect
	github.com/go-openapi/jsonreference v0.21.4 // indirect
	github.com/go-openapi/swag v0.25.4 // indirect
	github.com/go-openapi/swag/cmdutils v0.25.4 // indirect
	github.com/go-openapi/swag/conv v0.25.4 // indirect
	github.com/go-openapi/swag/fileutils v0.25.4 // indirect
	github.com/go-openapi/swag/jsonname v0.25.4 // indirect
	github.com/go-openapi/swag/jsonutils v0.25.4 // indirect
	github.com/go-openapi/swag/loading v0.25.4 // indirect
	github.com/go-openapi/swag/mangling v0.25.4 // indirect
	github.com/go-openapi/swag/netutils v0.25.4 // indirect
	github.com/go-openapi/swag/stringutils v0.25.4 // indirect
	github.com/go-openapi/swag/typeutils v0.25.4 // indirect
	github.com/go-openapi/swag/yamlutils v0.25.4 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.7.0-rc.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/gnostic-models v0.7.1 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/gorilla/websocket v1.5.4-0.20250319132907-e064f32e3674 // indirect
	github.com/harvester/go-common v0.0.0-20260119194217-0f17176ce67e // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v1.7.7 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kubernetes-csi/external-snapshotter/client/v4 v4.2.0 // indirect
	github.com/longhorn/go-common-libs v0.0.0-20250921030719-16313e7f30b3 // indirect
	github.com/longhorn/longhorn-manager v1.10.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/moby/sys/mountinfo v0.7.2 // indirect
	github.com/moby/sys/userns v0.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/onsi/gomega v1.40.0 // indirect
	github.com/opencontainers/cgroups v0.0.6 // indirect
	github.com/openshift/api v0.0.1 // indirect
	github.com/openshift/client-go v3.9.0+incompatible // indirect
	github.com/openshift/custom-resource-status v1.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.82.0 // indirect
	github.com/prometheus/client_golang v1.23.2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.66.1 // indirect
	github.com/prometheus/procfs v0.17.0 // indirect
	github.com/rancher/lasso v0.2.3 // indirect
	github.com/rancher/wrangler v1.1.2 // indirect
	github.com/rancher/wrangler/v3 v3.2.4 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.uber.org/mock v0.5.2 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.57.0 // indirect
	golang.org/x/exp v0.0.0-20250911091902-df9299821621 // indirect
	golang.org/x/net v0.60.0 // indirect
	golang.org/x/oauth2 v0.36.0 // indirect
	golang.org/x/sync v0.23.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/term v0.46.0 // indirect
	golang.org/x/text v0.42.0 // indirect
	golang.org/x/time v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	google.golang.org/grpc v1.83.2 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.13.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.34.3 // indirect
	k8s.io/client-go v12.0.0+incompatible // indirect
	k8s.io/cloud-provider v0.0.0 // indirect
	k8s.io/klog/v2 v2.140.0 // indirect
	k8s.io/kube-aggregator v0.33.1 // indirect
	k8s.io/kube-openapi v0.32.8 // indirect
	kubevirt.io/client-go v1.9.0 // indirect
	kubevirt.io/containerized-data-importer-api v1.64.0 // indirect
	kubevirt.io/controller-lifecycle-operator-sdk/api v0.2.4 // indirect
	kubevirt.io/kubevirt v1.9.0 // indirect
	kubevirt.io/virt-template-api v0.2.0 // indirect
	kubevirt.io/virt-template-client-go v0.2.0 // indirect
	sigs.k8s.io/controller-runtime v0.22.5 // indirect
	sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v6 v6.3.2-0.20260122202528-d9cc6641c482 // indirect
	sigs.k8s.io/yaml v1.6.0 // indirect
)
