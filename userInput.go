package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func getInputDate(msg string) time.Time {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("시작하는 날짜 및 시간을 입력하세요. (2022.09.01 15:00)")
	fmt.Println(msg)
	var t time.Time
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\x0d", "", -1)
		println(text)
		date, err := time.Parse("060102 1504", text)
		if err != nil {
			fmt.Println("잘못된 날짜형식, 다시 입력 해주세요 ")
			continue
		}
		t = date
		break
	}
	return t
}
