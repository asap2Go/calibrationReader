package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAlignmentByte
// - valid alignmentByte
// - empty alignmentByte
// - invalid alignmentByte
// - unexpected token

func TestParseAlignmentByte_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid alignmentByte
	tokenList = []string{emptyToken, "4"}
	tok := newTokenGenerator()
	_, err := parseAlignmentByte(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAlignmentByte_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty alignmentByte
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentByte(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentByte_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentByte
	tokenList = []string{emptyToken, beginA2mlToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentByte(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentByte_Invalid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentByte
	tokenList = []string{emptyToken, "9999999999999"}
	tok := newTokenGenerator()
	_, err := parseAlignmentByte(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
