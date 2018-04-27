package objloader

import (
	"bytes"
	"errors"
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

	if len(parts) != 4 {
		return nil, errors.New("Inocorrect format")
	}

	var x, y, z float64
	var err error

	if x, err = strconv.ParseFloat(string(parts[1]), 64); err != nil {
		return nil, err
	}

	if y, err = strconv.ParseFloat(string(parts[2]), 64); err != nil {
		return nil, err
	}

	if z, err = strconv.ParseFloat(string(parts[3]), 64); err != nil {
		return nil, err
	}

	return &Normal{
		index: -1,
		x:     x,
		y:     y,
		z:     z,
	}, nil
}
