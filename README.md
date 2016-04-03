# testable-vault

This is a dockerized [Hashicorp Vault](https://www.vaultproject.io) server intended for integration testing. 

This container runs a Vault server with a hardcoded `TEST_VAULT_ROOT_TOKEN` root token, which of course is **VERY INSECURE**. This server is also configured to run with TLS disabled (it's accessed via http, not https), and uses the ephemeral `inmem` secret backend, which means that it start with a clean slate every time it is run.
