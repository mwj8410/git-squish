package src

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	testGitDir()
	testGitChange()

	argsWithProg := os.Args

	var finalMessage string
	commits, err := strconv.Atoi(argsWithProg[1])

	for i := 0; int(i) < commits; i++ {
		finalMessage = getCommitMessage()
		squish1()
	}

	cmd := exec.Command("git", "commit", "-m", finalMessage)

	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s", out)
}

func getCommitMessage() string {
	cmd := exec.Command("git", "--no-pager", "log", "--format=%B", "-n", "1")

	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}
	return string(out)
}

func squish1 () {
	cmd := exec.Command("git", "reset", "--soft", "HEAD~1")

	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}
}

func testGitChange() {
	cmd := exec.Command("git" , "diff-index", "HEAD")

	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}

	if len(string(out)) != 0 {
		fmt.Print("Current directory appears to have pending changes!")
		os.Exit(1)
	}
}

func testGitDir () {
	cmd := exec.Command("find" , ".git")

	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}

	if strings.Contains(string(out), ".git") != true {
		fmt.Print("Current directory is not a git repo")
		os.Exit(1)
	}
}
