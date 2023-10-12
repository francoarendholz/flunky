package manage

import (
	_ "embed"
)

//go:embed groovy/approvePendingSignatures.groovy
var ApprovePendingSignaturesGroovy string

//go:embed groovy/decodeAllSecrets.groovy
var decodeAllSecretsGroovy string

//go:embed groovy/systemMessage.groovy
var setSystemMessageGroovy string
