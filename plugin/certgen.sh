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

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if ! [ -f ${SCRIPTDIR}/certgen ]; then
  echo "Error: certgen binary not found."
  exit 1
fi

case "${1:-"help"}" in
  "generate")
    if [[ $# < 2 ]]; then
      echo "Error: Please specify chart name"
      usage
      exit 1
    fi
    ${SCRIPTDIR}/certgen $@
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
