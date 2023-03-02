package documentsTcTestPackage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

func CreatePdf(data map[string]interface{}, pathpdf string) bool {

	// Prepare Data to File Pdf
	storefile := generateJsonFileToNode(data)
	command := formuleCommandLibraryNodejs(storefile, pathpdf)

	// Execute node
	result, err := exec.Command("node", command).Output()

	if err != nil {
		panic(err)
	}

	output := string(result)
	fmt.Println(output)
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

func formuleCommandLibraryNodejs(storefile string, pathpdf string) string {

	command, err := filepath.Abs("./modules/library-documents-tc-test/build/commands/create-pdf.js")
	//command, err := filepath.Abs("./node-pdfs/build/script-create-file.js")

	if err != nil {
		panic(err)
	}

	command += "-i " + storefile + " -o " + pathpdf

	return command
}
