package main

import (
	"encoding/json"
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
		content, metadata, err := internals.ReadPDF2(path)
		if err != nil {
			panic(fmt.Errorf("failed to read pdf: %w", err))
		}
		file, err := os.Create(internals.MakeOutputPath(path, DocumentPath, OutputPath, OutputExt))
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
