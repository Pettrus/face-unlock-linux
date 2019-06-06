package main

import (
	"bytes"
	"io"
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

		imgRecebida := <-li
		const modelDir = "/home/pettrus/go/src/go-unlock/models"
		const dataDir = "/home/pettrus/go/src/go-unlock/images"

		image, _ := os.Create("/home/pettrus/go/src/go-unlock/imagem.jpeg")
		defer image.Close()

		io.Copy(image, imgRecebida)

		rec, err := face.NewRecognizer(modelDir)
		if err != nil {
			log.Fatalln(err)
		}
		defer rec.Close()

		//------------

		dataImage := filepath.Join(dataDir, "base2.jpeg")

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

		testData := filepath.Join("/home/pettrus/go/src/go-unlock", "imagem.jpeg")
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
