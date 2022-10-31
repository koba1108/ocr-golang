package internals

import (
	"bytes"
	"github.com/dslipak/pdf"
	"path/filepath"
	"strings"
)

func ReplaceExt(filename, newExt string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))] + newExt
}

func MakeOutputPath(filePath, docPath, outputPath, ext string) string {
	return ReplaceExt(strings.Replace(filePath, docPath, outputPath, 1), ext)
}

func ReadFilenames(dir string) ([]string, error) {
	files, err := filepath.Glob(dir + "/*")
	if err != nil {
		return nil, err
	}
	var filenames []string
	for _, fn := range files {
		if !strings.HasSuffix(fn, ".gitkeep") {
			filenames = append(filenames, fn)
		}
	}
	return filenames, nil
}

func ReadPDF(path string) (string, int, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", 0, err
	}
	b, err := r.GetPlainText()
	if err != nil {
		return "", 0, err
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return "", 0, err
	}
	return buf.String(), r.NumPage(), nil
}
