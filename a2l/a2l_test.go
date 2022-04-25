package a2l

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestParseFromFile(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	a2lPath := "testing/ASAP2_Demo_V171_allKeywords.a2l"
	startTime := time.Now()
	a, err := ParseFromFile(a2lPath)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
	log.Info().Msg("time for parsing a2l test file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
}

func FuzzParseA2L(f *testing.F) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	a2lPath := "testing/ASAP2_Demo_V171_allKeywords.a2l"
	text, _ := readFileToString(a2lPath)
	f.Add(text)

	f.Fuzz(func(t *testing.T, orig string) {
		tg, err := buildTokenGeneratorFromString(orig)
		if err != nil {
			log.Err(err).Msg("could not create tokens from a2l file")
			log.Err(err).Msg(orig)
		}
		a, err := parseA2l(&tg)
		if err != nil {
			log.Err(err).Msg("failed parsing " + a.Project.Name + " with error:")
			log.Err(err).Msg(orig)
		}
	})
}

func BenchmarkParseFromFile(b *testing.B) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a2lPath := "testing/ASAP2_Demo_V171_allKeywords.a2l"
		startTime := time.Now()
		a, err := ParseFromFile(a2lPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
		log.Warn().Msg("time for parsing a2l bench file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
	}
}

//configureLogger adds a file logger, resets previous log file and does some formatting
func configureLogger() error {
	var err error
	var file *os.File
	file, err = os.Create("a2l_test.log")
	if err != nil {
		log.Error().Err(err).Msg("could not create calibration reader log-file")
		return err
	}
	fileWriter := zerolog.ConsoleWriter{Out: file, NoColor: true, TimeFormat: time.StampMicro}
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMicro}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	fileWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	log.Logger = zerolog.New(zerolog.MultiLevelWriter(fileWriter, consoleWriter)).With().Timestamp().Caller().Logger()
	return nil
}
