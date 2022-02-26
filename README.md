# calibrationReader
[![Go Report Card](https://goreportcard.com/badge/github.com/asap2Go/calibrationReader)](https://goreportcard.com/report/github.com/asap2Go/calibrationReader)  [![Codacy Badge](https://app.codacy.com/project/badge/Grade/e19560faf3484ccb88922ad3548b19ad)](https://www.codacy.com/gh/asap2Go/calibrationReader/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=asap2Go/calibrationReader&amp;utm_campaign=Badge_Grade)
 [![Codacy Badge](https://app.codacy.com/project/badge/Coverage/e19560faf3484ccb88922ad3548b19ad)](https://www.codacy.com/gh/asap2Go/calibrationReader/dashboard?utm_source=github.com&utm_medium=referral&utm_content=asap2Go/calibrationReader&utm_campaign=Badge_Coverage)
 [![Go](https://github.com/asap2Go/calibrationReader/actions/workflows/go-build-test-and-license.yml/badge.svg)](https://github.com/asap2Go/calibrationReader/actions/workflows/go-build-test-and-license.yml)
 [![CodeQL](https://github.com/asap2Go/calibrationReader/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/asap2Go/calibrationReader/actions/workflows/codeql-analysis.yml)
 [![Go Reference](https://pkg.go.dev/badge/github.com/asap2Go/calibrationReader.svg)](https://pkg.go.dev/github.com/asap2Go/calibrationReader)

 The calibrationReader package reads characteristics information from an a2l file and fills it with the data from a hex file.
 At least that is the plan (see TO-DO). 
 Currently it can parse a2l-files as well as the corresponding IntelHex32 or Motorola S19 files. 
 And it is quite fast at that. Currently a real world A2L(80MB) with its corresponding Hex File(10MB) will be parsed in less than a second.
 
 USAGE:
 cd, err := ReadCalibration(a2lFilePath, hexFilePath)
 
 parses a2l and hex file into datastructures.
 All relevant information e.g. Record Layouts, Measurements, Characterstics, etc. 
 is part of a module which is in turn part of the Project whithin the a2l data structure.
 An a2l file can contain several modules but usually only contains one. So indexOfModuleInProject (see below) can basically assumed to be 0.
 All datastructures that are directly below a module within the a2l data structure hierarchy are accessible through maps by their identifiers.
 
 To access System Constants:
 sc, exists := cd.a2l.Project.Module[indexOfModuleInProject].SystemConstants["NameOfSystemConstant"]
 
 To access Characteristics:
 c, exists := cd.a2l.Project.Module[indexOfModuleInProject].Characteristics["NameOfCharacteristic]
 
 and so on. 
 
 To access a specific memory location in the hex-file (contains a single byte):
 b, exists := cd.hex[12345]
 
 TO-DO:
 The Package still lacks the last bit of work which is implementing the methods for axis_pts, axis_descr, record_layout and fnc_values
 in order to understand the memory layout and position of a specific characteristic.
 This is somewhat of a convoluted mess in the a2l standard due to its historic growth and will be implemented when I have a little more spare time.
 
 I am a mechanical engineer, so any code you see might not be up to the standards of true/correct/modern/acceptable software development ;)
 Feedback is always appreciated.
 
 The only dependency outside the go standard library is currently zerolog.
