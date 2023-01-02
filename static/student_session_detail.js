const alertContainer = document.getElementById("alert-box");
const alertMsg = document.getElementById("alert");
const startBtn = document.getElementById('startSessionBtn');

const delay = ms => new Promise(res => setTimeout(res, ms));

startBtn.addEventListener('click', function (event) {
    event.preventDefault();
    initSession();
})

async function initSession() {
    var baseUrl = window.location.origin;
    var url = baseUrl+'/session/'+id;
    var newWindow = window.open(url, '_blank', "width="+screen.availWidth+",height="+screen.availHeight);
    newWindow.focus();
}

const alertSuccess = () => {
    alertMsg.innerHTML = "Success!";
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