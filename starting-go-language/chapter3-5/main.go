package main

import "fmt"

func main() {
	// 改行付きの標準出力
	fmt.Println("Hello Golang!!")
	// 複数引数（スペース区切りで出力される）
	fmt.Println("Hello", "Golang!!")
	// フォーマット文字列
	// 書式指定子
	// 数値: %s=文字列 %d=整数 %v=どんな型でもOK
	fmt.Printf("文字列=%s 数値=%d Any=%v \n", "文字", 100, [2]int{1,2})
	// パラメータ不足 MISSING
	fmt.Printf("%s=%s \n", "test")
	// パラメータ過剰 EXTRA
	fmt.Printf("%s=%s \n", "test", "test", "test")
}

