package main

import "github.com/bperezgo/backoffice-businesses/cmd/bootstrap"

func main() {
	if err := bootstrap.Run(); err != nil {
		panic(err)
	}
}
