package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]
	proj, ok := getProjPath(args)
	if !ok {
		var err error
		proj, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		args = append([]string{projPathArg, proj}, args...)
	}
	log.Printf("proj: '%s'", proj)

	version := getProjVersion(proj)
	log.Printf("version: '%s'", version)

	exe, ok := getUnityExePath(version)
	if !ok {
		log.Fatalf("Not found Unity %s", version)
	}
	log.Printf("exe: '%s'", exe)

	log.Printf("args: '%s'", strings.Join(args, "' '"))
	out, err := exec.Command(exe, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("output:\n====\n%s====\n", out)
}
