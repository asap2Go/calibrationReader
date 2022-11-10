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
	//A2L defines the Metadata for a given ECU-Project and its corresponding hex file
	A2l a2l.A2L
	//ModuleIndex defines which module within the a2l file is being used. Default is 0
	ModuleIndex uint8
	//Hex contains the flashable data for the ecu.
	//it is being simplified as a map that can be accessed by an address
	//represented by an integer and returns a byte as value.
	Hex map[uint32]byte
}

// ReadCalibration takes filepaths to the a2l file and the hex file,
// parses them in parallel and returns a CalibrationData struct
func ReadCalibration(a2lFilePath string, hexFilePath string) (CalibrationData, error) {
	var err error
	var cd CalibrationData
	//set ModuleIndex to zero as default as it covers 99% of use cases.
	cd.ModuleIndex = 0

	//set up channels for concurrent parsing of a2l and hex file as well as for the communication of potential parsing errors.
	var errChan = make(chan error, 2)
	var a2lChan = make(chan a2l.A2L, 1)
	var hexChan = make(chan map[uint32]byte, 1)

	//wait group to determine when both parsers have finished
	wgReaders := new(sync.WaitGroup)

	//initialize logging
	err = configureLogger()
	if err != nil {
		log.Err(err).Msg("could not create logger:")
		return cd, err
	}
	//Log Level per Default is warning as info and debug lead to excessive log files
	//and should only be used for debugging.
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	//start concurrent parsers as individual go routines
	wgReaders.Add(2)
	go readA2L(wgReaders, a2lChan, errChan, a2lFilePath)
	go readHex(wgReaders, hexChan, errChan, hexFilePath)
	//and wait until they're done
	wgReaders.Wait()
	//error channel is closed here while the other two channels are closed in the goroutines that use them exclusively.
	//this is safe because the a2l and hex channels are closed after a potential error message has been sent via the error channel.
	//meaning the waitgroup wgReaders blocks until all sending operations on the error channel are over.
	close(errChan)

	//check if any errors have occured within the readers
	//only the first error is returned

	var firstErr error
	if len(errChan) > 0 {
		for e := range errChan {
			if e != nil {
				firstErr = e
			}
			log.Err(e).Msg("parser encountered an error:")
		}
		return cd, firstErr
	}
	//in case no errors occured then read from the closed channels
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
		ce <- err //send an error via channel to signal it to the main thread
		close(ca)
	} else {
		ca <- a
		close(ca) //send the successfully parsed a2l structure to the main thread
		log.Info().Msg("parsed a2l file")
	}
}

// readHex is a helper function intended to be run in a separate go routine to call the hex parser
// in order to be able to parse hex and a2l in parallel
func readHex(wg *sync.WaitGroup, ch chan map[uint32]byte, ce chan error, hexFilePath string) {
	defer wg.Done()

	//check whether the hex or the s19 parser needs to be used.
	//probably improve this by putting the determination logic into a single unified hex parsing package

	//intel hex
	if strings.Contains(strings.ToLower(hexFilePath), ".hex") {
		h, err := ihex32.ParseFromFile(hexFilePath)
		if err != nil {
			log.Err(err).Msg("could not parse hex:")
			ce <- err //send an error via channel to signal it to the main thread
			close(ch)
		} else {
			ch <- h //send the successfully parsed hex map to the main thread
			close(ch)
			log.Info().Msg("parsed hex file")
		}
		//Motorola S19
	} else if strings.Contains(strings.ToLower(hexFilePath), ".s19") {
		h, err := srec19.ParseFromFile(hexFilePath)
		if err != nil {
			log.Err(err).Msg("could not parse hex:")
			ce <- err //send an error via channel to signal it to the main thread
			close(ch)
		} else {
			ch <- h //send the successfully parsed hex map to the main thread
			close(ch)
			log.Info().Msg("parsed hex file")
		}
	} else {
		err := errors.New("unsupported hex file type")
		log.Err(err).Msg("could not parse hex:")
		ce <- err //send an error via channel to signal it to the main thread
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
