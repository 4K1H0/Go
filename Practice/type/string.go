package main

import "fmt"

func main(){

  // string型のstr変数を宣言
  var str string

  // string型変数に値を代入
  str = "あ"

  // 文字列を「+」演算子で結合
  str = "あ" + "い"

  // 文字列を「+=」演算子で結合
  str += "う"

  // 結合結果を出力
  fmt.Println(str)
}
