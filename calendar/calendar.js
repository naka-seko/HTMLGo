/**
 * カレンダー作成スクリプト
 */
document.addEventListener('DOMContentLoaded', () => {
    const yearSelect = document.getElementById('year');
    const monthSelect = document.getElementById('month');
    const createButton = document.getElementById("generate-calendar");
    const holidaysApiUrl = "/holidays.json";
    const minYear = 1950;
    const maxYear = 2100;    // 静的ファイル配信を追加
    const currentYear = new Date().getFullYear();
    const currentMonth = new Date().getMonth() + 1;

    // 年プルダウン生成
    for (let y = minYear; y <= maxYear; y++) {
        const option = document.createElement('option');
        option.value = y;
        option.textContent = y;
        if (y === currentYear) option.selected = true;
        yearSelect.appendChild(option);
    }
    // 月プルダウン生成
    for (let m = 1; m <= 12; m++) {
        const option = document.createElement('option');
        option.value = m;
        option.textContent = m.toString().padStart(2, '0');
        if (m === currentMonth) option.selected = true;
        monthSelect.appendChild(option);
    }

    if (!createButton) return;

    createButton.addEventListener("click", async (event) => {
        event.preventDefault();
        const year = Number(yearSelect.value);
        const month = Number(monthSelect.value);
        try {
            const calendarData = await fetchCalendarData(year, month);
            renderCalendar(calendarData.calendar, year, month);
            const holidays = await fetchHolidaysData(holidaysApiUrl, year, month);
            colorHolidays(holidays);
        } catch (error) {
            console.error("エラー発生:", error);
        }
    });
});

/**
 * APIからカレンダーデータを取得
 */
async function fetchCalendarData(year, month) {
    const response = await fetch("/generate_calendar", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ year, month }),
    });
    if (!response.ok) throw new Error("APIエラー: " + response.statusText);
    return await response.json();
}

/**
 * カレンダーをHTMLテーブルとして描画
 */
function renderCalendar(calendar, year, month) {
    const table = document.createElement("table");
    table.style.borderCollapse = "collapse";

    // 曜日行
    const daysOfWeek = ["日", "月", "火", "水", "木", "金", "土"];
    const headerRow = document.createElement("tr");
    daysOfWeek.forEach(day => {
        const th = document.createElement("th");
        th.textContent = day;
        th.style.border = "1px solid #ddd";
        th.style.padding = "8px";
        headerRow.appendChild(th);
    });
    table.appendChild(headerRow);

    // 日付行
    calendar.forEach(week => {
        const row = document.createElement("tr");
        week.forEach(day => {
            const cell = document.createElement("td");
            cell.textContent = day || "";
            cell.className = "calendar-cell";
            cell.style.border = "1px solid #ddd";
            row.appendChild(cell);
        });
        table.appendChild(row);
    });

    const outputSection = document.getElementById("output-section");
    outputSection.innerHTML = "";
    outputSection.appendChild(table);
}

/**
 * 祝日データを取得
 */
async function fetchHolidaysData(holidaysApiUrl, year, month) {
    const response = await fetch(holidaysApiUrl);
    if (!response.ok) throw new Error("祝日データの取得に失敗しました");
    const data = await response.json();
    return data[String(year)]?.[String(month)] || [];
}

/**
 * カレンダーセルに祝日色を適用
 */
function colorHolidays(holidays) {
    const calendarCells = document.querySelectorAll(".calendar-cell");
    calendarCells.forEach(cell => {
        const day = parseInt(cell.textContent);
        if (!isNaN(day) && holidays.map(h => Number(h)).includes(day)) {
            cell.style.color = "red";
            cell.style.fontWeight = "bold";
            cell.style.backgroundColor = "#fdd";
            cell.title = "祝日";
        }
    });
}