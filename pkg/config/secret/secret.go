package secret

import _ "embed"

//go:embed secret.yaml
var SecretFile []byte
