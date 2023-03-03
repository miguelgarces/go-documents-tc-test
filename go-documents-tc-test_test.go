package documents_tc

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_CreatePdf_GeneratePdf(t *testing.T) {

	fmt.Println("Inicie la prueba")

	data := map[string]interface{}{
		"companyName":   "MELI-GO",
		"companyPhone":  "555-555-5555",
		"companyEmail":  "hello@meli.dev",
		"receiptNumber": "101445",
		"datePaid":      "1/4/2022",
		"paymentMethod": "Visa",
		"amount":        "$200.000",
	}

	path_output, err := filepath.Abs("./prueba.pdf")

	if err != nil {
		panic(err)
	}

	result := CreatePdf(data, path_output)

	if !result {
		t.Errorf("\"CreatePdf()\" FAILED, expected -> %v, got -> %v", true, result)
	} else {
		t.Logf("\"CreatePdf()\" SUCCEDED, expected -> %v, got -> %v", true, result)
	}
}
