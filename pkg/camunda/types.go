package camunda

import (
	"net/http"
	"net/url"
	"reflect"
)

// SortOrder defines sorting order for queries which allow to sort by columns
type SortOrder string

const (
	Asc  SortOrder = "asc"
	Desc SortOrder = "desc"
)

// used by count queries
type CountResult struct {
	Count int
}

type Link struct {
	Method *string `json:"method,omitempty"`
	Href   *string `json:"href,omitempty"`
	Rel    *string `json:"rel,omitempty"`
}

func NewVariableBuilder() *VariableBuilder {
	return &VariableBuilder{variables: &Variables{}}
}

type VariableBuilder struct {
	variables *Variables
}

func (vb *VariableBuilder) addStringVar(value string) *VariableBuilder {
	return vb
}

func (vb *VariableBuilder) addIntVar(value int) *VariableBuilder {
	return vb
}

func (vb *VariableBuilder) addLongVar(value int64) *VariableBuilder {
	return vb
}

func (vb *VariableBuilder) addDoubleVar(value float64) *VariableBuilder {
	return vb
}

func (vb *VariableBuilder) addFloatVar(value float32) *VariableBuilder {
	return vb
}

func (vb *VariableBuilder) addBooleanVar(value bool) *VariableBuilder {
	return vb
}

func (vb *VariableBuilder) addObjectVar(value interface{}) *VariableBuilder {
	return vb
}

func (vb *VariableBuilder) Build() *Variables {
	return vb.variables
}

type Variables struct {
	variables []*Variable
}

//A JSON object containing additional, value-type-dependent properties.
//For serialized variables of type Object, the following properties can be provided:
//
//objectTypeName: A string representation of the object's type name.
//serializationDataFormat: The serialization format used to store the variable.
//For serialized variables of type File, the following properties can be provided:
//
//filename: The name of the file. This is not the variable name but the name that will be used when downloading the file again.
//mimetype: The MIME type of the file that is being uploaded.
//encoding: The encoding of the file that is being uploaded.
type ValueInfo struct {
	// file
	Filename *string `json:"filename,omitempty"`
	MimeType *string `json:"mimetype,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	// object
	ObjectTypeName          *string `json:"objectTypeName,omitempty"`
	SerializationDataFormat *string `json:"serializationDataFormat,omitempty"`
}

type Variable struct {
	Type      *reflect.Kind `json:"type,omitempty"`  // The value type of the variable. (String, Boolean, Number, Object)
	Value     interface{}   `json:"value,omitempty"` // can be value of type String / Number / Boolean / Object
	ValueInfo *ValueInfo    `json:"valueInfo,omitempty"`
}

type PaginationQueryParams struct {
	firstResult int
	maxResults  int
}

type HttpConfiguration struct {
	Endpoint   *url.URL
	UserAgent  string
	BasicAuth  string
	HTTPClient *http.Client
}

type ClientOption interface {
	Apply(configuration *HttpConfiguration)
}
