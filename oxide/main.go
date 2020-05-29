package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/orsenkucher/nothing/encio"
)

func main() {
	if err := token(); err != nil {
		log.Fatalln(err)
	}
}

func token() error {
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	if *s == "" {
		return errors.New("[-s] -> encio must be handled")
	}

	key := encio.NewEncIO(*s)
	token, err := key.ReadFile("token")
	if err != nil {
		return errors.New("Can't read file")
	}
	fmt.Println(string(token))

	return nil
}
