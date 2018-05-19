package objloader_test

import (
	"testing"

	"github.com/kpelelis/go-engine/objloader"
)

func normalsAreEqual(a *objloader.Normal, b *objloader.Normal) bool {
	return (a == nil && b == nil) || (a.X == b.X && a.Y == b.Y && a.Z == b.Z)
}

func TestParseNormal(t *testing.T) {
	var normaltestcases = []struct {
		data   []byte
		Normal *objloader.Normal
	}{
		{[]byte("n 1 1 1"), &objloader.Normal{-1, 1, 1, 1}},
		{[]byte("n 0.25 1 1"), &objloader.Normal{-1, 0.25, 1, 1}},
		{[]byte("n 0.25 1 1 1"), nil},
		{[]byte("n 0.25 1 "), nil},
		{[]byte("0.25 1 "), nil},
		{[]byte("n foobar 2 1"), nil},
		{[]byte("n 0.25 foobar "), nil},
		{[]byte(""), nil},
	}
	for _, testcase := range normaltestcases {
		name := "parseNormal"
		t.Run(name, func(t *testing.T) {
			normal, err := objloader.ParseNormal(testcase.data)

			ok := (err == nil && normalsAreEqual(normal, testcase.Normal)) || (err != nil && normal == nil)

			if !ok {
				t.Errorf("Failed testcase %v with error %v", string(testcase.data), err)
			}
		})
	}
}
