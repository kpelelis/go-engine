package objloader

import (
	"bytes"
	"strconv"
)

type Normal struct {
	index int64
	x     float64
	y     float64
	z     float64
}

func parseNormal(buf []byte) (*Normal, error) {
	sep := []byte(" ")
	parts := bytes.Split(buf, sep)

	x, err := strconv.ParseFloat(string(parts[1]), 64)
	if err != nil {
		return nil, err
	}
	y, err := strconv.ParseFloat(string(parts[2]), 64)
	if err != nil {
		return nil, err
	}
	z, err := strconv.ParseFloat(string(parts[3]), 64)
	if err != nil {
		return nil, err
	}
	return &Normal{x: x, y: y, z: z}, nil
}
