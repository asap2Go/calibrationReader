package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseArraySize
// - valid array size
// - empty array size
// - unexpected token
// - invalid array size

func TestParseArraySize_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid array size
	tokenList = []string{emptyToken, "1"}
	tok := newTokenGenerator()
	_, err := parseArraySize(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseArraySize_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty array size
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseArraySize(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseArraySize_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid array size
	tokenList = []string{emptyToken, beginA2mlToken}
	tok := newTokenGenerator()
	_, err := parseArraySize(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseArraySize_InvalidArraySize_1(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid array size
	tokenList = []string{emptyToken, "a"}
	tok := newTokenGenerator()
	_, err := parseArraySize(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseArraySize_InvalidArraySize_2(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid array size
	tokenList = []string{emptyToken, "-25.3"}
	tok := newTokenGenerator()
	_, err := parseArraySize(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
