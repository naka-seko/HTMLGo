package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type RequestData struct {
    Year  int `json:"year"`
    Month int `json:"month"`
}

type ResponseData struct {
    Year     int           `json:"year"`
    Month    int           `json:"month"`
    Calendar [][]string    `json:"calendar"`
    Error    string        `json:"error,omitempty"`
}

func main() {
    http.HandleFunc("/generate_calendar", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusMethodNotAllowed)
            response := ResponseData{Error: "POSTメソッドのみ許可されています。"}
            _ = json.NewEncoder(w).Encode(response)
            return
        }

        defer r.Body.Close()
        var request RequestData
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&request)
        if err != nil || request.Year < 1950 || request.Year > 2099 || request.Month < 1 || request.Month > 12 {
            w.Header().Set("Content-Type", "application/json")
            response := ResponseData{Error: "入力が不正です。"}
            _ = json.NewEncoder(w).Encode(response)
            return
        }

        calendar := generateCalendar(request.Year, request.Month)
        response := ResponseData{Year: request.Year, Month: request.Month, Calendar: calendar}

        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, "サーバーエラー", http.StatusInternalServerError)
        }
    })

    fmt.Println("サーバーがhttp://localhost:8080で起動中...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("サーバー起動エラー: %v\n", err)
    }
}

func generateCalendar(year int, month int) [][]string {
    calendar := [][]string{}
    firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
    daysInMonth := firstDay.AddDate(0, 1, -1).Day()

    week := []string{}
    for i := 0; i < int(firstDay.Weekday()); i++ {
        week = append(week, "")
    }

    for day := 1; day <= daysInMonth; day++ {
        week = append(week, fmt.Sprintf("%d", day))
        if len(week) == 7 {
            calendar = append(calendar, week)
            week = []string{}
        }
    }

    if len(week) > 0 {
        for len(week) < 7 {
            week = append(week, "")
        }
        calendar = append(calendar, week)
    }
    return calendar
}
