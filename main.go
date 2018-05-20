package main

import (
	"fmt"
	"log"

	"github.com/kpelelis/go-engine/objloader"
)

func main() {
	reader, err := objloader.New("testfiles/test.obj")
	defer reader.Close()
	reader.Read()
	if err != nil {
		log.Fatalf("could not open file %q", err)
	}
	data, idx := reader.ExportIndexArrays()
	if err != nil {
		log.Fatalf("error while parsing wavefront file %q", err)
	}
	fmt.Println(data)
	fmt.Println(idx)
}
