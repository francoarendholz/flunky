package jobs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

func RunPipelineScript(pipelineScriptPath string) {

	// Generate a dynamic job name
	jobName := "flunky-" + time.Now().Format("20060102150405")

	// Get the pipeline script
	scriptBytes, err := os.ReadFile(pipelineScriptPath)
	if err != nil {
		fmt.Print(err)
	}

	scriptString := string(scriptBytes)

	escapedScript := strings.Replace(scriptString, "'", "\\'", -1)

	context := pongo2.Context{
		"jobName":        jobName,
		"pipelineScript": escapedScript,
	}

	compiledGroovy := base.CompileGroovy(context, runPipelineScriptGroovy)

	if viper.GetBool("verbose") {
		println("Running pipeline script remotely\n")
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
