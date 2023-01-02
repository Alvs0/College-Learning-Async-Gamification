const alertContainer = document.getElementById("alert-box");
const alertMsg = document.getElementById("alert");
const submitBtn = document.getElementById('submit-btn');

const delay = ms => new Promise(res => setTimeout(res, ms));

submitBtn.addEventListener('click', function (event) {
    event.preventDefault();
    addReward();
})

async function addReward() {
    name = document.getElementById("name").value;
    description = document.getElementById("description").value;
    quantity = document.getElementById("quantity").value;
    minimalLevel = document.getElementById("minimalLevel").value;
    requiredPoints = document.getElementById("requiredPoints").value;
    isActive = document.getElementById("isActive").value;
    image = document.getElementById("rewardImage");

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
        description: description,
        quantity: quantity,
        minimalLevel: minimalLevel,
        requiredPoints: requiredPoints,
        isActive: isActive,
        imageUrl: imageData["imageUrl"],
    };

    const res = await fetch("/reward_a/add", {
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