check_version_range(){
  minor_version="$(get_server_version | cut -d '.' -f 2)"
  [[ "$minor_version" == "3" ]]
}

# Cloud-config templates are stored as a configmap in Harvester versions prior
# to and including v1.3.x.
# From Harvester v1.4.x onwards, cloud-config templates are stored as secrets.
# This means that during an upgrade from v1.3.x to v1.4.x or newer, the
# existing cloud-config templates must be migrated.
migrate_cloudconfig_templates(){
  echo -n "migrating cloud-config templates... "
  user_templates="$(get_cloudconfig_templates user)"
  network_templates="$(get_cloudconfig_templates network)"

  for manifest in "${user_templates[@]}" "${network_templates[@]}" ; do
    secret="$(gen_cloudconfig_template_secret "$manifest")"
    kubectl create -f - <<EOF > /dev/null 2>&1
$secret
EOF
    kubectl delete -f - <<EOF > /dev/null 2>&1
$manifest
EOF
  done
  echo "done"
}

get_cloudconfig_templates(){
  local type="$1"

  kubectl get configmaps \
    --all-namespaces \
    -l harvesterhci.io/cloud-init-template="$type" \
    -o yaml \
  | yq eval \
    -I=0 \
    -o=json \
    '.items[]'
}

get_server_version(){
  kubectl get settings.harvesterhci.io server-version -o yaml | yq '.value'
}

gen_cloudconfig_template_secret(){
  local manifest="$1"

  local name
  name="$(yq '.metadata.name' <<< "$manifest")"

  local namespace
  namespace="$(yq '.metadata.namespace' <<< "$manifest")"

  local annotations
  annotations="$(yq '.metadata.annotations' <<< "$manifest")"

  local labels
  labels="$(yq '.metadata.labels' <<< "$manifest")"

  local data
  data="$(yq '.data.cloudInit' <<< "$manifest")"

  yq eval \
    ".stringData.cloudInit = \"$data\"" \
    <<EOF
---
apiVersion: v1
kind: Secret
metadata:
  name: $name
  namespace: $namespace
  annotations: $annotations
  labels: $labels
EOF
}

check_version_range \
  && migrate_cloudconfig_templates \
  || echo "skip migrating cloud-config templates"
