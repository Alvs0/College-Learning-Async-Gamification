const logoutBtn = document.getElementById('logout');

logoutBtn.addEventListener('click', function (event) {
    event.preventDefault();
    console.log("Clearing sesion storage");
    sessionStorage.clear();
})