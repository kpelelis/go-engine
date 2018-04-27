package objloader

import (
	"bytes"
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

	if len(parts) != 4 {
		return nil, nil
	}

	slashSep := []byte("/")
	var points [3]TrianglePoint
	for index, triangledata := range parts[1:] {
		tuple := bytes.Split(triangledata, slashSep)
		slashCnt := bytes.Count(triangledata, slashSep)

		var v1, v2, v3 int64
		var err error
		v1, err = strconv.ParseInt(string(tuple[0]), 10, 64)
		if err != nil {
			return nil, err
		}

		if len(tuple) > 1 && len(tuple[1]) > 0 {
			v2, err = strconv.ParseInt(string(tuple[1]), 10, 64)
		}

		if len(tuple) == 3 {
			v3, err = strconv.ParseInt(string(tuple[2]), 10, 64)
		}

		var vertexIndex, UVIndex, normalIndex int64

		vertexIndex = v1
		UVIndex = -1
		normalIndex = -1

		if slashCnt == 1 {
			UVIndex = v2
		}

		if slashCnt == 2 {
			normalIndex = v3
			if len(tuple[1]) > 0 {
				UVIndex = v2
			}
		}

		points[index] = TrianglePoint{
			vertexIndex: vertexIndex,
			UVIndex:     UVIndex,
			normalIndex: normalIndex,
		}
	}
	return &Triangle{Points: points}, nil
}
