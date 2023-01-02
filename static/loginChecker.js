window.onload = checkUser;
function checkUser() {
    var base_url = window.location.origin;
    if (!sessionStorage.getItem("user")) {
        window.location.href = base_url+"/login";
    }
}