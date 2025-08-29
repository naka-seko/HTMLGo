function toggleDarkMode() {
    const isDarkMode = document.body.classList.toggle("dark-mode");
    localStorage.setItem("darkMode", isDarkMode ? "enabled" : "disabled");
    searchBox.focus();
}

window.onload = function () {
    if (localStorage.getItem("darkMode") === "enabled") {
        document.body.classList.add("dark-mode");
    }
    displayHistory();
    searchBox.focus();
};

document.getElementById("meaningBox").addEventListener("focus", () => {
    meaningBox.lang = 'ja';
});
document.getElementById("meaningBox").addEventListener("blur", () => {
    meaningBox.lang = 'en';
});

async function requestToPHP(action, word, meaning = "") {
    const output = document.getElementById("output");
    output.textContent = "🍄 データを処理中です。少々お待ちください...";
    try {
        const res = await fetch("api.php", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ action, word, meaning })
        });
        if (!res.ok) {
            throw new Error(`サーバーエラー: ${res.status}`);
        }
        const data = await res.json();
        output.textContent = data.message;
        return data;
    } catch (err) {
        output.textContent = `エラーが発生しました: ${err.message}`;
        console.error("エラー:", err);
        throw Error;
    }
}

document.getElementById("searchBox").addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
        searchWord();
    }
});

document.getElementById("meaningBox").addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
        saveWord();
    }
});

function addToHistory(word) {
    if (!word) return;
    let history = JSON.parse(localStorage.getItem("history")) || [];
    if (!history.includes(word)) {
        history.push(word);
        localStorage.setItem("history", JSON.stringify(history));
    }
}

function displayHistory() {
    const historyItems = document.getElementById("historyItems");
    const history = JSON.parse(localStorage.getItem("history")) || [];
    historyItems.innerHTML = "";
    history.forEach(word => {
        const listItem = document.createElement("div");
        listItem.textContent = word;
        listItem.className = "history-item";
        historyItems.appendChild(listItem);
    });
}

function clearHistory() {
    localStorage.removeItem("history");
    displayHistory();
    searchBox.focus();
}

async function searchWord() {
    const searchBox = document.getElementById("searchBox");
    const meaningBox = document.getElementById("meaningBox");
    const word = searchBox.value.trim();
    if (!word) {
        const output = document.getElementById("output");
        output.textContent = "必要な単語を入力して下さい。";
        searchBox.focus();
        return;
    }
    try {
        const response = await requestToPHP("search", word, "");
        if (response.status === "success") {
            document.getElementById("output").textContent = response.message;
            addToHistory(word);
            displayHistory();
            searchBox.focus();
        } else {
            document.getElementById("output").textContent = response.message;
            meaningBox.focus();
        }
    } catch (err) {
        output.textContent = `エラーが発生しました: ${err.message}`;
        console.error("PHPリクエスト中にエラーが発生しました:", err);
        searchBox.focus();
    }
}

function deleteWord() {
    const searchBox = document.getElementById("searchBox");
    const word = searchBox.value.trim();
    if (!word) {
        const output = document.getElementById("output");
        output.textContent = "必要な単語を入力して下さい。";
    } else {
        requestToPHP("delete", word, "");
        searchBox.value = '';
    }
    searchBox.focus();
}

function saveWord() {
    const searchBox = document.getElementById("searchBox");
    const word = searchBox.value.trim();
    const meaningBox = document.getElementById("meaningBox");
    const meaning = meaningBox.value.trim();
    if (!word || !meaning) {
        const output = document.getElementById("output");
        output.textContent = "保存する英単語と日本語を入力して下さい。";
    } else {
        requestToPHP("save", word, meaning);
        meaningBox.value = '';
    }
    searchBox.focus();
}
