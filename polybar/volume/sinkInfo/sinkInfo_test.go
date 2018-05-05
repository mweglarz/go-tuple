package sinkInfo

import "testing"

var testInput string = `Sink #0
	State: SUSPENDED
	Name: alsa_output.pci-0000_01_00.1.hdmi-stereo
	Description: HDA NVidia Digital Stereo (HDMI)
	Driver: module-alsa-card.c
	Sample Specification: s16le 2ch 44100Hz
	Channel Map: front-left,front-right
	Owner Module: 6
	Mute: no
	Volume: front-left: 19661 /  30% / -31,37 dB,   front-right: 19661 /  30% / -31,37 dB
	        balance 0,00
	Base Volume: 65536 / 100% / 0,00 dB
	Monitor Source: alsa_output.pci-0000_01_00.1.hdmi-stereo.monitor
	Latency: 0 usec, configured 0 usec
	Flags: HARDWARE DECIBEL_VOLUME LATENCY SET_FORMATS
	Properties:
		alsa.resolution_bits = "16"
		device.api = "alsa"
		device.class = "sound"
		alsa.class = "generic"
		alsa.subclass = "generic-mix"
		alsa.name = "HDMI 0"
		alsa.id = "HDMI 0"
		alsa.subdevice = "0"
		alsa.subdevice_name = "subdevice #0"
		alsa.device = "3"
		alsa.card = "1"
		alsa.card_name = "HDA NVidia"
		alsa.long_card_name = "HDA NVidia at 0xf7080000 irq 17"
		alsa.driver_name = "snd_hda_intel"
		device.bus_path = "pci-0000:01:00.1"
		sysfs.path = "/devices/pci0000:00/0000:00:01.0/0000:01:00.1/sound/card1"
		device.bus = "pci"
		device.vendor.id = "10de"
		device.vendor.name = "NVIDIA Corporation"
		device.product.id = "0fba"
		device.string = "hdmi:1"
		device.buffering.buffer_size = "352768"
		device.buffering.fragment_size = "176384"
		device.access_mode = "mmap+timer"
		device.profile.name = "hdmi-stereo"
		device.profile.description = "Digital Stereo (HDMI)"
		device.description = "HDA NVidia Digital Stereo (HDMI)"
		alsa.mixer_name = "Nvidia GPU 72 HDMI/DP"
		alsa.components = "HDA:10de0072,10438569,00100100"
		module-udev-detect.discovered = "1"
		device.icon_name = "audio-card-pci"
	Ports:
		hdmi-output-0: HDMI / DisplayPort (priority: 5900, available)
	Active Port: hdmi-output-0
	Formats:
		pcm`

var testName string = "alsa_output.pci-0000_01_00.1.hdmi-stereo"

func TestParseData(t *testing.T) {
	expect := 30

	data := ([]byte)(testInput)
	sinkInfo := SinkInfo{testName}

	volume, err := sinkInfo.parseData(data)
	if err != nil {
		t.Fatal(err)
	}

	if volume != expect {
		t.Fatalf("expected %d%% volume, got %d%% volume", expect, volume)
	}
}

func TestFindVolumeLine(t *testing.T) {
	expect := "Volume: front-left: 19661 /  30% / -31,37 dB,   front-right: 19661 /  30% / -31,37 dB"
	sinkInfo := SinkInfo{testName}

	info, err := sinkInfo.truncateInfoToSinkInfo(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if info != expect {
		t.Fatalf("volume line don't match, expect \n%s\n*** got ***\n%s\n", expect, info)
	}
}

func TestFindVolumeEndIndex(t *testing.T) {
	offset := 222 + 33
	info := testInput[offset:]
	expect := 85
	sinkInfo := SinkInfo{testName}

	indx, err := sinkInfo.findVolumeEndIndex(info)
	if err != nil {
		t.Fatal(err)
	}

	if indx != expect {
		t.Fatalf("expecting to found end volume at index %d, method returns index %d, info %s", expect, indx, info)
	}
}

func TestFindVolumeStartIndex(t *testing.T) {
	info := testInput[33:]
	expect := 222
	sinkInfo := SinkInfo{testName}

	indx, err := sinkInfo.findVolumeStartIndex(info)
	if err != nil {
		t.Fatal(err)
	}

	if indx != expect {
		t.Fatalf("expecting to found volume at index %d, method returns index %d", expect, indx)
	}
}

func TestFindSinkInfoStart(t *testing.T) {
	expect := 33
	sinkInfo := SinkInfo{testName}

	indx, err := sinkInfo.findSinkInfoStartIndex(testInput)
	if err != nil {
		t.Fatal(err)
	}

	if indx != expect {
		t.Fatalf("expecting to found sink at index %d, method returns index %d", expect, indx)
	}
}

func TestFindSubstring(t *testing.T) {
	source := "hello world dupa"
	expect := 6

	indx, err := findSubstringInInfo(source, "world")
	if err != nil {
		t.Fatal(err)
	}

	if indx != expect {
		t.Fatalf("expecting to found pattern at index %d, method returns index %d", expect, indx)
	}
}
