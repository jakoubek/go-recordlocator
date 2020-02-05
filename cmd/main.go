package main

import (
	"fmt"

	"github.com/jakoubek/go-recordlocator"
)

var recloc *recordlocator.RecordLocator

func main() {
	fmt.Println("Test for Recordlocator")
	recloc = recordlocator.NewRecordLocator()

	str, _ := recloc.Encode(5325)
	fmt.Printf("The number %d encodes to %s\n", 5325, str)

	nr, _ := recloc.Decode("78G")
	fmt.Printf("The recordlocator %s decodes back to %d.\n", "78G", nr)

	for i := 1; i <= 100; i++ {
		encodeAndDecodeBack(int64(i))
	}

}

func encodeAndDecodeBack(nr int64) {
	str, _ := recloc.Encode(nr)
	back, _ := recloc.Decode(str)
	fmt.Println(nr, "=>", str, "=>", back)
}
