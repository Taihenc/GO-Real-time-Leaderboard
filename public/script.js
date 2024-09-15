const top3_div = document.querySelector('#top-3');
const leaderboard_table = document.querySelector('#leaderboard-table');

const players = [
    { name: 'John', score: 10, place: 1 },
    { name: 'Doe', score: 9, place: 2 },
    { name: 'Jane', score: 8, place: 3 },
    { name: 'Atom', score: 10, place: 4 },
    { name: 'Mars', score: 9, place: 5 },
    { name: 'Venus', score: 8, place: 6 },
    { name: 'Jupiter', score: 10, place: 7 },
];

const crownArray = [
    './assets/crown_first.svg',
    './assets/crown_second.svg',
    './assets/crown_third.svg'
]

const placeToColor = {
    1: "#FFD700",
    2: "#C0C0C0",
    3: "#CD7F32"
}

const placeToOrer = {
    1: "order-1",
    2: "order-0",
    3: "order-2"
}

const placeToPt = {
    1: "pt-0",
    2: "pt-3",
    3: "pt-5"
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

function updateLeaderboard() {
    for (let i = 0; i < players.length; i++) {
        if (i < 3) {
            top3_div.innerHTML += top3ToHTML(players[i]);
        } else {
            leaderboard_table.innerHTML += playerToTrHTML(players[i]);
        }
    }
}

updateLeaderboard();