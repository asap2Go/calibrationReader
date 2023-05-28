package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

// Specification of a binary blob object which has no further semantic interpretation in the MCD system.
// It is just an array of bytes without a conversion method or special record layout.
type blob struct {
	name                string
	nameSet             bool
	longIdentifier      string
	longIdentifierSet   bool
	address             string //uint32
	addressSet          bool
	size                uint32
	sizeSet             bool
	addressType         AddrTypeEnum
	addressTypeSet      bool
	annotation          []annotation
	calibrationAccess   calibrationAccessEnum
	displayIdentifier   DisplayIdentifier
	ecuAddressExtension ecuAddressExtension
	IfData              []IfData
	maxRefresh          MaxRefresh
	modelLink           modelLink
	symbolLink          symbolLink
}

func parseBlob(tok *tokenGenerator) (blob, error) {
	b := blob{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("blob annotation could not be parsed")
				break forLoop
			}
			b.annotation = append(b.annotation, buf)
			log.Info().Msg("blob annotation successfully parsed")
		case addressTypeToken:
			b.addressType, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("blob addressType could not be parsed")
				break forLoop
			}
			b.addressTypeSet = true
			log.Info().Msg("blob addressType successfully parsed")
		case calibrationAccessToken:
			b.calibrationAccess, err = parseCalibrationAccessEnum(tok)
			if err != nil {
				log.Err(err).Msg("blob calibrationAccess could not be parsed")
				break forLoop
			}
			log.Info().Msg("blob calibrationAccess successfully parsed")
		case displayIdentifierToken:
			b.displayIdentifier, err = parseDisplayIdentifier(tok)
			if err != nil {
				log.Err(err).Msg("blob displayIdentifier could not be parsed")
				break forLoop
			}
			log.Info().Msg("blob displayIdentifier successfully parsed")
		case ecuAddressExtensionToken:
			b.ecuAddressExtension, err = parseECUAddressExtension(tok)
			if err != nil {
				log.Err(err).Msg("blob ecuAddressExtension could not be parsed")
				break forLoop
			}
			log.Info().Msg("blob ecuAddressExtension successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("blob IfData could not be parsed")
				break forLoop
			}
			b.IfData = append(b.IfData, buf)
			log.Info().Msg("blob IfData successfully parsed")
		case maxRefreshToken:
			b.maxRefresh, err = parseMaxRefresh(tok)
			if err != nil {
				log.Err(err).Msg("blob maxRefresh could not be parsed")
				break forLoop
			}
			log.Info().Msg("blob maxRefresh successfully parsed")
		case modelLinkToken:
			b.modelLink, err = parseModelLink(tok)
			if err != nil {
				log.Err(err).Msg("blob modelLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("blob modelLink successfully parsed")
		case symbolLinkToken:
			b.symbolLink, err = parseSymbolLink(tok)
			if err != nil {
				log.Err(err).Msg("blob symbolLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("blob symbolLink successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("blob could not be parsed")
				break forLoop
			} else if tok.current() == endBlobToken {
				if !b.addressTypeSet {
					b.addressType = DIRECT
				}
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("blob could not be parsed")
				break forLoop
			} else if !b.nameSet {
				b.name = tok.current()
				b.nameSet = true
			} else if !b.longIdentifierSet {
				b.longIdentifier = tok.current()
				b.longIdentifierSet = true
			} else if !b.addressSet {
				b.address = tok.current()
				b.addressSet = true
			} else if !b.sizeSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("blob size could not be parsed")
					break forLoop
				}
				b.size = uint32(buf)
				b.sizeSet = true
				log.Info().Msg("blob size successfully parsed")
			}
		}
	}
	return b, err
}
