package manage

import (
	_ "embed"

	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

//go:embed groovy/approvePendingSignatures.groovy
var ApprovePendingSignaturesGroovy string

func ApprovePendingSignatures(force bool) {

	var context pongo2.Context

	if force == true {
		context = pongo2.Context{
			"force": "true",
		}
	} else {
		context = pongo2.Context{
			"force": "false",
		}
	}

	compiledGroovy := base.CompileGroovy(context, ApprovePendingSignaturesGroovy)

	if viper.GetBool("verbose") == true {
		println("Approving all pending script signatures.\n")
		println("Compiled Groovy:\n\n")
		println(compiledGroovy)
		println("\n\n")
	}

	base.PostScriptRequest("scriptText", compiledGroovy)

}
