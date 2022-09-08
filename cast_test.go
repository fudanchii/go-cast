package cast

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

type Str string

func (s Str) From(a int) Str {
	return Str(strconv.Itoa(a))
}

func TestCastIntoStrInt(t *testing.T) {
	num := 42

	str := Into[Str](num)

	if string(str) != "42" {
		t.Errorf("should equal 42, got %s", str)
	}
}

type JSON []byte

func (j JSON) From(m map[string]string) JSON {
	var (
		result []byte
		err    error
	)
	result, err = json.Marshal(m)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
		result = []byte("")
	}
	return JSON(result)
}

func TestCastMarshalJson(t *testing.T) {
	m := make(map[string]string)
	m["hello"] = "yay"

	result := Into[JSON](m)

	if string(result) != `{"hello":"yay"}` {
		t.Errorf("should result in correct json, got %s", result)
	}
}
