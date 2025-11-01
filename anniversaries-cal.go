package main

import (
	"flag"
	"fmt"
	"html"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
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
					Description: fmt.Sprintf("%s & %s: %d year anniversary as a %s", marriage.Partner1, marriage.Partner2, years, relationshipLabel),
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
				Description: fmt.Sprintf("%s & %s: 100 months as a %s anniversary", marriage.Partner1, marriage.Partner2, relationshipLabel),
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

func exportToHTML(dates []InterestingDate, filename string) error {
	var htmlBuilder strings.Builder
	
	// Write HTML header and CSS
	htmlBuilder.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Anniversaries Timeline</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        
        .container {
            max-width: 1200px;
            margin: 0 auto;
        }
        
        h1 {
            color: white;
            text-align: center;
            margin-bottom: 30px;
            font-size: 2.5em;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
        }
        
        .pinned-events {
            background: rgba(255, 255, 255, 0.95);
            border-radius: 12px;
            padding: 20px;
            margin-bottom: 20px;
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
            display: none;
        }
        
        .pinned-events.visible {
            display: block;
        }
        
        .pinned-events h2 {
            color: #764ba2;
            margin-bottom: 15px;
            font-size: 1.3em;
        }
        
        .pinned-event {
            background: linear-gradient(135deg, #ffd89b 0%, #19547b 100%);
            color: white;
            padding: 12px 15px;
            margin-bottom: 10px;
            border-radius: 8px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-weight: 500;
        }
        
        .pinned-event .relative-time {
            font-size: 0.9em;
            opacity: 0.95;
        }
        
        .timeline-container {
            background: rgba(255, 255, 255, 0.95);
            border-radius: 12px;
            padding: 30px;
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
            max-height: 70vh;
            overflow-y: auto;
            position: relative;
        }
        
        .timeline {
            position: relative;
            padding-left: 40px;
        }
        
        .timeline::before {
            content: '';
            position: absolute;
            left: 15px;
            top: 0;
            bottom: 0;
            width: 3px;
            background: linear-gradient(180deg, #667eea 0%, #764ba2 100%);
        }
        
        .timeline-item {
            position: relative;
            padding-bottom: 30px;
            cursor: pointer;
            transition: transform 0.2s;
        }
        
        .timeline-item:hover {
            transform: translateX(5px);
        }
        
        .timeline-item::before {
            content: '';
            position: absolute;
            left: -32px;
            top: 5px;
            width: 16px;
            height: 16px;
            border-radius: 50%;
            background: white;
            border: 3px solid #667eea;
            box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.1);
            transition: all 0.3s;
        }
        
        .timeline-item:hover::before {
            transform: scale(1.3);
            box-shadow: 0 0 0 8px rgba(102, 126, 234, 0.2);
        }
        
        .timeline-item.pinned::before {
            background: #ffd89b;
            border-color: #19547b;
            animation: pulse 2s infinite;
        }
        
        @keyframes pulse {
            0%, 100% {
                box-shadow: 0 0 0 4px rgba(255, 216, 155, 0.4);
            }
            50% {
                box-shadow: 0 0 0 8px rgba(255, 216, 155, 0.1);
            }
        }
        
        .timeline-item.today::before {
            background: #ff6b6b;
            border-color: #ee5a52;
            animation: todayPulse 1.5s infinite;
        }
        
        @keyframes todayPulse {
            0%, 100% {
                box-shadow: 0 0 0 4px rgba(255, 107, 107, 0.4);
            }
            50% {
                box-shadow: 0 0 0 10px rgba(255, 107, 107, 0.1);
            }
        }
        
        .timeline-item.past {
            opacity: 0.7;
        }
        
        .event-card {
            background: white;
            border-radius: 8px;
            padding: 20px;
            box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
            transition: all 0.3s;
        }
        
        .timeline-item:hover .event-card {
            box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
        }
        
        .event-date {
            font-size: 0.9em;
            color: #667eea;
            font-weight: 600;
            margin-bottom: 8px;
        }
        
        .event-description {
            color: #333;
            font-size: 1.1em;
            margin-bottom: 8px;
            line-height: 1.4;
        }
        
        .event-relative {
            font-size: 0.85em;
            color: #666;
        }
        
        .status-badge {
            display: inline-block;
            padding: 4px 12px;
            border-radius: 12px;
            font-size: 0.75em;
            font-weight: 600;
            margin-bottom: 8px;
        }
        
        .status-past {
            background: #e8eaf6;
            color: #5c6bc0;
        }
        
        .status-today {
            background: #ffebee;
            color: #e53935;
            animation: todayBadgePulse 2s infinite;
        }
        
        @keyframes todayBadgePulse {
            0%, 100% {
                background: #ffebee;
            }
            50% {
                background: #ffcdd2;
            }
        }
        
        .status-upcoming {
            background: #e8f5e9;
            color: #43a047;
        }
        
        .pin-button {
            background: transparent;
            border: 2px solid #667eea;
            color: #667eea;
            padding: 6px 12px;
            border-radius: 6px;
            cursor: pointer;
            font-size: 0.85em;
            font-weight: 600;
            transition: all 0.3s;
            margin-top: 8px;
        }
        
        .pin-button:hover {
            background: #667eea;
            color: white;
        }
        
        .pin-button.pinned {
            background: linear-gradient(135deg, #ffd89b 0%, #19547b 100%);
            border-color: #19547b;
            color: white;
        }
        
        ::-webkit-scrollbar {
            width: 10px;
        }
        
        ::-webkit-scrollbar-track {
            background: #f1f1f1;
            border-radius: 10px;
        }
        
        ::-webkit-scrollbar-thumb {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            border-radius: 10px;
        }
        
        ::-webkit-scrollbar-thumb:hover {
            background: linear-gradient(135deg, #764ba2 0%, #667eea 100%);
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>✨ Anniversaries Timeline ✨</h1>
        
        <div class="pinned-events" id="pinnedEvents">
            <h2>📌 Pinned Events</h2>
            <div id="pinnedList"></div>
        </div>
        
        <div class="timeline-container">
            <div class="timeline" id="timeline">
`)

	// Add timeline items
	for _, date := range dates {
		status := "upcoming"
		statusClass := "status-upcoming"
		itemClass := ""
		
		if date.DaysFromNow < 0 {
			status = "past"
			statusClass = "status-past"
			itemClass = "past"
		} else if date.DaysFromNow == 0 {
			status = "TODAY"
			statusClass = "status-today"
			itemClass = "today"
		}
		
		daysText := fmt.Sprintf("%d days", date.DaysFromNow)
		if date.DaysFromNow == 0 {
			daysText = "Today!"
		} else if date.DaysFromNow == 1 {
			daysText = "Tomorrow"
		} else if date.DaysFromNow == -1 {
			daysText = "Yesterday"
		} else if date.DaysFromNow > 0 {
			daysText = fmt.Sprintf("in %d days", date.DaysFromNow)
		} else {
			daysText = fmt.Sprintf("%d days ago", -date.DaysFromNow)
		}
		
		htmlBuilder.WriteString(fmt.Sprintf(`
                <div class="timeline-item %s" data-date="%s" data-days="%d">
                    <div class="event-card">
                        <span class="status-badge %s">%s</span>
                        <div class="event-date">%s</div>
                        <div class="event-description">%s</div>
                        <div class="event-relative">%s</div>
                        <button class="pin-button" onclick="togglePin(this)">📌 Pin</button>
                    </div>
                </div>
`, itemClass, html.EscapeString(date.Date.Format("2006-01-02")), date.DaysFromNow, statusClass, html.EscapeString(status), 
			html.EscapeString(date.Date.Format("January 2, 2006")), html.EscapeString(date.Description), html.EscapeString(daysText)))
	}

	// Write closing HTML and JavaScript
	htmlBuilder.WriteString(`
            </div>
        </div>
    </div>
    
    <script>
        const pinnedEvents = new Set();
        
        function togglePin(button) {
            const item = button.closest('.timeline-item');
            const date = item.dataset.date;
            const description = item.querySelector('.event-description').textContent;
            const days = parseInt(item.dataset.days);
            
            if (pinnedEvents.has(date)) {
                pinnedEvents.delete(date);
                item.classList.remove('pinned');
                button.classList.remove('pinned');
                button.textContent = '📌 Pin';
            } else {
                pinnedEvents.add(date);
                item.classList.add('pinned');
                button.classList.add('pinned');
                button.textContent = '✓ Pinned';
            }
            
            updatePinnedDisplay();
            savePinnedToStorage();
        }
        
        function updatePinnedDisplay() {
            const pinnedList = document.getElementById('pinnedList');
            const pinnedEventsContainer = document.getElementById('pinnedEvents');
            
            if (pinnedEvents.size === 0) {
                pinnedEventsContainer.classList.remove('visible');
                return;
            }
            
            pinnedEventsContainer.classList.add('visible');
            pinnedList.innerHTML = '';
            
            // Get currently visible items in viewport
            const timeline = document.querySelector('.timeline-container');
            const timelineRect = timeline.getBoundingClientRect();
            const items = document.querySelectorAll('.timeline-item');
            
            let visibleDays = null;
            items.forEach(item => {
                const rect = item.getBoundingClientRect();
                if (rect.top >= timelineRect.top && rect.top <= timelineRect.bottom) {
                    if (visibleDays === null) {
                        visibleDays = parseInt(item.dataset.days);
                    }
                }
            });
            
            // Display pinned events with relative time
            pinnedEvents.forEach(date => {
                // Use Array.from to safely search for elements by data attribute
                const items = Array.from(document.querySelectorAll('.timeline-item'));
                const item = items.find(el => el.dataset.date === date);
                if (!item) return;
                
                const description = item.querySelector('.event-description').textContent;
                const days = parseInt(item.dataset.days);
                
                let relativeText = '';
                if (visibleDays !== null && days !== visibleDays) {
                    const diff = days - visibleDays;
                    const absDiff = Math.abs(diff);
                    
                    if (absDiff === 0) {
                        relativeText = 'Same day';
                    } else if (absDiff === 1) {
                        relativeText = diff > 0 ? '+1 day' : '-1 day';
                    } else if (absDiff < 30) {
                        relativeText = diff > 0 ? '+' + absDiff + ' days' : '-' + absDiff + ' days';
                    } else if (absDiff < 365) {
                        const months = Math.round(absDiff / 30);
                        relativeText = diff > 0 ? '+' + months + ' month' + (months > 1 ? 's' : '') : '-' + months + ' month' + (months > 1 ? 's' : '');
                    } else {
                        const years = Math.round(absDiff / 365);
                        relativeText = diff > 0 ? '+' + years + ' year' + (years > 1 ? 's' : '') : '-' + years + ' year' + (years > 1 ? 's' : '');
                    }
                }
                
                const pinnedDiv = document.createElement('div');
                pinnedDiv.className = 'pinned-event';
                // Use textContent for safer DOM manipulation
                const descSpan = document.createElement('span');
                descSpan.textContent = description;
                const timeSpan = document.createElement('span');
                timeSpan.className = 'relative-time';
                timeSpan.textContent = relativeText;
                pinnedDiv.appendChild(descSpan);
                pinnedDiv.appendChild(timeSpan);
                pinnedList.appendChild(pinnedDiv);
            });
        }
        
        function savePinnedToStorage() {
            localStorage.setItem('pinnedEvents', JSON.stringify([...pinnedEvents]));
        }
        
        function loadPinnedFromStorage() {
            const saved = localStorage.getItem('pinnedEvents');
            if (saved) {
                const dates = JSON.parse(saved);
                dates.forEach(date => {
                    // Use Array.from to safely search for elements by data attribute
                    const items = Array.from(document.querySelectorAll('.timeline-item'));
                    const item = items.find(el => el.dataset.date === date);
                    if (item) {
                        pinnedEvents.add(date);
                        item.classList.add('pinned');
                        const button = item.querySelector('.pin-button');
                        button.classList.add('pinned');
                        button.textContent = '✓ Pinned';
                    }
                });
                updatePinnedDisplay();
            }
        }
        
        // Update pinned display on scroll
        document.querySelector('.timeline-container').addEventListener('scroll', () => {
            if (pinnedEvents.size > 0) {
                updatePinnedDisplay();
            }
        });
        
        // Scroll to today on load
        window.addEventListener('load', () => {
            loadPinnedFromStorage();
            
            const todayItem = document.querySelector('.timeline-item.today');
            if (todayItem) {
                todayItem.scrollIntoView({ behavior: 'smooth', block: 'center' });
            }
        });
    </script>
</body>
</html>`)

	// Write to file
	return os.WriteFile(filename, []byte(htmlBuilder.String()), 0644)
}

func main() {
	// Define command-line flags
	var configFile string
	var icalFile string
	var htmlFile string

	flag.StringVar(&configFile, "config", "anniversaries.yaml", "Path to config YAML file")
	flag.StringVar(&configFile, "c", "anniversaries.yaml", "Path to config YAML file (shorthand)")
	flag.StringVar(&icalFile, "ical", "", "Path to export iCal file (optional)")
	flag.StringVar(&htmlFile, "html", "", "Path to export HTML timeline file (optional)")
	flag.Parse()

	// If a positional argument is provided, use it as config file (for backwards compatibility)
	if flag.NArg() > 0 {
		configFile = flag.Arg(0)
	}

	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Printf("Error loading config file %s: %v\n", configFile, err)
		os.Exit(1)
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

	// Export to HTML if requested
	if htmlFile != "" {
		err := exportToHTML(interestingDates, htmlFile)
		if err != nil {
			fmt.Printf("Error exporting to HTML: %v\n", err)
		} else {
			fmt.Printf("Successfully exported to %s\n", htmlFile)
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
