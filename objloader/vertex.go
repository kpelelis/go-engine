package objloader

import (
	"bytes"
	"errors"
	"strconv"
)

type Vertex struct {
	index int64
	x     float64
	y     float64
	z     float64
	w     float64
}

func parseVertex(buf []byte) (*Vertex, error) {
	buf = bytes.TrimSpace(buf)
	sep := []byte(" ")
	parts := bytes.Split(buf, sep)

	var x, y, z, w float64
	var err error

	if len(parts) < 4 || len(parts) > 5 {
		return nil, errors.New("Incorrect Format")
	}

	if x, err = strconv.ParseFloat(string(parts[1]), 64); err != nil {
		return nil, err
	}

	if y, err = strconv.ParseFloat(string(parts[2]), 64); err != nil {
		return nil, err
	}

	if z, err = strconv.ParseFloat(string(parts[3]), 64); err != nil {
		return nil, err
	}

	w = -1
	if len(parts) == 5 {
		if w, err = strconv.ParseFloat(string(parts[4]), 64); err != nil {
			return nil, err
		}
	}

	return &Vertex{index: -1, x: x, y: y, z: z, w: w}, nil
}
