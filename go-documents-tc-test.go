package documents_tc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const sourcelibrary string = "library-documents-tc-test/dist/command.js"

// Create Pdf using library frontend rendering with react-df
func CreatePdf(data map[string]interface{}, pathpdf string) bool {

	// Prepare Data to File Pdf
	storefile := generateJsonFileToNode(data)

	// Prepare command
	cmd := formuleCommandLibraryNodejs(storefile, pathpdf)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Execute node
	err := cmd.Run()

	// Delete file .json
	deleteJsonFile(storefile)

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		panic(err)
	}

	return true
}

func generateJsonFileToNode(data map[string]interface{}) string {
	storefile, err := filepath.Abs("./data.json")

	if err != nil {
		panic(err)
	}

	file, err := json.Marshal(data)

	err = ioutil.WriteFile(storefile, file, 0644)

	if err != nil {
		panic(err)
	}

	return storefile
}

func formuleCommandLibraryNodejs(storefile string, pathpdf string) *exec.Cmd {

	path_root := RootDir()
	path_env_pkg := os.Getenv("PATH_PKG_NODE_FRONTEND")

	if path_env_pkg == "" {
		path_env_pkg = "pkg/"
	}

	command := path_root + "/" + path_env_pkg + sourcelibrary

	return exec.Command("node", command, "-i "+storefile, "-o "+pathpdf)
}

func deleteJsonFile(storefile string) {
	os.Remove(storefile)
}

func RootDir() string {
	dir, _ := os.Getwd()
	return dir
}
