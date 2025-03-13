### Go app to use a magnetic door switch to check if the door is open or closed.
Pico attached to breadboard

### To set up new Go app
`go mod init picoswitch`

### Install dependencies
`go mod tidy`

### Build the app that will run on the Pico
`tinygo build -target=pico -o main.uf2 .`

### Flash/transfer app to the Pico
- plug the pico into the pi with the bootsel button pressed. Once copied, the pico will reboot and run app.
`cp main.uf2 /media/rachel/RPI-RP2`

`tinygo monitor`
# picoswitch
