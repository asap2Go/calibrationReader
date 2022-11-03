/*reads characteristics information from a2l and fills it with the data from a hex file. At least that is the plan.
Currently it can parse a2l-files as well as the corresponding IntelHex32 or Motorola S19 files. And it is quite fast at that.
At the moment a real world A2L(80MB) with its corresponding Hex File(10MB) will be parsed in less than a second.
But it still lacks the last bit of work which is implementing the methods for
axis_pts, axis_descr, record_layout and fnc_values in order to understand the memory layout and position of a specific characteristic.
This is somewhat of a convoluted mess in the a2l standard due to its historic growth and will be implemented when I have a little more spare time.
The only dependency outside the go standard library is currently zerolog.*/

package calibrationReader

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/asap2Go/calibrationReader/a2l"

	"github.com/asap2Go/calibrationReader/ihex32"

	"github.com/asap2Go/calibrationReader/srec19"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// CalibrationData contains the parsed structs from the a2l as well as the byte data from the hex file
// that are parsed by ReadCalibration()
type CalibrationData struct {
	A2l a2l.A2L
	Hex map[uint32]byte
}

// ReadCalibration takes filepaths to the a2l file and the hex file,
// parses them in parallel and returns a CalibrationData struct
func ReadCalibration(a2lFilePath string, hexFilePath string) (CalibrationData, error) {
	var err error
	var cd CalibrationData
	var errChan = make(chan error, 2)
	var a2lChan = make(chan a2l.A2L, 1)
	var hexChan = make(chan map[uint32]byte, 1)
	wgReaders := new(sync.WaitGroup)

	err = configureLogger()
	if err != nil {
		log.Err(err).Msg("could not create logger:")
		return cd, err
	}
	wgReaders.Add(2)
	go readA2L(wgReaders, a2lChan, errChan, a2lFilePath)
	go readHex(wgReaders, hexChan, errChan, hexFilePath)

	wgReaders.Wait()
	close(errChan)

	//check if any errors have occured within the readers
	var firstErr error
	if len(errChan) > 0 {
		for e := range errChan {
			if e != nil {
				firstErr = e
			}
			log.Err(e).Msg("reader encountered an error:")
		}
		return cd, firstErr
	}
	cd.A2l = <-a2lChan
	cd.Hex = <-hexChan
	return cd, nil
}

// readA2L is a helper function intended to be run in a separate go routine to call the a2l parser
// in order to be able to parse hex and a2l in parallel
func readA2L(wg *sync.WaitGroup, ca chan a2l.A2L, ce chan error, a2lFilePath string) {
	defer wg.Done()
	a, err := a2l.ParseFromFile(a2lFilePath)
	if err != nil {
		log.Err(err).Msg("could not parse a2l:")
		ce <- err
		close(ca)
	} else {
		ca <- a
		close(ca)
		log.Info().Msg("parsed a2l file")
	}
}

// readHex is a helper function intended to be run in a separate go routine to call the hex parser
// in order to be able to parse hex and a2l in parallel
func readHex(wg *sync.WaitGroup, ch chan map[uint32]byte, ce chan error, hexFilePath string) {
	defer wg.Done()
	if strings.Contains(strings.ToLower(hexFilePath), ".hex") {
		h, err := ihex32.ParseFromFile(hexFilePath)
		if err != nil {
			log.Err(err).Msg("could not parse hex:")
			ce <- err
			close(ch)
		} else {
			ch <- h
			close(ch)
			log.Info().Msg("parsed hex file")
		}
	} else if strings.Contains(strings.ToLower(hexFilePath), ".s19") {
		h, err := srec19.ParseFromFile(hexFilePath)
		if err != nil {
			log.Err(err).Msg("could not parse hex:")
			ce <- err
			close(ch)
		} else {
			ch <- h
			close(ch)
			log.Info().Msg("parsed hex file")
		}
	} else {
		err := errors.New("unsupported hex file type")
		log.Err(err).Msg("could not parse hex:")
		ce <- err
		close(ch)
	}

}

// configureLogger adds a file logger, resets previous log file and does some formatting
func configureLogger() error {
	var err error
	var file *os.File
	file, err = os.Create("calibReader.log")
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

func (cd *CalibrationData) getObjectsByIdent(ident string) []interface{} {
	var calibrationObjects []interface{}
	var buf interface{}
	var exists bool

	for _, m := range cd.A2l.Project.Modules {
		buf, exists = m.AxisPts[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Characteristics[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.CompuMethods[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.CompuTabs[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.CompuVTabs[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.CompuVTabRanges[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Functions[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Groups[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Measurements[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.RecordLayouts[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Units[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
	}
	return calibrationObjects
}
