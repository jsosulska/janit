package preflight

import (
	"fmt"
	"log"
	"os"
)

//region Core Functionality
//Capitalize the Function to make it useable outside of this file.
func CheckRepository() err {
	// Is there a Github token &

	// Is GH Token valid?

	// For the tool to run, the repo needs both a .git and .github folders. This checks if the working directory has both.
	isValidRepo()
	if err != nil {
		log.Println(err)
	}
}

//endregion

//region Support Functionality
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
	//Read in the path. If path is empty, take current working directory and look for .git, .github, and .github/ISSUE_TEMPLATES folders.
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
		fmt.Println(gitNotInit)
	}

	//If true, continue. If false, prompt to create .github.
	getHub, err := exists(path + "/.github")
	if err != nil {
		os.Exit(1)
	}

	if getGit == false {
		fmt.Println(hubNotFound)
	}
	//If true, continue. If false, prompt to create .github/ISSUE_TEMPLATE.
	getHub, err := exists(path + "/.github/ISSUE_TEMPLATE")
	if err != nil {
		os.Exit(1)
	}

	if getGit == false {
		fmt.Println(templateNotFound)
	}

	os.Exit(0)
}

//endregion

//region Constants
const gitNotInit = `
Your current working directory is not a git repository.
Please initialize this as a git repository with the following command:
	git init
`

const hubNotFound = `
Your current working directory does not have a .github or ISSUE_TEMPLATE folder.
Please create the folder with the following command:
	mkdir -p .github/ISSUE_TEMPLATE
`

const templateNotFound = `
Your .github directory does not have a ISSUE_TEMPLATE folder.
Please create the folder with the following command:
	mkdir -p .github/ISSUE_TEMPLATE
`

//endregion
