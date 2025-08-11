let playerPos, suikaPos;

// ゲーム開始
function startGame() {
    document.getElementById('startBtn').style.display = "none";
    document.getElementById('game-area').style.display = "";
    initGame();
}

// ゲーム初期化
function initGame() {
    do {
        playerPos = { x: Math.floor(Math.random() * 5), y: Math.floor(Math.random() * 5) };
        suikaPos = { x: Math.floor(Math.random() * 5), y: Math.floor(Math.random() * 5) };
    } while (playerPos.x === suikaPos.x && playerPos.y === suikaPos.y);
    document.getElementById('game-log').innerHTML = "";
    enableButtons(true);
    document.getElementById('btn-restart').style.display = "none";
}

// プレイヤーとスイカの距離を計算
function calcDistance(pos1, pos2) {
    const diffX = pos1.x - pos2.x;
    const diffY = pos1.y - pos2.y;
    return Math.sqrt(diffX * diffX + diffY * diffY).toFixed(2);
}

// 移動処理
function move(direction) {
    switch (direction) {
        case 'n':
            playerPos.y = Math.max(0, playerPos.y - 1);
            break;
        case 's':
            playerPos.y = Math.min(4, playerPos.y + 1);
            break;
        case 'e':
            playerPos.x = Math.min(4, playerPos.x + 1);
            break;
        case 'w':
            playerPos.x = Math.max(0, playerPos.x - 1);
            break;
    }
    const distance = calcDistance(playerPos, suikaPos);
    const log = document.getElementById('game-log');
    if (playerPos.x === suikaPos.x && playerPos.y === suikaPos.y) {
        log.innerHTML += `<p>スイカを割りました！🎉</p>`;
        enableButtons(false);
        document.getElementById('btn-restart').style.display = "";
    } else {
        log.innerHTML += `<p>現在位置: (${playerPos.x}, ${playerPos.y})</p>`;
        log.innerHTML += `<p>スイカへの距離: ${distance} m</p>`;
    }
    // ログを自動スクロール
    autoScrollLog();
}

// ボタンの有効/無効
function enableButtons(flag) {
    document.getElementById('btn-n').disabled = !flag;
    document.getElementById('btn-s').disabled = !flag;
    document.getElementById('btn-e').disabled = !flag;
    document.getElementById('btn-w').disabled = !flag;
}

// ゲーム再スタート
function restartGame() {
    initGame();
}

// イベント登録
window.onload = function() {
    document.getElementById('btn-n').onclick = () => move('n');
    document.getElementById('btn-s').onclick = () => move('s');
    document.getElementById('btn-e').onclick = () => move('e');
    document.getElementById('btn-w').onclick = () => move('w');
    document.getElementById('btn-restart').onclick = restartGame;
}

// ログを自動スクロールする関数
function autoScrollLog() {
    const logContainer = document.querySelector('.log-container'); // クラスセレクターを使用
    if (logContainer) { // logContainerがnullでない場合に処理を実行
        logContainer.scrollTop = logContainer.scrollHeight;
    }
}
