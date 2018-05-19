package objloader

import (
	"bytes"
	"fmt"

	"github.com/kpelelis/go-engine/math"
)

type UV struct {
	Index int64
	U     float64
	V     float64
	W     float64
}

func ParseUV(buf []byte) (*UV, error) {
	buf = bytes.TrimSpace(buf)
	sep := []byte(" ")
	parts := bytes.Split(buf, sep)

	var u, v, w float64
	w = -1
	var err error

	if len(parts) < 3 || len(parts) > 4 {
		return nil, fmt.Errorf("incorrect format: %q", buf)
	}

	if err = math.ParseFloat64(parts[1], &u); err != nil {
		return nil, err
	}

	if err = math.ParseFloat64(parts[2], &v); err != nil {
		return nil, err
	}

	if len(parts) == 4 {
		if err = math.ParseFloat64(parts[3], &w); err != nil {
			return nil, err
		}
	}

	return &UV{
		Index: 1,
		U:     u,
		V:     v,
		W:     w,
	}, nil
}
