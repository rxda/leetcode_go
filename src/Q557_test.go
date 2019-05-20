package src

import "testing"

func BenchmarkReverseWords(b *testing.B) {
	in := "Let's take LeetCode contest Let's take LeetCode contest Let's take LeetCode contest"
	ans :="s'teL ekat edoCteeL tsetnoc s'teL ekat edoCteeL tsetnoc s'teL ekat edoCteeL tsetnoc"
	for i:=0;i<b.N;i++{
		if c:=ReverseWords(in);c!=ans {
			b.Error(c,"\n",ans)
		}
	}
}
