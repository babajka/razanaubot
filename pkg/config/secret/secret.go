package secret

import _ "embed"

type ConfigFile []byte

//go:embed secret.yaml
var Config ConfigFile
