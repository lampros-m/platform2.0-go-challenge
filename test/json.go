package test

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
)

// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// LooseJSONEq : Compares two json strings. Compares only the required key from expected object.
func LooseJSONEq(t TestingT, expected, got string) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	var exp interface{}
	err := json.Unmarshal([]byte(expected), &exp)
	if err != nil {
		t.Fatalf("Error in expected json: %s", err.Error())
		return
	}

	var g interface{}
	err = json.Unmarshal([]byte(got), &g)
	if err != nil {
		t.Fatalf("Error in got json: %s", err.Error())
		return
	}

	compareJSONField(t, "", exp, g)
}

func compareJSONField(t TestingT, path string, expected, got interface{}) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	if expected == nil && got == nil {
		return
	}

	et := reflect.TypeOf(expected)
	gt := reflect.TypeOf(got)

	if et != gt {
		if et == nil {
			t.Errorf("Wrong type of %s expected to be nil got %s", path, gt)
			return
		} else if gt == nil {
			errorField(t, path, expected, got)
			return
		}
		t.Errorf("Wrong type of %s expected %s got %#v", path, et, gt)
		return
	}

	switch tp := expected.(type) {
	case bool:
		expectJSONBool(t, path, tp, got.(bool))
	case float64:
		expectJSONFloat64(t, path, tp, got.(float64))
	case string:
		expectJSONString(t, path, tp, got.(string))
	case []interface{}:
		expectJSONArray(t, path, tp, got.([]interface{}))
	case map[string]interface{}:
		compareJSONMaps(t, path, tp, got.(map[string]interface{}))
	}
}

func compareJSONMaps(t TestingT, path string, expected, got map[string]interface{}) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	for k, e := range expected {
		currentPath := path + "." + k
		if g, found := got[k]; found {
			compareJSONField(t, currentPath, e, g)
		} else {
			t.Errorf("Expected to have key %s", currentPath)
		}
	}
}

func expectJSONBool(t TestingT, path string, expected, got bool) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	if expected != got {
		errorField(t, path, expected, got)
	}
}

func equalsFloat(a, b float64) bool {
	const e = .0001
	return math.Abs(a-b) > e
}

func expectJSONFloat64(t TestingT, path string, expected, got float64) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	if equalsFloat(expected, got) {
		errorField(t, path, expected, got)
	}
}

func expectJSONString(t TestingT, path string, expected, got string) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	if expected != got {
		errorField(t, path, expected, got)
	}
}

func expectJSONArray(t TestingT, path string, expected, got []interface{}) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	if len(expected) != len(got) {
		t.Errorf("Error in field %s array length expected %d got %d", path, len(expected), len(got))
		t.Errorf("Error %#v", got)
	}

	for i := range expected {
		if i >= len(got) {
			// error have already been printed for size difference
			return
		}
		currentPath := fmt.Sprintf("%s[%d]", path, i)
		compareJSONField(t, currentPath, expected[i], got[i])
	}
}

func errorField(t TestingT, path string, expected, got interface{}) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	t.Errorf("Error in field %s \n\texpected: %#v \n\tgot     : %#v", path, expected, got)
}
