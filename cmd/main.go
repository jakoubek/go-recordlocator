package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jakoubek/go-recordlocator"
)

var recloc *recordlocator.RecordLocator

var (
	cli struct {
		number int64
		rl     string
	}
)

func main() {

	flag.Int64Var(&cli.number, "number", 0, "Number to be encoded")
	flag.StringVar(&cli.rl, "rl", "", "Recordlocator to be decoded")

	flag.Parse()

	recloc = recordlocator.NewRecordLocator()

	if cli.number > 0 {

		rl, err := recloc.Encode(cli.number)
		if err != nil {
			log.Fatalf("Error encoding number %d to a recordlocator\n", cli.number)
		}
		fmt.Printf("%d => %s\n", cli.number, rl)

	}

	if cli.rl != "" {

		number, err := recloc.Decode(cli.rl)
		if err != nil {
			log.Fatalf("Error decoding number recordlocator %s back to a number\n", cli.rl)
		}
		fmt.Printf("%s => %d\n", cli.rl, number)

	}

}

func encodeAndDecodeBack(nr int64) {
	str, _ := recloc.Encode(nr)
	back, _ := recloc.Decode(str)
	fmt.Println(nr, "=>", str, "=>", back)
}
