package grammar

//constで定数を定義できる
//定数でiotaを使えば、簡単に連番の定数を作成できる

const (
	a = iota //0
	b //1
	c //2
	d //3
)

func Const1() int {
	return a + b + c + d
}
