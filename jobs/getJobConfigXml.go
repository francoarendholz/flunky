package jobs

import (
	_ "embed"

	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

func GetJobConfigXml(jobName string) {

	context := pongo2.Context{
		"jobName": jobName,
	}

	compiledGroovy := base.CompileGroovy(context, getJobConfigXml)

	if viper.GetBool("verbose") {
		println("Running Groovy script remotely\n")
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
