// game/suuji_answer.go
// 数字当てゲーム
// プレイヤーは1から100の間の数字を当てるゲーム。
// プレイヤーは数字を入力し、正解かどうかをフィードバックされる。
// 正解するまで繰り返す。
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("ようこそ。数字当てゲームへ！")
	rand.Seed(time.Now().UnixNano())
	numberToGuess := rand.Intn(100) + 1
	attempts := 0

	for {
		var guess int
		fmt.Print("１から１００の間で入力してね: ")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("数字（整数）を入れてください。")
			continue
		}
		attempts++
		if guess < numberToGuess {
			fmt.Println("小さいです！")
		} else if guess > numberToGuess {
			fmt.Println("大きいです！")
		} else {
			fmt.Printf("🎊ございます！ 入力された回数は %d です。\n", attempts)
			break
		}
	}
}
