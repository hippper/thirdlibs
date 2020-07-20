package utils

import (
	"bytes"
	json2 "encoding/json"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

func FormatJsonStr(instr string) string {
	var out bytes.Buffer
	json2.Indent(&out, []byte(instr), "", "  ")

	return "\n" + out.String() + "\n"
}

func FormatStruct(inst interface{}) string {
	instr := SerializeToJson(inst)
	return FormatJsonStr(instr)
}
