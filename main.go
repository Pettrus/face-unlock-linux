package main

import (
	"fmt"
	"os"
)

const path = "/lib/security/go-face-unlock/"
const permission = "auth sufficient pam_exec.so quiet stdout /lib/security/go-face-unlock/main"

const urlModel = "https://github.com/davisking/dlib-models/raw/master/"
const shapeModel = "shape_predictor_5_face_landmarks.dat.bz2"
const faceModel = "dlib_face_recognition_resnet_model_v1.dat.bz2"

func main() {
	param := ""

	if len(os.Args) > 1 {
		param = os.Args[1]

		if param == "install" {
			install()
		} else if param == "uninstall" {
			uninstall()
		} else if param == "add" {
			TakePicture(true, false)
		}
	} else {
		TakePicture(false, false)
	}
}

func install() {
	if _, err := os.Stat(path); err == nil {
		fmt.Println("Go face unlock is already installed on your system ;)")
		return
	}

	if _, err := os.Stat(path + "models"); os.IsNotExist(err) {
		os.MkdirAll(path+"models", os.ModePerm)
	}

	if _, err := os.Stat(path + "faces"); os.IsNotExist(err) {
		os.MkdirAll(path+"faces", os.ModePerm)
	}

	if _, err := os.Stat("/etc/pam.d"); os.IsNotExist(err) {
		fmt.Println("You don't have pam on your system.")
		return
	}

	if !Writable("/etc/pam.d") {
		fmt.Println("You don't have permission to write on pam.d, run again as sudo.")
		return
	}

	Wget(urlModel+shapeModel, path+"models/"+shapeModel)
	Wget(urlModel+faceModel, path+"models/"+faceModel)

	Bunzip2(path + "models/" + shapeModel)
	Bunzip2(path + "models/" + faceModel)

	CopyFile("main", path+"main")
	os.Chmod(path+"main", 1644)

	fmt.Println("Press 'Enter' when you are ready to take a picture, you need to be on a well lit room for best results.")
	fmt.Scanln()

	InsertStringToFile("/etc/pam.d/sudo", permission+"\n", 0)
	InsertStringToFile("/etc/pam.d/su", permission+"\n", 0)
	InsertStringToFile("/etc/pam.d/gdm-password", permission+"\n", 0)
	InsertStringToFile("/etc/pam.d/lightdm", permission+"\n", 0)
	InsertStringToFile("/etc/pam.d/gnome-screensaver", permission+"\n", 0)

	TakePicture(true, true)
}

func uninstall() {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Go face unlock is not currently installed :(")
		return
	}

	os.RemoveAll(path)
	RemoveStringFromFile("/etc/pam.d/sudo", permission)
	RemoveStringFromFile("/etc/pam.d/su", permission)
	RemoveStringFromFile("/etc/pam.d/gdm-password", permission)

	fmt.Println("Go face unlock removed with success! :(")
}
