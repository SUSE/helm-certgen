#!/bin/bash

set -eu

usage() {
cat << EOF
certgen is a Helm CLI plugin to help generate TLS certificates 
for the application/charts being deployed via Kubernetes Helm.

Usage:
  helm certgen CMD [CHART] 
  
Example:
  helm certgen generate suse/certgen-demo --namespace certgen-demo --name demo-web

Available Commands:
  generate    generates certificates using cert.yaml from chart definition

EOF
}


is_help() {
  case "$1" in
  "-h")
    return 0
    ;;
  "--help")
    return 0
    ;;
  "help")
    return 0
    ;;
  *)
    return 1
    ;;
esac
}



if [[ $# < 1 ]]; then
  usage
  exit 1
fi

if ! type "helm-certgen" > /dev/null; then
  echo "helm-certgen client needs to be installed"
  exit 1
fi

case "${1:-"help"}" in
  "generate")
    if [[ $# < 2 ]]; then
      echo "Error: Please specify chart name"
      usage
      exit 1
    fi
    helm-certgen $@
    ;;
  "help")
    usage
    ;;
  "--help")
    usage
    ;;
  "-h")
    usage
    ;;
  *)
    usage
    exit 1
    ;;
esac

exit 0
