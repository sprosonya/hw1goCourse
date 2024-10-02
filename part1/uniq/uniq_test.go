package uniq

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"uniq/settings"
)

func TestOK(t *testing.T) {
	tests := []struct {
		str      []string
		option   settings.Options
		expected []string // ожидаемый результат
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."}, settings.Options{Compare: settings.EqualWithCase}, []string{"I love music.", "", "I love music of Kartik.", "Thanks.", "I love music of Kartik."}},
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."}, settings.Options{CountNumOfRepeats: true, Compare: settings.EqualWithCase}, []string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks.", "2 I love music of Kartik."}},
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."}, settings.Options{WriteRepeatedLines: true, Compare: settings.EqualWithCase}, []string{"I love music.", "I love music of Kartik.", "I love music of Kartik."}},
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."}, settings.Options{OnlyUnicLines: true, Compare: settings.EqualWithCase}, []string{"", "Thanks."}},
		{[]string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."}, settings.Options{Compare: settings.EqualWithoutCase}, []string{"I LOVE MUSIC.", "", "I love MuSIC of Kartik.", "Thanks.", "I love music of kartik."}},
		{[]string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."}, settings.Options{Compare: settings.EqualWithCase, IgnoreNumOfFields: 1}, []string{"We love music.", "", "I love music of Kartik.", "Thanks."}},
		{[]string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."}, settings.Options{Compare: settings.EqualWithCase, IgnoreNumOfChars: 1}, []string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."}},
	}
	for i := range tests {
		assert.Equal(t, Uniq(tests[i].str, tests[i].option), tests[i].expected, "must be equal")
	}
}
