package objloader

import (
	"testing"
)

var normaltestcases = []struct {
	data   []byte
	Normal *Normal
}{
	{[]byte("n 1 1 1"), &Normal{-1, 1, 1, 1}},
	{[]byte("n 0.25 1 1"), &Normal{-1, 0.25, 1, 1}},
	{[]byte("n 0.25 1 1 1"), nil},
	{[]byte("n 0.25 1 "), nil},
	{[]byte("0.25 1 "), nil},
	{[]byte("n 0.25 foobar "), nil},
	{[]byte(""), nil},
}

func normalsAreEqual(a *Normal, b *Normal) bool {
	return (a == nil && b == nil) || (a.x == b.x && a.y == b.y && a.z == b.z)
}

func TestParseNormal(t *testing.T) {
	for _, testcase := range normaltestcases {
		name := "parseNormal"
		t.Run(name, func(t *testing.T) {
			normal, err := parseNormal(testcase.data)

			ok := (err == nil && normalsAreEqual(normal, testcase.Normal)) || (err != nil && normal == nil)

			if !ok {
				t.Errorf("Failed testcase %v with error %v", string(testcase.data), err)
			}
		})
	}
}
