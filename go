#!/bin/bash

set -e

echo "from go script JAMON JAMON AJMON  "

DC=docker-compose
D=DOCKER

R="\x1B[1;31m"
G="\x1B[1;32m"
W="\x1B[0m"

function info {
  echo -e "${G}${1}${W}"
}


function ci {
  info "Setting up CI dependencies"
  echo " Setting up CI dependencies"
  shift

  case "$1" in
      snap)
        DOCKER_COMPOSE_VERSION=1.5.2  # latest supported version on SNAPCI
        : ${SNAP_CI?} # Validate that we're in SNAPCI
        login_docker
        pip install ${DC}==${DOCKER_COMPOSE_VERSION}
      ;;
       *) echo $"Unknown ci command"
        helptext
        exit 1
  esac
}

function login {
  shift

  case "${1-docker}" in
      docker) login_docker
      ;;
      *) echo $"Unknown dev command"
        helptext
        exit 1
  esac
}

function login_docker {
  info "Logging into docker"

  : "${DOCKER_USERNAME?}"
  : "${DOCKER_PASSWORD?}"
  : "${DOCKER_EMAIL?}"

  ${D} login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD} -e ${DOCKER_EMAIL}
}
