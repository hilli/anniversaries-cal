package main

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
)

type Anniversary struct {
	Name string
	Date time.Time
}

func (Anniversary) String() string {
	// return *Anniversary.Name
	return "Anniversary"
}

func (a *Anniversary) CarbonDate() *carbon.Carbon {
	cd, _ := carbon.CreateFromDate(a.Date.Year(), a.Date.Month(), a.Date.Day(), "Europe/Copenhagen")
	return cd
	//.Year, time.July, 31, "Europe/Copenhagen")
}

func main() {
	my_dates := []Anniversary{
		{"My birth", time.Date(1973, time.July, 31, 0, 0, 0, 0, time.UTC)},
		{"Anniversary with my girlfriend", time.Date(2017, time.July, 4, 0, 0, 0, 0, time.UTC)},
		{"My girlfriends birthday", time.Date(1972, time.October, 7, 0, 0, 0, 0, time.UTC)},
		{"Anniversary of my first child", time.Date(2009, time.July, 6, 0, 0, 0, 0, time.UTC)},
		{"Anniversary of my second child", time.Date(2013, time.April, 11, 0, 0, 0, 0, time.UTC)},
		{"Anniversary of Birgitte & Ulriks Marriage", time.Date(2013, time.September, 19, 0, 0, 0, 0, time.UTC)},
	}
	now := carbon.Now()
	for _, date := range my_dates {
		if now.After(date.Date) {
			fmt.Printf("%s: %d years ago\n", date.Name, now.Year()-date.Date.Year())
		} else {
			fmt.Printf("%s: %d years to go\n", date.Name, date.Date.Year()-now.Year())
		}
		fmt.Printf("%s is 20000 days ago on the %s\n", date.Name, date.CarbonDate().AddDays(20000))
	}
	fmt.Printf("%d-%d-%d\n", now.Year(), now.Local().Month(), now.Day())
	fmt.Print(carbon.Now().DateString())
	// fmt.Print(carbon.Now().DiffForHumans(carbon.CreateFromDate(2013, time.September, 19)))

	carb, _ := carbon.CreateFromDate(my_dates[0].Date.Year(), my_dates[0].Date.Month(), my_dates[0].Date.Day(), "Europe/Copenhagen")
	diff := carb.DiffInDays(carbon.Now(), true)
	fmt.Println(diff)
	fmt.Println(diff)


}
