package pion

/*
for usb camera
echo $tt | go run main.go -video-src " v4l2src device=/dev/video0 !  jpegdec ! videoconvert ! videoscale ! video/x-raw, width=320, height=240 ! queue"

*/

/*
For Raspi Camera
git clone https://github.com/thaytan/gst-rpicamsrc.git
sudo apt-get install autoconf automake libtool libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev libraspberrypi-dev
./autogen --prefix=/usr --libdir=/usr/lib/arm-linux-gnueabihf/
make
sudo make install
echo $tt | go run main.go -video-src " rpicamsrc !  jpegdec ! videoconvert ! videoscale ! video/x-raw, width=320, height=240 ! queue"
*/
