package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	numberToGuess int
	attempts      int
	mu            sync.Mutex
)

type GuessRequest struct {
	Guess int `json:"guess"`
}

type GuessResponse struct {
	Result   string `json:"result"`
	Attempts int    `json:"attempts"`
}

// 数字当てゲームのメイン関数
func main() {
	rand.Seed(time.Now().UnixNano())
	resetGame()

	http.HandleFunc("/guess", guessHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// サーバーのポートを8080に設定
	fmt.Println("サーバーがhttp://localhost:8080で起動中...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("サーバー起動エラー: %v\n", err)
	}
}

func resetGame() {
	mu.Lock()
	defer mu.Unlock()
	numberToGuess = rand.Intn(100) + 1
	attempts = 0
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POSTのみ許可されています", http.StatusMethodNotAllowed)
		return
	}
	var req GuessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "リクエストが不正です", http.StatusBadRequest)
		return
	}
	mu.Lock()
	attempts++
	var result string
	if req.Guess < numberToGuess {
		result = "小さいです！"
	} else if req.Guess > numberToGuess {
		result = "大きいです！"
	} else {
		result = fmt.Sprintf("🎊ございます！ 入力された回数は %d です。", attempts)
		resetGame()
	}
	mu.Unlock()
	resp := GuessResponse{Result: result, Attempts: attempts}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
