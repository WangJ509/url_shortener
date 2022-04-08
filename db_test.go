package main

import (
	"testing"
	"time"
)

func TestCreateShortenURL(t *testing.T) {
	url := "test-shorten-url"
	originalURL := "www.google.com"
	expireTime := time.Now().Add(3 * time.Hour)
	err := CreateShortenURL(url, originalURL, expireTime)
	if err != nil {
		t.Fatal(err)
	}

	record, err := GetOriginalURL(url)
	if err != nil {
		t.Fatal(err)
	}

	if record != originalURL {
		t.Fatal("record inconsistent with original URL")
	}
}