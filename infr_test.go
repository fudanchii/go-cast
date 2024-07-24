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

	str := From[int, Str](num).Into()

	if string(str) != "42" {
		t.Errorf("should equal 42, got %s", str)
	}
}

func TestCastGlobalIntoStrInt(t *testing.T) {
	num := 121

	str := Into[Str](num)

	if string(str) != "121" {
		t.Errorf("should equal 121, got %s", str)
	}
}

func TestCastSlices(t *testing.T) {
	nums := []int{8, 42, 1337}
	testCases := []struct {
		Expected string
	}{
		{Expected: "8"},
		{Expected: "42"},
		{Expected: "1337"},
	}

	slices := IntoSliceOf[Str](nums)

	for idx, tc := range testCases {
		if tc.Expected != string(slices[idx]) {
			t.Errorf("should equal to \"%s\", got %s", tc.Expected, slices[idx])
		}
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

	result, err := TryFrom[map[string]string, JSON](m).TryInto()

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
