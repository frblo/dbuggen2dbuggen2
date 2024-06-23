package parser

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func extractDate(filename string) time.Time {
	date, err := time.Parse("2006-01-02T15", fmt.Sprintf("%vT01", filename[0:10]))
	if err != nil {
		date = time.Now()
		log.Println(err)
	}

	return date
}

func averageDate(dates []time.Time, issueMonthString string) time.Time {
	sum := int64(0)
	for _, date := range dates {
		sum += date.Unix()
	}

	mean := sum / int64(len(dates))
	meanDate := time.Unix(mean, 0)

	issueMonth, err := time.Parse("Jan 2006", issueMonthString)
	if err != nil {
		issueMonthString, err = translateDate(issueMonthString)
		if err != nil {
			log.Print(err)
			return meanDate
		}

		issueMonth, err = time.Parse("Jan 2006", issueMonthString)
		if err != nil {
			log.Print(err)
			return meanDate
		}
	}

	if meanDate.Month() != issueMonth.Month() || meanDate.Year() != issueMonth.Year() {
		zone, _ := time.LoadLocation("Europe/Stockholm")
		return time.Date(issueMonth.Year(), issueMonth.Month(), 0, 0, 0, 0, 0, zone)
	}

	return meanDate
}

func translateDate(dateString string) (string, error) {
	untranslatedMonth, year, found := strings.Cut(dateString, " ")

	if !found {
		return "", errors.New(`filename not formatted as "MONTH YEAR"`)
	}

	var month string
	switch strings.ToLower(untranslatedMonth) {
	case "januari":
		month = "Jan"
	case "februari":
		month = "Feb"
	case "mars":
		month = "Mar"
	case "april":
		month = "Apr"
	case "maj":
		month = "May"
	case "juni":
		month = "Jun"
	case "juli":
		month = "Jul"
	case "augusti":
		month = "Aug"
	case "oktober":
		month = "Oct"
	case "okt":
		month = "Oct"
	case "november":
		month = "nov"
	case "december":
		month = "dec"
	default:
		return "", errors.New(fmt.Sprintf("No translation for %v available", dateString))
	}

	return fmt.Sprintf("%v %v", month, year), nil
}
