package src

<<<<<<< HEAD
import (
	"strings"
)

func MostCommonWord(paragraph string, banned []string) string {
=======
import "strings"

func mostCommonWord(paragraph string, banned []string) string {
>>>>>>> 3302af24230eb8cd38a0443abbcf06003f17fdaf
	bytes := []byte(strings.ToLower(paragraph))
	var temp []byte
	var words []string
	for _,v :=range bytes{
		if (v<='Z' && v>= 'A') || (v<='z' && v>= 'a'){
			temp = append(temp, v)
		}else{
			if len(temp) != 0{
				words = append(words, string(temp))
				temp = []byte{}
			}
		}
	}
	if len(temp) != 0{
		words = append(words, string(temp))
		temp = []byte{}
	}
	wordsMap := make(map[string]int)
	for _,v := range words{
		if _,ok := wordsMap[v];!ok{
			wordsMap[v] = 1
		}else{
			wordsMap[v]++
		}
	}
	for _,v :=range banned{
		if _,ok := wordsMap[v];ok{
			wordsMap[v] = 0
		}
	}
	max:= 0
	res := ""
	for k,v :=range wordsMap{
		if v>max{
			res = k
			max = v
		}
	}
	return res
<<<<<<< HEAD
}
=======
}
>>>>>>> 3302af24230eb8cd38a0443abbcf06003f17fdaf
