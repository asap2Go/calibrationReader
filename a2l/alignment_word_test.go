package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAlignmentWord
// - valid alignmentWord
// - empty alignmentWord
// - invalid alignmentWord
// - unexpected token

func TestParseAlignmentWord_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid alignmentWord
	tokenList = []string{emptyToken, "0"}
	tok := newTokenGenerator()
	_, err := parseAlignmentWord(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAlignmentWord_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty alignmentWord
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentWord(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentWord_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentWord
	tokenList = []string{emptyToken, beginAnnotationToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentWord(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentWord_Invalid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentWord
	tokenList = []string{emptyToken, "3.876"}
	tok := newTokenGenerator()
	_, err := parseAlignmentWord(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
