package main

func main() {
	var input UserInput
	input.baseDate = getInputDate("시작하는 날짜 및 시간을 입력하세요. (220901 1500)")
	input.startDate = getInputDate("(시작일) 추출하고 싶은 날짜 및 시간을 입력하세요 (220901 1500)")
	input.endDate = getInputDate("(종료일) 추출하고 싶은 날짜 및 시간을 입력하세요 (220901 1500)")
	fileNames := readDir(INPUT_DIR)
	files := readFiles(fileNames)
	contents := parseFiles(files)
	var ols OutputLines
	for _, content := range contents {
		ols.lines = append(ols.lines, content.toOutputLine(input.baseDate))
	}
	ols.sort()
	writeFile("all.txt", ols.lines)
	writeFile("filtered.txt", ols.filtered(input))
}
