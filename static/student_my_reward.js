const alertContainer = document.getElementById("alert-box");
const alertMsg = document.getElementById("alert");

const delay = ms => new Promise(res => setTimeout(res, ms));
function generateUL() {
    var rewardObj = JSON.parse(rewardData);
    var rewardListDiv = document.createElement('div');
    rewardListDiv.className = 'rewards rewards-table';

    for (var idx = 0; idx < rewardObj.length; idx++) {
        var rewardDiv = document.createElement('div');
        rewardDiv.className = 'reward';

        var rewardImgDiv = document.createElement('div');
        rewardImgDiv.className = 'reward-img';

        var rewardIcon = document.createElement('img');

        if (rewardObj[idx].imageUrl) {
            rewardIcon.src = rewardObj[idx].imageUrl;
        } else {
            rewardIcon.src = "https://storage.googleapis.com/college-async-gamification/not_found.png"
        }

        rewardImgDiv.appendChild(rewardIcon);

        var rewardContentDiv = document.createElement('div');
        rewardContentDiv.className = 'reward-content';

        var rewardName = document.createElement('h3');
        rewardName.innerText = rewardObj[idx].name;

        var rewardQuantity = document.createElement('h4');
        rewardQuantity.innerText = 'Quantity: ' + rewardObj[idx].quantity;


        var useButton = document.createElement('button')
        useButton.type = "button";
        useButton.className = "custom-btn btn";
        let id = rewardObj[idx].id;
        useButton.onclick = function () {
            useReward(id);
        }
        useButton.innerHTML = 'Use Reward';

        rewardContentDiv.appendChild(rewardName);
        rewardContentDiv.appendChild(rewardQuantity);
        rewardContentDiv.appendChild(useButton);


        rewardDiv.appendChild(rewardImgDiv);
        rewardDiv.appendChild(rewardContentDiv);

        rewardListDiv.appendChild(rewardDiv);
    }

    return rewardListDiv;
}

async function useReward(id) {
    let bodyData = {
        rewardId: id,
    };

    const res = await fetch("/reward_use_st", {
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

document.getElementById('rewards').appendChild(generateUL());