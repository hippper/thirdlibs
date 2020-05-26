package utils

import (
	"bytes"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	validator "gopkg.in/go-playground/validator.v8"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var config *validator.Config = &validator.Config{TagName: "binding"}
var validate *validator.Validate = validator.New(config)

func init() {
	extra.RegisterFuzzyDecoders()
}

// Unserialize JSON
func UnserializeFromJSON(jsonstr string, st interface{}) error {
	d := json.NewDecoder(strings.NewReader(jsonstr))
	d.UseNumber()
	return d.Decode(st)
}

func DecodeJsonToStruct(msg string, req interface{}) error {
	err := UnserializeFromJSON(msg, req)
	if err != nil {
		return err
	}

	err = validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func SerializeToJSON(st interface{}) (string, error) {

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(st)

	return buf.String(), err
}
