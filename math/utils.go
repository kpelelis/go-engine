package math

import (
	"bytes"
	"strconv"
)

func ParseInt64(buf []byte, i *int64) error {
	buf = bytes.TrimSpace(buf)
	num, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		return err
	}
	*i = num
	return nil
}

func ParseInt32(buf []byte, i *int32) error {
	buf = bytes.TrimSpace(buf)
	num, err := strconv.ParseInt(string(buf), 10, 32)
	if err != nil {
		return err
	}
	*i = int32(num)
	return nil
}

func ParseFloat64(buf []byte, i *float64) error {
	buf = bytes.TrimSpace(buf)
	num, err := strconv.ParseFloat(string(buf), 64)
	if err != nil {
		return err
	}
	*i = num
	return nil
}

func ParseFloat32(buf []byte, i *float32) error {
	buf = bytes.TrimSpace(buf)
	num, err := strconv.ParseFloat(string(buf), 64)
	if err != nil {
		return err
	}
	*i = float32(num)
	return nil
}
