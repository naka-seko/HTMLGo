let playerPos, suikaPos;

// ゲーム開始
function startGame() {
    document.getElementById('startBtn').style.display = "none";

    document.getElementById('game-area').style.display = "";
    document.getElementById('user-input-area').style.display = "block";
    document.getElementById('user-input').style.display = "inline-block";
    document.getElementById('submitGuess').style.display = "inline-block";
    document.getElementById('game-log').innerHTML = "";
    enableButtons(true);
    document.getElementById('btn-restart').style.display = "none";
}

function submitGuess() {
    const input = document.getElementById('user-input');
    const value = Number(input.value);
    if (!value || value < 1 || value > 100) {
        showResult("1から100の数字を入力してください。");
        return;
    }
    fetch('/guess', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ guess: value })
    })
    .then(res => res.json())
    .then(data => {
        showResult(data.result);
        if (data.result.includes("ございます")) {
            document.getElementById('restartBtn').style.display = "";
            document.getElementById('guessBtn').disabled = true;
            document.getElementById('guessInput').disabled = true;
        }
    })
    .catch(() => showResult("通信エラー"));
}

function showResult(msg) {
    document.getElementById('result').textContent = msg;
}

// ゲーム再スタート
function restartGame() {
    document.getElementById('result').textContent = "";
    document.getElementById('guessInput').value = "";
    document.getElementById('restartBtn').style.display = "none";
    document.getElementById('guessBtn').disabled = false;
    document.getElementById('guessInput').disabled = false;
}

// ログを自動スクロールする関数
function autoScrollLog() {
    const logContainer = document.querySelector('.log-container'); // クラスセレクターを使用
    if (logContainer) { // logContainerがnullでない場合に処理を実行
        logContainer.scrollTop = logContainer.scrollHeight;
    }
}
    // ログを自動スクロール
    //autoScrollLog();
