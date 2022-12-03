package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print(GetMaxVisitors("0076_testdata_05"))
}

func GetMaxVisitors(filePath string) int {

	var n int
	visitorsIncrement := (new([24*60 + 1]int))[:]

	file, _ := os.Open(filePath)
	defer file.Close()

	reader := bufio.NewReader(file)

	nString, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.Trim(nString, "\n"))

	var line string
	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.Trim(line, "\n")

		comingTime, departureTime := getVisitorSchedule(line)

		visitorsIncrement[convertTimeToMinutes(comingTime)]++
		visitorsIncrement[convertTimeToMinutes(departureTime)+1]--
	}

	nVisitors := visitorsIncrement[0]
	nMaxVisitors := visitorsIncrement[0]

	for i := 1; i < len(visitorsIncrement)-1; i++ {
		nVisitors = nVisitors + visitorsIncrement[i]

		if nVisitors > nMaxVisitors {
			nMaxVisitors = nVisitors
		}
	}

	return nMaxVisitors
}

func getVisitorSchedule(line string) (string, string) {
	visitorSchedule := strings.Split(line, " ")
	return visitorSchedule[0], visitorSchedule[1]
}

func convertTimeToMinutes(time string) int {
	splittedTime := strings.Split(time, ":")
	hours, _ := strconv.Atoi(splittedTime[0])
	minutes, _ := strconv.Atoi(splittedTime[1])

	return 60*hours + minutes
}
