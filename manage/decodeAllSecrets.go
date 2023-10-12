package manage

import (
	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

func DecodeAllSecrets() {

	context := pongo2.Context{}

	compiledGroovy := base.CompileGroovy(context, decodeAllSecretsGroovy)

	if viper.GetBool("verbose") == true {
		println("Decoding all secrets in Jenkins\n")
		println("Compiled Groovy:\n\n")
		println(compiledGroovy)
		println("\n\n")
	}

	result, err := base.PostScriptRequest("scriptText", compiledGroovy)

	if err != nil {
		println(err)
	}

	println(result)

}
