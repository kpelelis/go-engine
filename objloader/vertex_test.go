package objloader_test

import (
	"testing"

	"github.com/kpelelis/go-engine/objloader"
)

func verticesAreEqual(a *objloader.Vertex, b *objloader.Vertex) bool {
	return (a == nil && b == nil) || (a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W)
}

func TestParseVertex(t *testing.T) {
	var vertextestcases = []struct {
		data   []byte
		Vertex *objloader.Vertex
	}{
		{[]byte("v 1 1 1"), &objloader.Vertex{-1, 1, 1, 1, -1}},
		{[]byte("v 1 1 1 1"), &objloader.Vertex{-1, 1, 1, 1, 1}},
		{[]byte("v 0.25 1 1 1"), &objloader.Vertex{-1, 0.25, 1, 1, 1}},
		{[]byte("v 0.25 1 1 1"), &objloader.Vertex{-1, 0.25, 1, 1, 1}},
		{[]byte("v 0.25 1 "), nil},
		{[]byte("0.25 1 "), nil},
		{[]byte("v 0.25 foobar "), nil},
		{[]byte(""), nil},
	}
	for _, testcase := range vertextestcases {
		name := "parseVertex"
		t.Run(name, func(t *testing.T) {
			vertex, err := objloader.ParseVertex(testcase.data)

			ok := (err == nil && verticesAreEqual(vertex, testcase.Vertex)) || (err != nil && vertex == nil)

			if !ok {
				t.Errorf("Failed testcase %v with error %v", string(testcase.data), err)
			}
		})
	}
}
