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
	"strconv"
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
	file, err := os.Create("calibReader.log")
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

// getRecordLayout tries to retrieve the record layout for a specified characteristic
func (cd *CalibrationData) getRecordLayout(c *a2l.Characteristic) (*a2l.RecordLayout, error) {
	if !c.DepositSet {
		err := errors.New("no deposit set in characteristic " + c.Name)
		log.Err(err).Msg("record layout not found")
		return nil, err
	}
	module := cd.A2l.Project.Modules[cd.ModuleIndex]
	rl, exists := module.RecordLayouts[c.Deposit]
	if !exists {
		err := errors.New("no record layout found for deposit identifier" + c.Deposit + " of characteristic " + c.Name)
		log.Err(err).Msg("record layout not found")
		return nil, err
	}
	return &rl, nil
}

// getNextAlignedAddress takes an adress of a record layout field and its datatype as well as a reference to the record layout itself.
// it computes whether the address given is aligned as defined by the alignemnts within the record layout for the given datatype.
// if the record layout does not provide an alignement the alignment from MOD_COMMON is used.
// if MOD_COMMON does not provide an alignment, then the default values (magic numbers below) are used.
func (cd *CalibrationData) getNextAlignedAddress(address uint32, dte a2l.DataTypeEnum, rl *a2l.RecordLayout) uint32 {
	switch dte {
	case a2l.UBYTE:
		if rl.AlignmentByte.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentByte.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentByte
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(1)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.SBYTE:
		if rl.AlignmentByte.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentByte.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentByte
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(1)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.UWORD:
		if rl.AlignmentWord.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentWord.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentWord
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(2)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.SWORD:
		if rl.AlignmentWord.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentWord.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentWord
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(2)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.ULONG:
		if rl.AlignmentLong.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentLong.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentLong
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(4)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.SLONG:
		if rl.AlignmentLong.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentLong.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentLong
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(4)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.AUint64:
		if rl.AlignmentInt64.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentInt64.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentInt64
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(8)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.AInt64:
		if rl.AlignmentInt64.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentInt64.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentInt64
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(8)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.Float16Ieee:
		if rl.AlignmentFloat16Ieee.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentFloat16Ieee.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentFloat16Ieee
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(2)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.Float32Ieee:
		if rl.AlignmentFloat32Ieee.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentFloat32Ieee.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentFloat32Ieee
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(4)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	case a2l.Float64Ieee:
		if rl.AlignmentFloat64Ieee.AlignmentBorderSet {
			modulo := address % uint32(rl.AlignmentFloat64Ieee.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
		a := cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon.AlignmentFloat64Ieee
		if a.AlignmentBorderSet {
			modulo := address % uint32(a.AlignmentBorder)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		} else {
			modulo := address % uint32(8)
			if modulo != 0 {
				return address + modulo
			} else {
				return address
			}
		}
	default:
		return address
	}
}

// getBytes gets a number of bites (length) from a given address and returns a byte slice
// if the address cannot be found an error is returned
func (cd *CalibrationData) getBytes(address uint32, length uint32) ([]byte, error) {
	nBytes := length / 8
	bytes := make([]byte, nBytes, 0)
	if length%8 != 0 {
		err := errors.New("unexpected number of bits " + strconv.Itoa(int(length)) + ". Expected a multiple of 8.")
		log.Error().Err(err).Msg("invalid length")
		return bytes, err
	}
	for i := 0; i < int(nBytes); i++ {
		b, exists := cd.Hex[address+uint32(i)]
		if !exists {
			err := errors.New("address " + strconv.Itoa(int(address+uint32(i))) + " not found in hex")
			log.Error().Err(err).Msg("invalid address")
			return bytes, err
		}
		bytes = append(bytes, b)
	}
	return bytes, nil
}

// GetAllValuesFromHex Reads all values for each characteristic and its corresponding record layout from the hex file and converts it to decimal, and physical values.
func (cd *CalibrationData) GetAllValuesFromHex() error {
	//for each characteristic get its record layout:
	for _, c := range cd.A2l.Project.Modules[cd.ModuleIndex].Characteristics {
		var cv CharacteristicValues
		cv.characteristic = &c
		cd.getValuesFromHex(&cv)
	}
	return nil
}

// getValuesFromHex Reads all values for ONE specific characteristic and its corresponding record layout from the hex file and converts it to decimal, and physical values.
func (cd *CalibrationData) getValuesFromHex(cv *CharacteristicValues) {
	rl, err := cd.getRecordLayout(cv.characteristic)
	if err != nil {
		log.Err(err).Msg("record layout for identifier '" + cv.characteristic.Deposit + "' not found")
	}

	//determine relative positions of each field of the record layout struct
	relPos, err := rl.GetRecordLayoutRelativePositions()
	if err != nil {
		log.Err(err).Msg("could not retrieve positions for record layout '" + cv.characteristic.Deposit + "'")
	}
	rl.RelativePositions = relPos

	//determine the start address of the characteristic
	curPos, err := cd.convertStringToUint32Address(cv.characteristic.Address)
	if err != nil {
		log.Err(err).Msg("could not convert address of characteristic '" + cv.characteristic.Name + "'")
	}

	for _, field := range cv.recordLayout.RelativePositions {
		switch field {
		case "AxisPtsX":
			cv.AxisX, err = cd.getAxisPointsX(rl, &curPos)
			if err != nil {
				log.Err(err).Msg("could not get values for X-Axis of characteristic '" + cv.characteristic.Name + "'")
			}
		case "AxisPtsY":
			cv.AxisY, err = cd.getAxisPointsY(rl, &curPos)
			if err != nil {
				log.Err(err).Msg("could not get values for Y-Axis of characteristic '" + cv.characteristic.Name + "'")
			}
		case "AxisPtsZ":
			cv.AxisZ, err = cd.getAxisPointsZ(rl, &curPos)
			if err != nil {
				log.Err(err).Msg("could not get values for Z-Axis of characteristic '" + cv.characteristic.Name + "'")
			}
		case "AxisPts4":
			cv.Axis4, err = cd.getAxisPoints4(rl, &curPos)
			if err != nil {
				log.Err(err).Msg("could not get values for 4-Axis of characteristic '" + cv.characteristic.Name + "'")
			}
		case "AxisPts5":
			cv.Axis5, err = cd.getAxisPoints5(rl, &curPos)
			if err != nil {
				log.Err(err).Msg("could not get values for 5-Axis of characteristic '" + cv.characteristic.Name + "'")
			}
		case "AxisRescaleX":
			//continue
		case "DistOpX":
		case "DistOpY":
		case "DistOpZ":
		case "DistOp4":
		case "DistOp5":
		case "FixNoAxisPtsX":
		case "FixNoAxisPtsY":
		case "FixNoAxisPtsZ":
		case "FixNoAxisPts4":
		case "FixNoAxisPts5":
		case "FncValues":
			//interesting part
		case "Identification":
		case "NoAxisPtsX":
		case "NoAxisPtsY":
		case "NoAxisPtsZ":
		case "NoAxisPts4":
		case "NoAxisPts5":
		case "NoRescaleX":
		case "OffsetX":
		case "OffsetY":
		case "OffsetZ":
		case "Offset4":
		case "Offset5":
		case "Reserved":
		case "RipAddrW":
		case "RipAddrX":
		case "RipAddrY":
		case "RipAddrZ":
		case "RipAddr4":
		case "RipAddr5":
		case "SrcAddrX":
		case "SrcAddrY":
		case "SrcAddrZ":
		case "SrcAddr4":
		case "SrcAddr5":
		case "ShiftOpX":
		case "ShiftOpY":
		case "ShiftOpZ":
		case "ShiftOp4":
		case "ShiftOp5":
		}
	}
}

func (cd *CalibrationData) getValue(curPos *uint32, dte a2l.DataTypeEnum, rl *a2l.RecordLayout) (interface{}, error) {
	buf, err := cd.getBytes(cd.getNextAlignedAddress(*curPos, dte, rl), uint32(dte.GetDatatypeLength()))
	if err != nil {
		log.Err(err).Msg("could not retrieve value as byteSlice")
		return nil, err
	}
	data, err := cd.convertByteSliceToDatatype(buf, rl.AxisPts5.Datatype)
	if err != nil {
		log.Err(err).Msg("could not convert byteSlice to dataType")
		return nil, err
	}
	return data, nil
}
