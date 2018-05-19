package objloader_test

import (
	"testing"

	"github.com/kpelelis/go-engine/objloader"
)

func UVsAreEqual(a *objloader.UV, b *objloader.UV) bool {
	return (a == nil && b == nil) || (a.U == b.U && a.V == b.V && a.W == b.W)
}

func TestParseUV(t *testing.T) {
	var uvtestcases = []struct {
		data []byte
		UV   *objloader.UV
	}{
		{[]byte("vt 1 1"), &objloader.UV{-1, 1, 1, -1}},
		{[]byte("vt 1 1 1"), &objloader.UV{-1, 1, 1, 1}},
		{[]byte("vt 0.25 1"), &objloader.UV{-1, 0.25, 1, -1}},
		{[]byte("vt 0.25 1 1"), &objloader.UV{-1, 0.25, 1, 1}},
		{[]byte("vt 0.25"), nil},
		{[]byte("0.25 1"), nil},
		{[]byte("vt 0.25 foobar "), nil},
		{[]byte(""), nil},
	}
	for _, testcase := range uvtestcases {
		name := "parseUV"
		t.Run(name, func(t *testing.T) {
			uv, err := objloader.ParseUV(testcase.data)

			ok := (err == nil && UVsAreEqual(uv, testcase.UV)) || (err != nil && uv == nil)

			if !ok {
				t.Errorf("Failed testcase %v with error %v", string(testcase.data), err)
			}
		})
	}
}
