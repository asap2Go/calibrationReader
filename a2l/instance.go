package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type instance struct {
	name                string
	nameSet             bool
	LongIdentifier      string
	longIdentifierSet   bool
	typeDefName         string
	typeDefNameSet      bool
	address             string
	addressSet          bool
	addressType         addrTypeEnum
	annotation          []annotation
	calibrationAccess   calibrationAccessEnum
	displayIdentifier   DisplayIdentifier
	ecuAddressExtension ecuAddressExtension
	ifData              []IfData
	matrixDim           MatrixDim
	maxRefresh          MaxRefresh
	modelLink           modelLink
	overwrite           overwrite
	readWrite           readWriteKeyword
	symbolLink          symbolLink
}

func parseInstance(tok *tokenGenerator) (instance, error) {
	i := instance{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("instance annotation could not be parsed")
				break forLoop
			}
			i.annotation = append(i.annotation, buf)
			log.Info().Msg("instance annotation successfully parsed")
		case addressTypeToken:
			i.addressType, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("instance addressType could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance addressType successfully parsed")
		case calibrationAccessToken:
			i.calibrationAccess, err = parseCalibrationAccessEnum(tok)
			if err != nil {
				log.Err(err).Msg("instance calibrationAccess could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance calibrationAccess successfully parsed")
		case displayIdentifierToken:
			i.displayIdentifier, err = parseDisplayIdentifier(tok)
			if err != nil {
				log.Err(err).Msg("instance displayIdentifier could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance displayIdentifier successfully parsed")
		case ecuAddressExtensionToken:
			i.ecuAddressExtension, err = parseECUAddressExtension(tok)
			if err != nil {
				log.Err(err).Msg("instance ecuAddressExtension could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance ecuAddressExtension successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("instance ifData could not be parsed")
				break forLoop
			}
			i.ifData = append(i.ifData, buf)
			log.Info().Msg("instance ifData successfully parsed")
		case matrixDimToken:
			i.matrixDim, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("instance matrixDim could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance matrixDim successfully parsed")
		case maxRefreshToken:
			i.maxRefresh, err = parseMaxRefresh(tok)
			if err != nil {
				log.Err(err).Msg("instance maxRefresh could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance maxRefresh successfully parsed")
		case modelLinkToken:
			i.modelLink, err = parseModelLink(tok)
			if err != nil {
				log.Err(err).Msg("instance modelLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance modelLink successfully parsed")
		case beginOverwriteToken:
			i.overwrite, err = parseOverwrite(tok)
			if err != nil {
				log.Err(err).Msg("instance overwrite could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance overwrite successfully parsed")
		case readWriteToken:
			i.readWrite, err = parseReadWrite(tok)
			if err != nil {
				log.Err(err).Msg("instance readWrite could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance readWrite successfully parsed")
		case symbolLinkToken:
			i.symbolLink, err = parseSymbolLink(tok)
			if err != nil {
				log.Err(err).Msg("instance symbolLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("instance symbolLink successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("instance could not be parsed")
				break forLoop
			} else if tok.current() == endInstanceToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("instance could not be parsed")
				break forLoop
			} else if !i.nameSet {
				i.name = tok.current()
				i.nameSet = true
				log.Info().Msg("instance name successfully parsed")
			} else if !i.longIdentifierSet {
				i.LongIdentifier = tok.current()
				i.longIdentifierSet = true
				log.Info().Msg("instance longIdentifier successfully parsed")
			} else if !i.typeDefNameSet {
				i.typeDefName = tok.current()
				i.typeDefNameSet = true
				log.Info().Msg("instance typeDefName successfully parsed")
			} else if !i.addressSet {
				i.address = tok.current()
				i.addressSet = true
				log.Info().Msg("instance address successfully parsed")
			}
		}
	}
	return i, err
}
