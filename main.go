package main

import (
	"os"
)

func main() {
	initialSetup := len(os.Args) > 1

	if initialSetup {
		const path = "/lib/security/go-unlock/"

		if _, err := os.Stat(path + "models"); os.IsNotExist(err) {
			os.MkdirAll(path+"models", os.ModePerm)
		}

		copyFile("main", path+"main")
		os.Chmod(path+"main", 1644)
		copyFile("models/dlib_face_recognition_resnet_model_v1.dat", path+"models/dlib_face_recognition_resnet_model_v1.dat")
		copyFile("models/shape_predictor_5_face_landmarks.dat", path+"models/shape_predictor_5_face_landmarks.dat")

		const permission = "auth sufficient pam_exec.so stdout /lib/security/go-unlock/main\n"
		InsertStringToFile("/etc/pam.d/sudo", permission, 0)
		InsertStringToFile("/etc/pam.d/su", permission, 0)
		InsertStringToFile("/etc/pam.d/gdm-password", permission, 0)
	}

	TakePicture(initialSetup)
}
