package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

const dictFile = "jisyo_fruit.txt"

var mu sync.Mutex

type Request struct {
	Action  string `json:"action"`
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
}

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api", apiHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("サーバーが http://localhost:8080 で起動中...")
	http.ListenAndServe(":8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POSTのみ許可されています", http.StatusMethodNotAllowed)
		return
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "リクエストが不正です", http.StatusBadRequest)
		return
	}
	word := strings.TrimSpace(req.Word)
	meaning := strings.TrimSpace(req.Meaning)
	mu.Lock()
	defer mu.Unlock()
	dict, _ := loadDictionary(dictFile)
	resp := Response{Message: "何も処理されませんでした。"}
	switch req.Action {
	case "search":
		if dict == nil {
			resp = Response{Status: "error", Message: "辞書データが正しくありません。"}
		} else if val, ok := dict[word]; ok {
			resp = Response{Status: "success", Message: fmt.Sprintf("%s の意味は「%s」です。", word, val)}
		} else {
			resp = Response{Status: "notfound", Message: fmt.Sprintf("お探しの単語 %s は見つかりませんでした。😢", word)}
		}
	case "save":
		if _, ok := dict[word]; ok {
			resp.Message = fmt.Sprintf("%s は辞書に存在します。更新保存しました。", word)
		} else {
			resp.Message = fmt.Sprintf("%s は辞書に存在しません。追加保存しました。", word)
		}
		dict[word] = meaning
		saveDictionary(dictFile, dict)
	case "delete":
		if _, ok := dict[word]; ok {
			delete(dict, word)
			saveDictionary(dictFile, dict)
			resp.Message = fmt.Sprintf("%s を削除しました。", word)
		} else {
			resp.Message = fmt.Sprintf("%s は辞書に存在しません。", word)
		}
	default:
		resp.Message = "不正なアクションです。"
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func loadDictionary(filename string) (map[string]string, error) {
	dict := make(map[string]string)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ",", 2)
		if len(parts) == 2 {
			dict[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return dict, nil
}

func saveDictionary(filename string, dict map[string]string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for k, v := range dict {
		fmt.Fprintf(writer, "%s,%s\n", k, v)
	}
	return writer.Flush()
}
