package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseArComponent
// - valid arComponent
// - empty arComponent
// - unexpected token

func TestParseArComponent_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid annotation
	tokenList = []string{beginArComponentToken, arPrototypeOfToken, "Test Prototype of", "ApplicationSwComponentType", endArComponentToken}
	tok := newTokenGenerator()
	_, err := parseArComponent(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseArComponent_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty annotation
	tokenList = []string{beginArComponentToken, emptyToken, endArComponentToken}
	tok := newTokenGenerator()
	_, err := parseArComponent(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseArComponent_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid annotation
	tokenList = []string{beginArComponentToken, beginA2mlToken, endArComponentToken}
	tok := newTokenGenerator()
	_, err := parseArComponent(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
