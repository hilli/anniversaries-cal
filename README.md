# Anniversaries Calendar

A Go program that generates interesting date calculations from a YAML configuration file containing people's birthdates, events, and marriages.

## Features

The program calculates and displays various "interesting dates" including:

### For People (Birthdates)
- **10,000 days birthday** - When someone reaches 10,000 days old
- **20,000 days birthday** - When someone reaches 20,000 days old
- **100 months birthday** - When someone reaches 100 months old
- **500 months birthday** - When someone reaches 500 months old
- **1,000 weeks birthday** - When someone reaches 1,000 weeks old
- **777 days birthday** - Lucky sevens milestone
- **12,345 days birthday** - Sequential digits milestone
- **π×1000 days birthday** - Pi-inspired milestone (3,141 days)
- **Fibonacci days birthdays** - Days matching Fibonacci numbers (233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946)
- **Half-age milestones** - When one person is exactly half the age of another person

### For Events
- **Powers of 2 days** - Calculates dates that are 2^N days (e.g., 2^10=1024, 2^11=2048, 2^12=4096 days) after an event
- **Round year anniversaries** - 5, 10, 25, 50, 75, and 100 year anniversaries
- **1 million minutes** - Exactly 1,000,000 minutes after an event
- **100,000 hours** - Exactly 100,000 hours after an event
- **Prime number days** - Days matching prime numbers (1009, 2003, 3001, 5003, 7001, 10007)

### For Marriages
- **Year anniversaries** - Tracks marriage anniversaries by year
- **Traditional anniversary names** - Paper (1st), Wood (5th), Tin (10th), Crystal (15th), China (20th), Silver (25th), Pearl (30th), Ruby (40th), Golden (50th), Diamond (60th)
- **100 months anniversary** - When a couple reaches 100 months of marriage
- **Days together milestones** - 1,000, 5,000, and 10,000 days together
- **Weeks together milestones** - 100 and 500 weeks together

## Usage

```bash
# Build the program
go build

# Run with default config file (anniversaries.yaml)
./anniversaries-cal

# Run with custom config file
./anniversaries-cal my-dates.yaml
```

## Configuration File Format

Create a YAML file with the following structure:

```yaml
people:
  - name: "Person Name"
    birthdate: "YYYY-MM-DD"

events:
  - name: "Event Name"
    date: "YYYY-MM-DD"

marriages:
  - partner1: "Person 1"
    partner2: "Person 2"
    date: "YYYY-MM-DD"
```

### Example Configuration

```yaml
people:
  - name: "Me"
    birthdate: "1973-07-31"
  - name: "My Girlfriend"
    birthdate: "1972-10-07"
  - name: "First Child"
    birthdate: "2009-07-06"

events:
  - name: "Moon Landing"
    date: "1969-07-20"
  - name: "Company Founded"
    date: "2020-01-15"

marriages:
  - partner1: "Birgitte"
    partner2: "Ulrik"
    date: "2013-09-19"
```

## Output

The program displays dates sorted chronologically with:
- Status: `[past]`, `[TODAY]`, or `[upcoming]`
- Date in YYYY-MM-DD format
- Description of the event
- Days from today (negative for past, positive for future)

Example output:
```
=== Interesting Dates Calendar ===

[past] 2025-01-15 - Company Founded: 5 year anniversary (in -287 days)
[past] 2025-08-24 - Company Founded: 2^11 days (2048 days) (in -66 days)
[past] 2025-09-19 - Birgitte & Ulrik: 12 year marriage anniversary (in -40 days)
[upcoming] 2026-09-19 - Birgitte & Ulrik: 13 year marriage anniversary (in 324 days)
[upcoming] 2027-05-29 - Birgitte & Ulrik: 5,000 days together (in 576 days)
[upcoming] 2027-07-11 - My Girlfriend's 20,000 days birthday (in 619 days)
[upcoming] 2028-01-13 - First Child's Fibonacci 6765 days birthday (in 805 days)
[upcoming] 2028-05-03 - Me's 20,000 days birthday (in 916 days)
[upcoming] 2031-06-12 - Company Founded: 100,000 hours (in 2052 days)
```

## Fallback Mode

If the configuration file is not found, the program falls back to displaying hardcoded anniversary data from the original implementation.

## Dependencies

- [github.com/uniplaces/carbon](https://github.com/uniplaces/carbon) - Date/time helper library
- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) - YAML parsing library
