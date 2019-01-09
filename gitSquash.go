package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)
func main() {
	// toDo: test that the cwd is a git repo
	// toDo: test that the current git repo has no pending changes

	argsWithProg := os.Args

	cmd := exec.Command("echo", argsWithProg[1])

	//cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf("%s\n", out)

	//parts := strings.Split(string(out), "\n--\n-- ")

	//fmt.Printf("combined out:\n%s\n", string(out))
	//fmt.Printf("%s", parts[0])

	//for i := range parts {
	//	entryParts := strings.Split(parts[i], "\n--\n\n")
	//	fmt.Printf("%s\n", entryParts[0])
	//
	//	metadata := strings.Split(entryParts[0], ": ")
	//	name:= metadata[1]
	//	fmt.Printf("%s\n", name)
	//}
}
