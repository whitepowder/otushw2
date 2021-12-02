package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("should be more")
	}
	dir, cmd := os.Args[1], os.Args[2:]
	env, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	output := RunCmd(cmd, env)
	os.Exit(output)
}
