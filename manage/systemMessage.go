package manage

import (
	_ "embed"

	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

//go:embed groovy/systemMessage.groovy
var setSystemMessageGroovy string

func SetSystemMessage(message string) {

	context := pongo2.Context{
		"message": message,
	}

	setMessageGroovy := base.CompileGroovy(context, setSystemMessageGroovy)

	if viper.GetBool("verbose") == true {
		println("Setting Jenkins System Message\n")
		println("Compiled Groovy:\n\n")
		println(setMessageGroovy)
		println("\n\n")
	}

	base.PostScriptRequest("scriptText", setMessageGroovy)

}
