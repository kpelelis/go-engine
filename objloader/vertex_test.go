package objloader

import (
	"testing"
)

var vertextestcases = []struct {
	data   []byte
	Vertex *Vertex
}{
	{[]byte("v 1 1 1"), &Vertex{-1, 1, 1, 1, -1}},
	{[]byte("v 1 1 1 1"), &Vertex{-1, 1, 1, 1, 1}},
	{[]byte("v 0.25 1 1 1"), &Vertex{-1, 0.25, 1, 1, 1}},
	{[]byte("v 0.25 1 1 1"), &Vertex{-1, 0.25, 1, 1, 1}},
	{[]byte("v 0.25 1 "), nil},
	{[]byte("0.25 1 "), nil},
	{[]byte("v 0.25 foobar "), nil},
	{[]byte(""), nil},
}

func verticesAreEqual(a *Vertex, b *Vertex) bool {
	return (a == nil && b == nil) || (a.x == b.x && a.y == b.y && a.z == b.z && a.w == b.w)
}

func TestParseVertex(t *testing.T) {
	for _, testcase := range vertextestcases {
		name := "parseVertex"
		t.Run(name, func(t *testing.T) {
			vertex, err := parseVertex(testcase.data)

			ok := (err == nil && verticesAreEqual(vertex, testcase.Vertex)) || (err != nil && vertex == nil)

			if !ok {
				t.Errorf("Failed testcase %v with error %v", string(testcase.data), err)
			}
		})
	}
}
