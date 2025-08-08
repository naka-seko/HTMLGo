// suika_wari.go
 // スイカ割ゲーム
 // プレイヤーはスイカの位置を知らされず、スイカまでの距離をヒントに移動して行く。
 // スイカの位置はランダムに決定される。
 // プレイヤーは北(n)、南(s)、東(e)、西(w)の方向に移動出来る。
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const BOARD_SIZE = 5

// x座標とy座標の生成
func generatePosition(size int) (int, int) {
	return rand.Intn(size), rand.Intn(size)
}

// ２点間の距離算出
func calcDistance(pos1, pos2 [2]int) float64 {
	diffX := float64(pos1[0] - pos2[0])
	diffY := float64(pos1[1] - pos2[1])
	return math.Sqrt(diffX*diffX + diffY*diffY)
}

// プレイヤーを移動
func movePosition(direction string, pos [2]int) [2]int {
	x, y := pos[0], pos[1]
	switch direction {
	case "n":
		y--
	case "s":
		y++
	case "w":
		x--
	case "e":
		x++
	}
	// ボード外に出ないように制限
	if x < 0 {
		x = 0
	}
	if x >= BOARD_SIZE {
		x = BOARD_SIZE - 1
	}
	if y < 0 {
		y = 0
	}
	if y >= BOARD_SIZE {
		y = BOARD_SIZE - 1
	}
	return [2]int{x, y}
}

// スイカ割ゲーム
func suikaWari() {
	rand.Seed(time.Now().UnixNano())
	suikaPos := [2]int{}
	playerPos := [2]int{}
	suikaPos[0], suikaPos[1] = generatePosition(BOARD_SIZE)
	playerPos[0], playerPos[1] = generatePosition(BOARD_SIZE)

	for suikaPos != playerPos {
		distance := calcDistance(suikaPos, playerPos)
		fmt.Printf("スイカへの距離: %.2f\n", distance)
		fmt.Print("n:北 s:南 e:東 w:西 > ")
		var c string
		fmt.Scanln(&c)
		playerPos = movePosition(c, playerPos)
	}
	fmt.Println("スイカを割りました！")
}

func main() {
	suikaWari()
}
