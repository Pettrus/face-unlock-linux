package main

import (
	"fmt"
	"os"
)

const path = "/lib/security/go-face-unlock/"
const permission = "auth sufficient pam_exec.so stdout /lib/security/go-face-unlock/main"

func main() {
	param := ""

	if len(os.Args) > 1 {
		param = os.Args[1]

		if param == "install" {
			install()
		} else if param == "uninstall" {
			uninstall()
		} else if param == "add" {
			TakePicture(true)
		}
	} else {
		TakePicture(false)
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

	CopyFile("main", path+"main")
	os.Chmod(path+"main", 1644)
	CopyFile("models/dlib_face_recognition_resnet_model_v1.dat", path+"models/dlib_face_recognition_resnet_model_v1.dat")
	CopyFile("models/shape_predictor_5_face_landmarks.dat", path+"models/shape_predictor_5_face_landmarks.dat")

	InsertStringToFile("/etc/pam.d/sudo", permission+"\n", 0)
	InsertStringToFile("/etc/pam.d/su", permission+"\n", 0)
	InsertStringToFile("/etc/pam.d/gdm-password", permission+"\n", 0)

	TakePicture(true)

	fmt.Println("Go face unlock installed with success! :)")
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
