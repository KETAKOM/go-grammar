package grammar

import "fmt"

//引数の型は必須
//返り値の型は必須
//複数の返り値を返すことができる
//関数もデータ型なので変数に格納可能
//即時関数のようにしてその場で実行することも可能

//引数なし関数
func SayHi() {
	fmt.Println("sayHi")
}

//引数あり関数
func SayName(name string) {
	fmt.Println(name)
}

//return関数
func GetMessage(name string) string {
	msg := "hei! my name is " + name
	return msg
}

//名前付きreturn関数
//関数内で使った変数名を返す
func GetHelloMessage(name string) (msg string) {
	msg = "Hello " + name
	return
}

//複数の返り値を返す
func Swap(a int, b int) (a_ int, b_ int) {
	a_ = a
	b_ = b

	return a_, b_
}

