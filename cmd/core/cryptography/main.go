package main

import (
	"fmt"

	"github.com/chrishayles/xrplgo/xrpl/core/cryptography"
)

func main() {

	c, err := cryptography.NewKeyPair(cryptography.ED25519, nil, false)

	if err != nil {
		fmt.Println("There was an error.")
		fmt.Println(err)
	} else {
		fmt.Println("S'all good.")
		fmt.Println(c)
	}

}
