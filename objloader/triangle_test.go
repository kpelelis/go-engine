package objloader

import (
	"testing"
)

var triangletestcases = []struct {
	data     []byte
	Triangle *Triangle
}{
	{[]byte("f 1 1 1"), &Triangle{-1, [3]TrianglePoint{
		TrianglePoint{-1, 1, -1, -1},
		TrianglePoint{-1, 1, -1, -1},
		TrianglePoint{-1, 1, -1, -1},
	}}},
	{[]byte("f 1//2 1/2/3 1/2"), &Triangle{-1, [3]TrianglePoint{
		TrianglePoint{-1, 1, -1, 2},
		TrianglePoint{-1, 1, 2, 3},
		TrianglePoint{-1, 1, 2, -1},
	}}},
	{[]byte("f 1/2/2 1/2 1/2"), &Triangle{-1, [3]TrianglePoint{
		TrianglePoint{-1, 1, 2, 2},
		TrianglePoint{-1, 1, 2, -1},
		TrianglePoint{-1, 1, 2, -1},
	}}},
	{[]byte("f 1/2"), nil},
	{[]byte("f 0.25 1 "), nil},
	{[]byte("0.25 1 "), nil},
	{[]byte("f 0.25 foobar "), nil},
	{[]byte(""), nil},
}

func trianglesAreEqual(a *Triangle, b *Triangle) bool {
	if a.index != b.index {
		return false
	}

	for i, pa := range a.Points {
		pb := b.Points[i]
		if pb.vertexIndex != pa.vertexIndex ||
			pa.UVIndex != pb.UVIndex ||
			pa.normalIndex != pb.normalIndex {
			return false
		}
	}
	return true
}

func TestParseTriangle(t *testing.T) {
	for _, testcase := range triangletestcases {
		name := "parseTriangle"
		t.Run(name, func(t *testing.T) {
			triangle, err := parseTriangle(testcase.data)

			ok := (err == nil && trianglesAreEqual(triangle, testcase.Triangle)) || (err != nil && triangle == nil)

			if !ok {
				t.Errorf("Failed testcase %v with error %v", string(testcase.data), err)
			}
		})
	}
}
