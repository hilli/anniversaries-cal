package main

import (
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
		if contains(d.Description, "half the age") {
			hasHalfAge = true
		}
		if contains(d.Description, "2^") {
			hasPowerOf2 = true
		}
		if contains(d.Description, "marriage anniversary") {
			hasMarriageAnniversary = true
		}
		if contains(d.Description, "10,000 days") {
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

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestLoadConfig(t *testing.T) {
	// Test that loading a non-existent file returns an error
	_, err := loadConfig("nonexistent.yaml")
	if err == nil {
		t.Error("Expected error loading non-existent file")
	}
}
