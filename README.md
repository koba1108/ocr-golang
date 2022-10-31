# ocr-golang
ocr-golang

## 1. PDF読み取り（dslipak/pdf）
URL: [https://github.com/dslipak/pdf](https://github.com/dslipak/pdf)

```shell
$ go run cmd/dslipak_pdf/main.go 
```


## 2. PDF読み取り その2（docconv)
URL: [https://github.com/sajari/docconv](https://github.com/sajari/docconv)

```shell
$ go run cmd/docconv/main.go 
```

### 下記エラーが起きた場合
```
error converting data: exec: "pdftotext": executable file not found in $PATH #46
```

対処方法：https://github.com/sajari/docconv/issues/46
