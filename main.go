package main

import (
	"fmt"
	"log"

	"github.com/kpelelis/go-engine/objloader"
)

func main() {
	reader, err := objloader.New("testfiles/test.obj")
	defer reader.Close()
	if err != nil {
		log.Fatalf("could not open file %q", err)
	}
	obj, err := reader.Read()
	if err != nil {
		log.Fatalf("error while parsing wavefront file %q", err)
	}
	var cube []float32
	for _, t := range obj.Triangles {
		for _, p := range t.Points {
			v := obj.Vertices[p.VertexIndex-1]
			uv := obj.UVs[p.UVIndex-1]
			cube = append(cube, []float32{
				float32(v.X), float32(v.Y), float32(v.Z), float32(uv.U), float32(uv.V),
			}...)
		}
	}
	fmt.Println(cube)
}
