package objloader

import (
	"bytes"
	"errors"
	"strconv"
)

type UV struct {
	index int64
	u     float64
	v     float64
	w     float64
}

func parseUV(buf []byte) (*UV, error) {
	buf = bytes.TrimSpace(buf)
	sep := []byte(" ")
	parts := bytes.Split(buf, sep)

	var u, v, w float64
	var err error

	if len(parts) < 3 || len(parts) > 4 {
		return nil, errors.New("Incorrect Format")
	}

	if u, err = strconv.ParseFloat(string(parts[1]), 64); err != nil {
		return nil, err
	}

	if v, err = strconv.ParseFloat(string(parts[2]), 64); err != nil {
		return nil, err
	}

	w = -1
	if len(parts) == 4 {
		if w, err = strconv.ParseFloat(string(parts[3]), 64); err != nil {
			return nil, err
		}
	}
	return &UV{index: 1, u: u, v: v, w: w}, nil
}
