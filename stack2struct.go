// stack2struct parses raw golang stack traces ([]byte) to a slice of well formated structs.
package stack2struct

import (
	"fmt"
	"strconv"
	"strings"
)

type stackTrace interface {
	AddEntry(lineNumber int, packageName string, fileName string, methodName string)
}

func Parse(trace []byte, stackTrace stackTrace) {

	lines := strings.Split(string(trace), "\n")

	var lineNumber int
	var fileName, packageName, methodName string

	for index, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		if index%2 == 0 {
			packageName, methodName = extractPackageName(line)
		} else {
			lineNumber, fileName = extractLineNumberAndFile(line)
			stackTrace.AddEntry(lineNumber, packageName, fileName, methodName)
		}
	}
}

func extractPackageName(line string) (packageName, methodName string) {
	packagePath, packageNameAndFunction := splitAtLastSlash(line)
	parts := strings.Split(packageNameAndFunction, ".")
	packageName = fmt.Sprintf("%s/%s", packagePath, parts[0])
	methodName = parts[1]
	return
}

func extractLineNumberAndFile(line string) (lineNumber int, fileName string) {
	_, fileAndLine := splitAtLastSlash(line)
	fileAndLine = removeSpaceAndSuffix(fileAndLine)
	parts := strings.Split(fileAndLine, ":")

	numberAsString := parts[1]
	number, _ := strconv.ParseUint(numberAsString, 10, 32)
	lineNumber = int(number)

	fileName = parts[0]
	return lineNumber, fileName
}

func splitAtLastSlash(line string) (left, right string) {
	parts := strings.Split(line, "/")
	right = parts[len(parts)-1]
	left = strings.Join(parts[:len(parts)-1], "/")
	return
}

func removeSpaceAndSuffix(line string) string {
	parts := strings.Split(line, " ")
	return strings.Join(parts[:len(parts)-1], " ")
}
