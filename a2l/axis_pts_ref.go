package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

/*If the addresses of the axis point values are separated from the table values in the
emulation memory and must be described by a special AXIS_PTS data record,
the data record is referenced by means of the keyword AXIS_PTS_REF.*/
type axisPtsRef struct {
	/*axisPoints name of the AXIS_PTS data record which describes the
	common axis points distribution (group axis points and record layout: see AXIS_PTS).
	Note: If the referenced axis is an element of an array or a component of a structure,
	the identifier to be used for the reference shall be the name built according to rules described at INSTANCE.
	Note: If (and only if) the axis description is utilized inside a structure definition and the common axis
	referenced by this keyword is a component of the same structure, then it is also allowed to refer to
	the axis by using the THIS keyword followed by a dot and the component name of the axis instead of using a concrete instance name.*/
	axisPoints    string
	axisPointsSet bool
}

func parseAxisPtsRef(tok *tokenGenerator) (axisPtsRef, error) {
	apr := axisPtsRef{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("axisPtsRef could not be parsed")
	} else if !apr.axisPointsSet {
		apr.axisPoints = tok.current()
		apr.axisPointsSet = true
		log.Info().Msg("axisPtsRef axisPoints successfully parsed")
	}
	return apr, err
}
