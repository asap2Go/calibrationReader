package calibrationReader

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
	"github.com/x448/float16"
)

func (cd *CalibrationData) getSystemConstant(ident string) (a2l.SystemConstant, error) {
	modPar := cd.A2l.Project.Modules[cd.ModuleIndex].ModPar
	s, exists := modPar.SystemConstants[ident]
	if !exists {
		err := errors.New("no system constant with name " + ident)
		log.Err(err).Msg("system constant not found")
		return s, err
	}
	return s, nil
}

func (cd *CalibrationData) getSystemConstantValue(ident string) (string, error) {
	sc, err := cd.getSystemConstant(ident)
	if err != nil {
		log.Err(err).Msg("could not get value of system constant")
		return "", err
	}
	var val string
	if !sc.ValueSet {
		err = errors.New("no value defined in system constant " + sc.Name)
		log.Err(err).Msg("could not get value of system constant")
		return "", err
	}
	return val, nil
}

// GetObjectByIdent returns an object with a given identifier that is defined within the a2l
// not all datastructures are checked. Only the most relevant ones
func (cd *CalibrationData) GetObjectByIdent(ident string) []interface{} {
	var calibrationObjects []interface{}
	var buf interface{}
	var exists bool

	m := cd.A2l.Project.Modules[cd.ModuleIndex]

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
	buf, exists = m.ModPar.SystemConstants[ident]
	if exists {
		calibrationObjects = append(calibrationObjects, buf)
	}
	buf, exists = m.Units[ident]
	if exists {
		calibrationObjects = append(calibrationObjects, buf)
	}
	return calibrationObjects
}

// hexToByteSlice converts at least a four character hexString to a slice of several bytes. fails if input is too short or not valid hex.
func hexToByteSlice(hexVal string) ([]byte, error) {
	decoded, err := hex.DecodeString(hexVal)
	if err != nil {
		log.Err(err)
	}
	return decoded, err
}

// convertStringToUint32Address is used to convert the adresses in string format in the characteristics to a uint32
func (cd *CalibrationData) convertStringToUint32Address(str string) (uint32, error) {
	var val uint32
	byteSlice, err := hexToByteSlice(strings.ReplaceAll(str, "0x", ""))
	if err != nil {
		log.Err(err).Msg("string '" + str + "' could not be parsed")
		return val, err
	}
	modCom := &cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon
	if modCom.ByteOrder.ByteOrder == a2l.MsbFirstMswLast || modCom.ByteOrder.ByteOrder == a2l.MsbLastMswFirst {
		err = errors.New("unexpected byte order")
		log.Err(err).Msg("byte order " + string(modCom.ByteOrder.ByteOrder) + "not implemented")
		return val, err
	}
	if !modCom.ByteOrder.ByteOrderSet || modCom.ByteOrder.ByteOrder == a2l.BigEndian || modCom.ByteOrder.ByteOrder == a2l.MsbFirst {
		val = binary.BigEndian.Uint32(byteSlice)
	} else {
		val = binary.LittleEndian.Uint32(byteSlice)
	}
	return val, nil
}

// converts a byteSlice into a a2l.DatatypeEnum datatype.
// if not enough bytes are supplied the conversion fails.
// if MsbFirstMswLast or MsbLastMswFirst are used as binary encoding then the conversion fails
// as those are not implemented
func (cd *CalibrationData) convertByteSliceToDatatype(byteSlice []byte, dte a2l.DataTypeEnum) (float64, error) {
	//bounds check
	if len(byteSlice) == 0 || len(byteSlice)*8 < int(dte.GetDatatypeLength()) {
		err := errors.New("byte slice holds " + fmt.Sprintf("%d", len(byteSlice)) + " bytes. " +
			strconv.Itoa(int(dte.GetDatatypeLength()/8)) + " bytes necessary to convert to datatype " + dte.String())
		log.Err(err).Msg("conversion failed")
		return 0.0, err
	}
	//check which byteorder is used
	modCom := &cd.A2l.Project.Modules[cd.ModuleIndex].ModCommon
	if modCom.ByteOrder.ByteOrder == a2l.MsbFirstMswLast || modCom.ByteOrder.ByteOrder == a2l.MsbLastMswFirst {
		err := errors.New("unexpected byte order")
		log.Err(err).Msg("byte order " + string(modCom.ByteOrder.ByteOrder) + "not implemented")
		return 0.0, err
	}
	if !modCom.ByteOrder.ByteOrderSet || modCom.ByteOrder.ByteOrder == a2l.BigEndian || modCom.ByteOrder.ByteOrder == a2l.MsbFirst {
		switch dte {
		case a2l.UBYTE:
			return float64(byteSlice[0]), nil
		case a2l.SBYTE:
			return float64(int8(byteSlice[0])), nil
		case a2l.UWORD:
			val := binary.BigEndian.Uint32(byteSlice)
			return float64(val), nil
		case a2l.SWORD:
			val := int32(binary.BigEndian.Uint32(byteSlice))
			return float64(val), nil
		case a2l.ULONG:
			val := binary.BigEndian.Uint64(byteSlice)
			return float64(val), nil
		case a2l.SLONG:
			val := int64(binary.BigEndian.Uint64(byteSlice))
			return float64(val), nil
		case a2l.AUint64:
			val := binary.BigEndian.Uint64(byteSlice)
			return float64(val), nil
		case a2l.AInt64:
			val := int64(binary.BigEndian.Uint64(byteSlice))
			return float64(val), nil
		case a2l.Float16Ieee:
			val := float16.Frombits(binary.BigEndian.Uint16(byteSlice))
			return float64(val), nil
		case a2l.Float32Ieee:
			val := math.Float32frombits(binary.BigEndian.Uint32(byteSlice))
			return float64(val), nil
		case a2l.Float64Ieee:
			val := math.Float64frombits(binary.BigEndian.Uint64(byteSlice))
			return float64(val), nil
		default:
			err := errors.New("unexpected datatype")
			log.Err(err).Msg("datatype " + dte.String() + " not implemented")
			return 0.0, err
		}
	} else {
		switch dte {
		case a2l.UBYTE:
			return float64(byteSlice[0]), nil
		case a2l.SBYTE:
			return float64(int8(byteSlice[0])), nil
		case a2l.UWORD:
			val := binary.LittleEndian.Uint32(byteSlice)
			return float64(val), nil
		case a2l.SWORD:
			val := int32(binary.LittleEndian.Uint32(byteSlice))
			return float64(val), nil
		case a2l.ULONG:
			val := binary.LittleEndian.Uint64(byteSlice)
			return float64(val), nil
		case a2l.SLONG:
			val := int64(binary.LittleEndian.Uint64(byteSlice))
			return float64(val), nil
		case a2l.AUint64:
			val := binary.LittleEndian.Uint64(byteSlice)
			return float64(val), nil
		case a2l.AInt64:
			val := int64(binary.LittleEndian.Uint64(byteSlice))
			return float64(val), nil
		case a2l.Float16Ieee:
			val := float16.Frombits(binary.LittleEndian.Uint16(byteSlice))
			return float64(val), nil
		case a2l.Float32Ieee:
			val := math.Float32frombits(binary.LittleEndian.Uint32(byteSlice))
			return float64(val), nil
		case a2l.Float64Ieee:
			val := math.Float64frombits(binary.LittleEndian.Uint64(byteSlice))
			return float64(val), nil
		default:
			err := errors.New("unexpected datatype")
			log.Err(err).Msg("datatype " + dte.String() + " not implemented")
			return 0.0, err
		}
	}
}
