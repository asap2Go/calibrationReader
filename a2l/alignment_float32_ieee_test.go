package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAlignmentFloat32Ieee
// - valid alignmentFloat32Ieee
// - empty alignmentFloat32Ieee
// - invalid alignmentFloat32Ieee
// - unexpected token

func TestParseAlignmentFloat32Ieee_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid alignmentFloat32Ieee
	tokenList = []string{emptyToken, "4"}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat32Ieee(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAlignmentFloat32Ieee_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty alignmentFloat32Ieee
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat32Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentFloat32Ieee_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentFloat32Ieee
	tokenList = []string{emptyToken, beginA2mlToken}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat32Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAlignmentFloat32Ieee_Invalid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid alignmentFloat32Ieee
	tokenList = []string{emptyToken, "9999999999999"}
	tok := newTokenGenerator()
	_, err := parseAlignmentFloat32Ieee(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
