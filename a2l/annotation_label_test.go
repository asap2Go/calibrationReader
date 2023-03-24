package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAnnotationLabel
// - valid annotationLabel
// - empty annotationLabel
// - unexpected token

func TestParseAnnotationLabel_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid annotationLabel
	tokenList = []string{emptyToken, "TEST"}
	tok := newTokenGenerator()
	_, err := parseAnnotationLabel(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAnnotationLabel_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty annotationLabel
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAnnotationLabel(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAnnotationLabel_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid annotationLabel
	tokenList = []string{emptyToken, beginA2mlToken}
	tok := newTokenGenerator()
	_, err := parseAnnotationLabel(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
