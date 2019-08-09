package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/Kagami/go-face"
)

const directory = "/lib/security/go-face-unlock/"

func AddFace(imgBuffer *bytes.Buffer, initialSetup bool) {
	rec, err := face.NewRecognizer(directory + "models")

	if err != nil {
		log.Fatalln(err)
	}
	defer rec.Close()

	face, err := rec.Recognize(imgBuffer.Bytes())
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	if len(face) == 1 {
		SaveFaceDescriptions(face[0].Descriptor)

		if initialSetup {
			fmt.Println("Go face unlock installed with success! :)")
		} else {
			fmt.Println("New face added")
		}

		os.Exit(0)
	} else {
		if len(face) > 1 {
			fmt.Println("There can only be one person on the picture")
		} else {
			fmt.Println("No face found")
		}

		if initialSetup {
			fmt.Println("Trying again...")
		}
	}
}

func IdentifyFace(imgBuffer *bytes.Buffer) {
	for i := 0; i < 4; i++ {
		rec, err := face.NewRecognizer(directory + "models")
		if err != nil {
			log.Fatalln(err)
		}
		defer rec.Close()

		//------------

		testf, err := rec.RecognizeSingle(imgBuffer.Bytes())
		if err != nil {
			log.Fatalln(err)
		}

		//------------

		var faceDesc []face.Descriptor

		for _, file := range ReturnFilesOnFolder(directory + "faces") {
			lines, err := File2lines(directory + "faces/" + file.Name())

			if err != nil {
				log.Fatalf("Error reading face descriptions")
			}

			var descriptions [128]float32

			for i, line := range lines {
				num, err := strconv.ParseFloat(line, 32)
				if err != nil {
					log.Fatalf("Error converting face description")
				}
				descriptions[i] = float32(num)
			}

			faceDesc = append(faceDesc, descriptions)
		}

		if testf != nil {
			id := compareFaces(faceDesc, testf.Descriptor, 0.6)
			if id < 0 {
				log.Fatalln("didn't find known face")
			}

			//Face found, exit successfully
			os.Exit(0)
		}
	}

	fmt.Println("Face not recognized, falling back to password.")

	os.Exit(1)
}

func compareFaces(samples []face.Descriptor, comp face.Descriptor, tolerance float32) int {
	res := faceDistance(samples, comp)
	r := -1
	v := float32(1)

	for i, s := range res {
		t := euclideanNorm(s)
		if t < tolerance && t < v {
			v = t
			r = i
		}
	}

	return r
}

func faceDistance(samples []face.Descriptor, comp face.Descriptor) []face.Descriptor {
	res := make([]face.Descriptor, len(samples))

	for i, s := range samples {
		for j, _ := range s {
			res[i][j] = samples[i][j] - comp[j]
		}
	}

	return res
}

func euclideanNorm(f face.Descriptor) float32 {
	var s float32
	for _, v := range f {
		s = s + v*v
	}

	return float32(math.Sqrt(float64(s)))
}
