package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

//check access type. DIRECT is the most used. Just read value from a given address.
//in case other access types are set this gets more complicated as either offsets or pointers are leveraged to
//define the position of the calibration objects.
//for VALUE Type: just read one value at curPos
//for higher level objects: read the number of elements defined by the matrix dim or NoAxisPts fields with the direction (row, column, alternate, ...) specified.
//for applicable objects check whether STATIC_RECORD_LAYOUT and STATIC_ ADDRESS_OFFSET are set
//to determine how to read the FNC_Values correctly, this can lead to hard to detect errors when not implemented.

func (cv *CharacteristicValues) getFncValues(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	val, err := cv.getValuesByCharacteristicType(cd, rl, curPos)
	if err != nil {
		log.Err(err).Msg("could not retrieve fncValues value")
		return nil, err
	}

	return val, err
}

func (cv *CharacteristicValues) getValuesByCharacteristicType(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	if !cv.characteristic.TypeSet {
		err := errors.New("characteristic type not set")
		log.Err(err).Msg("could not determine type of characteristic " + cv.characteristic.Name)
		return nil, err
	}
	if !rl.FncValues.DatatypeSet {
		err := errors.New("fncValues datatype not set")
		log.Err(err).Msg("could not determine datatype of FncValues of characteristic " + cv.characteristic.Name)
		return nil, err
	}
	switch cv.characteristic.Type {
	case a2l.ASCII:
	case a2l.Curve:
	case a2l.Map:
	case a2l.Cuboid:
	case a2l.Cube4:
	case a2l.Cube5:
	case a2l.ValBlk:
	case a2l.Value:
		//ToDo:
	case a2l.Derived:
	case a2l.ExtendedSi:
	default:
		err := errors.New("characteristic type not defined")
		log.Err(err).Msg("could not determine type of characteristic " + cv.characteristic.Name)
		return nil, err
	}

	return nil, nil
}
