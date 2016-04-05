# testable-vault

This is a dockerized [Hashicorp Vault](https://www.vaultproject.io) server intended for integration testing. 

This container runs a Vault server with a hardcoded `TEST_VAULT_ROOT_TOKEN` root token, which of course is **VERY INSECURE**. This server is also configured to run with TLS disabled (it's accessed via http, not https), and uses the ephemeral `inmem` secret backend, which means that it start with a clean slate every time it is run.

This image also supports initializing each vault instance with some fixture data on boot. This provides an easy way to start up a vault instance with secrets already in place. You provide fixture data by attaching a `/etc/vault_fixtures.json` volume to the container. You can look at the tests for the image for an example of how this works.

## Usage

### from Docker Hub
There are images for recent versions of Vault pushed to Docker Hub. To start up a testable vault server just do:
```
docker run -ti moredip/testable-vault:0.5.0
```

### pre-configured secrets
At launch the vault server will look for a file at `/etc/vault_fixtures.json`. If that file is present (i.e. via volume mount) then it will be used to pre-populate the server with an initial set of secrets. Here's an example showing the file format:
``` json
{
  "secret/an-api-key": {
    "key": "12141511",
    "root_url": "http://example.com"
  },
  "secret/another-secret": {
    "value": "the secret value"
  }
}
```

### building from source
If you want to build an image locally for some reason you can do something like:

```
docker build --build-arg=VAULT_VERSION=0.5.0 -t testable-vault:0.5.0 .
```

