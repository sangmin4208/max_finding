package main

import (
	"fmt"
	"os"
	"strings"
)

func writeFile(fileName string, lines []*OutputLine) {
	result := []string{}
	for _, line := range lines {
		result = append(result, line.toPlainText())
	}
	err := os.WriteFile(getOutputPath(fileName), []byte(strings.Join(result, "\n")), 0777)
	if err != nil {
		println(err.Error())
		panic("파일 작성 실패, 폴더 확인")
	}
}

func getOutputPath(fileName string) string {
	return fmt.Sprintf("%v/%v", OUPUT_PATH, fileName)
}
