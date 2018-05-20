package objloader_test

import (
	"testing"

	"github.com/kpelelis/go-engine/objloader"
)

func facesAreEqual(a *objloader.Face, b *objloader.Face) bool {
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

func TestParseFace(t *testing.T) {
	var facesTestcases = []struct {
		data []byte
		Face *objloader.Face
	}{
		{[]byte("f 1 1 1"), &objloader.Face{-1, []objloader.FacePoint{
			objloader.FacePoint{1, -1, -1},
			objloader.FacePoint{1, -1, -1},
			objloader.FacePoint{1, -1, -1},
		}}},
		{[]byte("f 1//2 1/2/3 1/2"), &objloader.Face{-1, []objloader.FacePoint{
			objloader.FacePoint{1, -1, 2},
			objloader.FacePoint{1, 2, 3},
			objloader.FacePoint{1, 2, -1},
		}}},
		{[]byte("f 1//2 1/2/3 1/2 1/2/3"), &objloader.Face{-1, []objloader.FacePoint{
			objloader.FacePoint{1, -1, 2},
			objloader.FacePoint{1, 2, 3},
			objloader.FacePoint{1, 2, -1},
			objloader.FacePoint{1, 2, 3},
		}}},
		{[]byte("f 1/2/2 1/2 1/2"), &objloader.Face{-1, []objloader.FacePoint{
			objloader.FacePoint{1, 2, 2},
			objloader.FacePoint{1, 2, -1},
			objloader.FacePoint{1, 2, -1},
		}}},
		{[]byte("f 1/2"), nil},
		{[]byte("f 0.25 1 "), nil},
		{[]byte("0.25 1 "), nil},
		{[]byte("f 0.25 foobar "), nil},
		{[]byte(""), nil},
	}
	for _, testcase := range facesTestcases {
		name := "parseFace"
		t.Run(name, func(t *testing.T) {
			face, err := objloader.ParseFace(testcase.data)

			ok := (err == nil && facesAreEqual(face, testcase.Face)) || err != nil

			if !ok {
				t.Errorf("Failed testcase %v with error %v", string(testcase.data), err)
			}
		})
	}
}
