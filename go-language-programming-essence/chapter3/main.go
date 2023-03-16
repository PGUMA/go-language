package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main () {
	// constによる定数宣言
	const x = 1
	//x = 2 代入不可
	const y = x + 2
	const z = x + 0.999999

	// iota

	type Animal int
	type Fruit int

	const (
		APPLE Fruit = iota
		BANANA
	)

	const (
		MONKEY Animal = iota
		CAW
	)

	const fruits = BANANA
	const animal = MONKEY

	// 型が違うので比較不可としてくれる
	//fmt.Println("fruits = animal result is", banana == monkey)

	// if分
	// カッコ不要
	if user, err := userName(); err == nil {
		fmt.Println(user)
	}

	// switch文
	// breakは不要、次のケースにも渡したい場合はfallthroughキーワードを宣言する
	switch  {
	case fruits == BANANA :
	case fruits == APPLE:
	default:
	}

	// 構造体
	fmt.Printf("%v\n", User {"test", 1})
	userJsonByte, _ := json.Marshal(User {"test", 1})
	// 非公開フィールドはjson化されない点に注意
	fmt.Printf("%v\n", string(userJsonByte))
}

func userName() (string, error) {
	return "", errors.New("Oops")
}

type User struct {
	Name string
	sex int
}
