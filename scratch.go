package main

import (
	"fmt"
	"os"
)

func exists(path string) (bool, error) {
	// _ means get rid of the response, just tell me if it errors.
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	getGit, err := exists(path + "/.git")
	if err != nil {
		os.Exit(1)
	}

	getHub, err := exists(path + "/.github")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Here is getGit; %t", getGit))
	fmt.Println(fmt.Sprintf("Here is getHub; %t", getHub))
	os.Exit(0)
}
