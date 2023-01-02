function generateUL() {
    var adminObj = JSON.parse(adminData);
    var adminListDiv = document.createElement('div');
    adminListDiv.className = 'admins admins-table';

    for (var idx = 0; idx < adminObj.length; idx++) {
        var adminDiv = document.createElement('div');
        adminDiv.className = 'admin';
        adminDiv.style = "cursor:pointer;";
        let id = adminObj[idx].id;
        adminDiv.onclick = function () {
            redirectToAdminDetail(id);
        };

        var adminImgDiv = document.createElement('div');
        adminImgDiv.className = 'admin-img';

        var adminIcon = document.createElement('img');

        if (adminObj[idx].profileImageUrl) {
            adminIcon.src = adminObj[idx].profileImageUrl;
        } else {
            adminIcon.src = "https://storage.googleapis.com/college-async-gamification/not_found.png"
        }

        adminImgDiv.appendChild(adminIcon);

        var adminContentDiv = document.createElement('div');
        adminContentDiv.className = 'admin-content';

        var adminName = document.createElement('h3');
        adminName.innerText = adminObj[idx].name;

        var adminEmail = document.createElement('h5');
        adminEmail.innerText = adminObj[idx].email;

        var adminPhoneNumber = document.createElement('h5');
        adminPhoneNumber.innerText = adminObj[idx].phoneNumber;

        var adminBirthDate = document.createElement('h5');
        adminBirthDate.innerText = adminObj[idx].birthDate;

        adminContentDiv.appendChild(adminName);
        adminContentDiv.appendChild(adminEmail);
        adminContentDiv.appendChild(adminPhoneNumber);
        adminContentDiv.appendChild(adminBirthDate);

        adminDiv.appendChild(adminImgDiv);
        adminDiv.appendChild(adminContentDiv);

        adminListDiv.appendChild(adminDiv);
    }

    return adminListDiv;
}

function redirectToAdminDetail(id) {
    var baseUrl = window.location.origin;
    window.location.href = baseUrl + "/admin_detail_sa/" + id;
}

document.getElementById('admins').appendChild(generateUL());