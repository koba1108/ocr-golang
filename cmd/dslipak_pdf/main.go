package main

import (
	"bytes"
	"fmt"
	"github.com/dslipak/pdf"
)

const DocumentPath = "documents/pdf"

// ファイルを配置してね
const (
	Pdf1 = "a.pdf"
	Pdf2 = "b.pdf"
)

func main() {
	content, err := readPdf(fmt.Sprintf("%s/%s", DocumentPath, Pdf2))
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func readPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return "", err
	}
	return buf.String(), nil
}
