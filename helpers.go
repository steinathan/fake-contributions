package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// parseCommitsPerDay parses a string with comma-separated integers.
func parseCommitsPerDay(commitsPerDay int) ([]int, error) {
	var intervals []int
	for _, v := range splitAndTrim(strconv.Itoa(commitsPerDay), ",") {
		interval, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		intervals = append(intervals, interval)
	}
	return intervals, nil
}

// splitAndTrim splits a string by a separator and trims each part.
func splitAndTrim(s, sep string) []string {
	splitStr := strings.Split(s, sep)
	for i, v := range splitStr {
		splitStr[i] = strings.TrimSpace(v)
	}
	return splitStr
}

// parseDateOrDefault parses a date string or returns a default value.
func parseDateOrDefault(dateStr string, defaultValue time.Time) time.Time {
	if dateStr == "" {
		return defaultValue
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return defaultValue
	}
	return date
}

// createCommitDateList  creates a list of commit dates from a list of intervals.
func createCommitDateList(intervals []int, workdaysOnly bool, startDate, endDate time.Time) []time.Time {
	var commitDateList []time.Time
	currentDate := startDate

	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		if !workdaysOnly || !isWeekend(currentDate) {
			numCommits := intervals[getRandomIntInclusive(0, len(intervals)-1)]
			for i := 0; i < numCommits; i++ {
				commitDate := setRandomHours(currentDate)
				commitDateList = append(commitDateList, commitDate)
			}
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return commitDateList
}

// isWeekend  returns true if the date is a weekend.
func isWeekend(date time.Time) bool {
	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
}

// getRandomIntInclusive returns a random number between min and max inclusive.
func getRandomIntInclusive(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// setRandomHours sets random hours for a given date.
func setRandomHours(date time.Time) time.Time {
	hour := getRandomIntInclusive(9, 16)
	minute := getRandomIntInclusive(0, 59)
	second := getRandomIntInclusive(0, 59)
	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, second, 0, date.Location())
}
