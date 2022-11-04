package a2l

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
	"unicode"

	"github.com/rs/zerolog/log"
)

var (
	//useMultithreading generally enables multithreaded parsing of a2l-file.
	//will be automatically deactivated if the tokenizer recognizes that multiple modules are present.
	//deactivate for debugging.
	useMultithreading = true
	//numProc is used to set the amount of goroutines in case useMultithreading is true.
	//numProc = runtime.NumCPU() * 2 has proven to be reliably fast for different cpu models.
	//factors above 4 will generally lead to severe performance degredation due to channel and locking overhead.
	numProc = runtime.NumCPU()
)

// A2L is the main struct returned by the a2l package.
// it contains all datatypes parsed from the .a2l file
type A2L struct {
	Asap2Version asap2Version
	A2mlVersion  a2mlVersion
	Project      Project
}

// toHashMap recursively generates a HashMap of all structs and their values in the a2l struct
// potentially useful if goPy is used to build a python library of calibrationReader
func (a *A2L) toHashMap() (map[string]interface{}, error) {
	var m map[string]interface{}
	bytes, err := json.Marshal(a)
	if err != nil {
		log.Err(err).Msg("could not serialize a2l struct into json")
		return m, err
	}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		log.Err(err).Msg("could not deserialize json into hashmap")
		return m, err
	}
	return m, err
}

// ParseFromFile is the main exported function to be called from a2l package.
// it takes an .a2l file and parses it
func ParseFromFile(filepath string) (A2L, error) {
	var err error
	var text string
	var tg tokenGenerator
	var a A2L

	startTime := time.Now()
	//read a2l file from a text file
	text, err = readFileToString(filepath)
	if err != nil {
		log.Err(err).Msg("a2l file could not be read:")
		return a, err
	}
	//divide a2l file into individual tokens
	tg, err = buildTokenGeneratorFromString(text)
	if err != nil {
		log.Err(err).Msg("could not create tokens from a2l file:")
		return a, err
	}
	//then parse the tokens created by the tokenizer
	a, err = parseA2l(&tg)
	if err != nil {
		log.Err(err).Msg("failed parsing with error:")
		return a, err
	}

	log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Info().Msg("time for parsing file: " + fmt.Sprint(elapsed.Milliseconds()))
	return a, nil
}

// parseA2l handles the parsing of the a2l struct.
// as opposed to ParseFromFile which also handles creation of the tokenizer and file reading, etc.
func parseA2l(tok *tokenGenerator) (A2L, error) {
	a2l := A2L{}
	var err error
forLoop:
	for {
		switch tok.current() {
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
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("a2l could not be parsed")
				break forLoop
			}
		}
		tok.next()
	}
	return a2l, err
}

// readFileToString opens and reads a file, then returns a string value
func readFileToString(filepath string) (string, error) {
	bytesString, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	cleanedStr := cleanBytesString(bytesString)
	return cleanedStr, nil
}

func cleanBytesString(bs []byte) string {
	//remove escaped characters
	re := regexp.MustCompile(`\\.`)
	escapedBytesString := re.ReplaceAll(bs, []byte("$1"))
	//convert byte array to string
	str := string(escapedBytesString)
	//remove unprintable chars at the start and end of the a2l file
	str = strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
	return str
}
