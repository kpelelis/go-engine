package objloader

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Interface for reading .obj files. Read is medium independent (e.g. Hard Disk,
// Network). Close should terminate all IO handlers.
type WavefrontReader interface {
	Read() error
	Close() error
}

// File Reader implements the WavefrontReader interface by manipulation files.
type FileReader struct {
	fd   *os.File
	objs []*Obj
}

// Read in FileReader, accepts a filename and reads it in chunks equal to the
// page size of the OS (TODO: Test different buffer sizes).
func (fr *FileReader) Read() error {
	pagesize := 64 * os.Getpagesize()
	buf := make([]byte, pagesize)

	for {
		_, err := fr.fd.Read(buf)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		// Find the end of a line and send it to the line parser
		// TODO: What happens if we are in the middle of the line while the buffer
		// is completely read?
		lineStart := 0
		var o *Obj
		for lineEnd, char := range buf {
			if char == '\n' && lineEnd > lineStart {
				// Remove trailing and leading space
				line := buf[lineStart:lineEnd]
				line = bytes.TrimSpace(line)
				spaceSep = []byte(" ")

				lineParts := bytes.Split(line, spaceSep)

				instruction := string(lineParts[0])

				// TODO: Find best practices in go for altering slices in
				// functions since they are passed by value

				if len(fr.objs) > 0 {
					o = fr.objs[len(fr.objs)-1]
				}

				switch {
				// Comment
				case instruction == "#":
					break

				case instruction == "o":
					name := string(bytes.Split(line, spaceSep)[1])
					fmt.Println(name)
					fr.objs = append(fr.objs, &Obj{Name: name})

				// Vertex
				case instruction == "v":
					vertex, err := ParseVertex(line)
					if err != nil {
						return err
					}
					vertex.Index = int32(len(o.Vertices) + 1)
					o.Vertices = append(o.Vertices, *vertex)

				// UV
				case instruction == "vt":
					uv, err := ParseUV(line)
					if err != nil {
						return err
					}
					uv.Index = int32(len(o.UVs) + 1)
					o.UVs = append(o.UVs, *uv)

				// Normal
				case instruction == "vn":
					normal, err := ParseNormal(line)
					if err != nil {
						return err
					}
					normal.Index = int32(len(o.Normals) + 1)
					o.Normals = append(o.Normals, *normal)

				// Triangle
				case instruction == "f":
					face, err := ParseFace(line)
					if err != nil {
						return err
					}
					face.Index = int32(len(o.Faces) + 1)
					o.Faces = append(o.Faces, *face)

				case instruction == "mtllib":
					fmt.Println("Parsing Material")

				case instruction == "usemtl":
					o.Material = string(lineParts[1])
				}
				lineStart = lineEnd + 1
			}
		}

	}
	return nil
}

func (fr *FileReader) Close() error {
	fr.fd.Close()
	return nil
}

// TODO: Test
func (fr *FileReader) ExportToFloat32Array() [][]float32 {
	var ret [][]float32
	for _, obj := range fr.objs {
		var objFlt32Arr []float32
		for _, face := range obj.Faces {
			for _, p := range face.Points {
				v := obj.Vertices[p.VertexIndex-1]
				uv := obj.UVs[p.UVIndex-1]
				n := obj.Normals[p.NormalIndex-1]
				objFlt32Arr = append(objFlt32Arr, []float32{
					v.X, v.Y, v.Z, uv.U, uv.V, n.X, n.Y, n.Z,
				}...)
				fmt.Printf("%v, %v, %v, %v, %v,\n", v.X, v.Y, v.Z, uv.U, uv.V)
			}
		}
		ret = append(ret, objFlt32Arr)
	}
	return ret
}

// TODO: Test
func (fr *FileReader) ExportIndexArrays() ([][]float32, [][]uint32) {
	var indexBuffers [][]uint32
	var vertexDataBuffers [][]float32

	for _, obj := range fr.objs {
		vertexCache := make(map[string]uint32)
		var objFlt32Arr []float32
		var objIndexBuf []uint32
		var objIndex uint32 = 0

		for _, face := range obj.Faces {
			for _, p := range face.Points {
				v := obj.Vertices[p.VertexIndex-1]
				uv := obj.UVs[p.UVIndex-1]
				n := obj.Normals[p.NormalIndex-1]

				// Check if we have the vertex in the cache
				vertexKey := fmt.Sprintf("%v.%v.%v", p.VertexIndex, p.UVIndex, p.NormalIndex)
				if val, ok := vertexCache[vertexKey]; ok {
					// If so, append the number in the index buffer
					objIndexBuf = append(objIndexBuf, val)
					continue
				}

				// Else append the data
				objFlt32Arr = append(objFlt32Arr, []float32{v.X, v.Y, v.Z, uv.U, uv.V, n.X, n.Y, n.Z}...)
				// And then add the unique index and increment it
				objIndexBuf = append(objIndexBuf, objIndex)
				vertexCache[vertexKey] = objIndex

				objIndex++

				//fmt.Printf("%v, %v, %v, %v, %v,\n", v.X, v.Y, v.Z, uv.U, uv.V)
			}
		}
		vertexDataBuffers = append(vertexDataBuffers, objFlt32Arr)
		indexBuffers = append(indexBuffers, objIndexBuf)
	}
	return vertexDataBuffers, indexBuffers
}

func New(filename string) (*FileReader, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fr := &FileReader{fd: fd}
	return fr, nil
}
