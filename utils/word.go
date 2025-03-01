package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/Esword618/unioffice/document"
)

// GetWordContent gets the content of a word file
//
// 注意：只支持 docx
// 返回结果是纯文本，不包含格式
func GetWordContent(file string) (string, error) {
	doc, err := document.Open(file)
	if err != nil {
		return "", err
	}

	var content string
	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			content += run.Text() + "\n"
		}
	}
	return content, nil
}

// Markdown2Word converts markdown to word
// 备注：这个效果不好，不建议使用
// func Markdown2Word(markdown string) (bytes.Buffer, error) {
// 	var buf bytes.Buffer
// 	md := goldmark.New()
// 	if err := md.Convert([]byte(markdown), &buf); err != nil {
// 		return bytes.Buffer{}, err
// 	}

// 	htmlContent := buf.String()

// 	doc := document.New()
// 	para := doc.AddParagraph()
// 	run := para.AddRun()
// 	run.AddText(htmlContent)

// 	var wordFileBuffer bytes.Buffer
// 	if err := doc.Save(&wordFileBuffer); err != nil {
// 		return bytes.Buffer{}, err
// 	}
// 	return wordFileBuffer, nil
// }

// Word2Markdown converts word to markdown
//
// 返回 markdown 内容
func Word2Markdown(wordFilePath string) (string, error) {
	cmd := exec.Command("pandoc", wordFilePath, "-f", "docx", "-t", "markdown")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// Markdown2Word converts markdown to word
//
// 返回 word 文件内容
func Markdown2Word(markdownContent string) (bytes.Buffer, error) {
	tempMdFile, err := ioutil.TempFile("", "temp*.md")
	if err != nil {
		return bytes.Buffer{}, err
	}
	defer os.Remove(tempMdFile.Name())

	if _, err = tempMdFile.Write([]byte(markdownContent)); err != nil {
		tempMdFile.Close()
		return bytes.Buffer{}, err
	}
	tempMdFile.Close()

	tempWordFile := "output" + RandomString(12) + ".docx"
	cmd := exec.Command("pandoc", tempMdFile.Name(), "-o", tempWordFile)
	if err = cmd.Run(); err != nil {
		return bytes.Buffer{}, err
	}
	defer os.Remove(tempWordFile)

	file, err := os.Open(tempWordFile)
	if err != nil {
		return bytes.Buffer{}, err
	}
	defer file.Close()

	var wordFileBuffer bytes.Buffer
	if _, err = io.Copy(&wordFileBuffer, file); err != nil {
		return bytes.Buffer{}, err
	}
	return wordFileBuffer, nil
}
