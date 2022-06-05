package jsondto

type API interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
	MarshalIndent(interface{}, string, string) ([]byte, error)
}

var AsObject API = api

func Marshal(v interface{}) ([]byte, error) {
	return api.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return api.Unmarshal(data, v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return api.MarshalIndent(v, prefix, indent)
}
