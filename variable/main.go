package main

import(
	"fmt"
)

func main() {
	//変数
	//明示的な変数の定義
	//数値型
	var n int
	n = 5
	//文字列型
	var m string
	m = "aaa"
	//複数の数値型の変数を定義
	var a, b, c int
	a = 1
	b ,c = 2
	

	//int型の変数x, yとstring型の変数nameを定義
	var (
		x, y int
		name string
	)
	
	//暗黙的な変数の定義
	//【暗黙的には[変数名] := [変数の値]もしくは1
	// var [変数名] = [変数の値]のように定義する】
	//型推論が行われ、型指定の必要がない

	i := 1
	b := true
	f := 3.14
	s := "abc"

	var ii = 1
	var (
		bb = true
		ff = 2222
		ss = "abss"
	)


	//関数内で定義された変数はローカル変数である
	//関数外で定義された変数はパッケージ変数である
	var local_variable = 1 //ローカル変数
}
var package_variable = 1 //パッケージ変数
