package main

import (
	"io/ioutil"
	"log"
	"os"
	".."
)

func main() {
	//load data
	cdg_file_data, err := ioutil.ReadFile("../cdg/JVC - Nursery - Do-Re-Mi.cdg")
	if err != nil {
		log.Fatal("Couldn't read .cdg file")
	}

	png := karaoke.ScreenshotAtTime(cdg_file_data, 11)

    file, err := os.OpenFile(
		"screenshot.png",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
    )
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytesWritten, err := file.Write(png)
	if err != nil {
	  log.Fatal(err)
	}

	log.Printf("Wrote %d bytes to %s\n", bytesWritten, "screenshot.png")
}