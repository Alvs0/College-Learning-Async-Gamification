const delay = ms => new Promise(res => setTimeout(res, ms));
document.addEventListener("visibilitychange", function () {
    if (document.hidden === true) {
        points = points - 5;
        document.title = "Im away";
    } else {
        document.title = "Im here";
    }
})

var player, isPlaying = false;

function onYouTubeIframeAPIReady() {
    player = new YT.Player('video', {
        height: '600',
        width: '800',
        videoId: videoId,
        events: {
            onReady: onPlayerReady,
            onStateChange: onPlayerStateChange
        }
    });
}

function onPlayerReady(event) {
    event.target.playVideo();
}

function onPlayerStateChange(event) {
    if (event.data == YT.PlayerState.PLAYING) {
        isPlaying = true;
        startSession();
    } else if (event.data == YT.PlayerState.PAUSED) {
        isPlaying = false;
    } else if (event.data === 0) {
        isPlaying = false;
        endSession();
    }
}

var points = 0;

async function startSession() {
    var randomImage = new Array();

    randomImage[0] = "https://storage.googleapis.com/college-async-gamification/shark.png";
    randomImage[1] = "https://storage.googleapis.com/college-async-gamification/turtle.png";
    randomImage[2] = "https://storage.googleapis.com/college-async-gamification/x-ray_frame.png";
    randomImage[3] = "https://storage.googleapis.com/college-async-gamification/octopus.png";
    randomImage[4] = "https://storage.googleapis.com/college-async-gamification/sea_horse.png";
    randomImage[5] = "https://storage.googleapis.com/college-async-gamification/bird.png";

    while (isPlaying === true) {
        randomDelay = randomizer(10000, 20000);
        await delay(randomDelay);
        imageContainer = document.getElementById("imageContainer");
        var number = Math.floor(Math.random() * randomImage.length);

        const randImg = document.createElement('img');
        randImg.src = randomImage[number]
        randImg.className = "randomImage"
        randImg.onclick = function () {
            imageContainer.removeChild(randImg)
            points += 2;
        }

        var randX = randomizer(0, 50);
        var randY = randomizer(0, 50);

        randImg.style.marginLeft = randX + "%";
        randImg.style.marginTop = randY + "%";

        imageContainer.appendChild(randImg);
        points = points - 1;
        setTimeout(() => imageContainer.removeChild(randImg), 2000);
        await delay(5000);
    }
}

function randomizer(a, b) {
    return Math.floor(Math.random() * b) + a;
}

async function endSession() {
    if (points < 0) {
        return;
    }

    let bodyData = {
        points: points,
    };

    const res = await fetch("/session_add_point", {
        method: "POST",
        body: JSON.stringify(bodyData)
    })

    const response = await res.text();

    let data = JSON.parse(response);
}