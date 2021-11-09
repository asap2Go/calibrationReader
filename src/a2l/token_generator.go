package a2l

import (
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
)

var tokenList = make([]string, 2000000)
var moduleCount int

type tokenGenerator struct {
	index int
}

func (tg *tokenGenerator) current() string {
	if tg.index <= len(tokenList)-1 {
		return tokenList[tg.index]
	} else {
		return emptyToken
	}
}

func (tg *tokenGenerator) next() string {
	if tg.index < len(tokenList)-1 {
		tg.index++
		return tokenList[tg.index]
	} else {
		return emptyToken
	}

}

func (tg *tokenGenerator) previous() {
	tg.index--
}

func buildTokenGeneratorFromString(str string) (tokenGenerator, error) {
	//Split text file into lines and the lines into words separated by whitespace
	var lineContents [][]string
	lines := strings.Split(str, "\n")
	for _, l := range lines {
		if strings.TrimSpace(l) != emptyToken {
			lineContents = append(lineContents, strings.Fields(l))
		}
	}
	tokenList = buildTokenList(lineContents)
	tg := tokenGenerator{}
	tg.index = 0
	if moduleCount > 1 {
		useMultithreading = false
	}
	return tg, nil
}

func buildTokenList(str [][]string) []string {
	var err error
	var tl []string
	currentOuterIndex := 0
	currentInnerIndex := 0
	firstRun := true
	t := emptyToken
	//build a list of all valid tokens and push them onto tl.
	for {
		t, err = buildNextValidToken(&currentOuterIndex, &currentInnerIndex, str, firstRun)
		firstRun = false
		if err != nil {
			tl = append(tl, emptyToken)
			return tl
		}
		if t != emptyToken {
			tl = append(tl, t)
		} else {
			tl = append(tl, emptyToken)
			return tl
		}
	}
}

func buildNextValidToken(currentOuterIndex *int, currentInnerIndex *int, str [][]string, firstRun bool) (string, error) {
	var err error
	var t string
	if !firstRun {
		t, err = moveToNextRawValue(currentOuterIndex, currentInnerIndex, str)
		if err != nil {
			return emptyToken, err
		}
	} else {
		t = str[*currentOuterIndex][*currentInnerIndex]
		firstRun = false
	}
start:
	if t == emptyToken {
		return emptyToken, errors.New("reached eof")
	} else if strings.Contains(t, beginLineCommentToken) {
		t, err = skipLineComment(currentOuterIndex, currentInnerIndex, str)
		if err != nil {
			return emptyToken, err
		}
		goto start //if you use recursion then defer will stack up until you finally return and then jump x number of valid tokens.
	} else if strings.Contains(t, beginMultilineCommentToken) {
		t, err = skipMultilineComment(currentOuterIndex, currentInnerIndex, str)
		if err != nil {
			return emptyToken, err
		}
		goto start //if you use recursion then defer will stack up until you finally return and then jump x number of valid tokens.
	} else if strings.Contains(t, quotationMarkToken) {
		textInQuotationMarks, err := getTextInQuotationMarks(currentOuterIndex, currentInnerIndex, str)
		if err != nil {
			return emptyToken, err
		}
		return textInQuotationMarks, nil
	} else if strings.Contains(t, slashToken) {
		twoWordedKeyword, err := getTwoWordedToken(currentOuterIndex, currentInnerIndex, str)
		if err != nil {
			return emptyToken, err
		}
		return twoWordedKeyword, nil
	} else {
		return t, nil
	}
}

func skipLineComment(currentOuterIndex *int, currentInnerIndex *int, str [][]string) (string, error) {
	if strings.Contains(str[*currentOuterIndex][*currentInnerIndex], beginLineCommentToken) {
		//move to the next line and reset inner index to 0 (first word in new line)
		if *currentOuterIndex < len(str)-1 {
			*currentOuterIndex++
			*currentInnerIndex = 0
			return str[*currentOuterIndex][*currentInnerIndex], nil
		} else {
			return emptyToken, errors.New("reached eof")
		}

	} else {
		return emptyToken, errors.New("no comment found")
	}
}

func skipMultilineComment(currentOuterIndex *int, currentInnerIndex *int, str [][]string) (string, error) {
	var err error
	t := str[*currentOuterIndex][*currentInnerIndex]
	if strings.Contains(t, beginMultilineCommentToken) && !strings.Contains(t, endMultilineCommentToken) {
		//Get the next token in a Loop
		for {
			t, err = moveToNextRawValue(currentOuterIndex, currentInnerIndex, str)
			//Until the delimiter of the multiline comment is found
			if err != nil {
				return emptyToken, err
			}
			if strings.Contains(t, endMultilineCommentToken) {
				t, err = moveToNextRawValue(currentOuterIndex, currentInnerIndex, str)
				return t, err
			} else if t == emptyToken {
				log.Err(err).Msg("exit multi line comment due to EOF")
				err = errors.New("reached eof")
				return emptyToken, err
			}

		}
	} else if strings.Contains(t, beginMultilineCommentToken) && strings.Contains(t, endMultilineCommentToken) {
		t, err = moveToNextRawValue(currentOuterIndex, currentInnerIndex, str)
		return t, err
	} else {
		err = errors.New("reached eof")
		return emptyToken, err
	}
}

func getTextInQuotationMarks(currentOuterIndex *int, currentInnerIndex *int, str [][]string) (string, error) {
	//Get the next token in a Loop
	var err error
	t := str[*currentOuterIndex][*currentInnerIndex]
	if strings.Count(t, quotationMarkToken) == 0 {
		err = errors.New("no quotation mark found")
		return emptyToken, err
	} else if strings.Count(t, quotationMarkToken)%2 != 0 {
		//in case there is an odd number of quoatation marks we continue
		var numQuot int
		textInQuotationMarks := t
		for {
			t, err = moveToNextRawValue(currentOuterIndex, currentInnerIndex, str)
			if err != nil {
				return emptyToken, err
			}
			textInQuotationMarks = textInQuotationMarks + spaceToken + t
			//Until there is an even number of quotation marks
			numQuot = strings.Count(textInQuotationMarks, quotationMarkToken)
			if numQuot%2 == 0 {
				return textInQuotationMarks, nil
			}
		}
	} else {
		//in case there is an even number of quoatation marks we only need the current token
		textInQuotationMarks := t
		return textInQuotationMarks, nil
	}
}

func moveToNextRawValue(currentOuterIndex *int, currentInnerIndex *int, str [][]string) (string, error) {
	//If there are still tokens left in the current line
	if len(str[*currentOuterIndex])-1 > *currentInnerIndex && len(str[*currentOuterIndex]) > 0 {
		//move to the next token of the current line
		*currentInnerIndex++
		//If there are no tokens left in the current line
		return str[*currentOuterIndex][*currentInnerIndex], nil
	} else if len(str[*currentOuterIndex])-1 == *currentInnerIndex && len(str)-1 == *currentOuterIndex && str[*currentOuterIndex][*currentInnerIndex] != emptyToken {
		str[*currentOuterIndex] = append(str[*currentOuterIndex], emptyToken)
		*currentInnerIndex++
		//reached EOF
		return emptyToken, errors.New("reached eof")
	} else if len(str[*currentOuterIndex])-1 == *currentInnerIndex && len(str)-1 > *currentOuterIndex {
		//then go the next line and move to its first token
		*currentOuterIndex++
		*currentInnerIndex = 0
		if len(str[*currentOuterIndex]) > 0 {
			return str[*currentOuterIndex][*currentInnerIndex], nil
		} else {
			return moveToNextRawValue(currentOuterIndex, currentInnerIndex, str)
		}

	} else {
		return emptyToken, errors.New("unexpected error in token generation while trying to moveToNextRawValue")
	}
}

func isKeyword(str string) bool {
	isFound := false
	//look whether the given string is contained in the list of
	//valid a2l tokens as defined in tokens.go
	for i := 0; i < len(keywordList); i++ {
		if str == keywordList[i] {
			isFound = true
			break
		}
	}
	return isFound
}

func getTwoWordedToken(currentOuterIndex *int, currentInnerIndex *int, str [][]string) (string, error) {
	var err error
	t := str[*currentOuterIndex][*currentInnerIndex]
	if strings.Contains(t, "/begin") || strings.Contains(t, "/end") {
		twoWordedToken := t
		t, err = moveToNextRawValue(currentOuterIndex, currentInnerIndex, str)
		if err != nil {
			return emptyToken, err
		}
		twoWordedToken = twoWordedToken + spaceToken + t
		if twoWordedToken == beginModuleToken {
			//count the number of modules so the program can decide whether it is allowed to parse multithreaded.
			moduleCount++
		}
		return twoWordedToken, nil

	} else {
		err = errors.New("no /begin or /end token found")
		return emptyToken, err
	}
}
