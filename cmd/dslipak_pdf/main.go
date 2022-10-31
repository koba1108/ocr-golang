package main

import (
	"fmt"
	"github.com/koba1108/ocr-golang/internals"
	"log"
	"os"
	"strings"
)

const (
	DocumentPath = "documents/pdf"
	OutputPath   = "outputs/pdf"
	OutputExt    = ".txt"
)

func main() {
	filePaths, err := internals.ReadFilenames(DocumentPath)
	if err != nil {
		panic(fmt.Errorf("failed to read filenames: %w", err))
	}
	for _, path := range filePaths {
		if !strings.HasSuffix(path, ".pdf") {
			continue
		}
		content, page, err := internals.ReadPDF(path)
		if err != nil {
			panic(fmt.Errorf("failed to read pdf: %w", err))
		}
		file, err := os.Create(internals.MakeOutputPath(path, DocumentPath, OutputPath, OutputExt))
		if err != nil {
			panic(fmt.Errorf("failed to create file: %w", err))
		}
		_, err = file.WriteString(content)
		if err != nil {
			panic(fmt.Errorf("failed to write to file: %w", err))
		}
		_ = file.Close()
		log.Printf("%s created: page = %d", path, page)
	}
}
