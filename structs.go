package main

import (
	"fmt"
	"sort"
	"time"
)

type Content struct {
	regione string
	lines   []*InputLine
}

type InputLine struct {
	index int
	date  time.Time
	value float64
}
type OutputLine struct {
	regione       string
	maxValue      string
	date          time.Time
	formattedDate string
}

type OutputLines struct {
	lines []*OutputLine
}

type UserInput struct {
	baseDate  time.Time
	startDate time.Time
	endDate   time.Time
}

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

func (c *Content) getMaxValue() *InputLine {
	rows := c.lines
	sort.Slice(rows, func(i int, j int) bool {
		return rows[i].value > rows[j].value
	})
	return rows[0]
}

func (c *Content) filtered(userInput UserInput) []*InputLine {
	lines := []*InputLine{}
	for _, line := range c.lines {
		// if(line.date )
		if line.date.Equal(userInput.startDate) ||
			line.date.Equal(userInput.endDate) ||
			(line.date.After(userInput.startDate) && line.date.Before(userInput.endDate)) {
			lines = append(lines, line)
		}

	}

	return lines
}

func (ol *OutputLine) toPlainText() string {
	// Incheon              2022  09  04  22 10     2.206
	return fmt.Sprintf("%v  %v  %v", ol.regione, ol.formattedDate, ol.maxValue)
}

func (ols OutputLines) sort() {
	sort.Slice(ols.lines, func(i, j int) bool {
		first, second := ols.lines[i], ols.lines[j]
		return OrderOfRegions[first.regione] < OrderOfRegions[second.regione]
	})

}
