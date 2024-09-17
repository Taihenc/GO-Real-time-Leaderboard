const submit_div = document.querySelector('#submit-div');

if (getUsername()) {
    submit_div.innerHTML = `
    <form onsubmit='submitScore(); return false'>
        <div id='player-input' class='w-1/2 mx-auto flex justify-center gap-2'>
            <div class="flex">
                <span
                    class="inline-flex items-center px-2 text-sm text-gray-900 bg-gray-200 border rounded-e-0 border-gray-300 border-e-0 rounded-s-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                    <img src='./assets/game.svg' class='w-6 h-6'>
                </span>
                <input type="text" id="game"
                    class="rounded-none rounded-e-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Game">
            </div>
            <div class="flex">
                <span
                    class="inline-flex items-center px-3 text-sm text-gray-900 bg-gray-200 border rounded-e-0 border-gray-300 border-e-0 rounded-s-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                    <img src='./assets/star.png' class='w-4 h-4'>
                </span>
                <input type="text" id="score"
                    class="rounded-none rounded-e-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Score">
            </div>
            <button type="submit"
                class="text-white end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Submit</button>
        </div>
    </form>
    `;
}

function submitScore() {
    const game = document.getElementById('game').value;
    const score = document.getElementById('score').value;
    const username = getUsername();

    fetch('/submit-score', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "Game": game,
            "Playername": username,
            "Score": parseInt(score)
        }),
    })
        .then((response) =>
            leaderboardResponseHandler(response)
        );

}

function leaderboardResponseHandler(res) {
    if (res.status === 200) {
        window.location.reload();
    } else {
        res.text().then((text) => alert(text));
    }
}