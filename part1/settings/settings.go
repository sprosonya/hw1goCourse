package settings

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	WriteRepeatedLines bool
	CountNumOfRepeats  bool
	OnlyUnicLines      bool
	Compare            func(string, string, int, int) bool
	IgnoreNumOfChars   int
	IgnoreNumOfFields  int
	IgnoreCase         bool
}

func SetInput() (io.Reader, error) {
	if flag.Arg(0) != "" {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		return os.Stdin, nil
	}
}
func SetOutput() (io.Writer, error) {
	if flag.Arg(1) != "" {
		file, err := os.OpenFile(flag.Arg(1), os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		return os.Stdout, nil
	}
}
func checkRightUsage(options Options) error {
	var booleans []bool = []bool{options.WriteRepeatedLines, options.CountNumOfRepeats, options.OnlyUnicLines}
	count := 0
	for _, value := range booleans {
		if value {
			count++
		}
	}
	if count > 1 {
		usagesMas := rightUsage(options, booleans)
		var usages string
		for i := range usagesMas {
			usages += "\n" + usagesMas[i]
		}
		return fmt.Errorf("use only one option: -d or -c or -u \ntry one of this usages:" + usages)
	}
	return nil
}
func GetOptions() (Options, error) {
	var flagD bool
	var flagC bool
	var flagU bool
	var flagI bool
	var flagF int
	var flagS int
	flag.BoolVar(&flagD, "d", false, "only repeat")
	flag.BoolVar(&flagC, "c", false, "count")
	flag.BoolVar(&flagU, "u", false, "not repeat")
	flag.BoolVar(&flagI, "i", false, "ignore letter case")
	flag.IntVar(&flagF, "f", 0, "ignore first n fields")
	flag.IntVar(&flagS, "s", 0, "ignore first n chars")
	flag.Parse()
	options := Options{
		WriteRepeatedLines: flagD,
		CountNumOfRepeats:  flagC,
		OnlyUnicLines:      flagU,
		IgnoreNumOfChars:   flagS,
		IgnoreNumOfFields:  flagF,
		IgnoreCase:         flagI,
	}
	err := checkRightUsage(options)
	if err != nil {
		return Options{}, err
	}
	if flagI {
		options.Compare = EqualWithoutCase
	} else {
		options.Compare = EqualWithCase
	}
	return options, nil
}
func rightUsage(options Options, booleans []bool) []string {
	str := "go run main.go "
	var flags []string = []string{"-d", "-c", "-u"}
	var strings []string = make([]string, 0)
	for i, value := range booleans {
		if value {
			strings = append(strings, str)
			strings[i] += flags[i]
		}
	}
	if options.IgnoreNumOfChars > 0 {
		for i := range strings {
			strings[i] = strings[i] + " -s " + strconv.Itoa(options.IgnoreNumOfChars)
		}
	}
	if options.IgnoreNumOfFields > 0 {
		for i := range strings {
			strings[i] = strings[i] + " -f " + strconv.Itoa(options.IgnoreNumOfFields)
		}
	}
	if options.IgnoreCase {
		for i := range strings {
			strings[i] = strings[i] + " -i "
		}
	}
	if flag.Arg(0) != "" {
		for i := range strings {
			strings[i] = strings[i] + " " + flag.Arg(0)
		}
	}
	if flag.Arg(1) != "" {
		for i := range strings {
			strings[i] = strings[i] + " " + flag.Arg(1)
		}
	}
	return strings
}
func cutFieldsAndChars(str string, ignoredFields int, ignoredChars int) string {
	words := strings.Fields(str)
	if len(words) < ignoredFields {
		return ""
	}
	newStr := strings.Join(words[ignoredFields:], " ")
	if len(newStr) < ignoredChars {
		return ""
	}
	return strings.Join(words[ignoredFields:], " ")[ignoredChars:]
}

func EqualWithoutCase(str1 string, str2 string, ignoredFields int, ignoredChars int) bool {
	newStr1 := cutFieldsAndChars(str1, ignoredFields, ignoredChars)
	newStr2 := cutFieldsAndChars(str2, ignoredFields, ignoredChars)
	return strings.ToLower(newStr1) == strings.ToLower(newStr2)
}

func EqualWithCase(str1 string, str2 string, ignoredFields int, ignoredChars int) bool {
	newStr1 := cutFieldsAndChars(str1, ignoredFields, ignoredChars)
	newStr2 := cutFieldsAndChars(str2, ignoredFields, ignoredChars)
	return newStr1 == newStr2
}
