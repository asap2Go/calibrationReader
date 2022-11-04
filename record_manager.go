package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
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

func (cd *CalibrationData) getCharacteristicValueBinary(c a2l.Characteristic) (interface{}, error) {
	var err error
	rl, err := cd.getRecordLayout(c)
	if err != nil {
		return nil, err
	}
	log.Debug().Msg("record layout " + rl.Name + "found")
	return err, nil
}

func (cd *CalibrationData) getCharacteristicValueDecimal() (interface{}, error) {
	return nil, nil
}

func (cd *CalibrationData) getCharacteristicValueDisplay() (interface{}, error) {
	return nil, nil
}

func (cd *CalibrationData) getCharacteristicValueHex() (interface{}, error) {
	return nil, nil
}

// get record layout for a specified characteristic
func (cd *CalibrationData) getRecordLayout(c a2l.Characteristic) (a2l.RecordLayout, error) {
	var err error
	var rl a2l.RecordLayout
	module := cd.A2l.Project.Modules[cd.ModuleIndex]
	if !c.DepositSet {
		err = errors.New("no deposit set in characteristic " + c.Name)
		log.Err(err).Msg("record layout not found")
		return rl, err
	}
	var exists bool
	rl, exists = module.RecordLayouts[c.Deposit]
	if !exists {
		err = errors.New("no record layout found for deposit identifier" + c.Deposit + " of characteristic " + c.Name)
		log.Err(err).Msg("record layout not found")
		return rl, err
	}
	return rl, nil
}
