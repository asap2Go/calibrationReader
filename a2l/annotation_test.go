package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAnnotation
// - valid annotation
// - empty annotation
// - unexpected token

func TestParseAnnotation_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid annotation
	tokenList = []string{beginAnnotationToken, annotationLabelToken, "Test Label", annotationOriginToken, "Test Origin", beginAnnotationTextToken, "Test", "Text", endAnnotationTextToken, endAnnotationToken}
	tok := newTokenGenerator()
	_, err := parseAnnotation(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAnnotation_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty annotation
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAnnotation(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAnnotation_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid annotation
	tokenList = []string{emptyToken, beginAxisDescrToken}
	tok := newTokenGenerator()
	_, err := parseAnnotation(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
