package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAlignmentFloat64Ieee
// - valid alignmentFloat64Ieee
// - empty alignmentFloat64Ieee
// - invalid alignmentFloat64Ieee
// - unexpected token

func TestParseAlignmentFloat64Ieee_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid alignmentFloat64Ieee
	tokenList = []string{emptyToken, "8"}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat64Ieee(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAlignmentFloat64Ieee_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty alignmentFloat64Ieee
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat64Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentFloat64Ieee_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentFloat64Ieee
	tokenList = []string{emptyToken, beginA2mlToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat64Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentFlaot64Ieee_Invalid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentFloat64Ieee
	tokenList = []string{emptyToken, "-50"}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat64Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
