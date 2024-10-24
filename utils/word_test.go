package utils

import (
	"bytes"
	"testing"
)

func TestGetWordContent(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case1", args{"../uploads/中央组织部召开改进推动高质量发展的政绩考核工作座谈会_1qnLNrJKuu_1729490438.docx"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWordContent(tt.args.file)
			t.Logf("get %s", got)
			t.Logf("err %s", err)
		})
	}
}

func TestMarkdown2Word(t *testing.T) {
	type args struct {
		markdown string
	}
	tests := []struct {
		name    string
		args    args
		want    bytes.Buffer
		wantErr bool
	}{
		{"case1", args{"# 一级标题\n## 二级标题\n### 三级标题\n#### 四级标题\n##### 五级标题\n###### 六级标题\n\n**加粗**\n\n*斜体*\n\n~~删除线~~\n\n`行内代码`\n\n"}, bytes.Buffer{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Markdown2Word(tt.args.markdown)
			t.Logf("get %s", got.String())
			t.Logf("err %s", err)
		})
	}
}

func TestWord2Markdown(t *testing.T) {
	type args struct {
		wordFilePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", args{"../uploads/写一篇干部选拔任用工作情况报告.docx"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Word2Markdown(tt.args.wordFilePath)
			t.Logf("get %s", got)
			t.Logf("err %v", err)
		})
	}
}
