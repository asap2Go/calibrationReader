package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

/*This keyword shall be used if the function is derived from an AUTOSAR
"SwComponentType" or "SwComponentPrototype" as a link into the AUTOSAR model,
e.g. for referencing the documentation of the component.
The "SwComponentType" shall appear as function to give a link to the documentation of
the component type and as container for any existing shared parameters.
The "SwComponentPrototype" shall appear as function to give a link to the documentation
of the component instance and as container for any provided or received or local
variables (MEASUREMENTs) and calibration parameters (CHARACTERISTICs). It shall
reference it’s "SwComponentType" using the keyword AR_PROTOTYPE_OF .*/
type arComponent struct {
	/*Specifies the kind of "SwComponentType" in the
	AUTOSAR system. The list of possible component types
	can be found in the AUTOSAR standard “Software
	Component Template”. Most common component types are:
	“ApplicationSwComponentType”,
	“ParameterSwComponentType”,
	”NvBlockSwComponentType” and
	“CompositionSwComponentType”.*/
	componentType    string
	componentTypeSet bool
	arPrototypeOf    arPrototypeOf
}

func parseArComponent(tok *tokenGenerator) (arComponent, error) {
	ac := arComponent{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case arPrototypeOfToken:
			ac.arPrototypeOf, err = parseArPrototypeOf(tok)
			if err != nil {
				log.Err(err).Msg("arComponent arPrototypeOf could not be parsed")
				break forLoop
			}
			log.Info().Msg("arComponent arPrototypeOf successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("arComponent could not be parsed")
				break forLoop
			} else if tok.current() == endArComponentToken {
				break forLoop
			} else if !ac.componentTypeSet {
				ac.componentType = tok.current()
				ac.componentTypeSet = true
				log.Info().Msg("arComponent componentType successfully parsed")
			}
		}
	}
	return ac, err
}
