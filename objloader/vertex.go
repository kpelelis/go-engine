package objloader

import (
	"bytes"
	"fmt"

	"github.com/kpelelis/go-engine/math"
)

type Vertex struct {
	Index int32
	X     float32
	Y     float32
	Z     float32
	W     float32
}

func ParseVertex(buf []byte) (*Vertex, error) {
	buf = bytes.TrimSpace(buf)
	sep := []byte(" ")
	parts := bytes.Split(buf, sep)

	if len(parts) < 4 || len(parts) > 5 {
		return nil, fmt.Errorf("incorrect format: %q", buf)
	}

	var x, y, z, w float32
	w = -1

	var err error

	if err = math.ParseFloat32(parts[1], &x); err != nil {
		return nil, err
	}

	if err = math.ParseFloat32(parts[2], &y); err != nil {
		return nil, err
	}

	if err = math.ParseFloat32(parts[3], &z); err != nil {
		return nil, err
	}

	if len(parts) == 5 {
		if err = math.ParseFloat32(parts[4], &w); err != nil {
			return nil, err
		}
	}

	return &Vertex{
		Index: -1,
		X:     x,
		Y:     y,
		Z:     z,
		W:     w,
	}, nil
}
