// Package perl is a simple parser for Perl data structures. It will convert a given
// perl data structure into JSON.
package perl

import (
	"log"
	"os/exec"
)

const (
	perlCmd = "perl"
)

// ToJSON will convert a given perl structure into JSON. It takes a perl variable
// name (e.g. "%Info") and the data declaration (e.g. "%Info = ("Thing"=> 1)") and
// outputs it to JSON bytes.
// NOTE: Data given will NOT be escaped, so it should only be from a trusted source
// or escaped before using this function.
func ToJSON(variable, data string) []byte {
	var (
		cmdOut []byte
		err    error
	)

	args := []string{"-MJSON", "-e", "use JSON; " + data + " ; my $json = encode_json \\" + variable + "; print $json"}
	if cmdOut, err = exec.Command(perlCmd, args...).Output(); err != nil {
		log.Println("Error interpreting perl data")
		log.Println("  args:", args)
		log.Fatal(err)
	}

	return cmdOut
}
