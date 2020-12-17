package preflight

import (
	"fmt"
	"log"
	"os"
)

//Capitalize the Function to make it useable outside of this file.
func CheckRepository() err {
	// Is there a Github token & is it valid

	// For the tool to run, the repo needs both a .git and .github folders. This checks if the working directory has both.
	isValidRepo()
	if err != nil {
		log.Println(err)
	}
	// Is there a .github folder in the repo

	// Is there <labels>.json in the repo

	// Is there a folder called ISSUETEMPLATES

}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	// _ means get rid of the response, just tell me if it errors.
	_, err := os.Stat(path)
	//Error handling
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	//Otherwise
	return true, nil
}

func isValidRepo() (bool, error) {
	//Read in the path. If path is empty, take current directory and look for .git or .github folder.
	path, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	//If true, continue. If false, prompt for git init.
	getGit, err := exists(path + "/.git")
	if err != nil {
		os.Exit(1)
	}

	if getGit == false {
		fmt.Println(gitInit)
	}

	//If true, continue. If false, prompt to create .github.
	getHub, err := exists(path + "/.github")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Here is getGit; %t", getGit))
	fmt.Println(fmt.Sprintf("Here is getHub; %t", getHub))
	os.Exit(0)
}

const gitInit = `
Your current working directory is not a git repository.
Please initialize this as a git repository with the following command:
	git init
`
