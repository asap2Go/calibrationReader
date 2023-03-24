package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAlignmentInt64
// - valid alignmentInt64
// - empty alignmentInt64
// - invalid alignmentInt64
// - unexpected token

func TestParseAlignmentInt64_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid alignmentInt64
	tokenList = []string{emptyToken, "4"}
	tok := newTokenGenerator()
	_, err := parseAlignmentInt64(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAlignmentInt64_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty alignmentInt64
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentInt64(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentInt64_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentInt64
	tokenList = []string{emptyToken, beginA2mlToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentInt64(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentInt64_Invalid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentInt64
	tokenList = []string{emptyToken, "2.5"}
	tok := newTokenGenerator()
	_, err := parseAlignmentInt64(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
