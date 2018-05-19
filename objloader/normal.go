package objloader

import (
	"bytes"
	"errors"

	"github.com/kpelelis/go-engine/math"
)

type Normal struct {
	Index int64
	X     float64
	Y     float64
	Z     float64
}

func ParseNormal(buf []byte) (*Normal, error) {
	sep := []byte(" ")
	parts := bytes.Split(buf, sep)

	if len(parts) != 4 {
		return nil, errors.New("Inocorrect format")
	}

	var x, y, z float64
	var err error

	if err = math.ParseFloat64(parts[1], &x); err != nil {
		return nil, err
	}

	if err = math.ParseFloat64(parts[2], &y); err != nil {
		return nil, err
	}

	if err = math.ParseFloat64(parts[3], &z); err != nil {
		return nil, err
	}

	return &Normal{
		Index: -1,
		X:     x,
		Y:     y,
		Z:     z,
	}, nil
}
