package testhelpers

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WriteFile - writes a file to the specified path with the specified content
func WriteFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0777)
	if err != nil {
		fmt.Println("Could not write test cert")
		os.Exit(1)
	}
}

// RemoveFile - removes the scpecified file
func RemoveFile(path string) {
	err1 := os.Remove(path)
	if err1 != nil {
		fmt.Println("Could not remove the test file because: ", err1)
		os.Exit(1)
	}
}
