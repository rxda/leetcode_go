package src

var keyMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}
var result []string
func LetterCombinations(digits string) []string {
	keys := []rune(digits)
	for k,v :=range keys{
		if k==0 {
			result = append(result, keyMap[string(v)]...)
		}else{
			temp := []string{}
			for _,v1 :=range result{
				for _,v2 :=range keyMap[string(v)]{
					temp = append(temp, v1+v2)
				}
			}
			result =temp
		}
	}
	return result
}
