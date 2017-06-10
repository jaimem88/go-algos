package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var hourStr string
	fmt.Scanf("%s", &hourStr)

	fmt.Println(convertTo24Hr(hourStr))
}

//In: 07:05:45PM
//Out: 19:05:45
func convertTo24Hr(hour string) string {
	am := regexp.MustCompile("[0-9]{2}AM")
	pm := regexp.MustCompile("[0-9]{2}PM")
	seconds := regexp.MustCompile("[0-9]{2}")
	units := strings.Split(hour, ":")
	newHr := ""

	if am.MatchString(units[2]) {
		// means AM
		if units[0] == "12" {
			newHr = "00"
		} else {
			newHr = units[0]
		}

	} else if pm.MatchString(units[2]) {
		hr, _ := strconv.Atoi(units[0])
		if hr != 12 {
			hr += 12
		}
		newHr = strconv.Itoa(hr)
	}
	return newHr + ":" + units[1] + ":" + seconds.FindString(units[2])
}
