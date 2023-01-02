function generateUL() {
    var sessionObj = JSON.parse(sessionData);
    var sessionListDiv = document.createElement('div');
    sessionListDiv.className = 'sessions sessions-table';

    for (var idx = 0; idx < sessionObj.length; idx++) {
        var sessionDiv = document.createElement('div');
        sessionDiv.className = 'session';
        sessionDiv.style = "cursor:pointer;";
        let id = sessionObj[idx].id;
        sessionDiv.onclick = function () {
            redirectToRewardDetail(id);
        };

        var sessionImgDiv = document.createElement('div');
        sessionImgDiv.className = 'session-img';

        var sessionIcon = document.createElement('img');

        if (sessionObj[idx].imageUrl) {
            sessionIcon.src = sessionObj[idx].imageUrl;
        } else {
            sessionIcon.src = "https://storage.googleapis.com/college-async-gamification/not_found.png"
        }

        sessionImgDiv.appendChild(sessionIcon);

        var sessionContentDiv = document.createElement('div');
        sessionContentDiv.className = 'session-content';

        var sessionName = document.createElement('h3');
        sessionName.innerText = sessionObj[idx].name;

        sessionContentDiv.appendChild(sessionName);

        sessionDiv.appendChild(sessionImgDiv);
        sessionDiv.appendChild(sessionContentDiv);

        sessionListDiv.appendChild(sessionDiv);
    }

    return sessionListDiv;
}

function redirectToRewardDetail(id) {
    var baseUrl = window.location.origin;
    window.location.href = baseUrl + "/session_detail_a/" + id;
}

document.getElementById('sessions').appendChild(generateUL());