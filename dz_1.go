package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	user := os.Getenv("USER")
	if user == "" {
		user = "Guest"
	}

	name := flag.String("name", "student", "Your name")
	lang := flag.String("lang", "RU", "Language: RU or EN")
	flag.Parse()

	args := flag.Args()

	version := runtime.Version()

	fmt.Printf(" username: %s\n Args = %v\n Hello, %s! Lang=%s\n Go version = %v\n", user, args, *name, *lang, version)
}
