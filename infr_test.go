package infr

import (
	"encoding/json"
	"strconv"
	"testing"
)

type Str string

func (s Str) From(a int) Str {
	return Str(strconv.Itoa(a))
}

func TestCastMethodIntoStrInt(t *testing.T) {
	num := 42

	str := FI[int, Str]{Source: num}.Into()

	if string(str) != "42" {
		t.Errorf("should equal 42, got %s", str)
	}
}

func TestCastGlobalIntoStrInt(t *testing.T) {
	num := 121

	str := Into[Str](num)

	if string(str) != "121" {
		t.Errorf("should equal 42, got %s", str)
	}
}

type JSON []byte

func (j JSON) TryFrom(m map[string]string) (JSON, error) {
	var (
		result []byte
		err    error
	)
	result, err = json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return JSON(result), nil
}

func TestCastMethodTryIntoMarshalJson(t *testing.T) {
	m := make(map[string]string)
	m["hello"] = "yay"

	mp := TFI[map[string]string, JSON]{Source: m}
	result, err := mp.TryInto()

	if err != nil || string(result) != `{"hello":"yay"}` {
		t.Errorf("should result in correct json, got %s", result)
	}
}

func TestCastGlobalTryIntoMarshalJson(t *testing.T) {
	m := make(map[string]string)
	m["hello"] = "world"

	result, err := TryInto[JSON](m)

	if err != nil || string(result) != `{"hello":"world"}` {
		t.Errorf("should result in correct json, got %s", result)
	}
}
