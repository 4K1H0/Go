package main

import "fmt"

func main(){

  // int型からmyInteger型を宣言
  type myInteger int

  // myInteger型の変数を作成
  var i myInteger = 123

  // int型と同じように演算が可能
  i += 1

  // 出力
  fmt.Println(i)

  // 新しい構造体の作成
  type myStruct struct {
    int_a int
    int_b int
  }
}
