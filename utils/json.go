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

func SerializeToJson(st interface{}) string {
	ba, _ := json.Marshal(st)
	jsonstr := string(ba)

	return jsonstr
}

func EncodeToJson(st interface{}) string {

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.Encode(st)

	return buf.String()
}

func DecodeFromJson(jsonstr string, st interface{}) error {
	d := json.NewDecoder(strings.NewReader(jsonstr))
	d.UseNumber()
	return d.Decode(st)
}

// add validate
func DecodeJsonToStruct(msg string, req interface{}) error {
	err := DecodeFromJson(msg, req)
	if err != nil {
		return err
	}

	err = validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func FormatJsonStr(instr string) string {
	var out bytes.Buffer
	json.Indent(&out, []byte(instr), "", "  ")

	return "\n" + out.String() + "\n"
}

func FormatStruct(inst interface{}) string {
	instr := SerializeToJson(inst)
	return FormatJsonStr(instr)
}

