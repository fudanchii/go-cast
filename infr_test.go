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

type RawJSON []byte

type JSONStruct map[string]any

func (j RawJSON) TryFrom(m JSONStruct) (RawJSON, error) {
	var (
		result []byte
		err    error
	)
	result, err = json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return RawJSON(result), nil
}

func (j JSONStruct) TryFrom(s RawJSON) (JSONStruct, error) {
	err := json.Unmarshal(s, &j)

	return j, err
}

func TestCastMethodTryIntoMarshalJson(t *testing.T) {
	m := make(JSONStruct)
	m["hello"] = "yay"

	result, err := TryFrom[JSONStruct, RawJSON](m).TryInto()

	if err != nil || string(result) != `{"hello":"yay"}` {
		t.Errorf("should result in correct json, got %s", result)
	}
}

func TestCastGlobalTryIntoMarshalJson(t *testing.T) {
	m := make(JSONStruct)
	m["hello"] = "world"

	result, err := TryInto[RawJSON](m)

	if err != nil || string(result) != `{"hello":"world"}` {
		t.Errorf("should result in correct json, got %s", result)
	}
}

func TestCastGlobalTryIntoSliceOfMarshalJson(t *testing.T) {
	cases := []JSONStruct{
		{
			"ping": "pong",
		},
		{
			"hello": "world",
		},
	}

	results := []RawJSON{
		[]byte(`{"ping":"pong"}`),
		[]byte(`{"hello":"world"}`),
	}

	actuals, err := TryIntoSliceOf[RawJSON](cases)

	if err != nil {
		t.Errorf("should return no error, got error : %s", err)
	}

	if string(actuals[0]) != string(results[0]) {
		t.Errorf("should returns correct json.")
	}

	if string(actuals[1]) != string(results[1]) {
		t.Errorf("should returns correct json.")
	}

	{
		cases := []struct {
			Raw   []RawJSON
			Error bool
		}{
			{
				Raw: []RawJSON{[]byte(`{"ping":"pong"}`)},
			},
			{
				Raw:   []RawJSON{[]byte(`{`)},
				Error: true,
			},
		}

		for _, c := range cases {
			actuals, err := TryIntoSliceOf[JSONStruct](c.Raw)

			if c.Error && err == nil {
				t.Errorf("should results in error.")
				continue
			}

			if (!c.Error) && err != nil {
				t.Errorf("should not results in error.")
				continue
			}

			if !c.Error && actuals == nil {
				t.Errorf("should not results in nil.")
			}
		}
	}
}

func TestCopyAsRef(t *testing.T) {
	c := 42
	d := CopyAsRef(c)
	e := &c
	*e = 3
	if c != 3 {
		t.Errorf("should affected by change from a pointer")
	}

	if *d == 3 {
		t.Errorf("should not affected by change to a copy")
	}
}
