package uniq

import (
	"strconv"
	"uniq/settings"
)

type lineInfo struct {
	str   string
	count int
}

func Uniq(masOfLines []string, options settings.Options) []string {
	var lineInfoMas []lineInfo = make([]lineInfo, 0)
	var ignoredChars int = options.IgnoreNumOfChars
	var ignoredFields int = options.IgnoreNumOfFields

	for i := 0; i < len(masOfLines); i++ {
		str := masOfLines[i]
		if i == 0 {
			lineInfoMas = append(lineInfoMas, lineInfo{str, 1})
			continue
		}
		prev := masOfLines[i-1]
		if options.Compare(str, prev, ignoredFields, ignoredChars) {
			lineInfoMas[len(lineInfoMas)-1].count += 1
		} else {
			lineInfoMas = append(lineInfoMas, lineInfo{str, 1})
		}
	}

	var resultMas []string = make([]string, 0)

	if options.WriteRepeatedLines {
		for i := 0; i < len(lineInfoMas); i++ {
			num := lineInfoMas[i].count
			if num > 1 {
				resultMas = append(resultMas, lineInfoMas[i].str)
			}
		}
	} else if options.OnlyUnicLines {
		for i := 0; i < len(lineInfoMas); i++ {
			num := lineInfoMas[i].count
			if num == 1 {
				resultMas = append(resultMas, lineInfoMas[i].str)
			}
		}
	} else if options.CountNumOfRepeats {
		for i := 0; i < len(lineInfoMas); i++ {
			num := lineInfoMas[i].count
			resultMas = append(resultMas, strconv.Itoa(num)+" "+lineInfoMas[i].str)
		}
	} else {
		for i := 0; i < len(lineInfoMas); i++ {
			resultMas = append(resultMas, lineInfoMas[i].str)
		}
	}
	return resultMas
}
