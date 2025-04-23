package jsoncolor

import (
	"encoding/json"
	"github.com/fatih/color"
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

func Test_MatchesGoForScalarFieldAndForListWithElements(t *testing.T) {
	color.NoColor = true // color not important to this test

	v := struct {
		Number int   `json:"number"`
		List   []int `json:"list"`
	}{
		Number: 123,
		List:   []int{1, 2, 3},
	}
	goJsonBuf, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	jsonColorBuf, err := MarshalIndent(v, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	goJsonStr := string(goJsonBuf)
	jsonColorStr := string(jsonColorBuf)

	if goJsonStr != jsonColorStr {
		t.Errorf("jsoncolor output does not match go's json output. jsoncolor and go respectively:\n%q\n%q\n", jsonColorStr, goJsonStr)
	}
}
