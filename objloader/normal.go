package objloader

import (
	"bytes"
	"errors"

	"github.com/kpelelis/go-engine/math"
)

type Normal struct {
	Index int32
	X     float32
	Y     float32
	Z     float32
}

func ParseNormal(buf []byte) (*Normal, error) {
	sep := []byte(" ")
	parts := bytes.Split(buf, sep)

	if len(parts) != 4 {
		return nil, errors.New("Inocorrect format")
	}

	var x, y, z float32
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

	return &Normal{
		Index: -1,
		X:     x,
		Y:     y,
		Z:     z,
	}, nil
}
