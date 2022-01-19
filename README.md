# calibrationReader
[![Go Report Card](https://goreportcard.com/badge/github.com/asap2Go/calibrationReader)](https://goreportcard.com/report/github.com/asap2Go/calibrationReader)

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/12a206a67d5e4789943701f757b49f6d)](https://www.codacy.com/gh/asap2Go/calibrationReader/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=asap2Go/calibrationReader&amp;utm_campaign=Badge_Grade)

[![Go](https://github.com/asap2Go/calibrationReader/actions/workflows/go-build-test-and-license.yml/badge.svg)](https://github.com/asap2Go/calibrationReader/actions/workflows/go-build-test-and-license.yml)

[![CodeQL](https://github.com/asap2Go/calibrationReader/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/asap2Go/calibrationReader/actions/workflows/codeql-analysis.yml)

[![Coverage Status](https://coveralls.io/repos/github/asap2Go/calibrationReader/badge.svg?branch=main)](https://coveralls.io/github/asap2Go/calibrationReader?branch=main)

[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/12a206a67d5e4789943701f757b49f6d)](https://www.codacy.com/gh/asap2Go/calibrationReader/dashboard?utm_source=github.com&utm_medium=referral&utm_content=asap2Go/calibrationReader&utm_campaign=Badge_Coverage)

 reads characteristics information from a2l and fills it with the data from a hex file.
 At least that is the plan. 
 Currently it can parse a2l-files as well as the corresponding IntelHex32 or Motorola S19 files. 
 And it is quite fast at that. Currently a real world A2L(80MB) with its corresponding Hex File(10MB) will be parsed in less than a second.
 
 But it still lacks the last bit of work which is implementing the methods for axis_pts, axis_descr, record_layout and fnc_values
 in order to understand the memory layout and position of a specific characteristic.
 This is somewhat of a convoluted mess in the a2l standard due to its historic growth and will be implemented when I have a little more spare time.
 
 I am a mechanical engineer, so any code you see might not be up to the standards of true/correct/modern/acceptable software development ;)
 Feedback is always appreciated.
 
 The only dependency outside the go standard library is currently zerolog.
