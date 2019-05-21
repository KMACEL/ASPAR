package videostreaming

/*
echo $tt | go run main.go -video-src " v4l2src device=/dev/video1 !  jpegdec ! videoconvert ! videoscale ! video/x-raw, width=320, height=240 ! queue"

*/
