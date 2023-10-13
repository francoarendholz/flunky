package jobs

import (
	_ "embed"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

func GetJobsTree() {

	context := pongo2.Context{}

	compiledGroovy := base.CompileGroovy(context, getJobsTree)

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

	result = removeResultFromOutput(result)
	println(result)

}

func removeResultFromOutput(output string) string {
	return strings.Split(output, "Result:")[0]
}
