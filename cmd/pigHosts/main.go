package main

import (
	"fmt"
	"reflect"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `
  Usage: arguments_example [-vqrh] [FILE] ...
    arguments_example (--left | --right) CORRECTION FILE
Process FILE and optionally apply correction to either left-hand side or
  right-hand side.
Arguments:
  FILE        File to process
Options:
  -h --help
  `

	arguments, _ := docopt.ParseDoc(usage)

	fmt.Printf("%v", reflect.TypeOf(arguments["FILE"]))

}
