package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"time"

	ics "github.com/arran4/golang-ical"
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
	Partner1         string `yaml:"partner1"`
	Partner2         string `yaml:"partner2"`
	Date             string `yaml:"date"`
	RelationshipType string `yaml:"relationship_type,omitempty"` // Optional: "marriage", "couple", "partnership", etc.
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

		// 1,000 weeks birthday
		date1kWeeks := birthdate.AddDate(0, 0, 1000*7)
		if date1kWeeks.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date1kWeeks.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's 1,000 weeks birthday", person.Name),
				Date:        date1kWeeks,
				DaysFromNow: daysFromNow,
			})
		}

		// 777 days birthday (lucky sevens)
		date777 := birthdate.AddDate(0, 0, 777)
		if date777.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date777.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's lucky 777 days birthday", person.Name),
				Date:        date777,
				DaysFromNow: daysFromNow,
			})
		}

		// 12,345 days birthday (sequential digits)
		date12345 := birthdate.AddDate(0, 0, 12345)
		if date12345.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date12345.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's 12,345 days birthday", person.Name),
				Date:        date12345,
				DaysFromNow: daysFromNow,
			})
		}

		// Pi thousand days (3141 days, approximately pi * 1000)
		date3141 := birthdate.AddDate(0, 0, 3141)
		if date3141.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date3141.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's π×1000 days birthday (3,141 days)", person.Name),
				Date:        date3141,
				DaysFromNow: daysFromNow,
			})
		}

		// 500 months birthday
		date500m := birthdate.AddDate(0, 500, 0)
		if date500m.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date500m.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's 500 months birthday", person.Name),
				Date:        date500m,
				DaysFromNow: daysFromNow,
			})
		}

		// 20,000 days birthday
		date20k := birthdate.AddDate(0, 0, 20000)
		if date20k.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date20k.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's 20,000 days birthday", person.Name),
				Date:        date20k,
				DaysFromNow: daysFromNow,
			})
		}

		// 30,000 days birthday
		date30k := birthdate.AddDate(0, 0, 30000)
		if date30k.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date30k.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s's 30,000 days birthday", person.Name),
				Date:        date30k,
				DaysFromNow: daysFromNow,
			})
		}

		// Yearly birthdays up to 100 years
		for years := 1; years <= 100; years++ {
			birthdayDate := birthdate.AddDate(years, 0, 0)
			if birthdayDate.After(now.AddDate(-1, 0, 0)) && birthdayDate.Before(now.AddDate(2, 0, 0)) {
				daysFromNow := int64(birthdayDate.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s's %d year birthday", person.Name, years),
					Date:        birthdayDate,
					DaysFromNow: daysFromNow,
				})
			}
		}

		// Fibonacci days (1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946)
		fibonacciDays := []int{233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946}
		for _, fib := range fibonacciDays {
			dateFib := birthdate.AddDate(0, 0, fib)
			if dateFib.After(now.AddDate(-1, 0, 0)) {
				daysFromNow := int64(dateFib.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s's Fibonacci %d days birthday", person.Name, fib),
					Date:        dateFib,
					DaysFromNow: daysFromNow,
				})
			}
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

		// Round years (5, 10, 25, 50, 75, 100 years)
		roundYears := []int{5, 10, 25, 50, 75, 100}
		for _, years := range roundYears {
			dateAtYear := eventDate.AddDate(years, 0, 0)
			if dateAtYear.After(now.AddDate(-1, 0, 0)) && dateAtYear.Before(now.AddDate(10, 0, 0)) {
				daysFromNow := int64(dateAtYear.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s: %d year anniversary", event.Name, years),
					Date:        dateAtYear,
					DaysFromNow: daysFromNow,
				})
			}
		}

		// Yearly events up to 100 years
		for years := 1; years <= 100; years++ {
			yearDate := eventDate.AddDate(years, 0, 0)
			if yearDate.After(now.AddDate(-1, 0, 0)) && yearDate.Before(now.AddDate(2, 0, 0)) {
				// Skip if already covered by round years
				isRoundYear := false
				for _, roundYear := range roundYears {
					if years == roundYear {
						isRoundYear = true
						break
					}
				}
				if !isRoundYear {
					daysFromNow := int64(yearDate.Sub(now).Hours() / 24)
					dates = append(dates, InterestingDate{
						Description: fmt.Sprintf("%s: %d year anniversary", event.Name, years),
						Date:        yearDate,
						DaysFromNow: daysFromNow,
					})
				}
			}
		}

		// 1 million minutes
		dateMillionMin := eventDate.Add(time.Minute * 1000000)
		if dateMillionMin.After(now.AddDate(-1, 0, 0)) && dateMillionMin.Before(now.AddDate(10, 0, 0)) {
			daysFromNow := int64(dateMillionMin.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s: 1 million minutes", event.Name),
				Date:        dateMillionMin,
				DaysFromNow: daysFromNow,
			})
		}

		// 100,000 hours
		date100kHours := eventDate.Add(time.Hour * 100000)
		if date100kHours.After(now.AddDate(-1, 0, 0)) && date100kHours.Before(now.AddDate(10, 0, 0)) {
			daysFromNow := int64(date100kHours.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s: 100,000 hours", event.Name),
				Date:        date100kHours,
				DaysFromNow: daysFromNow,
			})
		}

		// Prime number days (1009, 2003, 3001, 5003, 7001, 10007)
		primeDays := []int{1009, 2003, 3001, 5003, 7001, 10007}
		for _, prime := range primeDays {
			datePrime := eventDate.AddDate(0, 0, prime)
			if datePrime.After(now.AddDate(-1, 0, 0)) && datePrime.Before(now.AddDate(10, 0, 0)) {
				daysFromNow := int64(datePrime.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s: %d days (prime)", event.Name, prime),
					Date:        datePrime,
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

		// Determine relationship type label
		relationshipLabel := "marriage"
		if marriage.RelationshipType != "" {
			relationshipLabel = marriage.RelationshipType
		}

		// Calculate anniversaries for next 100 years
		for years := 1; years <= 100; years++ {
			anniversaryDate := marriageDate.AddDate(years, 0, 0)
			if anniversaryDate.After(now.AddDate(-1, 0, 0)) && anniversaryDate.Before(now.AddDate(2, 0, 0)) {
				daysFromNow := int64(anniversaryDate.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s & %s: %d year %s anniversary", marriage.Partner1, marriage.Partner2, years, relationshipLabel),
					Date:        anniversaryDate,
					DaysFromNow: daysFromNow,
				})
			}
		}

		// Bronze anniversary (12.5 years = 4562.5 days, we'll use 4563 days)
		bronzeDate := marriageDate.AddDate(0, 0, 4563)
		if bronzeDate.After(now.AddDate(-1, 0, 0)) && bronzeDate.Before(now.AddDate(2, 0, 0)) {
			daysFromNow := int64(bronzeDate.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s & %s: Bronze anniversary (12.5 years)", marriage.Partner1, marriage.Partner2),
				Date:        bronzeDate,
				DaysFromNow: daysFromNow,
			})
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

		// 1,000 days together
		date1kDays := marriageDate.AddDate(0, 0, 1000)
		if date1kDays.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date1kDays.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s & %s: 1,000 days together", marriage.Partner1, marriage.Partner2),
				Date:        date1kDays,
				DaysFromNow: daysFromNow,
			})
		}

		// 5,000 days together
		date5kDays := marriageDate.AddDate(0, 0, 5000)
		if date5kDays.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date5kDays.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s & %s: 5,000 days together", marriage.Partner1, marriage.Partner2),
				Date:        date5kDays,
				DaysFromNow: daysFromNow,
			})
		}

		// 10,000 days together
		date10kDays := marriageDate.AddDate(0, 0, 10000)
		if date10kDays.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date10kDays.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s & %s: 10,000 days together", marriage.Partner1, marriage.Partner2),
				Date:        date10kDays,
				DaysFromNow: daysFromNow,
			})
		}

		// 100 weeks together
		date100Weeks := marriageDate.AddDate(0, 0, 100*7)
		if date100Weeks.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date100Weeks.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s & %s: 100 weeks together", marriage.Partner1, marriage.Partner2),
				Date:        date100Weeks,
				DaysFromNow: daysFromNow,
			})
		}

		// 500 weeks together
		date500Weeks := marriageDate.AddDate(0, 0, 500*7)
		if date500Weeks.After(now.AddDate(-1, 0, 0)) {
			daysFromNow := int64(date500Weeks.Sub(now).Hours() / 24)
			dates = append(dates, InterestingDate{
				Description: fmt.Sprintf("%s & %s: 500 weeks together", marriage.Partner1, marriage.Partner2),
				Date:        date500Weeks,
				DaysFromNow: daysFromNow,
			})
		}

		// Traditional anniversary names for special years
		specialYears := map[int]string{
			1:  "Paper",
			5:  "Wood",
			10: "Tin",
			15: "Crystal",
			20: "China",
			25: "Silver",
			30: "Pearl",
			40: "Ruby",
			50: "Golden",
			60: "Diamond",
		}
		for year, name := range specialYears {
			specialDate := marriageDate.AddDate(year, 0, 0)
			if specialDate.After(now.AddDate(-1, 0, 0)) && specialDate.Before(now.AddDate(2, 0, 0)) {
				daysFromNow := int64(specialDate.Sub(now).Hours() / 24)
				dates = append(dates, InterestingDate{
					Description: fmt.Sprintf("%s & %s: %s anniversary (%d years)", marriage.Partner1, marriage.Partner2, name, year),
					Date:        specialDate,
					DaysFromNow: daysFromNow,
				})
			}
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
	data, err := os.ReadFile(filename)
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

func exportToIcal(dates []InterestingDate, filename string) error {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)
	cal.SetName("Interesting Dates Calendar")
	cal.SetDescription("Anniversaries, birthdays, and special events")

	for _, id := range dates {
		event := cal.AddEvent(fmt.Sprintf("event-%d", id.Date.Unix()))
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(time.Now())
		event.SetStartAt(id.Date)
		event.SetEndAt(id.Date)
		event.SetSummary(id.Description)
		event.SetAllDayStartAt(id.Date)
	}

	// Write to file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return cal.SerializeTo(f)
}

func main() {
	// Define command-line flags
	var configFile string
	var icalFile string
	
	flag.StringVar(&configFile, "config", "anniversaries.yaml", "Path to config YAML file")
	flag.StringVar(&configFile, "c", "anniversaries.yaml", "Path to config YAML file (shorthand)")
	flag.StringVar(&icalFile, "ical", "", "Path to export iCal file (optional)")
	flag.Parse()

	// If a positional argument is provided, use it as config file (for backwards compatibility)
	if flag.NArg() > 0 {
		configFile = flag.Arg(0)
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
	sort.Slice(interestingDates, func(i, j int) bool {
		return interestingDates[i].Date.Before(interestingDates[j].Date)
	})

	// Export to iCal if requested
	if icalFile != "" {
		err := exportToIcal(interestingDates, icalFile)
		if err != nil {
			fmt.Printf("Error exporting to iCal: %v\n", err)
		} else {
			fmt.Printf("Successfully exported to %s\n", icalFile)
		}
	}

	// Display interesting dates
	fmt.Println("=== Interesting Dates Calendar ===")
	fmt.Println()
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
