package a2l

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	//useMultithreading enables multithreaded parsing of a2l-file. Will be deactivated if multiple modules are recognized
	useMultithreading = true
	//numProc is used to set the amount of goroutines in case useMultithreading is true.
	numProc = runtime.NumCPU() * 2
)

type A2L struct {
	Asap2Version asap2Version
	A2mlVersion  a2mlVersion
	Project      Project
}

func ParseFromFile(filepath string) (A2L, error) {
	var logFile *os.File
	var err error
	var f string
	var a A2L
	startTime := time.Now()
	logFile, err = createLogger()
	if err != nil {
		log.Err(err).Msg("a2l log-file could not be created:")
		return a, err
	}
	f, err = readFileToString(filepath)
	if err != nil {
		log.Err(err).Msg("a2l test-file could not be read:")
		return a, err
	}
	var tg tokenGenerator
	tg, _ = buildTokenGeneratorFromString(f)
	log.Info().Msg("created tokenizer")
	//log.Info().Msg("created tok")
	a, err = parseA2l(&tg)
	if err != nil {
		log.Err(err).Msg("failed parsing with error:")
		return a, err
	}
	log.Info().Str("prject name", a.Project.name).Msg("finished parsing:")
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Info().Msg("time for parsing file: " + fmt.Sprint(elapsed.Milliseconds()))
	//in case there are no error delete log file
	os.Remove(logFile.Name())
	return a, nil
}

func parseA2l(tok *tokenGenerator) (A2L, error) {
	a2l := A2L{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case asap2VersionToken:
			a2l.Asap2Version, err = parseASAP2Version(tok)
			if err != nil {
				log.Err(err).Msg("a2l asap2Version could not be parsed")
				break forLoop
			} else {
				log.Info().Msg("a2l asap2Version successfully parsed")
			}
		case a2mlVersionToken:
			a2l.A2mlVersion, err = parseA2MLVersion(tok)
			if err != nil {
				log.Err(err).Msg("a2l a2mlVersion could not be parsed")
				break forLoop
			}
			log.Info().Msg("a2l a2mlVersion successfully parsed")
		case beginProjectToken:
			a2l.Project, err = parseProject(tok)
			if err != nil {
				log.Err(err).Msg("a2l project could not be parsed")
				break forLoop
			}
			log.Info().Msg("a2l project successfully parsed")
		default:
			if tok.current() == emptyToken {
				log.Info().Msg("a2l parsed")
				break forLoop
			}
		}
	}
	return a2l, err
}

func readFileToString(filepath string) (string, error) {
	bytesString, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	text := string(bytesString)
	return text, nil
}

func createLogger() (*os.File, error) {
	var err error
	var dir string
	var file *os.File
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	dir, err = os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("could not find current working directory")
		return file, err
	}
	file, err = os.CreateTemp(dir, "a2l.*.log")
	if err != nil {
		log.Error().Err(err).Msg("could not create a2l log-file")
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
