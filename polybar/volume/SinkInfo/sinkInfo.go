package sinkInfo

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func NewWithName(name string) *SinkInfo {
	return &SinkInfo{Name: name}
}

type SinkInfo struct {
	Name string
}

func (self SinkInfo) GetVolume() (int, error) {
	output, err := exec.Command("pactl", "list", "sinks").Output()
	if err != nil {
		return 0, err
	}
	return self.parseData(output)
}

func (self SinkInfo) parseData(data []byte) (int, error) {

	info, err := self.truncateInfoToSinkInfo(string(data))
	if err != nil {
		return 0, err
	}

	r, err := regexp.Compile("[0-9]+%")
	if err != nil {
		return 0, err
	}

	result := r.FindString(info)
	if result == "" {
		return 0, errors.New("unknown volume")
	}
	result = strings.TrimSuffix(result, "%")

	volume, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return volume, nil
}

func (self SinkInfo) truncateInfoToSinkInfo(info string) (string, error) {

	sinkStartIndex, err := self.findSinkInfoStartIndex(info)
	if err != nil {
		return "", err
	}

	info = info[sinkStartIndex:]
	volumeStartIndex, err := self.findVolumeStartIndex(info)
	if err != nil {
		return "", err
	}

	info = info[volumeStartIndex:]
	volumeEndIndex, err := self.findVolumeEndIndex(info)
	if err != nil {
		return "", err
	}

	info = info[:volumeEndIndex]
	return info, nil
}

func (self SinkInfo) findSinkInfoStartIndex(info string) (int, error) {
	return findSubstringInInfo(info, self.Name)
}

func (self SinkInfo) findVolumeStartIndex(info string) (int, error) {
	return findSubstringInInfo(info, "Volume")
}

func (self SinkInfo) findVolumeEndIndex(info string) (int, error) {
	return findSubstringInInfo(info, "\n")
}

func findSubstringInInfo(info, pattern string) (int, error) {
	indx := strings.Index(info, pattern)
	if indx == -1 {
		return 0, errors.New("no " + pattern + " in info")
	}

	return indx, nil
}
