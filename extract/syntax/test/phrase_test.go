package test

import (
	"github.com/auho/go-nlp/extract/syntax"
	"strings"
	"testing"
)

var str2 = "电影/n 好看/a ，/x 特效/n 好/a 。/x 经济/n 实惠/vn 、/x 动力/n 不错/a 、/x 油耗/n 低/a 。/x 环境/n 不错/a ，/x 进去/v 就/d 有/v 暖气/n ，/x 叫/v 的/uj 11/m 号/m 技师/n ，/x 服务/vn 确实/ad 不错/a ，/x 95/m 后/t 妹子/n ，/x 技术/n 好/a ，/x 挺不错/i 的/uj 体验/n 。/x"

func Test_Phrase(t *testing.T) {
	beaus := strings.Split(str2, " ")

	a := syntax.NewPhrase()
	a.Matchmaker("n a")

	res := a.Seek(beaus)
	//t.Log(res)
	if len(res) == 6 {
		t.Log("ok")
	} else {
		t.Error("error")
	}

	a.Matchmaker("ad a")

	res = a.Seek(beaus)
	//t.Log(res)
	if len(res) == 7 {
		t.Log("ok")
	} else {
		t.Error("error")
	}
}

func Benchmark_Phrase_Seek(b *testing.B) {
	b.StopTimer()
	a := syntax.NewPhrase()

	a.Matchmaker("n a")
	beaus := strings.Split(str2, " ")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		a.Seek(beaus)
	}
}
