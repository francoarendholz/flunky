package manage

import (
	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

func SetSystemMessage(message string) {

	context := pongo2.Context{
		"message": message,
	}

	compiledGroovy := base.CompileGroovy(context, setSystemMessageGroovy)

	if viper.GetBool("verbose") == true {
		println("Setting Jenkins System Message\n")
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
