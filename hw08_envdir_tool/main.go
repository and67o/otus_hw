package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("not1 enough args")
	}

	dir := os.Args[1]

	env, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[2:]
	resCode := RunCmd(cmd, env)

	os.Exit(resCode)
}
