// This is a sample code using kytea package.
// You need to prepare the model file which KyTea needs.
package main

import (
	"fmt"
	"log"

	"github.com/tetsuok/kytea-go"
)

func main() {
	// TODO: Edit the path to model file
	tagger, err := kytea.Create("./model.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer tagger.Destroy()

	util, err := tagger.CreateStringUtil()
	if err != nil {
		log.Fatal(err)
	}
	defer util.Destroy()

	sent, err := util.CreateSentence("太郎は花子のケーキを食べた。")
	if err != nil {
		log.Fatal(err)
	}
	defer sent.Destroy()

	tagger.CalculateWS(sent)
	tagger.CalculateAllTags(sent)
	for i := 0; i < sent.NumWords(); i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%v/%v", util.SurfaceAt(sent, i), util.ReadingAt(sent, i))
	}
	fmt.Printf("\n")
}
