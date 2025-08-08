package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"encoding/json"
	"net/http"
)

const BOARD_SIZE = 5

func generatePosition(size int) (int, int) {
	return rand.Intn(size), rand.Intn(size)
}

func calcDistance(pos1, pos2 [2]int) float64 {
	diffX := float64(pos1[0] - pos2[0])
	diffY := float64(pos1[1] - pos2[1])
	return math.Sqrt(diffX*diffX + diffY*diffY)
}

// If you want to use the web server part, define GameState and gameHandler like this:

type GameState struct {
	PlayerPos [2]int   `json:"playerPos"`
	SuikaPos  [2]int   `json:"suikaPos"`
	Distance  float64  `json:"distance"`
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	suikaPos := [2]int{}
	playerPos := [2]int{}
	for {
		suikaPos[0], suikaPos[1] = generatePosition(BOARD_SIZE)
		playerPos[0], playerPos[1] = generatePosition(BOARD_SIZE)
		if suikaPos != playerPos {
			break
		}
	}
	distance := calcDistance(suikaPos, playerPos)
	state := GameState{
		PlayerPos: playerPos,
		SuikaPos:  suikaPos,
		Distance:  distance,
	}
	w.Header().Set("Content-Type", "application/json")
	if playerPos == suikaPos {
		fmt.Fprintln(w, "スイカを割りました！")
	} else {
		json.NewEncoder(w).Encode(state)
	}
}

func main() {
	// 静的ファイルサーバー（index.html, suika_wari.jsなど）
	http.Handle("/", http.FileServer(http.Dir(".")))

    // サーバーのポートを8080に設定
	fmt.Println("サーバーがhttp://localhost:8080で起動中...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("サーバー起動エラー: %v\n", err)
	}
}
