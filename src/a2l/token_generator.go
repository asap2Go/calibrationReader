package a2l

import (
	"errors"
	"strings"
	"sync/atomic"

	"github.com/rs/zerolog/log"
)

const expectedNumberOfTokens = 3000000

var tokenList []string
var moduleCount uint32

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
	var locTokens []chan []string
	tokenList = make([]string, 0, expectedNumberOfTokens)
	lines := strings.Split(str, "\n")
	for i := 0; i < numProc; i++ {
		//calculate start and end for the slices each go routine is given to parse
		start := (len(lines) / numProc) * i
		end := start + (len(lines) / numProc)
		if i+1 == numProc {
			//integer divisions might round up or down, so we make sure that we get the "real" end here
			end = len(lines)
		}
		c := make(chan []string, 1)
		locTokens = append(locTokens, c)
		go tokenBuilderRoutine(lines[start:end], c)
	}
	//collect token lists from channels
	for _, c := range locTokens {
		for t := range c {
			tokenList = append(tokenList, t...)
		}
	}
	tokenList = append(tokenList, emptyToken)
	tg := tokenGenerator{}
	tg.index = 0
	/*Set multithreading flag in accordance with the number of modules
	Currently there is no support for parsing multi-module files with multithreading enabled due to complexity reasons
	and minimal benefit due to overhead of detecting module borders and parser coordination / channel creation.*/
	if moduleCount > 1 {
		useMultithreading = false
		log.Info().Uint32("number of modules", moduleCount).Msg("multiple modules detected, parsing singlethreaded")
	} else {
		useMultithreading = true
		log.Info().Msg("only one module detected, parsing multithreaded")
	}
	return tg, nil
}

func tokenBuilderRoutine(lines []string, c chan []string) {
	//initialize fields with the expected capacity of tokens it will have to hold in order to avoid reallocations
	fields := make([][]string, 0, expectedNumberOfTokens/numProc)
	for _, l := range lines {
		if strings.TrimSpace(l) != emptyToken {
			fields = append(fields, strings.Fields(l))
		}
	}
	c <- buildTokenList(fields)
	close(c)
}

func buildTokenList(str [][]string) []string {
	var err error
	tl := make([]string, 0, expectedNumberOfTokens/numProc)
	currentOuterIndex := 0
	currentInnerIndex := 0
	firstRun := true
	t := emptyToken
	//build a list of all valid tokens and push them onto tl.
	for {
		t, err = buildNextValidToken(&currentOuterIndex, &currentInnerIndex, str, firstRun)
		firstRun = false
		if err != nil {
			return tl
		}
		if t != emptyToken {
			if strings.Contains(t, "*/") {
				tl = make([]string, 0, expectedNumberOfTokens/numProc)
			} else {
				tl = append(tl, t)
			}
		} else {
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
	} else if strings.Contains(t, slashToken) && !strings.Contains(t, endMultilineCommentToken) {
		twoWordedKeyword, err := getTwoWordedToken(currentOuterIndex, currentInnerIndex, str)
		if err != nil {
			return emptyToken, err
		}
		return twoWordedKeyword, nil
	} else {
		return t, nil
	}
}

//skipLineComment if a "//" is detected the token generator moves to the next line (currentOuterIndex++)
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

//skipLineComment if a "/*" is detected then token generator gets the next token until a "*/" is reached.
//it will then return the next raw value token after the comment
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

//moveToNextRawValue returns back the next valid, white space separated value.
//in case the line ends it restarts on the next line
//in case there are no lines left and no words within the last line it will return an empty token
//which is used to signal the parser that the eof has been reached.
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

//isKeyword is used in the matrixDim parser to detected when there are no dimensions left to parse
//this is necessary because not every version of the a2l standard has a clear rule about how many dimensions should be expected
//or are necessary to define ("1 0 0" and "1" are both valid descriptions for a curve)
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

//getTwoWordedToken handles keywords that contain a / like e.g. "/begin CHARACTERISTIC"
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
			//atomic is used so several goroutines can update the variable
			atomic.AddUint32(&moduleCount, 1)
		}
		return twoWordedToken, nil

	} else {
		err = errors.New("no /begin or /end token found")
		return emptyToken, err
	}
}
