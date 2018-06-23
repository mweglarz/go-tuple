package big

import "testing"

func TestSample(t *testing.T) {

	inputs := []string{"ab", "bb", "hefg", "dhck", "dkhc"}
	outputs := []string{"ba", "no answer", "hegf", "dhkc", "hcdk"}

	for i := 0; i < len(inputs); i++ {
		if result := BiggerIsGreater(inputs[i]); result != outputs[i] {
			t.Fatalf("expected %s, got %s", outputs[i], result)
		}
	}
}

func TestCase1(t *testing.T) {

	input := "lgeuvf"
	output := "lgevfu"
	// wrong result lgfeuv

	if result := BiggerIsGreater(input); result != output {
		t.Fatalf("expected %s, got %s", output, result)
	}
}
