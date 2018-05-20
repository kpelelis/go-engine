package objloader

import (
	"bytes"
	"fmt"

	"github.com/kpelelis/go-engine/math"
)

type FacePoint struct {
	VertexIndex int32
	UVIndex     int32
	NormalIndex int32
}

type Face struct {
	Index  int32
	Points []FacePoint
}

var spaceSep = []byte(" ")
var slashSep = []byte("/")

func ParseFace(buf []byte) (*Face, error) {
	parts := bytes.Split(buf, spaceSep)

	if len(parts) < 4 {
		return nil, fmt.Errorf("Incorrect face format %q", string(buf))
	}

	var points []FacePoint
	for _, faceData := range parts[1:] {
		tuple := bytes.Split(faceData, slashSep)

		if len(tuple) < 1 || len(tuple) > 3 {
			return nil, fmt.Errorf("incorrect format: %q", faceData)
		}

		var vertexIndex, UVIndex, normalIndex int32 = -1, -1, -1
		var err error

		if err = math.ParseInt32(tuple[0], &vertexIndex); err != nil {
			return nil, err
		}

		// This should happen when we have v1/v2 or v1/v2/v3
		if len(tuple) > 1 && len(tuple[1]) > 0 {
			if err = math.ParseInt32(tuple[1], &UVIndex); err != nil {
				return nil, err
			}
		}

		// This should be parsed if we have v1/v2/v3 or v1//v3
		if len(tuple) == 3 {
			if err = math.ParseInt32(tuple[2], &normalIndex); err != nil {
				return nil, err
			}
		}

		points = append(points, FacePoint{
			VertexIndex: vertexIndex,
			UVIndex:     UVIndex,
			NormalIndex: normalIndex,
		})
	}
	return &Face{
		Index:  -1,
		Points: points,
	}, nil
}
