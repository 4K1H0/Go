// 論理演算型のコード

// パッケージ、bool.goをmainパッケージにグルーピングする。
package main

// fmtパッケージをインポートする。
import "fmt"

// main関数を定義
func main(){
    
  // bool型の変数を宣言
  var bool_b bool

  // bool型の変数にリテラル定数trueとfalseを代入
  bool_b = true
  bool_b = false

  // bool型の変数に論理演算した結果を代入
  // true or false → true
  bool_b = true || false

  // 出力
  fmt.Println(bool_b)
}
