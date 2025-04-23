package jsoncolor

import (
	"encoding/json"
	"log"
	"testing"
)

// https://github.com/nwidger/jsoncolor/issues/7
func Test_Issue7(t *testing.T) {
	v := struct {
		Null   *string `json:"null"`
		True   bool    `json:"true"`
		False  bool    `json:"false"`
		Number int     `json:"number"`
		String string  `json:"string"`
	}{
		True:   true,
		False:  false,
		Number: 123,
		String: "string",
	}
	goJsonBuf, err := json.MarshalIndent(v.String, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	jsonColorBuf, err := MarshalIndent(v.String, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	goJsonStr := string(goJsonBuf)
	jsonColorStr := string(jsonColorBuf)

	if goJsonStr != jsonColorStr {
		t.Errorf("jsoncolor output %q does not match go's json output %s", jsonColorStr, goJsonStr)
	}
}
