document.addEventListener('DOMContentLoaded', function () {
    var currentSound = null;
    var currentTrackElement = null;

    document.querySelectorAll('.track').forEach(function (trackElement) {
        var playButton = trackElement.querySelector('.play-button');
        var pauseButton = trackElement.querySelector('.pause-button');
        var progressBar = trackElement.querySelector('.progress');
        var audioSrc = playButton.getAttribute('data-src');

        playButton.addEventListener('click', function () {
            if (currentSound && currentTrackElement !== trackElement) {
                currentSound.stop();
            }

            if (!currentSound || currentTrackElement !== trackElement) {
                currentSound = new Howl({
                    src: [audioSrc],
                    html5: true,
                    onplay: function () {
                        requestAnimationFrame(updateProgress);
                    }
                });
                currentTrackElement = trackElement;
            }

            currentSound.play();
        });

        pauseButton.addEventListener('click', function () {
            if (currentSound) {
                currentSound.pause();
            }
        });

        function updateProgress() {
            var seek = currentSound.seek() || 0;
            progressBar.style.width = (((seek / currentSound.duration()) * 100) || 0) + '%';
            if (currentSound.playing()) {
                requestAnimationFrame(updateProgress);
            }
        }
    });
});
