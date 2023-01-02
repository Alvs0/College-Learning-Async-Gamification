var check = function () {
    if (document.getElementById('password').value ===
        document.getElementById('confirm-password').value) {
        document.getElementById('message').style.color = 'green';
        document.getElementById('message').innerHTML = 'matching';
    } else {
        document.getElementById('message').style.color = 'red';
        document.getElementById('message').innerHTML = 'not matching';
    }
}

const alertContainer = document.getElementById("alert-box");
const alertMsg = document.getElementById("alert");
const submitBtn = document.getElementById('submit-btn');

const delay = ms => new Promise(res => setTimeout(res, ms));

submitBtn.addEventListener('click', function (event) {
    event.preventDefault();
    register();
})

async function register() {
    rawBirthDate = new Date(document.getElementById('birthday').value);
    birthDate = rawBirthDate.toString("DD MM YYYY");
    name = document.getElementById("name").value;
    email = document.getElementById("email").value;
    collegeName = document.getElementById("collegeName").value;
    phoneNumber = document.getElementById("phone-number").value;
    password = document.getElementById('password').value
    image = document.getElementById("profileImage");

    if (!image.files[0]) {
        alertBox("Please upload an image to use as Profile Icon");

        return
    }

    let formData = new FormData()
    formData.append('file', image.files[0])

    const imageUploadRes = await fetch("/upload_file", {
        method: "POST",
        body: formData,
    })

    const imageUploadResponse = await imageUploadRes.text();
    let imageData = JSON.parse(imageUploadResponse);

    if (!imageData["imageUrl"]) {
        alertBox("Failed to upload image, Please Try Again");
    }

    let bodyData = {
        name: name,
        email: email,
        phoneNumber: phoneNumber,
        password: password,
        birthDate: birthDate,
        collegeName: collegeName,
        imageUrl: imageData["imageUrl"],
    };

    const res = await fetch("/register", {
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