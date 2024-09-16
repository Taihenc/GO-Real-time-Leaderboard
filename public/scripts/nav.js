const login_btn = document.querySelector('#login-btn');
const register_btn = document.querySelector('#register-btn');

const login_dropdown = document.querySelector('#login-dropdown');
const register_dropdown = document.querySelector('#register-dropdown');

const login_register_div = document.querySelector('#login-register-div');

login_btn.addEventListener('click', () => {
    if (!login_dropdown.classList.contains('invisible')) {
        login_dropdown.classList.add('invisible');
        login_dropdown.classList.remove('opacity-100');
        return;
    }
    login_dropdown.classList.remove('invisible');
    login_dropdown.classList.add('opacity-100');

    if (!register_dropdown.classList.contains('invisible')) {
        register_dropdown.classList.add('invisible');
        register_dropdown.classList.remove('opacity-100');
    }
});

register_btn.addEventListener('click', () => {
    if (!register_dropdown.classList.contains('invisible')) {
        register_dropdown.classList.add('invisible');
        register_dropdown.classList.remove('opacity-100');
        return;
    }
    register_dropdown.classList.remove('invisible');
    register_dropdown.classList.add('opacity-100');

    if (!login_dropdown.classList.contains('invisible')) {
        login_dropdown.classList.add('invisible');
        login_dropdown.classList.remove('opacity-100');
    }
});

const username = getUsername();
if (username) {

    const capitalize = username.charAt(0).toUpperCase() + username.slice(1);

    login_register_div.innerHTML = `
        <span class="text-white">
            Welcome, ${capitalize}
        </span>
        <img class="w-10 h-10 rounded-full border"
            src="${"https://api.dicebear.com/9.x/micah/svg?seed=" + username}" alt="Profile">
        <button id="logout-btn" onclick="logout()" class="px-6 py-2 text-gray-400 hover:text-gray-700 hover:bg-gray-300 border border-white rounded-lg select-none">Logout</button>
    `;
}