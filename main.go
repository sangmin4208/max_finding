package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const INPUT_DIR = "input"
const OUPUT_PATH = "output/output.txt"

type Content struct {
	regione string
	rows    []Row
}

type Row struct {
	index int
	value float64
}
type WriteRow struct {
	regione  string
	maxValue string
	date     string
}

type Input struct {
	date     time.Time
	interval int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input Input
	fmt.Println("날짜 입력 (2022.09.01 15:00)")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\x0d", "", -1)
		println(text)
		testDate, err := time.Parse("2006.01.02 15:04", text)
		if err != nil {
			fmt.Println("잘못된 날짜형식 다시 입력 ")
			continue
		}
		input.date = testDate
		break
	}
	fmt.Println("시간 간격(분) 입력 (숫자만)")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\x0d", "", -1)
		interval, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("숫자만 입력 ")
			continue
		}
		input.interval = interval
		break
	}

	contents := readFiles()
	writeRows := []WriteRow{}

	for _, content := range contents {
		writeRows = append(writeRows, parseWriteRow(content, input))
	}

	sortWriteRows(writeRows)
	writeFile(writeRows)
}

func writeFile(writeRows []WriteRow) {
	result := []string{}
	for _, row := range writeRows {
		result = append(result, fmt.Sprintf("%v %v %v", row.date, row.regione, row.maxValue))
	}
	err := ioutil.WriteFile(OUPUT_PATH, []byte(strings.Join(result, "\n")), 0777)
	if err != nil {
		println(err.Error())
		panic("파일 작성 실패, 폴더 확인")
	}
}

func sortWriteRows(writeRows []WriteRow) {
	sort.Slice(writeRows, func(i, j int) bool {
		first, second := writeRows[i], writeRows[j]
		return orderOfRegions[first.regione] < orderOfRegions[second.regione]
	})

}
func parseWriteRow(content Content, input Input) WriteRow {
	maxRow := getMaxRow(content)
	date := input.date.Add(time.Duration(time.Minute * time.Duration(input.interval*(maxRow.index))))
	return WriteRow{
		regione:  content.regione,
		maxValue: fmt.Sprintf("%v", maxRow.value),
		date:     date.Format("2006.01.02 15:04:05"),
	}
}

func getMaxRow(content Content) Row {
	rows := content.rows
	sort.Slice(rows, func(i int, j int) bool {
		return rows[i].value > rows[j].value
	})
	return rows[0]
}

func readFiles() []Content {
	files, err := ioutil.ReadDir(INPUT_DIR)
	if err != nil {
		panic("input 폴더가 없음")
	}
	result := []Content{}
	for _, file := range files {
		result = append(result, parseFiles(file))
	}
	return result
}
func parseFiles(file fs.FileInfo) Content {
	regione := parseName(file.Name())
	readFile, err := os.Open(
		fmt.Sprintf("%v/%v", INPUT_DIR, file.Name()),
	)
	if err != nil {
		panic("파일 읽기 실패")
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	rows := []Row{}
	for fileScanner.Scan() {
		rows = append(rows, parseRow(fileScanner.Text()))
	}
	return Content{
		regione: regione,
		rows:    rows,
	}
}
func parseName(fileName string) string {
	// fort.63_Diff_ANHEUNG(ANHG).dat
	_, result, _ := strings.Cut(fileName, "fort.63_Diff_")
	// ANHEUNG(ANHG).dat
	result = strings.Split(result, "(")[0]
	return result
}

func parseRow(row string) Row {
	row = strings.TrimSpace(row)
	splited := strings.Split(row, " ")
	i, v := splited[0], strings.TrimSpace(strings.Join(splited[1:], ""))
	idx, err := strconv.Atoi(i)
	if err != nil {
		panic("인덱스가 이상함")
	}
	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic("값이 이상함")
	}
	return Row{idx, value}
}

var orderOfRegions = map[string]int{
	"INCHEON":          0,
	"ANSAN":            1,
	"PYEONGTAEK":       2,
	"DAESAN":           3,
	"ANHEUNG":          4,
	"BORYEONG":         5,
	"EOCHEONGDO":       6,
	"JANGHANG":         7,
	"GUNSAN_OUTERPORT": 8,
	"WIDO":             9,
	"YEONGGWANG":       10,
	"MOKPO":            11,
	"HEUKSANDO":        12,
	"JINDO":            13,
	"WANDO":            14,
	"GEOMUNDO":         15,
	"GOHEUNG_BALPO":    16,
	"YEOSU":            17,
	"TONGYEONG":        18,
	"GEOJEDO":          19,
	"MASAN":            20,
	"BUSAN":            21,
	"JEJU":             22,
	"SEOGWIP":          23,
	"SEONGSANPO":       24,
	"MOSEULOP":         25,
	"CHUJADO":          26,
	"ULSAN":            27,
	"POHANG":           28,
	"HUPO":             29,
	"MUKHO":            30,
	"SOKCHO":           31,
	"ULLEUNGDO":        32,
}

// MUKHO
// SEOGIPO
// INCHEON
// TONGYEONG
// GOHEUNG_BALPO
// ANSAN
// PYEONGTAEK
// DAESAN
// ANHEUNG
// BORYEONG
// EOCHEONGDO
// JANGHANG
// GUNSAN_OUTERPORT
// WIDO
// YEONGGWANG
// MOKPO
// HEUKSANDO
// JINDO
// WANDO
// GEOMUNDO
// YEOSU
// GEOJEDO
// MASAN
// BUSAN
// JEJU
// SEONGSANPO
// MOSEULOP
// CHUJADO
// ULSAN
// POHANG
// HUPO
// SOKCHO
// ULLEUNGDO
