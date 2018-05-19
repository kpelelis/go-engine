package objloader

import (
	"bytes"
	"fmt"

	"github.com/kpelelis/go-engine/math"
)

type TrianglePoint struct {
	VertexIndex int64
	UVIndex     int64
	NormalIndex int64
}

type Triangle struct {
	Index  int64
	Points []TrianglePoint
}

var spaceSep = []byte(" ")

func ParseTriangle(buf []byte) (*Triangle, error) {
	parts := bytes.Split(buf, spaceSep)

	slashSep := []byte("/")
	var points []TrianglePoint
	for _, triangledata := range parts[1:] {
		tuple := bytes.Split(triangledata, slashSep)

		if len(tuple) < 1 || len(tuple) > 3 {
			return nil, fmt.Errorf("incorrect format: %q", triangledata)
		}

		var vertexIndex, UVIndex, normalIndex int64 = -1, -1, -1
		var err error

		if err = math.ParseInt64(tuple[0], &vertexIndex); err != nil {
			return nil, err
		}

		// This should happen when we have v1/v2 or v1/v2/v3
		if len(tuple) > 1 && len(tuple[1]) > 0 {
			if err = math.ParseInt64(tuple[1], &UVIndex); err != nil {
				return nil, err
			}
		}

		// This should be parsed if we have v1/v2/v3 or v1//v3
		if len(tuple) == 3 {
			if err = math.ParseInt64(tuple[2], &normalIndex); err != nil {
				return nil, err
			}
		}

		points = append(points, TrianglePoint{
			VertexIndex: vertexIndex,
			UVIndex:     UVIndex,
			NormalIndex: normalIndex,
		})
	}
	return &Triangle{
		Index:  -1,
		Points: points,
	}, nil
}
