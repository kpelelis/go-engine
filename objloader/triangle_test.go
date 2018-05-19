package objloader_test

import (
	"testing"

	"github.com/kpelelis/go-engine/objloader"
)

func trianglesAreEqual(a *objloader.Triangle, b *objloader.Triangle) bool {
	if a.Index != b.Index {
		return false
	}

	for i, pa := range a.Points {
		pb := b.Points[i]
		if pb.VertexIndex != pa.VertexIndex ||
			pa.UVIndex != pb.UVIndex ||
			pa.NormalIndex != pb.NormalIndex {
			return false
		}
	}
	return true
}

func TestParseTriangle(t *testing.T) {
	var triangletestcases = []struct {
		data     []byte
		Triangle *objloader.Triangle
	}{
		{[]byte("f 1 1 1"), &objloader.Triangle{-1, [3]objloader.TrianglePoint{
			objloader.TrianglePoint{1, -1, -1},
			objloader.TrianglePoint{1, -1, -1},
			objloader.TrianglePoint{1, -1, -1},
		}}},
		{[]byte("f 1//2 1/2/3 1/2"), &objloader.Triangle{-1, [3]objloader.TrianglePoint{
			objloader.TrianglePoint{1, -1, 2},
			objloader.TrianglePoint{1, 2, 3},
			objloader.TrianglePoint{1, 2, -1},
		}}},
		{[]byte("f 1/2/2 1/2 1/2"), &objloader.Triangle{-1, [3]objloader.TrianglePoint{
			objloader.TrianglePoint{1, 2, 2},
			objloader.TrianglePoint{1, 2, -1},
			objloader.TrianglePoint{1, 2, -1},
		}}},
		{[]byte("f 1/2"), nil},
		{[]byte("f 0.25 1 "), nil},
		{[]byte("0.25 1 "), nil},
		{[]byte("f 0.25 foobar "), nil},
		{[]byte(""), nil},
	}
	for _, testcase := range triangletestcases {
		name := "parseTriangle"
		t.Run(name, func(t *testing.T) {
			triangle, err := objloader.ParseTriangle(testcase.data)

			ok := (err == nil && trianglesAreEqual(triangle, testcase.Triangle)) || (err != nil && triangle == nil)

			if !ok {
				t.Errorf("Failed testcase %v with error %v", string(testcase.data), err)
			}
		})
	}
}
