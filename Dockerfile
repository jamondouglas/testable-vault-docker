# THIS DOCKERFILE BASED ON https://github.com/voxxit/dockerfiles/blob/master/vault/Dockerfile

FROM voxxit/base:alpine

# only needed so we can wget the vault release from an https url
RUN apk --update add openssl 

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

COPY entrypoint /usr/local/bin/
RUN chmod u+x /usr/local/bin/entrypoint

EXPOSE 8200

ENTRYPOINT [ "entrypoint" ]
