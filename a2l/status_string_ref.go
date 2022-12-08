package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

/*
StatusStringRef is used to split up the value range of ECU internal values into a numerical
and a verbal part. The verbal part can be used to visualize status information (e.g.
"Sensor not connected"). It is used at COMPU_METHOD.
The conversion table referenced by the keyword STATUS_STRING_REF must not define
a default value.
Note: The MC-System at first checks whether the internal value is in the range defined
at the STATUS_STRING_REF conversion table. In this case, the tool displays
the corresponding text. Otherwise, the MC-System uses the regular computation
method.
The MC-System must not respect the limits when evaluating the
STATUS_STRING_REF.
Note: PHYS Values defined by STATUS_STRING_REF are not selectable for
calibration. To ensure this, the values referenced by STATUS_STRING_REF
must be outside the ECU internal limits of all calibration objects
(CHARACTERISTIC, AXIS_PTS) using the corresponding conversion method.
*/
type StatusStringRef struct {
	//ConversionTable is a reference to a verbal conversion table (COMPU_VTAB or COMPU_VTAB_RANGE)
	ConversionTable    string
	ConversionTableSet bool
}

func parseStatusStringRef(tok *tokenGenerator) (StatusStringRef, error) {
	ssr := StatusStringRef{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("statusStringRef could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("statusStringRef could not be parsed")
	} else if !ssr.ConversionTableSet {
		ssr.ConversionTable = tok.current()
		ssr.ConversionTableSet = true
		log.Info().Msg("statusStringRef conversionTable successfully parsed")
	}
	return ssr, err
}
