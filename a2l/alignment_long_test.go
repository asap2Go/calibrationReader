package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAlignmentLong
// - valid alignmentLong
// - empty alignmentLong
// - invalid alignmentLong
// - unexpected token

func TestParseAlignmentLong_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid alignmentLong
	tokenList = []string{emptyToken, "8"}
	tok := newTokenGenerator()
	_, err := parseAlignmentLong(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAlignmentLong_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty alignmentLong
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentLong(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentLong_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentLong
	tokenList = []string{emptyToken, beginAxisDescrToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentLong(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentLong_Invalid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentLong
	tokenList = []string{emptyToken, "3.9"}
	tok := newTokenGenerator()
	_, err := parseAlignmentLong(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
