let playerPos, suikaPos;

// ゲーム開始
function startGame() {
    document.getElementById('startBtn').style.display = "none";

    document.getElementById('game-area').style.display = "";
    document.getElementById('user-input-area').style.display = "block";
    document.getElementById('user-input').style.display = "inline-block";
    document.getElementById('submitGuess').style.display = "inline-block";
    document.getElementById('game-log').innerHTML = "";
    document.getElementById('btn-restart').style.display = "none";
}

// 数字当てゲーム実行
async function submitGuess() {
    const startBtn = document.getElementById('startBtn');
    const gameArea = document.getElementById('game-area');
    const userInputArea = document.getElementById('user-input-area');
    startBtn.style.display = "none";
    gameArea.style.display = "";
    userInputArea.style.display = "block";
    const input = document.getElementById('user-input');
    const value = Number(input.value);

    if (!value || value < 1 || value > 100) {
        updateGameLog("1から100の数字を入力して下さい。");
        return;
    }
    try {
        const res = await fetch('/guess', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ guess: value })
        });
        const data = await res.json();
        updateGameLog(data.result);
        console.log(data);
        if (data.result.includes("ございます")) {
            document.getElementById('btn-restart').style.display = "";
            document.getElementById('guessBtn').disabled = true;
            document.getElementById('guessInput').disabled = true;
        }
    } catch (e) {
        updateGameLog("通信エラー");
    }
}

// メッセージログ出力
function updateGameLog(message) {
    const gameLog = document.getElementById('game-log');
    gameLog.style.display = "block";  // 表示する
    gameLog.innerHTML += `<p>${message}</p>`;  // 新しいログを追加
    // ログを自動スクロール
    autoScrollLog();
}

// ゲーム再スタート
function restartGame() {
    document.getElementById('guessBtn').disabled = false;
    document.getElementById('guessInput').disabled = false;
    document.getElementById('game-log').textContent = "";
    document.getElementById('guessInput').value = "";
    document.getElementById('restartBtn').style.display = "none";
}
// ログを自動スクロールする関数
function autoScrollLog() {
    const logContainer = document.querySelector('log-container'); // クラスセレクターを使用
    if (logContainer) { // logContainerがnullでない場合に処理を実行
        logContainer.scrollTop = logContainer.scrollHeight;
    }
}
