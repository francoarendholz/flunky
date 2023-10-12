package jobs

import (
	"encoding/xml"
	"reflect"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/flosch/pongo2/v6"
	"github.com/francoarendholz/flunky/base"
	"github.com/spf13/viper"
)

type XPathSearchStruct struct {
	SearchString string
	Type         reflect.Type
}

func ConvertJobParams(jobName string) {

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

	resultXml, err := base.PostScriptRequest("scriptText", compiledGroovy)

	if err != nil {
		println(err)
	}

	// Remove the first line. Workaround as only XML 1.0 is supported but Jenkins returns XML 1.1
	// As soon as Go supports XML 1.1, this workaround should be removed
	resultXml = removeFirstLine(resultXml)

	doc, err := xmlquery.Parse(strings.NewReader(resultXml))
	if err != nil {
		panic(err)
	}

	// XPath queries to find the different parameter types
	queryStructMap := []XPathSearchStruct{
		{
			SearchString: "//properties/hudson.model.ParametersDefinitionProperty/*/org.biouno.unochoice.ChoiceParameter",
			Type:         reflect.TypeOf(UnoChoiceParameter{}),
		},
		{
			SearchString: "//properties/hudson.model.ParametersDefinitionProperty/*/org.biouno.unochoice.CascadeChoiceParameter",
			Type:         reflect.TypeOf(UnoCascadeChoiceParameter{}),
		},
		{
			SearchString: "//properties/hudson.model.ParametersDefinitionProperty/*/hudson.model.BooleanParameterDefinition",
			Type:         reflect.TypeOf(BooleanParameter{}),
		},
		{
			SearchString: "//properties/hudson.model.ParametersDefinitionProperty/*/hudson.model.ChoiceParameterDefinition",
			Type:         reflect.TypeOf(ChoiceParameter{}),
		},
		{
			SearchString: "//properties/hudson.model.ParametersDefinitionProperty/*/com.syhuang.hudson.plugins.listgitbranchesparameter.ListGitBranchesParameterDefinition",
			Type:         reflect.TypeOf(ListGitBranchesParameter{}),
		},
	}

	// Create a slice of interfaces to store the parameters
	parameters := []interface{}{}

	// Loop through the XPath queries
	for _, xPathQuery := range queryStructMap {
		// Loop through the XML nodes found by the XPath query
		for _, n := range xmlquery.Find(doc, xPathQuery.SearchString) {
			// Create a new instance of the struct type
			parameter := reflect.New(xPathQuery.Type).Interface()
			// Unmarshal the XML node into the struct
			xml.Unmarshal([]byte(n.OutputXML(true)), parameter)
			parameters = append(parameters, parameter)
		}
	}

	println(generateJenkinsPipelineCode(parameters))

}

func removeFirstLine(input string) string {
	lines := strings.SplitAfterN(input, "", 2)
	if len(lines) == 2 {
		return lines[1]
	}
	return ""
}

func generateJenkinsPipelineCode(parameters []interface{}) string {

	parameterOutput := `
properties([
	parameters([
`
	// Loop through the parameters and generate the Jenkins Pipeline code
	for _, param := range parameters {
		switch p := param.(type) {
		// Active Choices Choice Parameter
		case *UnoChoiceParameter:
			context := pongo2.Context{
				"name":            p.Name,
				"choiceType":      p.ChoiceType,
				"description":     p.Description,
				"filterable":      strings.ToLower(strconv.FormatBool(p.Filterable)),
				"filterLength":    p.FilterLength,
				"scriptClass":     "GroovyScript",
				"script":          p.Script.SecureScript.Script,
				"scriptSandbox":   strings.ToLower(strconv.FormatBool(p.Script.SecureScript.Sandbox)),
				"fallbackScript":  p.Script.SecureFallbackScript.Script,
				"fallbackSandbox": strings.ToLower(strconv.FormatBool(p.Script.SecureFallbackScript.Sandbox)),
			}

			compiledGroovy := base.CompileGroovy(context, unoChoiceParameter)
			parameterOutput += compiledGroovy + ",\n"

		// Active Choices Reactive Parameter
		case *UnoCascadeChoiceParameter:
			context := pongo2.Context{
				"name":                 p.Name,
				"choiceType":           p.ChoiceType,
				"description":          p.Description,
				"filterable":           strings.ToLower(strconv.FormatBool(p.Filterable)),
				"filterLength":         p.FilterLength,
				"scriptClass":          "GroovyScript",
				"script":               p.Script.SecureScript.Script,
				"scriptSandbox":        strings.ToLower(strconv.FormatBool(p.Script.SecureScript.Sandbox)),
				"fallbackScript":       p.Script.SecureFallbackScript.Script,
				"fallbackSandbox":      strings.ToLower(strconv.FormatBool(p.Script.SecureFallbackScript.Sandbox)),
				"referencedParameters": p.ReferencedParameters,
			}

			compiledGroovy := base.CompileGroovy(context, unoCascadeChoiceParameter)
			parameterOutput += compiledGroovy + ",\n"

		case *ListGitBranchesParameter:
			context := pongo2.Context{
				"name":               p.Name,
				"description":        p.Description,
				"branchFilter":       p.BranchFilter,
				"selectedValue":      p.SelectedValue,
				"tagFilter":          p.TagFilter,
				"sortMode":           p.SortMode,
				"remoteURL":          p.RemoteURL,
				"credentialsId":      p.CredentialsId,
				"listSize":           p.ListSize,
				"defaultValue":       p.DefaultValue,
				"type":               p.Type,
				"quickFilterEnabled": strings.ToLower(strconv.FormatBool(p.QuickFilterEnabled)),
			}

			compiledGroovy := base.CompileGroovy(context, listGitBranchesParameter)
			parameterOutput += compiledGroovy + ",\n"

		// Boolean Parameter
		case *BooleanParameter:
			context := pongo2.Context{
				"name":         p.Name,
				"description":  p.Description,
				"defaultValue": strings.ToLower(strconv.FormatBool(p.DefaultValue)),
			}

			compiledGroovy := base.CompileGroovy(context, booleanParameter)
			parameterOutput += compiledGroovy + ",\n"

		// Choice Parameter
		case *ChoiceParameter:
			context := pongo2.Context{
				"name":        p.Name,
				"description": p.Description,
				"choices":     strings.Join(p.Choices.Choices, "\\n"),
			}

			compiledGroovy := base.CompileGroovy(context, choiceParameter)
			parameterOutput += compiledGroovy + ",\n"
		}
	}

	parameterOutput += `
	])	
])
	`

	return parameterOutput
}
