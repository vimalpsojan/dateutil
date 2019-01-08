package main

import (
	"fmt"
	"github.com/srlbv/dateutil"
	"time"
)

func main() {
	from := time.Date(1993, 4, 8, 22, 50, 42, 0, time.UTC)
	to := time.Date(2019, 1, 9, 00, 40, 39, 0, time.UTC)
	checkDate(from, to)
}

func checkDate(from, to time.Time) {
	dif := dateutil.Difference(from, to)
	fmt.Println(dif)
	fmt.Println(from)
	fmt.Println(to)
	fmt.Printf("%d Years %d Months %d Days %d H %d M %d S\n", dif.Years, dif.Months, dif.Days, dif.Hours, dif.Minutes, dif.Seconds)
	check := from.AddDate(dif.Years, dif.Months, dif.Days)
	check = check.Add(time.Hour * time.Duration(dif.Hours))
	check = check.Add(time.Minute * time.Duration(dif.Minutes))
	check = check.Add(time.Second * time.Duration(dif.Seconds))
	fmt.Println(check)

	if to.Equal(check) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Error")
	}
}
