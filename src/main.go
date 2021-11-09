package main

import (
	"asap2Go/calibrationReader/a2l"
	"asap2Go/calibrationReader/ihex32"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var err error
	var a a2l.A2L
	var h ihex32.Hex
	var logFile *os.File

	logFile, err = configureLogger()
	if err != nil {
		log.Err(err).Msg("could not create logger:")
	}
	a, err = a2l.ParseFromFile("/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171.a2l")
	if err != nil {
		log.Err(err).Msg("could not parse a2l:")
	}
	log.Info().Msg("parsed a2l file")
	h, err = ihex32.ParseFromFile("/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171.hex")
	if err != nil {
		log.Err(err).Msg("could not parse hex:")
	}
	log.Info().Msg("parsed hex file")
	c, exists := a.Project.Module[0].Characteristics["ASAM.C.SCALAR.UBYTE.IDENTICAL"]
	if exists {
		log.Info().Str("Long Identifier of Characteristic ASAM.C.SCALAR.UBYTE.IDENTICAL", c.LongIdentifier)
		log.Info().Int("Number of Bytes in Hex", len(h.DataBytes))
	}
	os.Remove(logFile.Name())
}

//configureLogger adds a file logger with a temporary file and some formatting
func configureLogger() (*os.File, error) {
	var err error
	var dir string
	var file *os.File
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	dir, err = os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("could not find current working directory")
		return file, err
	}
	file, err = os.CreateTemp(dir, "calibReader.*.log")
	if err != nil {
		log.Error().Err(err).Msg("could not create calibration reader log-file")
		return file, err
	}
	fileWriter := zerolog.New(file).With().Logger()
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	consoleWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	multi := zerolog.MultiLevelWriter(fileWriter, consoleWriter)
	log.Logger = zerolog.New(multi).With().Timestamp().Caller().Logger()
	return file, err
}
