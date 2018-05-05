package sinkVolume

import (
	"os/exec"
	"strconv"
)

func New(name string, step int) *SinkVolume {
	return &SinkVolume{name, step}
}

type SinkVolume struct {
	Name string
	Step int
}

func (self SinkVolume) VolumeUp() {
	strStep := "+" + strconv.Itoa(self.Step) + "%"
	self.changeVolume(strStep)
}

func (self SinkVolume) VolumeDown() {
	strStep := "-" + strconv.Itoa(self.Step) + "%"
	self.changeVolume(strStep)
}

func (self SinkVolume) changeVolume(step string) {
	cmd := exec.Command("pactl", "set-sink-volume", self.Name, step)
	_ = cmd.Run()
}
