package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

/*This model_link can be used to reference a software model object the
CHARACTERISTIC, AXIS_PTS, BLOB, INSTANCE or MEASUREMENT was derived from.
Note:
The model_link does not have to be unique and is not referenced elsewhere in the A2L file.
But it is recommended that the model_link is unique
in order to avoid confusion in the user interface of the MCD system.
Note:
In the model world, the objects may have different names than in the calibration world.
To allow a round trip (model, A2L file, data exchange file, model)
the model name must be transferred to the calibration world.
The A2L file generation may add the model object name to the A2L file via the
MODEL_LINK additionally to the ASAP2 object name.
or further process steps the MC-System may transfer the MODEL_LINK to data exchan*/
type modelLink struct {
	//parameter name in the software model object
	modelName    string
	modelNameSet bool
}

func parseModelLink(tok *tokenGenerator) (modelLink, error) {
	ml := modelLink{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("modelLink could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("modelLink could not be parsed")
	} else if !ml.modelNameSet {
		ml.modelName = tok.current()
		ml.modelNameSet = true
		log.Info().Msg("modelLink symbolName successfully parsed")
	}
	return ml, err
}
