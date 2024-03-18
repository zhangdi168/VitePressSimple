package mytest

import (
	"testing"
	"wailstemplate/application/services/system"
)

func TestCopy(t *testing.T) {
	sys1 := system.NewSystemService()
	sys1.CopyPath(`E:\cach\my1\package.json`, `E:\cach\myBak\package.json`, true)
}
