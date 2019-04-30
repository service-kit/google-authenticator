package main

import (
	"fmt"
	"github.com/service-kit/google-authenticator/auth"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "must specify key to use")
		os.Exit(1)
	}
	secret := os.Args[1]
	fmt.Println(auth.GetOneTimePassword(secret))
}