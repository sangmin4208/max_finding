package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func (c *Content) toOutputLine(t time.Time) *OutputLine {
	mv := c.getMaxValue()
	date := t.Add(time.Duration(time.Minute * time.Duration(10*(mv.index+1))))
	return &OutputLine{
		regione:       c.regione,
		maxValue:      fmt.Sprintf("%v", mv.value),
		date:          date,
		formattedDate: date.Format("2006 01 02 15 04"),
	}
}

func (c *Content) getMaxValue() InputLine {
	rows := c.lines
	sort.Slice(rows, func(i int, j int) bool {
		return rows[i].value > rows[j].value
	})
	return rows[0]
}

func (ol *OutputLine) toPlainText() string {
	// Incheon              2022  09  04  22 10     2.206
	return fmt.Sprintf("%v  %v  %v", ol.regione, ol.formattedDate, ol.maxValue)
}

func (ols OutputLines) filtered(userInput UserInput) []*OutputLine {
	lines := []*OutputLine{}
	for _, line := range ols.lines {
		// if(line.date )
		if line.date.Equal(userInput.startDate) ||
			line.date.Equal(userInput.endDate) ||
			(line.date.After(userInput.startDate) && line.date.Before(userInput.endDate)) {
			lines = append(lines, line)
		}

	}

	return lines
}
func (ols OutputLines) sort() {
	sort.Slice(ols.lines, func(i, j int) bool {
		first, second := ols.lines[i], ols.lines[j]
		return OrderOfRegions[first.regione] < OrderOfRegions[second.regione]
	})

}

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
