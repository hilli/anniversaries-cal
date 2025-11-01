# Anniversaries Calendar

A Go program that generates interesting date calculations from a YAML configuration file containing people's birthdates, events, and marriages.

## Features

The program calculates and displays various "interesting dates" including:

### For People (Birthdates)
- **10,000 days birthday** - When someone reaches 10,000 days old
- **20,000 days birthday** - When someone reaches 20,000 days old
- **30,000 days birthday** - When someone reaches 30,000 days old
- **Yearly birthdays** - Annual birthdays up to 100 years
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
- **Yearly anniversaries** - Annual anniversaries up to 100 years
- **1 million minutes** - Exactly 1,000,000 minutes after an event
- **100,000 hours** - Exactly 100,000 hours after an event
- **Prime number days** - Days matching prime numbers (1009, 2003, 3001, 5003, 7001, 10007)

### For Marriages
- **Year anniversaries** - Tracks marriage/relationship anniversaries up to 100 years
- **Bronze anniversary** - 12.5 years (4,563 days) together
- **Traditional anniversary names** - Paper (1st), Wood (5th), Tin (10th), Crystal (15th), China (20th), Silver (25th), Pearl (30th), Ruby (40th), Golden (50th), Diamond (60th)
- **Custom relationship types** - Optional ability to label relationships (e.g., "couple", "partnership" instead of "marriage")
- **100 months anniversary** - When a couple reaches 100 months together
- **Days together milestones** - 1,000, 5,000, and 10,000 days together
- **Weeks together milestones** - 100 and 500 weeks together

## Usage

```bash
# Build the program
go build

# Run with default config file (anniversaries.yaml)
./anniversaries-cal

# Run with custom config file using -c flag
./anniversaries-cal -c my-dates.yaml

# Run with custom config file using --config flag
./anniversaries-cal --config my-dates.yaml

# Run with custom config file (backwards compatible)
./anniversaries-cal my-dates.yaml

# Export to iCal format
./anniversaries-cal --ical anniversaries.ics

# Export to HTML timeline format
./anniversaries-cal --html timeline.html

# Combine config and exports
./anniversaries-cal -c my-dates.yaml --ical my-calendar.ics --html my-timeline.html
```

### Command-Line Options

- `-c, --config <file>`: Specify the path to a YAML configuration file (default: `anniversaries.yaml`)
- `--ical <file>`: Export all interesting dates to an iCal file for use with calendar applications (optional)
- `--html <file>`: Export all interesting dates to an interactive HTML timeline (optional)

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
    relationship_type: "marriage"  # Optional: "couple", "partnership", etc.
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
  - partner1: "Alice"
    partner2: "Bob"
    date: "2018-05-20"
    relationship_type: "couple"  # Optional: custom relationship label
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

### HTML Timeline Export

The `--html` flag generates an interactive HTML timeline with the following features:

- **Beautiful visual design**: A sleek, gradient-themed interface with smooth animations
- **Scrollable timeline**: Events are displayed chronologically with past events at the top and future events below
- **Pin functionality**: Click the pin button on any event to pin it to the top of the page
- **Relative time indicators**: When you scroll through the timeline, pinned events show their relative distance (in days, months, or years) from the currently visible events
- **Persistent pins**: Pinned events are saved in browser local storage and persist across page reloads
- **Responsive design**: Works on desktop and mobile devices
- **Self-contained**: All CSS and JavaScript are embedded in the HTML file - no external dependencies required

The timeline automatically scrolls to "today" on page load, making it easy to see upcoming events. Past events are shown with reduced opacity to help focus on what's coming next.

![Timeline Example](https://github.com/user-attachments/assets/4f2c06b5-efd4-47e1-a9f9-1cbd914b5aaf)
![Pinned Events with Relative Time](https://github.com/user-attachments/assets/aae6e9b0-94f4-450e-94e0-ac63f5ce283c)

## Dependencies

- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) - YAML parsing library
- [github.com/arran4/golang-ical](https://github.com/arran4/golang-ical) - iCal/ICS file generation library
