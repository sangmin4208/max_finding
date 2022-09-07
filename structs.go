package main

import "time"

type Content struct {
	regione string
	lines   []InputLine
}

type InputLine struct {
	index int
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
