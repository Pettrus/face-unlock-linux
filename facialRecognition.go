package main

import (
	"bytes"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/Kagami/go-face"
)

func IdentifyFace(li chan *bytes.Buffer) {
	numOfTry := 0

	for {
		numOfTry++

		if numOfTry > 5 {
			os.Exit(1)
		}

		const directory = "/lib/security/go-face-unlock/"

		rec, err := face.NewRecognizer(directory + "models")
		if err != nil {
			log.Fatalln(err)
		}
		defer rec.Close()

		//------------

		dataImage := filepath.Join(directory, "base.jpeg")

		faces, err := rec.RecognizeFile(dataImage)
		if err != nil {
			log.Fatalln(err)
		}

		var samples []face.Descriptor
		var totalF []int32
		for i, f := range faces {
			samples = append(samples, f.Descriptor)
			totalF = append(totalF, int32(i))
		}

		//-------------

		testData := filepath.Join(directory, "image.jpeg")
		testf, err := rec.RecognizeSingleFile(testData)
		if err != nil {
			log.Fatalln(err)
		}

		if testf != nil {
			id := compareFaces(samples, testf.Descriptor, 0.6)
			if id < 0 {
				log.Fatalln("didn't find known face")
			}

			//Face found, exit successfully
			os.Exit(0)
		}
	}
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
