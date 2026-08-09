package api

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const layout = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	date, err := time.Parse(layout, dstart)
	if err != nil {
		return "", errors.New("dstart parse error")
	}
	if repeat == "" {
		return "", errors.New("repeat cannot be empty")
	}
	elements := strings.Split(repeat, " ")
	s := elements[0]
	if len(elements) != 2 {
		if s != "y" {
			return "", errors.New("unknown repeat format: wrong number of elements")
		}
	}
	if s != "d" && s != "y" {
		return "", errors.New("unknown repeat format: wrong time elements")
	}

	d := 0
	if s == "d" {
		d, err = strconv.Atoi(elements[1])
		if err != nil {
			return "", errors.New("wrong repeat format: can't parse number")
		}
	}
	if s == "d" {
		if d < 1 || d > 400 {
			return "", errors.New("wrong repeat format: number must be between 1 and 400")
		}
	}

	for {
		// fmt.Println("Hello, For-Cycle", "now", now, "dstart: ", dstart, "repeat:", repeat)
		if s == "y" {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		if s == "d" {
			date = date.AddDate(0, 0, d)
			if afterNow(date, now) {
				break
			}
		}
	}

	return date.Format(layout), nil
}

func afterNow(date, now time.Time) bool {
	if date.After(now) {
		return true
	}
	return false
}
