package manage

import (
	_ "embed"

	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

//go:embed groovy/decodeAllSecrets.groovy
var decodeAllSecretsGroovy string

func DecodeAllSecrets() {

	context := pongo2.Context{}

	compiledGroovy := base.CompileGroovy(context, decodeAllSecretsGroovy)

	if viper.GetBool("verbose") == true {
		println("Decoding all secrets in Jenkins\n")
		println("Compiled Groovy:\n\n")
		println(compiledGroovy)
		println("\n\n")
	}

	base.PostScriptRequest("scriptText", compiledGroovy)

}
