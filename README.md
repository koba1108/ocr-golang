# ocr-golang
ocr-golang

## 1. PDF読み取り（dslipak/pdf）
URL: [https://github.com/dslipak/pdf](https://github.com/dslipak/pdf)

```shell
$ go run cmd/dslipak_pdf/main.go 
```


## 2. PDF, Docx読み取り（docconv)
URL: [https://github.com/sajari/docconv](https://github.com/sajari/docconv)

```shell
$ go run cmd/docconv/main.go 
```

### 下記エラーが起きた場合
```
error converting data: exec: "pdftotext": executable file not found in $PATH #46
```

対処方法：https://github.com/sajari/docconv/issues/46


## 3. Docx, Excel, PPT読み取り（unidoc/unioffice)
URL: [https://github.com/unidoc/unioffice](https://github.com/unidoc/unioffice)
※ そもそも従量課金制のソフトウェアだった
@see: https://github.com/unidoc/unioffice-examples/blob/master/document/edit-document/main.go#L16-L22

```shell
$ go run cmd/unidoc/main.go

Unlicensed version of UniOffice
- Get a trial license on https://unidoc.io
panic: failed to read pdf: unioffice license required
```
