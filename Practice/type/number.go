package main

import "fmt"

func main(){

  // int型の変数宣言
  var int_i int = 12345

  // int型からint64bit型へ変換(互換性無しのためキャスト)
  var int64_i int64 = int64(int_i)

  // 出力
  fmt.Println(int64_i)
}
