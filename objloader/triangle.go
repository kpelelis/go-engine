package objloader

import (
	"bytes"
	"errors"
	"strconv"
)

type TrianglePoint struct {
	index       int64
	vertexIndex int64
	UVIndex     int64
	normalIndex int64
}

type Triangle struct {
	index  int64
	Points [3]TrianglePoint
}

func parseTriangle(buf []byte) (*Triangle, error) {
	spaceSep := []byte(" ")
	parts := bytes.Split(buf, spaceSep)

	// A triangle should have 3 parts (4 including the f prefix)
	if len(parts) != 4 {
		return nil, errors.New("Incorrect format")
	}

	slashSep := []byte("/")
	var points [3]TrianglePoint
	for index, triangledata := range parts[1:] {
		tuple := bytes.Split(triangledata, slashSep)

		if len(tuple) < 1 || len(tuple) > 3 {
			return nil, errors.New("Incorrect format")
		}

		slashCnt := bytes.Count(triangledata, slashSep)

		var v1, v2, v3 int64
		var err error

		if v1, err = strconv.ParseInt(string(tuple[0]), 10, 64); err != nil {
			return nil, err
		}

		// This should happen when we have v1/v2 or v1/v2/v3
		if len(tuple) > 1 && len(tuple[1]) > 0 {
			if v2, err = strconv.ParseInt(string(tuple[1]), 10, 64); err != nil {
				return nil, err
			}
		}

		// This should be parsed if we have v1/v2/v3 or v1//v3
		if len(tuple) == 3 {
			if v3, err = strconv.ParseInt(string(tuple[2]), 10, 64); err != nil {
				return nil, err
			}
		}

		var vertexIndex, UVIndex, normalIndex int64

		vertexIndex = v1
		UVIndex = -1
		normalIndex = -1

		// If we have just one slash, we have 2 parts (vertex and UV)
		if slashCnt == 1 {
			UVIndex = v2
		}

		// If we have 2 slashes we either have 3 parts (vertex, UV and normals) or
		// 2 (vertex and normals)
		if slashCnt == 2 {
			normalIndex = v3
			if len(tuple[1]) > 0 {
				UVIndex = v2
			}
		}

		points[index] = TrianglePoint{
			index:       -1,
			vertexIndex: vertexIndex,
			UVIndex:     UVIndex,
			normalIndex: normalIndex,
		}
	}
	return &Triangle{
		index:  -1,
		Points: points,
	}, nil
}
