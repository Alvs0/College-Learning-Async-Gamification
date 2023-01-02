function generateUL() {
    var studentObj = JSON.parse(studentData);
    var studentListDiv = document.createElement('div');
    studentListDiv.className = 'students students-table';

    for (var idx = 0; idx < studentObj.length; idx++) {
        var studentDiv = document.createElement('div');
        studentDiv.className = 'student';
        studentDiv.style = "cursor:pointer;";
        let id = studentObj[idx].id;
        studentDiv.onclick = function () {
            redirectToRewardDetail(id);
        };

        var studentImgDiv = document.createElement('div');
        studentImgDiv.className = 'student-img';

        var studentIcon = document.createElement('img');

        if (studentObj[idx].profileImageUrl) {
            studentIcon.src = studentObj[idx].profileImageUrl;
        } else {
            studentIcon.src = "https://storage.googleapis.com/college-async-gamification/not_found.png"
        }

        studentImgDiv.appendChild(studentIcon);

        var studentContentDiv = document.createElement('div');
        studentContentDiv.className = 'student-content';

        var studentName = document.createElement('h3');
        studentName.innerText = studentObj[idx].name;

        studentContentDiv.appendChild(studentName);

        studentDiv.appendChild(studentImgDiv);
        studentDiv.appendChild(studentContentDiv);

        studentListDiv.appendChild(studentDiv);
    }

    return studentListDiv;
}

function redirectToRewardDetail(id) {
    var baseUrl = window.location.origin;
    window.location.href = baseUrl + "/student_detail_a/" + id;
}

document.getElementById('students').appendChild(generateUL());