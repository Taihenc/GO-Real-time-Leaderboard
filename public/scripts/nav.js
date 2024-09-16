const login_btn = document.querySelector('#login-btn');
const register_btn = document.querySelector('#register-btn');

const login_dropdown = document.querySelector('#login-dropdown');
const register_dropdown = document.querySelector('#register-dropdown');

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