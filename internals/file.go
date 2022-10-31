package internals

import (
	"bytes"
	"code.sajari.com/docconv"
	"github.com/dslipak/pdf"
	"github.com/unidoc/unioffice/common/license"
	unidoc "github.com/unidoc/unioffice/document"
	"path/filepath"
	"strings"
)

func ReplaceExt(filename, newExt string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))] + newExt
}

func MakeOutputPath(filePath, docPath, outputPath, SdkName, ext string) string {
	return ReplaceExt(strings.Replace(filePath, docPath, outputPath, 1), "."+SdkName+ext)
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

// ReadPDF is used: dslipak/pdf
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

// ReadDocument is used: docconv
func ReadDocument(path string) (string, map[string]string, error) {
	r, err := docconv.ConvertPath(path)
	if err != nil {
		return "", nil, err
	}
	return r.Body, r.Meta, nil
}

func SetUnicodeLicenseKey() error {
	// UnicodeLicenseKey is methods in config.go (ignored)
	return license.SetMeteredKey(GetUnicodeLicenseKey())
}

// ReadDocx is used: unidoc
func ReadDocx(path string) (string, error) {
	// GetUnicodeLicenseKey is methods in config.go (ignored)
	r, err := unidoc.Open(path)
	if err != nil {
		return "", err
	}
	var res string
	var ps []unidoc.Paragraph
	for _, p := range r.Paragraphs() {
		ps = append(ps, p)
	}
	for _, p := range ps {
		for _, r := range p.Runs() {
			res += r.Text()
		}
	}
	return res, nil
}
