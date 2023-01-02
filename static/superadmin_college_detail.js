const alertContainer = document.getElementById("alert-box");
const alertMsg = document.getElementById("alert");
const deleteBtn = document.getElementById('deleteCollegeBtn');

const delay = ms => new Promise(res => setTimeout(res, ms));

deleteBtn.addEventListener('click', function (event) {
    event.preventDefault();
    deleteCollege();
})

async function deleteCollege() {
    let bodyData = {
        id: id,
    };

    const res = await fetch("/college_delete_sa", {
        method: "POST",
        body: JSON.stringify(bodyData)
    })

    const response = await res.text();

    let data = JSON.parse(response);

    if (!data["message"]) {
        alertSuccess();
        await delay(1500);
        window.location.href = data["redirectTo"];
    } else {
        alertBox(data["message"]);
    }
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