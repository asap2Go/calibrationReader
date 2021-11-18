package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*a2mlVersion exists in order to declare what kind of BLOBs should be generated from
the AML parts. Since ASAP2 version 1.3.1 a specification for the storage layout of the
BLOBs exist. The keyword is optional. When the keyword is omitted, or the version
number is below 1.3.1 then the old BLOB format is used. When the A2ML version number
is 1.3.1, then the new format must be generated.
The A2ML version can be expressed by two numerals:
- VersionNo
- UpgradeNo
where ‘VersionNo’ represents the main version number and ‘UpgradeNo’ the upgrade
number (fractional part of version number).
This keyword will not be interpreted semantically anymore.*/
type a2mlVersion struct {
	//versionNo contains the Version number of AML part
	versionNo    uint16
	versionNoSet bool
	//upgradeNo contains the Upgrade number of AML part
	upgradeNo    uint16
	upgradeNoSet bool
}

func parseA2MLVersion(tok *tokenGenerator) (a2mlVersion, error) {
	av := a2mlVersion{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("a2mlVersion could not be parsed")
			break forLoop
		} else if !av.versionNoSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("a2mlVersion versionNo could not be parsed")
				break forLoop
			}
			av.versionNo = uint16(buf)
			av.versionNoSet = true
			log.Info().Msg("a2mlVersion versionNo successfully parsed")
		} else if !av.upgradeNoSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("a2mlVersion upgradeNo could not be parsed")
				break forLoop
			}
			av.upgradeNo = uint16(buf)
			av.upgradeNoSet = true
			log.Info().Msg("a2mlVersion upgradeNo successfully parsed")
			break forLoop
		}
	}
	return av, err
}
