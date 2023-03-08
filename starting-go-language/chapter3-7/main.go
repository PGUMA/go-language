package main

import "fmt"

func main() {
	// 符号なし整数として扱いたい場合は整数リテラルからの型推論では不可能
	i := 1
	ui := uint32(i)
	fmt.Println(ui)

	// 数値系の表現を超えた演算に注意
	// オーバフローした場合はエラーとならずにオーバーラップ（循環）してしまう
	a := 256
	aa := uint8(a)
	b := aa + uint8(1)
	fmt.Println(b)

	// rune型（ユニコードのポイントを表現した型）
	rp := ' '
	fmt.Println(rp)

	// raw文字列型
	raw := `
	\n
	`
	fmt.Println(raw)
}
