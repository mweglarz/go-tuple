package main

import (
	"flag"
	"fmt"

	info "tuple-mw.com/polybar/mw-volume/sinkInfo"
	volume "tuple-mw.com/polybar/mw-volume/sinkVolume"
)

func main() {
	step := flag.Uint("step", 5, "volume up/down step")
	volUp := flag.Bool("up", false, "volume up")
	volDown := flag.Bool("down", false, "volume down")
	flag.Parse()

	if len(flag.Args()) < 1 {
		panic("must provide sink name")
	}

	name := flag.Args()[0]

	defer func() {
		info := info.NewWithName(name)
		vol, err := info.GetVolume()
		if err == nil {
			fmt.Printf("%d", vol)
		} else {
			fmt.Println("-1")
		}
	}()

	if *volUp {
		vol := volume.New(name, int(*step))
		vol.VolumeUp()

	} else if *volDown {
		vol := volume.New(name, int(*step))
		vol.VolumeDown()
	}
}
