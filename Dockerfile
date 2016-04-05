# This Dockerfile is loosely based on voxxit/vault - https://github.com/voxxit/dockerfiles/blob/master/vault/Dockerfile

FROM voxxit/base:alpine

MAINTAINER "Pete Hodgson"
LABEL version="0.1"
LABEL description="A [Hashicorp Vault](https://www.vaultproject.io) server, pre-configured with a hardcoded root token. Intended for integration testing."

# only need openssl so we can wget the vault release from an https url
RUN apk --update add openssl jq

ARG VAULT_VERSION
ENV VAULT_VERSION ${VAULT_VERSION:-0.5.2}

RUN cd /tmp \
	&& wget -O vault.zip https://releases.hashicorp.com/vault/${VAULT_VERSION}/vault_${VAULT_VERSION}_linux_amd64.zip \
  && unzip vault.zip \
  && mv vault /usr/local/bin/ \
  && rm -f /tmp/vault*

COPY vault_inmem_config.hcl /etc/vault.hcl

COPY vault_init /usr/local/bin/
RUN chmod u+x /usr/local/bin/vault_init

COPY load_fixture_into_vault /usr/local/bin/
RUN chmod u+x /usr/local/bin/load_fixture_into_vault

COPY entrypoint /usr/local/bin/
RUN chmod u+x /usr/local/bin/entrypoint

EXPOSE 8200

ENTRYPOINT [ "entrypoint" ]
