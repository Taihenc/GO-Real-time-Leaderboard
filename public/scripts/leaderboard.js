const player_input = document.querySelector('#player-input');
const top3_div = document.querySelector('#top-3');
const leaderboard_table = document.querySelector('#leaderboard-table');

var players = [];

// flow: get username -> get game list -> get game in query(if there is) -> get player in leaderboard
(async () => {
    if (getUsername()) {
        player_input.innerHTML += `
        <div class="flex w-[15rem]">
            <span
                class="inline-flex items-center px-3 text-sm text-gray-900 bg-gray-200 border rounded-e-0 border-gray-300 border-e-0 rounded-s-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                <img src='./assets/star.png' class='w-4 h-4'>
            </span>
            <input type="text" id="score"
                class="rounded-none rounded-e-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Score">
        </div>
        <button type="submit"
            class="text-white end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            Submit
        </button>
        `;
    }
    await getGameList();
    setSelectedGameFromQuery();
    addEventListenersForGameSelect();
    getPlayerInLeaderboard();
})();

async function getGameList() {
    const gameSelect = document.querySelector('#game-select');

    await fetch('/gamelist')
        .then((res) => res.json())
        .then((data) => {
            if (!data || data.length === 0) {
                // alert('No data found for this game');
                return;
            }
            for (let i = 0; i < data.length; i++) {
                if (i === 0) {
                    gameSelect.innerHTML = "";
                    gameSelect.innerHTML += `<option value="${data[i]}" selected>${data[i]}</option>`;
                    continue;
                }
                gameSelect.innerHTML += `<option value="${data[i]}">${data[i]}</option>`;
            }
        });
}

function setSelectedGameFromQuery() {
    const urlParams = new URLSearchParams(window.location.search);
    const selectedGame = urlParams.get('game');

    const gameSelect = document.querySelector('#game-select');
    for (let i = 0; i < gameSelect.options.length; i++) {
        if (gameSelect.options[i].value === selectedGame) {
            gameSelect.selectedIndex = i;
            break;
        }
    }
}

function addEventListenersForGameSelect() {
    const gameSelect = document.querySelector('#game-select');
    gameSelect.addEventListener('change', () => {
        window.history.pushState({}, '', '/?game=' + gameSelect.value);
        getPlayerInLeaderboard(gameSelect.value);
    });
}

function submitScore() {
    const gameSelect = document.querySelector('#game-select');
    const score = document.getElementById('score').value;
    const username = getUsername();

    fetch('/submit-score', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "Game": gameSelect.value,
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

function getPlayerInLeaderboard(game) {
    fetch('/leaderboard' + '?game=' + document.querySelector('#game-select').value)
        .then((res) => res.json())
        .then((data) => {
            if (!data || data.length === 0) {
                updateLeaderboard([]);
                return;
            }
            let players = []
            for (let i = 0; i < data.length; i++) {
                players.push({
                    name: data[i].PlayerName,
                    score: data[i].Score,
                    place: i + 1
                });
            }
            updateLeaderboard(players);
        });
}

/**
 * 
 * @param {*} player {name: string, score: number}
 * @param {*} place {1, 2, 3}
 * @returns 
 */
function top3ToHTML(player) {
    return `
    <div class="${placeToPt[player.place]} ${placeToOrer[player.place]}">
        <div class="relative flex justify-center flex-col items-center">
            <div class="flex z-[10] flex-col items-center justify-center sm:translate-y-7 translate-y-4">
                <img src='${crownArray[player.place - 1]}' alt='' />
                <img class="rounded-full sm:w-[108px] sm:h-[108px] w-[50px] h-[50px] mask mask-circle"
                    src="${"https://api.dicebear.com/9.x/micah/svg?seed=" + player.name}"
                    alt="Profile" style="border: 3px solid ${placeToColor[player.place]};">
            </div>
            <div
                class="bg-[#252525]  rounded-[12px] text-center gap-x-[8px] max-sm:h-[82px] sm:h-[117px] sm:w-[110px]  max-sm:w-[55px] flex flex-col items-center justify-center font-semibold tracking-[0.42px] sm:text-[14px] text-[10px]">
                <h2 class="text-[#fff] line-clamp-1">${player.name}</h2>
                <h2 style="color: ${placeToColor[player.place]};">score: ${player.score}</h2>
                <h2 class="absolute max-sm:bottom-2 bottom-3" style="color: ${placeToColor[player.place]};">#${player.place}</h2>
            </div>
        </div>
    </div>
    `
}

function playerToTrHTML(player) {
    return `
    <tr
        class="odd:bg-gray-900  even:bg-gray-800 border-b border-gray-700">
        <th scope="row" class="px-6 py-4 font-medium whitespace-nowrap text-white">
            ${player.place}
        </th>
        <td class="px-6 py-4 text-white">
            ${player.score}
        </td>
        <th scope="row" class="flex items-center px-6 py-4 whitespace-nowrap text-white">
            <img class="w-10 h-10 rounded-full border"
                src="${"https://api.dicebear.com/9.x/micah/svg?seed=" + player.name}" alt="Profile">
            <div class="ml-5 ps-3">
                <div class="text-base font-semibold">${player.name}</div>
            </div>
        </th>
    </tr>
    `
}

function updateLeaderboard(players) {
    top3_div.innerHTML = "";
    leaderboard_table.innerHTML = "";
    for (let i = 0; i < players.length; i++) {
        if (i < 3) {
            top3_div.innerHTML += top3ToHTML(players[i]);
        } else {
            leaderboard_table.innerHTML += playerToTrHTML(players[i]);
        }
    }
}