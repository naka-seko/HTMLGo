package main
// IGNORE: このコードはGo言語で書かれています。Goのバージョン1.18以降が必要です。
// IGNORE: このコードは、指定された年と月のカレンダーを生成するAPIサーバーです。

// IGNORE: このコードは、HTTP POSTリクエストを受け取り、指定された年と月のカレンダーを生成して返します。
// IGNORE: また、静的ファイルサーバーとしてindex.htmlやcalendar.jsなどのファイルも提供します。
import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// RequestDataは、クライアントからのリクエストデータを表します。
// Yearは年、Monthは月を表します。
type RequestData struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

// ResponseDataは、APIのレスポンスデータを表します。
// Yearは年、Monthは月、Calendarはカレンダーの2次元配列を表します。
// Errorはエラーメッセージを含む場合があります。
type ResponseData struct {
	Year     int        `json:"year"`
	Month    int        `json:"month"`
	Calendar [][]string `json:"calendar"`
	Error    string     `json:"error,omitempty"`
}

// main関数は、HTTPサーバーを起動し、/generate_calendarエンドポイントを設定します。
func main() {
	http.HandleFunc("/generate_calendar", func(w http.ResponseWriter, r *http.Request) {
		// リクエストメソッドがPOSTでない場合は405 Method Not Allowedを返す
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ResponseData{Error: "POSTメソッドのみ許可されています。"})

			return
		}
		// リクエストボディをデコードして、年と月を取得
		defer r.Body.Close()
		var request RequestData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		// 年と月の値が不正な場合は400 Bad Requestを返す
		// 年は1950年から2099年、月は1から12の範囲でなければならない
		if err != nil || request.Year < 1950 || request.Year > 2099 || request.Month < 1 || request.Month > 12 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ResponseData{Error: "入力が不正です。"})

			return
		}
		// 年と月が有効な場合はカレンダーを生成してレスポンスを返す
		calendar := generateCalendar(request.Year, request.Month)
		response := ResponseData{Year: request.Year, Month: request.Month, Calendar: calendar}
		w.Header().Set("Content-Type", "application/json")
		// レスポンスをJSON形式でエンコードして返す
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "サーバーエラー", http.StatusInternalServerError)
		}
	})

	// 静的ファイルサーバー（index.html, calendar.js, holidays.jsonなど）
	http.Handle("/", http.FileServer(http.Dir(".")))

	// サーバーをポート8081で起動
	fmt.Println("サーバーがhttp://localhost:8081で起動中...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("サーバー起動エラー: %v\n", err)
	}
}

// 指定された年・月のカレンダーを2次元配列で返す
func generateCalendar(year int, month int) [][]string {
	calendar := [][]string{}
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	daysInMonth := firstDay.AddDate(0, 1, -1).Day()

	// カレンダーのヘッダーを追加
	week := []string{}
	for i := 0; i < int(firstDay.Weekday()); i++ {
		week = append(week, "")
	}
	// 月の最初の日から最後の日までループして日付を追加
	for day := 1; day <= daysInMonth; day++ {
		week = append(week, fmt.Sprintf("%d", day))
		if len(week) == 7 {
			calendar = append(calendar, week)
			week = []string{}
		}
	}

	// 最後の週が7日未満の場合は空の文字列で埋める
	// もし最後の週が空でない場合はカレンダーに追加
	if len(week) > 0 {
		for len(week) < 7 {
			week = append(week, "")
		}
		calendar = append(calendar, week)
	}
	return calendar
}

// --- IGNORE ---
