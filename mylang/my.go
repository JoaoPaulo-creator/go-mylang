package mylang

import (
	"fmt"
	"my-lang/scanner"
)

func Run(file string) {
	scnr := scanner.NewScanner(file)
	fmt.Println(scnr.ScanToken())
}
