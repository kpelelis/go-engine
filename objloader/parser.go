package objloader

import (
	"bytes"
	"io"
	"os"
)

// Interface for reading .obj files. Read is medium independent (e.g. Hard Disk,
// Network). Close should terminate all IO handlers.
type WavefrontReader interface {
	Read() (*Obj, error)
	Close() error
}

// File Reader implements the WavefrontReader interface by manipulation files.
type FileReader struct {
	fd *os.File
}

// Read in FileReader, accepts a filename and reads it in chunks equal to the
// page size of the OS (TODO: Test different buffer sizes).
func (fr *FileReader) Read() (*Obj, error) {
	pagesize := os.Getpagesize()
	buf := make([]byte, pagesize)

	var obj Obj

	for {
		_, err := fr.fd.Read(buf)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		// Find the end of a line and send it to the line parser
		// TODO: What happens if we are in the middle of the line while the buffer
		// is completely read?
		lineStart := 0
		for lineEnd, char := range buf {
			if char == '\n' && lineEnd > lineStart {
				fr.parseLine(buf[lineStart:lineEnd], &obj)
				lineStart = lineEnd + 1
			}
		}

	}
	return &obj, nil
}

func (fr *FileReader) Close() error {
	fr.fd.Close()
	return nil
}

func (fr *FileReader) parseLine(buf []byte, o *Obj) error {
	// Remove trailing and leading space
	buf = bytes.TrimSpace(buf)

	// Comment
	if buf[0] == '#' {
		return nil
	}

	// Vertex
	if buf[0] == 'v' && buf[1] == ' ' {
		vertex, err := parseVertex(buf)
		if err != nil {
			return err
		}
		vertex.index = int64(len(o.Vertices) + 1)
		o.Vertices = append(o.Vertices, *vertex)
	}

	// UV
	if buf[0] == 'v' && buf[1] == 't' {
		uv, err := parseUV(buf)
		if err != nil {
			return err
		}
		uv.index = int64(len(o.UVs) + 1)
		o.UVs = append(o.UVs, *uv)
	}

	// Normal
	if buf[0] == 'v' && buf[1] == 'n' {
		normal, err := parseNormal(buf)
		if err != nil {
			return err
		}
		normal.index = int64(len(o.Normals) + 1)
		o.Normals = append(o.Normals, *normal)
	}

	// Triangle
	if buf[0] == 'f' && buf[1] == ' ' {
		triangle, err := parseTriangle(buf)
		if err != nil {
			return err
		}
		triangle.index = int64(len(o.Triangles) + 1)
		o.Triangles = append(o.Triangles, *triangle)
	}

	return nil
}

func NewWavefrontReader(filename string) (WavefrontReader, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fr := &FileReader{fd: fd}
	return fr, nil
}
