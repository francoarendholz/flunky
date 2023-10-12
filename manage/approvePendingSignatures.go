package manage

import (
	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

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

	result, err := base.PostScriptRequest("scriptText", compiledGroovy)

	if err != nil {
		println(err)
	}

	println(result)

}
