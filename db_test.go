package main

import "testing"

func TestParseDatabaseURL(t *testing.T) {
	testURL := "postgres://postgres@127.0.0.1:5432/test_db"
	parsedURL := parseDatabaseURL(testURL)
	expectedParsedURL := "host=127.0.0.1 port=5432 user=postgres dbname=test_db sslmode=disable"
	if parsedURL != expectedParsedURL {
		t.Errorf("Database URL was incorrect. Expected %v but got %v", expectedParsedURL, parsedURL)
	}
}
