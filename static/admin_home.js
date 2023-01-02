let toggler = document.querySelector(".toggler");

window.addEventListener("click", event => {
    if(event.target.className == "toggler" || event.target.className == "toggle") {
        document.body.classList.toggle("show-nav");
    } else if (event.target.className == "overlay") {
        document.body.classList.remove("show-nav");
    }

    if(document.body.className == "show-nav") {
        toggler.innerHTML = "&laquo";
    } else {
        toggler.innerHTML = "&raquo";
    }
});