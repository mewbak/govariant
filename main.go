//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"log"
	"os"
)

var config struct {
	PkgName  string
	TypeName string
	Variants []string
}

var usage = `
govariant SUM_TYPE VARIANT_ONE VARIANT_TWO [OTHER_VARIANTS...]

SUM_TYPE is the name of the generated type. The VARIANTS are the variant types;
there must be at least two of them.
`[1:]

func init() {
	switch len(os.Args) {
	case 0, 1:
		log.Println("result type name not specified")
		fmt.Print(usage)
		os.Exit(1)
	case 2:
		log.Println("no variants specified")
		fmt.Print(usage)
		os.Exit(1)
	case 3:
		log.Println("only one variant specified")
		fmt.Print(usage)
		os.Exit(1)
	}

	config.TypeName = os.Args[1]
	config.Variants = os.Args[2:]
}

func main() {
}