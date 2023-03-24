package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseArPrototypeOf
// - valid arPrototypeOf
// - empty arPrototypeOf
// - unexpected token

func TestParseArPrototypeOf_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid annotation
	tokenList = []string{arPrototypeOfToken, "Test Prototype of", "ApplicationSwComponentType", endArComponentToken}
	tok := newTokenGenerator()
	_, err := parseArPrototypeOf(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseArPrototypeOf_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty annotation
	tokenList = []string{arPrototypeOfToken, emptyToken, endArComponentToken}
	tok := newTokenGenerator()
	_, err := parseArPrototypeOf(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseArPrototypeOf_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid annotation
	tokenList = []string{arPrototypeOfToken, beginA2mlToken, endArComponentToken}
	tok := newTokenGenerator()
	_, err := parseArPrototypeOf(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
