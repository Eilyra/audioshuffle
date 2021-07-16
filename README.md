# audioshuffle

A simple tool to allow streamers to play music from a directory and adjust levels for themselves as well as the stream separately by simply using the mixer tools built into software like OBS.

## Usage

1. Download the binary for your platform along with the index.html file from the releases (.exe is for Windows, without extension is for Linux)
1. Put them in the same folder
1. Create a folder named "music" within that folder
1. Put the music you want to play within the "music" folder
1. Launch the binary, it creates a web server listening on port 9119
1. Add http://localhost:9119/ as a browser source in OBS or your streaming software of choice
    1. Customize the CSS as you see fit, would recommend at the very least increasing font size
    1. Enable "Control audio via OBS" to be able to control the audio in OBS
1. Enable audio monitoring in the OBS mixer or use a plugin like [Audio Monitor](https://obsproject.com/forum/resources/audio-monitor.1186/) if you would like to hear the music too
