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

	fmt.Println("Result: " + out.String())
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

	command, err := filepath.Abs("./modules/library-documents-tc-test/build/commands/create-pdf.js")

	if err != nil {
		panic(err)
	}

	return exec.Command("node", command, "-i "+storefile, "-o "+pathpdf)
}

func deleteJsonFile(storefile string) {
	os.Remove(storefile)
}
