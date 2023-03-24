package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAnnotationText
// - valid annotationText
// - empty annotationText
// - unexpected token

func TestParseAnnotationText_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid annotationText
	tokenList = []string{emptyToken, "TEST", endAnnotationTextToken}
	tok := newTokenGenerator()
	_, err := parseAnnotationText(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAnnotationText_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty annotationText
	tokenList = []string{emptyToken, emptyToken, endAnnotationTextToken}
	tok := newTokenGenerator()
	_, err := parseAnnotationText(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAnnotationText_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid annotationText
	tokenList = []string{emptyToken, beginAxisDescrToken, endAnnotationTextToken}
	tok := newTokenGenerator()
	_, err := parseAnnotationText(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
