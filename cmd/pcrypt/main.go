package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/unix4fun/passwd"
)

// a simple example of passwd package usage.
func main() {
	var p *passwd.Profile

	scryptFlag := flag.Bool("scrypt", false, "use scrypt derivation")
	argonFlag := flag.Bool("argon", false, "use argon derivation")
	bcryptFlag := flag.Bool("bcrypt", false, "use bcrypt derivation")
	checkFlag := flag.String("check", "", "hash to compare")

	flag.Parse()
	argv := flag.Args()

	switch {
	case *argonFlag:
		p = passwd.New(passwd.Argon2idCommon)
	case *scryptFlag:
		p = passwd.New(passwd.ScryptMin)
	case *bcryptFlag:
		p = passwd.New(passwd.BcryptCommon)
	default:
		fmt.Printf("no derivation, so nothing to do")
		os.Exit(0)
	}

	fmt.Printf("argv[%d]: %q\n", len(argv), argv)
	if len(argv) > 0 {
		for idx, passwordStr := range argv {
			h, err := p.Hash([]byte(passwordStr))
			if err != nil {
				log.Fatalf("hashing error: %v\n", err)
			}
			fmt.Printf("[%d] password: %s\nhashed: %s\n", idx, passwordStr, h)
			if len(*checkFlag) > 0 {
				fmt.Printf("[%d] is %s the passwd? %v\n", idx, passwordStr, passwd.Compare([]byte(*checkFlag), []byte(passwordStr)))
			}
		}
	}

	os.Exit(0)
}
