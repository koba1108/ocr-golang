package main

import (
	"encoding/json"
	"fmt"
	"github.com/koba1108/ocr-golang/internals"
	"log"
	"os"
)

const (
	SdkName        = "unidoc"
	DocxPath       = "documents/docx"
	DocxOutputPath = "outputs/docx"
	OutputExt      = ".txt"
)

func main() {
	filePaths, err := internals.ReadFilenames(DocxPath)
	if err != nil {
		panic(fmt.Errorf("failed to read filenames: %w", err))
	}
	for _, path := range filePaths {
		content, err := internals.ReadDocx(path)
		if err != nil {
			panic(fmt.Errorf("failed to read pdf: %w", err))
		}
		file, err := os.Create(internals.MakeOutputPath(path, DocxPath, DocxOutputPath, SdkName, OutputExt))
		if err != nil {
			panic(fmt.Errorf("failed to create file: %w", err))
		}
		if _, err = file.WriteString(content); err != nil {
			panic(fmt.Errorf("failed to write to file: %w", err))
		}
		_ = file.Close()

		j, _ := json.Marshal("")
		log.Printf("%s created: metadata = %v", path, string(j))
	}
}