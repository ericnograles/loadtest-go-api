package main

import "testing"

func TestParseLocalDatabaseURL(t *testing.T) {
	testURL := "postgres://postgres@127.0.0.1:5432/test_db"
	parsedURL := parseDatabaseURL(testURL)
	expectedParsedURL := "host=127.0.0.1 port=5432 user=postgres dbname=test_db sslmode=disable"
	if parsedURL != expectedParsedURL {
		t.Errorf("Database URL was incorrect. Expected %v but got %v", expectedParsedURL, parsedURL)
	}
}

func TestParseRemoteDatabaseURL(t *testing.T) {
	testURL := "postgres://postgres:password@some-ec2-server.aws.com:5432/test_db"
	parsedURL := parseDatabaseURL(testURL)
	expectedParsedURL := "host=some-ec2-server.aws.com port=5432 user=postgres password=password dbname=test_db sslmode=require"
	if parsedURL != expectedParsedURL {
		t.Errorf("Database URL was incorrect. Expected %v but got %v", expectedParsedURL, parsedURL)
	}
}
