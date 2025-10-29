package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"

	"github.com/uniplaces/carbon"
	"gopkg.in/yaml.v3"
)

type Person struct {
	Name      string `yaml:"name"`
	Birthdate string `yaml:"birthdate"`
}

type Event struct {
	Name string `yaml:"name"`
	Date string `yaml:"date"`
}

type Marriage struct {
	Partner1 string `yaml:"partner1"`
	Partner2 string `yaml:"partner2"`
	Date     string `yaml:"date"`
}

type Config struct {
	People    []Person   `yaml:"people"`
	Events    []Event    `yaml:"events"`
	Marriages []Marriage `yaml:"marriages"`
}

type InterestingDate struct {
	Description string
	Date        time.Time
	DaysFromNow int64
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func calculateInterestingDates(config Config) []InterestingDate {
	var dates []InterestingDate
	now := time.Now()

	// Process people - calculate interesting birthdays
	for _, person := range config.People {
		birthdate, err := parseDate(person.Birthdate)
		if err != nil {
			fmt.Printf("Error parsing date for %s: %v\n", person.Name, err)
			continue
		}

		// 10,000 days birthday
		date10k := birthdate.AddDate(0, 0, 10000)
		if date10k.After(now.AddDate(-1, 0, 0)) { // Show if within past year or future
			daysFromNow := int64(date10k.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's 10,000 days birthday", person.Name),
				Date:        date10k,
				DaysFromNow: daysFromNow,
			})
		}

		// 100 months birthday
		date100m := birthdate.AddDate(0, 100, 0)
		if date100m.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date100m.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's 100 months birthday", person.Name),
				Date:        date100m,
				DaysFromNow: daysFromNow,
			})
		}
	}

	// Process events - calculate powers of 2 days
	for _, event := range config.Events {
		eventDate, err := parseDate(event.Date)
		if err != nil {
			fmt.Printf("Error parsing date for %s: %v\n", event.Name, err)
			continue
		}

		// Calculate powers of 2 days (2^10 through 2^20)
		for power := 10; power <= 20; power++ {
			days := int(math.Pow(2, float64(power)))
			dateAtPower := eventDate.AddDate(0, 0, days)
			if dateAtPower.After(now.AddDate(-1, 0, 0)) && dateAtPower.Before(now.AddDate(10, 0, 0)) {
				daysFromNow := int64(dateAtPower.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s: 2^%d days (%d days)", event.Name, power, days),
					Date:        dateAtPower,
					DaysFromNow: daysFromNow,
				})
			}
		}
	}

	// Process marriages - calculate year anniversaries
	for _, marriage := range config.Marriages {
		marriageDate, err := parseDate(marriage.Date)
		if err != nil {
			fmt.Printf("Error parsing date for marriage: %v\n", err)
			continue
		}

		// Calculate anniversaries for next 50 years
		yearsSince := now.Year() - marriageDate.Year()
		for years := yearsSince; years <= yearsSince+10; years++ {
			if years <= 0 {
				continue
			}
			anniversaryDate := marriageDate.AddDate(years, 0, 0)
			if anniversaryDate.After(now.AddDate(-1, 0, 0)) && anniversaryDate.Before(now.AddDate(2, 0, 0)) {
				daysFromNow := int64(anniversaryDate.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s & %s: %d year marriage anniversary", marriage.Partner1, marriage.Partner2, years),
					Date:        anniversaryDate,
					DaysFromNow: daysFromNow,
				})
			}
		}

		// 100 months anniversary
		date100m := marriageDate.AddDate(0, 100, 0)
		if date100m.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date100m.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s & %s: 100 months marriage anniversary", marriage.Partner1, marriage.Partner2),
				Date:        date100m,
				DaysFromNow: daysFromNow,
			})
		}
	}

	// Calculate when one person is half the age of another
	for i, person1 := range config.People {
		for j, person2 := range config.People {
			if i >= j {
				continue
			}
			birthdate1, err1 := parseDate(person1.Birthdate)
			birthdate2, err2 := parseDate(person2.Birthdate)
			if err1 != nil || err2 != nil {
				continue
			}

			// Determine who is younger and who is older
			var youngerName, olderName string
			var youngerBirth, olderBirth time.Time
			
			if birthdate1.After(birthdate2) {
				youngerName = person1.Name
				olderName = person2.Name
				youngerBirth = birthdate1
				olderBirth = birthdate2
			} else {
				youngerName = person2.Name
				olderName = person1.Name
				youngerBirth = birthdate2
				olderBirth = birthdate1
			}

			// The date when younger person is half the age of older person
			// Age_younger = Age_older / 2
			// (date - youngerBirth) = (date - olderBirth) / 2
			// 2 * (date - youngerBirth) = date - olderBirth
			// 2*date - 2*youngerBirth = date - olderBirth
			// date = 2*youngerBirth - olderBirth
			ageDiff := youngerBirth.Sub(olderBirth)
			halfAgeDate := youngerBirth.Add(ageDiff)
			if halfAgeDate.After(now.AddDate(-1, 0, 0)) && halfAgeDate.Before(now.AddDate(10, 0, 0)) {
				daysFromNow := int64(halfAgeDate.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s is half the age of %s", youngerName, olderName),
					Date:        halfAgeDate,
					DaysFromNow: daysFromNow,
				})
			}
		}
	}

	return dates
}

func loadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	configFile := "anniversaries.yaml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Printf("Error loading config file %s: %v\n", configFile, err)
		fmt.Println("Using default data...")

		// Fallback to hardcoded data
		my_dates := []struct {
			Name string
			Date time.Time
		}{
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
		}
		return
	}

	// Calculate interesting dates
	interestingDates := calculateInterestingDates(*config)

	// Sort dates by how soon they occur
	// Simple bubble sort for demonstration
	for i := 0; i < len(interestingDates)-1; i++ {
		for j := 0; j < len(interestingDates)-i-1; j++ {
			if interestingDates[j].Date.After(interestingDates[j+1].Date) {
				interestingDates[j], interestingDates[j+1] = interestingDates[j+1], interestingDates[j]
			}
		}
	}

	// Display interesting dates
	fmt.Println("=== Interesting Dates Calendar ===\n")
	for _, id := range interestingDates {
		status := "upcoming"
		if id.DaysFromNow < 0 {
			status = "past"
		} else if id.DaysFromNow == 0 {
			status = "TODAY"
		}
		fmt.Printf("[%s] %s - %s (in %d days)\n",
			status,
			id.Date.Format("2006-01-02"),
			id.Description,
			id.DaysFromNow,
		)
	}
}
