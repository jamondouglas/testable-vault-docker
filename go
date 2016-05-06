DC=docker-compose

function ci {
  info "Setting up CI dependencies"

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
