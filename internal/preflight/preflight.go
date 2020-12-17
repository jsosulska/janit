package preflight

import "os"

//Capitalize the Function to make it useable outside of this file.
func CheckRepository() err {
	// Is there a Github token & is it valid

	// Is the current folder path a git repo

	// Is there a .github folder in the repo

	// Is there <labels>.json in the repo

	// Is there a folder called ISSUETEMPLATES

}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	// _ means get rid of the response, just tell me if it errors.
	_, err := os.Stat(path)
    if os.IsNotExist(err) { return false, nil }
	return false, err
	
	if err == nil { return true, nil }
}

func isValidRepo (path string) ( bool, error) {
	//Read in the path. If path is empty, take current directory and look for .git or .github folder.
	if path == "" {
		path := os.Getwd()
		return exists(path + “/.git”) && exists(path + “/.github”)
	}

	//If path includes .git or .github, read the path exactly.

	//If path is anything else, use that as a prefix to search for .git/.github
}