package test

import (
	"testing"
	"strings"
	"github.com/auho/go-nlp/extract/syntax"
)

var str = "电影/n 好看/a ，/x 特效/n 好/a 。/x 经济/n 实惠/vn 、/x 动力/n 不错/a 、/x 油耗/n 低/a 。/x 环境/n 不错/a ，/x 进去/v 就/d 有/v 暖气/n ，/x 叫/v 的/uj 11/m 号/m 技师/n ，/x 服务/vn 确实/ad 不错/a ，/x 95/m 后/t 妹子/n ，/x 技术/n 好/a ，/x 挺不错/i 的/uj 体验/n 。/x"

func TestHeadAndTailLimited(t *testing.T) {
	a := syntax.NewHeadAndTailLimited()
	a.Matchmaker("n", "a", "4")

	beaus := strings.Split(str, " ")

	res := a.Seek(beaus)
	if len(res) == 7 {
		t.Log("ok")
	} else {
		t.Error("error")
	}
}

func BenchmarkHeadAndTailLimited_SeekOne(b *testing.B) {
	b.StopTimer()
	a := syntax.NewHeadAndTailLimited()

	a.Matchmaker("n", "a", "4")
	beaus := strings.Split(str, " ")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		a.Seek(beaus)
	}
}
