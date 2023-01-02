const alertContainer = document.getElementById("alert-box");
const alertMsg = document.getElementById("alert");
const submitBtn = document.getElementById('submit-btn');

const delay = ms => new Promise(res => setTimeout(res, ms));

window.onload = clearSession;
function clearSession() {
    window.sessionStorage.clear();
}

submitBtn.addEventListener('click', function (event) {
    event.preventDefault();
    login();
})

async function login() {
    email = document.getElementById("email").value;
    password = document.getElementById("password").value;

    let bodyData = {
        email: email,
        password: password,
        sessionStorage: sessionStorage.getItem("loggedUser"),
    };

    const res = await fetch("/login", {
        method: "POST",
        body: JSON.stringify(bodyData)
    })

    const response = await res.text();

    let data = JSON.parse(response);

    if (!data["message"]) {
        sessionStorage.setItem("user", data["loggedUser"]);
        alertSuccess();
        await delay(1500);
        window.location.href = data["redirectTo"];
    } else {
        alertBox(data["message"]);
    }
}

const alertSuccess = () => {
    alertMsg.innerHTML = "Login success!";
    alertContainer.style.top = `5%`;
    alertContainer.style.background = `rgb(119, 255, 119)`;
    setTimeout(() => {
        alertContainer.style.top = null;
    }, 2000);
}

const alertBox = (data) => {
    alertMsg.innerHTML = data;

    alertContainer.style.top = `5%`;
    alertContainer.style.background = `rgb(255, 119, 119)`;
    setTimeout(() => {
        alertContainer.style.top = null;
    }, 2000);
}