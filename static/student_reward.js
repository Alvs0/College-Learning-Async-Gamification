function generateUL() {
    var rewardObj = JSON.parse(rewardData);
    var rewardListDiv = document.createElement('div');
    rewardListDiv.className = 'rewards rewards-table';

    for (var idx = 0; idx < rewardObj.length; idx++) {
        var rewardDiv = document.createElement('div');
        rewardDiv.className = 'reward';
        rewardDiv.style = "cursor:pointer;";
        let id = rewardObj[idx].id;
        rewardDiv.onclick = function () {
            redirectToRewardDetail(id);
        };

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

        rewardContentDiv.appendChild(rewardName);

        rewardDiv.appendChild(rewardImgDiv);
        rewardDiv.appendChild(rewardContentDiv);

        rewardListDiv.appendChild(rewardDiv);
    }

    return rewardListDiv;
}

function redirectToRewardDetail(id) {
    var baseUrl = window.location.origin;
    window.location.href = baseUrl + "/reward_detail_st/" + id;
}

document.getElementById('rewards').appendChild(generateUL());