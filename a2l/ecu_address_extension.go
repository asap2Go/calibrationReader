package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
This keyword is used to specify additional address information. For instance it can be
used, to distinguish different address spaces of an ECU (multi-micro controller devices).
ECU_ADDRESS_EXTENSION is an optional keyword of MEASUREMENT, AXIS_PTS
and CHARACTERISTIC.
Some calibration interfaces, such as CCP and XCP need an address extension to
access ECU data. To avoid the need for additional IF_DATA section at calibration
and measurement objects, the keyword ECU_ADDRESS_EXTENSION has been
introduced
*/
type ecuAddressExtension struct {
	extension    int16
	extensionSet bool
}

func parseECUAddressExtension(tok *tokenGenerator) (ecuAddressExtension, error) {
	eae := ecuAddressExtension{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("ecuAddressExtension could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("ecuAddressExtension could not be parsed")
	} else if !eae.extensionSet {
		var buf int64
		buf, err = strconv.ParseInt(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("ecuAddressExtension extension could not be parsed")
		}
		eae.extension = int16(buf)
		eae.extensionSet = true
		log.Info().Msg("ecuAddressExtension extension successfully parsed")
	}
	return eae, err
}
