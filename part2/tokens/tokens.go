package tokens

var priority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"~": 3,
}

func IsOperator(s string) bool {
	_, exists := priority[s]
	return exists
}

func GetPriority(op string) int {
	return priority[op]
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func IsNumber(s string) bool {
	for i := range s {
		if !IsDigit(s[i]) {
			return false
		}
	}
	return true
}
