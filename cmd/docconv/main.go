package main

import (
	"encoding/json"
	"fmt"
	"github.com/koba1108/ocr-golang/internals"
	"log"
	"os"
)

const (
	SdkName        = "docconv"
	DocxPath       = "documents/docx"
	DocxOutputPath = "outputs/docx"
	PdfPath        = "documents/pdf"
	PdfOutputPath  = "outputs/pdf"
	OutputExt      = ".txt"
)

func main() {
	readDocuments(DocxPath, DocxOutputPath)
	readDocuments(PdfPath, PdfOutputPath)
}

func readDocuments(docPath, outputPath string) {
	filePaths, err := internals.ReadFilenames(docPath)
	if err != nil {
		panic(fmt.Errorf("failed to read filenames: %w", err))
	}
	for _, path := range filePaths {
		content, metadata, err := internals.ReadDocument(path)
		if err != nil {
			panic(fmt.Errorf("failed to read pdf: %w", err))
		}
		file, err := os.Create(internals.MakeOutputPath(path, docPath, outputPath, SdkName, OutputExt))
		if err != nil {
			panic(fmt.Errorf("failed to create file: %w", err))
		}
		if _, err = file.WriteString(content); err != nil {
			panic(fmt.Errorf("failed to write to file: %w", err))
		}
		_ = file.Close()

		j, _ := json.Marshal(metadata)
		log.Printf("%s created: metadata = %v", path, string(j))
	}
}
