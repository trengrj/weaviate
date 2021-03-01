package contextionary

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	inputVersionRegexString   = `^.*-v(?P<Major>[0-9]+)\.(?P<Minor>[0-9]+)\.(?P<Patch>[0-9]+)$`
	minimumVersionRegexString = `^(?P<Major>[0-9]+)\.(?P<Minor>[0-9])+\.(?P<Patch>[0-9]+)$`
)

func extractVersionAndCompare(input, requiredMin string) (bool, error) {
	inputRegexp := regexp.MustCompile(inputVersionRegexString)
	minimumRegexp := regexp.MustCompile(minimumVersionRegexString)

	if ok := inputRegexp.MatchString(input); !ok {
		return false, fmt.Errorf("unexpected input version tag: %s", input)
	}

	if ok := minimumRegexp.MatchString(requiredMin); !ok {
		return false, fmt.Errorf("unexpected threshold version tag: %s", requiredMin)
	}

	inputMatches := inputRegexp.FindAllStringSubmatch(input, 4)
	inputMajor, _ := strconv.Atoi(inputMatches[0][1])
	inputMinor, _ := strconv.Atoi(inputMatches[0][2])
	inputPatch, _ := strconv.Atoi(inputMatches[0][3])

	minimumMatches := minimumRegexp.FindAllStringSubmatch(requiredMin, 4)
	minimumMajor, _ := strconv.Atoi(minimumMatches[0][1])
	minimumMinor, _ := strconv.Atoi(minimumMatches[0][2])
	minimumPatch, _ := strconv.Atoi(minimumMatches[0][3])

	return compareSemver(inputMajor, inputMinor, inputPatch, minimumMajor, minimumMinor, minimumPatch), nil
}

func compareSemver(iMaj, iMin, iPat, rMaj, rMin, rPat int) bool {
	if iMaj > rMaj {
		return true
	}

	if iMaj < rMaj {
		return false
	}

	if iMin > rMin {
		return true
	}

	if iMin < rMin {
		return false
	}

	if iPat > rPat {
		return true
	}

	if iPat < rPat {
		return false
	}

	return true
}
