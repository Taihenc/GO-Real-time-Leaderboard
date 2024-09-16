function getUsername() {
    const token = localStorage.getItem('token');
    if (!token) {
        return null;
    }
    const payload = token.split('.')[1];
    const base64 = payload.replace('-', '+').replace('_', '/');
    const user = JSON.parse(window.atob(base64));
    return user.username;
}

function register() {
    const username = document.querySelector('#register-username').value;
    const password = document.querySelector('#register-password').value;
    const passwordConfirm = document.querySelector('#register-password-confirm').value;

    if (password !== passwordConfirm) {
        alert('Passwords do not match');
        return;
    }

    fetch('/auth/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
    })
        .then((res) => {
            responseHandler(res);
        })
}

function login() {
    const username = document.querySelector('#login-username').value;
    const password = document.querySelector('#login-password').value;

    fetch('/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
    })
        .then((res) => {
            responseHandler(res);
        })
}

function responseHandler(res) {
    if (res.status === 200) {
        res.body.getReader().read().then(({ value }) => {
            const token = new TextDecoder().decode(value);
            localStorage.setItem('token', token);
            window.location.reload();
        });
    } else {
        res.text().then((text) => alert(text));
    }
}