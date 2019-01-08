package dateutil

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDifference(t *testing.T) {
	from := time.Date(2018, 1, 31, 5, 0, 0, 0, time.UTC)
	to := time.Date(2019, 2, 1, 1, 30, 59, 0, time.UTC)
	if !checkDate(from, to) {
		t.Error("Incorrect, Did't get same to when difference added with from time")
	}

	from = time.Date(2019, 1, 31, 5, 0, 0, 0, time.UTC)
	to = time.Date(2019, 2, 1, 1, 30, 59, 0, time.UTC)
	if !checkDate(from, to) {
		t.Error("Incorrect, Did't get same to when difference added with from time")
	}

	from = time.Date(1993, 4, 8, 22, 55, 49, 0, time.UTC)
	to = time.Date(2019, 4, 22, 10, 30, 59, 0, time.UTC)
	if !checkDate(from, to) {
		t.Error("Incorrect, Did't get same to when difference added with from time")
	}

	from = time.Date(1993, 12, 6, 0, 55, 49, 0, time.UTC)
	to = time.Date(2019, 4, 22, 10, 30, 59, 0, time.UTC)
	if !checkDate(from, to) {
		t.Error("Incorrect, Did't get same to when difference added with from time")
	}

	from = time.Date(2019, 12, 22, 10, 30, 59, 0, time.UTC)
	to = time.Date(2019, 12, 22, 10, 30, 59, 0, time.UTC)
	if !checkDate(from, to) {
		t.Error("Incorrect, Did't get same to when difference added with from time")
	}

	from = time.Date(2019, 12, 22, 10, 30, 59, 0, time.UTC)
	tempT := from.Add(time.Duration(rand.Int()))
	to = time.Date(tempT.Year(), tempT.Month(), tempT.Day(), tempT.Hour(), tempT.Minute(), tempT.Second(), 0, time.UTC)
	if !checkDate(from, to) {
		t.Error("Incorrect, Did't get same to when difference added with from time")
	}

	to = time.Date(2019, 12, 22, 10, 30, 59, 0, time.UTC)
	tempT = to.Add(time.Duration(rand.Int()) * -1)
	from = time.Date(tempT.Year(), tempT.Month(), tempT.Day(), tempT.Hour(), tempT.Minute(), tempT.Second(), 0, time.UTC)
	if !checkDate(from, to) {
		t.Error("Incorrect, Did't get same to when difference added with from time")
	}
}

func checkDate(from, to time.Time) bool {
	dif := Difference(from, to)
	fmt.Println(dif)
	fmt.Println(from)
	fmt.Println(to)
	fmt.Printf("%d Years %d Months %d Days %d H %d M %d S\n", dif.Years, dif.Months, dif.Days, dif.Hours, dif.Minutes, dif.Seconds)
	check := from.AddDate(dif.Years, dif.Months, dif.Days)
	check = check.Add(time.Hour * time.Duration(dif.Hours))
	check = check.Add(time.Minute * time.Duration(dif.Minutes))
	check = check.Add(time.Second * time.Duration(dif.Seconds))
	fmt.Println(check)
	return to.Equal(check)
}
