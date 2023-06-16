package gostructs

import (
	"reflect"

	"github.com/buxizhizhoum/inflection"
)

type (
	Decoder struct {
		resource interface{}
		config   *DecoderConfig
	}
	DecoderConfig struct {
		ShouldSnakeCase bool
	}
	DecodedResult struct {
		Name       string                 `json:"name"`
		Attributes map[string]interface{} `json:"attributes"`
	}
)

func NewDecoder(config *DecoderConfig) (decoder *Decoder, err error) {
	decoder = &Decoder{
		config: config,
	}
	return
}

func (decoder *Decoder) convertKey(key string) (result string) {
	result = key
	if decoder.config.ShouldSnakeCase {
		result = inflection.Underscore(key)
	}
	return
}

func (decoder *Decoder) Decode(resource interface{}) (result *DecodedResult, err error) {
	var (
		resourceType  reflect.Type
		resourceValue reflect.Value
	)

	resourceType = reflect.TypeOf(resource)
	resourceValue = reflect.ValueOf(resource)

	result = &DecodedResult{
		Name:       decoder.convertKey(resourceType.Name()),
		Attributes: make(map[string]interface{}),
	}

	for idx := 0; idx < resourceType.NumField(); idx++ {
		var (
			attrKey string
			attrVal interface{}
		)
		attrKey = resourceType.Field(idx).Name
		attrVal = resourceValue.FieldByName(attrKey).Interface()

		result.Attributes[decoder.convertKey(attrKey)] = attrVal
	}
	return
}
