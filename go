#!/bin/bash

set -e

echo "from go script JAMON JAMON AJMON  "

DC=docker-compose
D=docker

R="\x1B[1;31m"
G="\x1B[1;32m"
W="\x1B[0m"

function helptext {
    echo "Usage: ./go <command> [sub-command]"
    echo ""
    echo "Available commands are:"
    echo "    flyway [cmd] Run database migrations"
    echo "        baseline Create a new database"
    echo "        migrate  Run database migrations"
    echo "        info     See current status"
    echo "        *        See: (https://flywaydb.org/)"
    echo "    sbt [cmd]    SBT commands (http://www.scala-sbt.org/)"
    echo "    test [type]  Run tests"
    echo "       all       Default, run unit and integration tests"
    echo "       unit      Run unit tests"
    echo "       it        Run integration tests"
    echo "       smoke     Run smoke tests"
    echo "    build [type] Build docker image(s)"
    echo "       all       Default, build everything"
    echo "       service   Build the service"
    echo "       migration Build the docker migration image"
    echo "       coverage  Code coverage"
    echo "    push [type]  Push docker image(s) to dockerhub"
    echo "       all       Default, build everything"
    echo "       coverage  Code coverage"
    echo "       service   Default, build everything"
    echo "       migration Default, build everything"
    echo "    deploy [env] Deploy image from DockerHub"
    echo "    dev          Development command"
    echo "       start     Start service locally"
    echo "       docs      Start local docs server"
    echo "       stop      Stop service locally"
    echo "       clean     Clean data"
    echo "    login docker Login to docker"
    echo "    ci           CI-specific commands"
    echo "      snap       Setup dependencies for SnapCI"
}

function info {
  echo -e "${G}${1}${W}"
}


function ci {
  info "Setting up CI dependencies"
  shift

  case "$1" in
      snap)
        DOCKER_COMPOSE_VERSION=1.5.2  # latest supported version on SNAPCI
        : ${SNAP_CI?} # Validate that were in SNAPCI
        pip -V
        login_docker
        pip -v install ${DC}==${DOCKER_COMPOSE_VERSION}
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

case "$1" in
    help) helptext
    ;;
    flyway) flyway $@
    ;;
    sbt) sbt_ $@
    ;;
    build) build $@
    ;;
    test) test_ $@
    ;;
    dev) dev $@
    ;;
    login) login $@
    ;;
    ci) ci $@
    ;;
    push) push $@
    ;;
    deploy) deploy $@
    ;;
    *)
      helptext
      error $"Usage: $0 {help|flyway|sbt|build|test|dev|ci|push|deploy}"
      exit 1
esac
