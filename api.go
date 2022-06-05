package jsondto

import jsoniter "github.com/json-iterator/go"

var (
	api = jsoniter.Config{
		EscapeHTML:                    true,
		ObjectFieldMustBeSimpleString: true,
		CaseSensitive:                 true,
	}.Froze()
)

func init() {
	api.RegisterExtension(nillableTypesEmptyOnNilExt{})
}
