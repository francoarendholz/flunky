package jobs

import (
	_ "embed"
)

//go:embed groovy/getJobConfigXml.groovy
var getJobConfigXml string

//go:embed groovy/runPipelineScript.groovy
var runPipelineScriptGroovy string

//go:embed templates/UnoChoiceParameter.groovy
var unoChoiceParameter string

//go:embed templates/UnoCascadeChoiceParameter.groovy
var unoCascadeChoiceParameter string

//go:embed templates/UnoDynamicReferenceParameter.groovy
var unoDynamicReferenceParameter string

//go:embed templates/BooleanParameter.groovy
var booleanParameter string

//go:embed templates/ChoiceParameter.groovy
var choiceParameter string

//go:embed templates/ListGitBranchesParameter.groovy
var listGitBranchesParameter string

//go:embed templates/TextParameter.groovy
var textParameter string

//go:embed templates/StringParameter.groovy
var stringParameter string

//go:embed templates/PasswordParameter.groovy
var passwordParameter string
