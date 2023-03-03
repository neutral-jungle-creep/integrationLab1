package main

import (
	"flag"
	"github.com/lukasjarosch/go-docx"
)

var filePath, outputPath string

func init() {
	flag.StringVar(&filePath, "template", "template.docx", "path to the template docx file")
	flag.StringVar(&outputPath, "out", "output.docx", "path to the output docx")
	flag.Parse()
}

func main() {
	replaceMap := docx.PlaceholderMap{
		"key":                         "REPLACE some more",
		"key-with-dash":               "REPLACE",
		"key-with-dashes":             "REPLACE",
		"key with space":              "REPLACE",
		"key_with_underscore":         "REPLACE",
		"multiline":                   "REPLACE",
		"key.with.dots":               "REPLACE",
		"mixed-key.separator_styles#": "REPLACE",
		"yet-another_placeholder":     "REPLACE",
		"foo":                         "bar",
		"newlinetester":               "hello1\nhello2\nhello3",
	}

	doc, err := docx.Open(filePath)
	if err != nil {
		panic(err)
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		panic(err)
	}

	err = doc.WriteToFile(outputPath)
	if err != nil {
		panic(err)
	}
}
