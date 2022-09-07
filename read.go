package main

import (
	"fmt"
	"os"
)

func readDir(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic("input 폴더가 없음")
	}

	result := []string{}
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result
}

func readFiles(fileNames []string) []*os.File {
	result := []*os.File{}
	for _, fileName := range fileNames {
		result = append(result, readFile(fileName))
	}
	return result
}

func readFile(fileName string) *os.File {

	file, err := os.Open(
		fmt.Sprintf("%v/%v", INPUT_DIR, fileName),
	)
	if err != nil {
		panic("파일 읽기 실패")
	}

	return file
}
