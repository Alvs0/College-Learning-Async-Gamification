const alertContainer = document.getElementById("alert-box");
const alertMsg = document.getElementById("alert");
const submitBtn = document.getElementById('submit-btn');

const delay = ms => new Promise(res => setTimeout(res, ms));

submitBtn.addEventListener('click', function (event) {
    event.preventDefault();
    addCollege();
})

async function addCollege() {
    name = document.getElementById("name").value;
    address = document.getElementById("address").value;
    image = document.getElementById("collegeIcon");

    if (!image.files[0]) {
        alertBox("Please upload an image to use as College Icon");

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
        address: address,
        imageUrl: imageData["imageUrl"],
    };

    const res = await fetch("/college_sa/add", {
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