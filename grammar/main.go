package main

import (
	"fmt"
	"./grammar"
)

func main() {

	/*ここから定数、関数*/
	var test_const = grammar.Const1()
	fmt.Println(test_const)

	grammar.SayHi()
	grammar.SayName("taro yamada")
	fmt.Println(grammar.GetMessage("masami iwaki"))
	fmt.Println(grammar.GetHelloMessage("tonoma"))
	fmt.Println(grammar.Swap(4, 5))

	//関数もデータの型なので変数に代入可能
	//その際は変数名いらない
	kansu := func(aa int, bb int) (int, int) {
		return bb , aa
	}
	
	fmt.Println(kansu(3, 4))
	
	//即時関数的なことも可能
	func(msg string) {
		fmt.Println(msg)
	}("sokujikansu")

	/*関数終わり*/
	
	/*配列開始*/
	//宣言と代入を分ける
	var arr_a [5]int  //arr_a[0]-arr_a[4]
	arr_a[2] = 3
	arr_a[4] = 10
	fmt.Println(arr_a)

	//宣言と代入を同時にする
	arr_b := [3]int{3,6,8}
	fmt.Println(arr_b)
	
	//配列の個数が未定の場合
	arr_c := [...]int{1,3,4,5,7,7}	
	fmt.Println(arr_c)

	/*配列終わり*/

	/*スライス開始*/
	/*
	■golangの配列ではダメなところ
	1.配列はあくまでも複数の要素を持った値→C言語なら要素0のポインタになったりした
	2.関数に配列を渡す場合には値を丸ごと渡すことになるので、メモリ的に非効率になる
	3.配列の要素数は固定になっていて、動的に変化できない
	*/
	
	/*
	■golangのスライス
	・golangでは配列よりもスライスをよく使う
	・配列の一部、または全部を指し示す参照型のデータ
	・スライスは配列への参照なので値を変更すると元の配列も値が変更される
	*/
	
	sl_a := [5]int{1, 2, 3, 4, 5}
	sl_b := sl_a[2:4]
	sl_c := sl_a[2:]
	sl_d := sl_a[:4]
	sl_e := sl_a[:]

	fmt.Println(sl_b)
	fmt.Println(sl_c)
	fmt.Println(sl_d)
	fmt.Println(sl_e)

	//スライスの長さを取得する
	fmt.Println(len(sl_c))

	//スライスの最大容量を取得
	fmt.Println(cap(sl_c))

	//make関数でいきなりスライスを作成する
	s1 := make([]int, 3) //[000]
	
	//いきなり値を割り当てたスライスを作成する
	s2 := []int{11, 22, 33}	//配列の宣言と似ている
	fmt.Println(s1)
	fmt.Println(s2)
	
	//appendでスライスの末尾に要素を追加
	s3 := append(s2, 8, 2, 10)
	fmt.Println(s3)
	fmt.Println(s2)

	//copyでスライスをコピー
	//返り血は要素数
	s_cp := make([]int, len(s3))
	s4 := copy(s_cp, s3)
	fmt.Println(s4)
	/*スライス終わり*/

	/*マップ*/
	//ハッシュ、連想配列、オブジェクト
	//mapを宣言する
	m1 := make(map[string]int)
	//mapに値を設定
	m1["yamada"] = 120
	m1["iwaki"] = 110
	
	//値を指定しながらmapを宣言する
	m2 := map[string]string{"yamada": "120kg", "iwaki": "110kg"}
	fmt.Println(m1)
	fmt.Println(m2)

	//キーの存在を調べる
	vv, vvv := m2["iwaki"]
	fmt.Println(vv, vvv) //"120kg, true"

	vvvv, vvvvv := m2["satonaka"]
	fmt.Println(vvvv, vvvvv) //false

	//要素を削除する
	delete(m2, "yamada")
	fmt.Println(m2)	//map["iwaki": "110kg"]

	/*マップ終わり*/

	/*条件分岐開始*/
	//if
	//golangの特徴として、if文内でしか使わない変数はif文の中で定義できる
	//普通の条件分岐
	score := 61
	
	if score > 80 {
		fmt.Println("grate")
	} else if score <= 80 && score > 60 {
		fmt.Println("good")
	} else {
		fmt.Println("soso")
	}

	//golang特有の条件分岐
	if score2 := 100; score2 > 80 {
		fmt.Println("grate")
	} else if score2 <= 80 && score2 > 60 {
		fmt.Println("good")
	} else {
		fmt.Println("soso")
	}

	//switch
	sw_a := "yellow"

	switch sw_a {
	case "red":
		fmt.Println("stop")
	case "yellow":
		fmt.Println("caution")
	default: 
		fmt.Println("GO")
	} 

	/*条件分岐終わり*/

	/*繰り返し開始*/
	//for文
	//for文しかない→while文はない
	
	//普通の繰り返し
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 8 {
			break
		}
		fmt.Println(i)
	}

	//while文的にも作れる
	jouken := 0
	for jouken < 10 {
		fmt.Println(jouken)
		jouken++
	}

	//range
	//配列、スライス、mapの要素分だけ処理を繰り返す
	//スライス
	ra_s := []string {"one", "two", "three"}

	for key, value := range ra_s {
		fmt.Println(key, value)
	}

	//ブランク修飾子
	// _にしておくと何が入ってもそれを廃棄してくれる
	for _, value2 := range ra_s {
		fmt.Println(value2)
	}	

	//マップ
	ra_map := map[string]string{"yamada": "tarou", "iwaki": "masami", "satonaka": "satoru"}

	for key3, value3 := range ra_map {
		fmt.Println(key3, value3)
	}
	/*繰り返し終わり*/

	/*構造体開始*/
	//■構造体の定理とフィールド
	//複数の値を意味のあるまとまりとして新しい型を定義することができる
	
	//構造体を定義
	type user struct {
		name string
		age int
		sex int
		address string	
	}

	//newでuser構造体分の領域を確保して、初期化して、そのポインタを返す
	u1 := new(user)
	
	u1.name = "nobita"
	u1.age = 12
	u1.sex = 1
	u1.address = "saitama"

	fmt.Println(u1)

	//ポインタが返ってきているので、下記の書き方でもOK
	u2 := new(user)
	
	(*u2).name = "sizuka"
	(*u2).age = 12
	(*u2).sex = 2
	(*u2).address = "tokyo"

	fmt.Println(u2)

	//初期化するときに直接値を与えることも可能
	//こちらの場合はポイントでない構造体のデータが入る
	u3 := user{name: "doraemon", age: 9999, sex: 0, address: "21seiki"}

	fmt.Println(u3)


	//メソッド
	//golangのメソッドは構造体などのデータ型に紐付いた関数になっていて、データの型の定義の中ではなく、データの型とは別に書く
	//構造体とメソッド(関数)の紐付けにはレシーバを使う
	/*構造体終わり*/

	
}
