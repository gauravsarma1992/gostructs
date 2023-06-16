# GoStructs

GoStructs is a package to provide information of a struct
using reflect


```bash
	decoder, _ := NewDecoder(&DecoderConfig{ShouldSnakeCase: true})
	result, _ := decoder.Decode(srv)
```

will return a result with the following struct
```go
	DecodedResult struct {
		Name       string                 `json:"name"`
		Attributes map[string]interface{} `json:"attributes"`
	}
```
