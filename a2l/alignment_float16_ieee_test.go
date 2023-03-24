package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAlignmentFloat16Ieee
// - valid alignmentFloat16Ieee
// - empty alignmentFloat16Ieee
// - invalid alignmentFloat16Ieee
// - unexpected token

func TestParseAlignmentFloat16Ieee_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid alignmentFloat16Ieee
	tokenList = []string{emptyToken, "4"}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat16Ieee(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAlignmentFloat16Ieee_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty alignmentFloat16Ieee
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat16Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentFloat16Ieee_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentFloat16Ieee
	tokenList = []string{emptyToken, beginA2mlToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat16Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentFloat16Ieee_Invalid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentFloat16Ieee
	tokenList = []string{emptyToken, "9999999999999"}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat16Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
