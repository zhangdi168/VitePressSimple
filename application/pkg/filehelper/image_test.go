package filehelper

import (
	"testing"
	"wails3templatevuets/app/mypkg/mydev"
)

func TestFind(t *testing.T) {
	mydev.StartTrack()
	imgPath := FindFirstImage(`D:\data\camera\GrpcCamera\GrpcCamera\GrpcCamera`)
	println(imgPath)
	mydev.EndTrack()
}
