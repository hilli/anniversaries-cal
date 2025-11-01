package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Time
		wantErr  bool
	}{
		{"2020-01-15", time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC), false},
		{"1973-07-31", time.Date(1973, 7, 31, 0, 0, 0, 0, time.UTC), false},
		{"invalid", time.Time{}, true},
	}

	for _, tt := range tests {
		got, err := parseDate(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("parseDate(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && !got.Equal(tt.expected) {
			t.Errorf("parseDate(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestCalculateInterestingDates(t *testing.T) {
	config := Config{
		People: []Person{
			{Name: "Parent", Birthdate: "1970-01-01"},
			{Name: "Child", Birthdate: "2000-01-01"},
		},
		Events: []Event{
			{Name: "Test Event", Date: "2024-01-01"},
		},
		Marriages: []Marriage{
			{Partner1: "Alice", Partner2: "Bob", Date: "2020-06-15"},
		},
	}

	dates := calculateInterestingDates(config)

	// We should get some interesting dates
	if len(dates) == 0 {
		t.Error("Expected some interesting dates, got none")
	}

	// Check that we have various types of dates
	hasHalfAge := false
	hasPowerOf2 := false
	hasMarriageAnniversary := false
	has10kDays := false

	for _, d := range dates {
		if strings.Contains(d.Description, "half the age") {
			hasHalfAge = true
		}
		if strings.Contains(d.Description, "2^") {
			hasPowerOf2 = true
		}
		if strings.Contains(d.Description, "marriage anniversary") {
			hasMarriageAnniversary = true
		}
		if strings.Contains(d.Description, "10,000 days") {
			has10kDays = true
		}
	}

	if !hasHalfAge {
		t.Error("Expected half-age calculation")
	}
	if !hasPowerOf2 {
		t.Error("Expected power of 2 calculation")
	}
	if !hasMarriageAnniversary {
		t.Error("Expected marriage anniversary")
	}
	if !has10kDays {
		t.Error("Expected 10,000 days birthday")
	}
}

func TestLoadConfig(t *testing.T) {
	// Test that loading a non-existent file returns an error
	_, err := loadConfig("nonexistent.yaml")
	if err == nil {
		t.Error("Expected error loading non-existent file")
	}
}

func TestRelationshipType(t *testing.T) {
	config := Config{
		Marriages: []Marriage{
			{Partner1: "Alice", Partner2: "Bob", Date: "2020-01-01", RelationshipType: "couple"},
			{Partner1: "Charlie", Partner2: "Diana", Date: "2020-01-01"}, // No relationship type
		},
	}

	dates := calculateInterestingDates(config)

	// Check that we have couple anniversary
	hasCoupleAnniversary := false
	hasMarriageAnniversary := false

	for _, d := range dates {
		if strings.Contains(d.Description, "couple anniversary") {
			hasCoupleAnniversary = true
		}
		if strings.Contains(d.Description, "marriage anniversary") {
			hasMarriageAnniversary = true
		}
	}

	if !hasCoupleAnniversary {
		t.Error("Expected couple anniversary")
	}
	if !hasMarriageAnniversary {
		t.Error("Expected marriage anniversary")
	}
}

func TestBronzeAnniversary(t *testing.T) {
	// Test bronze anniversary (12.5 years = 4563 days)
	config := Config{
		Marriages: []Marriage{
			{Partner1: "Alice", Partner2: "Bob", Date: "2013-06-01"},
		},
	}

	dates := calculateInterestingDates(config)

	hasBronze := false
	for _, d := range dates {
		if strings.Contains(d.Description, "Bronze anniversary") {
			hasBronze = true
			break
		}
	}

	if !hasBronze {
		t.Error("Expected Bronze anniversary (12.5 years)")
	}
}

func Test30kDaysBirthday(t *testing.T) {
	config := Config{
		People: []Person{
			{Name: "Test Person", Birthdate: "1945-01-01"},
		},
	}

	dates := calculateInterestingDates(config)

	has30kDays := false
	for _, d := range dates {
		if strings.Contains(d.Description, "30,000 days birthday") {
			has30kDays = true
			break
		}
	}

	if !has30kDays {
		t.Error("Expected 30,000 days birthday")
	}
}

func TestYearlyBirthdays(t *testing.T) {
	config := Config{
		People: []Person{
			{Name: "Test Person", Birthdate: "2000-01-01"},
		},
	}

	dates := calculateInterestingDates(config)

	hasYearlyBirthday := false
	for _, d := range dates {
		if strings.Contains(d.Description, "year birthday") {
			hasYearlyBirthday = true
			break
		}
	}

	if !hasYearlyBirthday {
		t.Error("Expected yearly birthdays")
	}
}

func TestYearlyEventAnniversaries(t *testing.T) {
	config := Config{
		Events: []Event{
			{Name: "Test Event", Date: "2024-01-01"},
		},
	}

	dates := calculateInterestingDates(config)

	hasYearlyAnniversary := false
	for _, d := range dates {
		if strings.Contains(d.Description, "year anniversary") {
			hasYearlyAnniversary = true
			break
		}
	}

	if !hasYearlyAnniversary {
		t.Error("Expected yearly event anniversaries")
	}
}

func TestExportToIcal(t *testing.T) {
	dates := []InterestingDate{
		{
			Description: "Test Event",
			Date:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			DaysFromNow: 100,
		},
	}

	tmpFile := os.TempDir() + "/test-calendar.ics"
	err := exportToIcal(dates, tmpFile)
	if err != nil {
		t.Errorf("exportToIcal failed: %v", err)
	}

	// Check that file was created
	if _, err := os.Stat(tmpFile); err != nil {
		t.Errorf("iCal file was not created: %v", err)
	}

	// Clean up
	os.Remove(tmpFile)
}

func TestExportToHTML(t *testing.T) {
	dates := []InterestingDate{
		{
			Description: "Test Event 1",
			Date:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			DaysFromNow: 100,
		},
		{
			Description: "Test Event 2",
			Date:        time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC),
			DaysFromNow: -50,
		},
	}

	tmpFile := os.TempDir() + "/test-timeline.html"
	err := exportToHTML(dates, tmpFile)
	if err != nil {
		t.Errorf("exportToHTML failed: %v", err)
	}

	// Check that file was created
	if _, err := os.Stat(tmpFile); err != nil {
		t.Errorf("HTML file was not created: %v", err)
	}

	// Read file and check for expected content
	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Errorf("Failed to read HTML file: %v", err)
	}

	htmlContent := string(content)
	
	// Check for essential HTML elements
	if !strings.Contains(htmlContent, "<!DOCTYPE html>") {
		t.Error("HTML file missing DOCTYPE declaration")
	}
	if !strings.Contains(htmlContent, "Anniversaries Timeline") {
		t.Error("HTML file missing title")
	}
	if !strings.Contains(htmlContent, "Test Event 1") {
		t.Error("HTML file missing first event")
	}
	if !strings.Contains(htmlContent, "Test Event 2") {
		t.Error("HTML file missing second event")
	}
	if !strings.Contains(htmlContent, "timeline-item") {
		t.Error("HTML file missing timeline-item class")
	}
	if !strings.Contains(htmlContent, "togglePin") {
		t.Error("HTML file missing pin functionality")
	}

	// Clean up
	os.Remove(tmpFile)
}
