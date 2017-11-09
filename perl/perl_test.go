package perl

import (
	"fmt"
)

func ExampleToJSON() {
	data := ToJSON("%Info", `my %Info = ("DUlastValue" => 29)`)
	fmt.Println(string(data))
	// Output: {"DUlastValue":29}
}
