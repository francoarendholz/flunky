package jobs

import (
	_ "embed"

	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

//go:embed groovy/getJobConfigXml.groovy
var getJobConfigXml string

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

	base.PostScriptRequest("scriptText", compiledGroovy)

}
