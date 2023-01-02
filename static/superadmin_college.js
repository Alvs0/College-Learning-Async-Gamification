function generateUL() {
    var collegeObj = JSON.parse(collegeData);
    var collegeListDiv = document.createElement('div');
    collegeListDiv.className = 'colleges colleges-table';

    for(var idx = 0; idx < collegeObj.length; idx++) {
        var collegeDiv = document.createElement('div');
        collegeDiv.className = 'college';
        collegeDiv.style = "cursor:pointer;";
        let id = collegeObj[idx].id;
        collegeDiv.onclick = function() {redirectToCollegeDetail(id);};

        var collegeImgDiv = document.createElement('div');
        collegeImgDiv.className = 'college-img';

        var collegeIcon = document.createElement('img');

        if (collegeObj[idx].imageUrl) {
            collegeIcon.src = collegeObj[idx].imageUrl;
        } else {
            collegeIcon.src = "https://storage.googleapis.com/college-async-gamification/not_found.png"
        }

        collegeImgDiv.appendChild(collegeIcon);

        var collegeContentDiv = document.createElement('div');
        collegeContentDiv.className = 'college-content';

        var collegeName = document.createElement('h3');
        collegeName.innerText = collegeObj[idx].name;

        collegeContentDiv.appendChild(collegeName);

        collegeDiv.appendChild(collegeImgDiv);
        collegeDiv.appendChild(collegeContentDiv);

        collegeListDiv.appendChild(collegeDiv);
    }

    return collegeListDiv;
}

function redirectToCollegeDetail(id) {
    var baseUrl = window.location.origin;
    window.location.href = baseUrl+"/college_detail_sa/"+ id;
}

document.getElementById('colleges').appendChild(generateUL());