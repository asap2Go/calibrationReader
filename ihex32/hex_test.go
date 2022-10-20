package ihex32

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestParseFromFile(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	hexPath := "testing/ASAP2_Demo_V171.hex"
	h, err := ParseFromFile(hexPath)
	if err == nil {
		if len(h) != 137324 {
			err = errors.New("wrong length of dataset")
			t.Fatalf("failed parsing with error: %s.", err)
		}
	} else {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func FuzzParseHex(f *testing.F) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	hexPath := "testing/ASAP2_Demo_V171.hex"
	text, _ := readFileToString(hexPath)
	var errList []error
	f.Add(text)

	f.Fuzz(func(t *testing.T, orig string) {
		//split the text into lines
		lines := strings.Split(orig, "\r\n")
		if len(lines) == 1 {
			//in case unix line terminator is used.
			lines = strings.Split(orig, "\n")
		}
		h, err := parseHex(lines)
		if err != nil {
			exists := false
			for _, e := range errList {
				if err == e {
					exists = true
					break
				}
			}
			if !exists {
				errList = append(errList, err)
				fmt.Println(len(h))
				log.Err(err).Msg("could not parse hex-file with length " + strconv.Itoa(len(h)))
				log.Err(err).Msg(orig)
			}
		}
	})
}

func BenchmarkParseFromFile(b *testing.B) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	//b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		startTime := time.Now()
		hexPath := "testing/ASAP2_Demo_V171.hex"
		h, err := ParseFromFile(hexPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
			log.Info().Msg("length of data in hex file: " + fmt.Sprint(len(h)))
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Warn().Msg("time for parsing ihex32 bench file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
	}
}

//configureLogger adds a file logger, resets previous log file and does some formatting
func configureLogger() error {
	var err error
	var file *os.File
	file, err = os.Create("ihex32_test.log")
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
