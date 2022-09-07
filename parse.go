package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseFiles(files []*os.File, t time.Time) []Content {
	result := []Content{}
	for _, file := range files {
		newContent := parseFile(file, t)
		result = append(result, newContent)
	}

	return result
}

func parseFile(file *os.File, t time.Time) Content {
	regione := parseName(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	rows := []*InputLine{}
	skipLine(scanner)
	for scanner.Scan() {
		rows = append(rows, parseLine(len(rows), scanner.Text(), t))
	}
	return Content{
		regione: regione,
		lines:   rows,
	}
}

func skipLine(scanner *bufio.Scanner) {
	scanner.Scan()
}

func parseName(file *os.File) string {
	// READFile/ANHEUNG_surge.dat
	fileName := file.Name()
	fileName = strings.Split(fileName, "/")[1]
	return strings.Split(fileName, "_")[0]
}

func parseLine(idx int, row string, t time.Time) *InputLine {
	row = strings.TrimSpace(row)
	splited := strings.Split(row, " ")
	v := strings.TrimSpace(strings.Join(splited[1:], ""))
	value, err := strconv.ParseFloat(v, 64)
	date := t.Add(time.Duration(time.Minute * time.Duration(10*(idx+1))))

	if err != nil {
		println(err.Error())
		panic("값이 이상함")
	}
	return &InputLine{idx, date, value}
}
