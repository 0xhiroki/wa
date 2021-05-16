package main

import (
	"flag"
	"fmt"

	"wa/name"
)

func main() {
	rng, err := name.DefaultRNG()
	if err != nil {
		panic(err)
	}

	familyNamePrefix := flag.String("family", "", "A family name prefix")
	givenNamePrefix := flag.String("given", "", "A given name prefix")

	flag.Parse()

	c := &name.Config{}
	if familyNamePrefix != nil {
		c.FamilyNameStartsWith = *familyNamePrefix
	}
	if givenNamePrefix != nil {
		c.GivenNameStartsWith = *givenNamePrefix
	}

	randomFullname := name.GenerateWithConfig(rng, c)
	fmt.Println(randomFullname)
}
