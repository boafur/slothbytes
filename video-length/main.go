package main

import (
	"strconv"
	"strings"
	test "video-length/testing_suite"
)

func toSeconds(timeString string) (int, error) {
	minutes, err := strconv.Atoi(strings.Split(timeString, ":")[0])
	seconds, err := strconv.Atoi(strings.Split(timeString, ":")[1])

	if seconds >= 60 {
		return -1, nil
	}

	return seconds + minutes*60, err
}

func main() {
	tests := make(map[string]int) // since we use a map, tests will not be run in the same order as we declare them :)
	tests["1:00"] = 60
	tests["01:00"] = 60
	tests["13:56"] = 836
	tests["10:60"] = -1
	tests["121:49"] = 7309
	for timeString, shouldEqual := range tests {
		output, err := toSeconds(timeString)
		test.LogAssert(timeString, output, shouldEqual, err)
	}
}
