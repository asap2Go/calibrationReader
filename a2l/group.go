package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type group struct {
	groupName              string
	groupNameSet           bool
	groupLongIdentifier    string
	groupLongIdentifierSet bool
	annotation             []annotation
	functionList           FunctionList
	ifData                 []IfData
	refCharacteristic      refCharacteristic
	refMeasurement         refMeasurement
	root                   rootKeyword
	subGroup               subGroup
}

func parseGroup(tok *tokenGenerator) (group, error) {
	g := group{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("group annotation could not be parsed")
				break forLoop
			}
			g.annotation = append(g.annotation, buf)
			log.Info().Msg("group annotation successfully parsed")
		case beginFunctionListToken:
			g.functionList, err = parseFunctionList(tok)
			if err != nil {
				log.Err(err).Msg("group functionList could not be parsed")
				break forLoop
			}
			log.Info().Msg("group functionList successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("group ifData could not be parsed")
				break forLoop
			}
			g.ifData = append(g.ifData, buf)
			log.Info().Msg("group ifData successfully parsed")
		case beginRefCharacteristicToken:
			g.refCharacteristic, err = parseRefCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("group refCharacteristic could not be parsed")
				break forLoop
			}
			log.Info().Msg("group refCharacteristic successfully parsed")
		case beginRefMeasurementToken:
			g.refMeasurement, err = parseRefMeasurement(tok)
			if err != nil {
				log.Err(err).Msg("group refMeasurement could not be parsed")
				break forLoop
			}
			log.Info().Msg("group refMeasurement successfully parsed")
		case rootToken:
			g.root, err = parseRoot(tok)
			if err != nil {
				log.Err(err).Msg("group root could not be parsed")
				break forLoop
			}
			log.Info().Msg("group root successfully parsed")
		case beginSubGroupToken:
			g.subGroup, err = parseSubGroup(tok)
			if err != nil {
				log.Err(err).Msg("group subGroup could not be parsed")
				break forLoop
			}
			log.Info().Msg("group subGroup successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("group could not be parsed")
				break forLoop
			} else if tok.current() == endGroupToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("group could not be parsed")
				break forLoop
			} else if !g.groupNameSet {
				g.groupName = tok.current()
				g.groupNameSet = true
				log.Info().Msg("group groupName successfully parsed")
			} else if !g.groupLongIdentifierSet {
				g.groupLongIdentifier = tok.current()
				g.groupLongIdentifierSet = true
				log.Info().Msg("group groupLongIdentifier successfully parsed")
			}
		}
	}
	return g, err
}
