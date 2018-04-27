package main

import (
	"fmt"
	"github.com/kpelelis/go-engine/objloader"
)

func main() {
	reader, err := objloader.NewWavefrontReader("testfiles/test.obj")
	if err != nil {
		fmt.Println(err)
	}
	obj, err := reader.Read()
	fmt.Println(obj)
	reader.Close()
}
