package jobs

import "encoding/xml"

type UnoChoiceParameter struct {
	XMLName          xml.Name     `xml:"org.biouno.unochoice.ChoiceParameter"`
	Plugin           string       `xml:"plugin,attr"`
	Name             string       `xml:"name"`
	Description      string       `xml:"description"`
	RandomName       string       `xml:"randomName"`
	VisibleItemCount int          `xml:"visibleItemCount"`
	Script           GroovyScript `xml:"script"`
	ProjectName      string       `xml:"projectName"`
	ProjectFullName  string       `xml:"projectFullName"`
	ChoiceType       string       `xml:"choiceType"`
	Filterable       bool         `xml:"filterable"`
	FilterLength     int          `xml:"filterLength"`
}

type UnoCascadeChoiceParameter struct {
	XMLName              xml.Name      `xml:"org.biouno.unochoice.CascadeChoiceParameter"`
	Name                 string        `xml:"name"`
	Description          string        `xml:"description"`
	RandomName           string        `xml:"randomName"`
	VisibleItemCount     int           `xml:"visibleItemCount"`
	Script               GroovyScript  `xml:"script"`
	ProjectName          string        `xml:"projectName"`
	ProjectFullName      string        `xml:"projectFullName"`
	Parameters           LinkedHashMap `xml:"parameters"`
	ReferencedParameters string        `xml:"referencedParameters"`
	ChoiceType           string        `xml:"choiceType"`
	Filterable           bool          `xml:"filterable"`
	FilterLength         int           `xml:"filterLength"`
}

type UnoDynamicReferenceParameter struct {
	XMLName              xml.Name      `xml:"org.biouno.unochoice.DynamicReferenceParameter"`
	Name                 string        `xml:"name"`
	Description          string        `xml:"description"`
	RandomName           string        `xml:"randomName"`
	VisibleItemCount     int           `xml:"visibleItemCount"`
	Script               GroovyScript  `xml:"script"`
	ProjectName          string        `xml:"projectName"`
	ProjectFullName      string        `xml:"projectFullName"`
	Parameters           LinkedHashMap `xml:"parameters"`
	ReferencedParameters string        `xml:"referencedParameters"`
	ChoiceType           string        `xml:"choiceType"`
	OmitValueField       bool          `xml:"omitValueField"`
}

type GroovyScript struct {
	XMLName              xml.Name             `xml:"script"`
	SecureScript         SecureScript         `xml:"secureScript"`
	SecureFallbackScript SecureFallbackScript `xml:"secureFallbackScript"`
}

type SecureScript struct {
	XMLName xml.Name `xml:"secureScript"`
	Plugin  string   `xml:"plugin,attr"`
	Script  string   `xml:"script"`
	Sandbox bool     `xml:"sandbox"`
}

type SecureFallbackScript struct {
	XMLName xml.Name `xml:"secureFallbackScript"`
	Plugin  string   `xml:"plugin,attr"`
	Script  string   `xml:"script"`
	Sandbox bool     `xml:"sandbox"`
}

type ChoiceParameter struct {
	XMLName     xml.Name    `xml:"hudson.model.ChoiceParameterDefinition"`
	Name        string      `xml:"name"`
	Description string      `xml:"description"`
	Choices     StringArray `xml:"choices"`
}

type BooleanParameter struct {
	XMLName      xml.Name `xml:"hudson.model.BooleanParameterDefinition"`
	Name         string   `xml:"name"`
	Description  string   `xml:"description"`
	DefaultValue bool     `xml:"defaultValue"`
}

type ListGitBranchesParameter struct {
	XMLName            xml.Name `xml:"com.syhuang.hudson.plugins.listgitbranchesparameter.ListGitBranchesParameterDefinition"`
	Name               string   `xml:"name"`
	Description        string   `xml:"description"`
	RemoteURL          string   `xml:"remoteURL"`
	CredentialsId      string   `xml:"credentialsId"`
	TagFilter          string   `xml:"tagFilter"`
	BranchFilter       string   `xml:"branchFilter"`
	SortMode           string   `xml:"sortMode"`
	SelectedValue      string   `xml:"selectedValue"`
	QuickFilterEnabled bool     `xml:"quickFilterEnabled"`
	ListSize           int      `xml:"listSize"`
	DefaultValue       string   `xml:"defaultValue"`
	Type               string   `xml:"type"`
}

type LinkedHashMap struct {
	XMLName xml.Name `xml:"parameters"`
}

type StringArray struct {
	XMLName xml.Name `xml:"choices"`
	Choices []string `xml:"a>string"`
}

type TextParameter struct {
	XMLName      xml.Name `xml:"hudson.model.TextParameterDefinition"`
	Name         string   `xml:"name"`
	Description  string   `xml:"description"`
	DefaultValue string   `xml:"defaultValue"`
}

type StringParameter struct {
	XMLName      xml.Name `xml:"hudson.model.StringParameterDefinition"`
	Name         string   `xml:"name"`
	Description  string   `xml:"description"`
	DefaultValue string   `xml:"defaultValue"`
}

type PasswordParameter struct {
	XMLName      xml.Name `xml:"hudson.model.PasswordParameterDefinition"`
	Name         string   `xml:"name"`
	Description  string   `xml:"description"`
	DefaultValue string   `xml:"defaultValue"`
}
