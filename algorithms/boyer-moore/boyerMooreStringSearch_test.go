package boyerMooreStringSearch

import (
	"fmt"
	"testing"
)

func TestSearch_Pattern_Returns_PatternIndexOrNil(t *testing.T) {
	s := "Hello, world"
	pattern := "world"
	doTesting(s, pattern, t)
}

func Test2(t *testing.T) {
	s := ` Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean et congue leo, quis imperdiet nulla. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas bibendum a sapien non venenatis. Nam vel finibus sem, sit amet suscipit turpis. Fusce accumsan dui nec arcu rutrum, sit amet lacinia leo condimentum. Vestibulum id nisl eget erat condimentum sodales. Morbi at lectus varius, hendrerit eros vitae, pulvinar lacus. Suspendisse porttitor elementum lectus, in aliquam orci mattis non. Phasellus condimentum purus a diam vulputate, vestibulum finibus dolor vulputate. Donec tincidunt urna posuere mauris interdum, id scelerisque mi viverra. Integer eget tempus augue, non eleifend tellus. Nulla ornare aliquam sapien consectetur mollis. Maecenas hendrerit enim ac erat eleifend, a maximus augue fringilla.

Sed imperdiet faucibus dolor, sit amet lobortis quam tempor feugiat. Etiam ut dolor et dui aliquet porttitor. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec vitae ex sit amet leo laoreet maximus. Sed pulvinar libero a urna porta, non vulputate tellus mattis. Nam convallis, lorem vitae varius efficitur, nunc ante vulputate massa, ornare eleifend enim libero vel ipsum. Donec eleifend eros non ante cursus, vel tempus nisi congue. Nunc at enim ut augue aliquet elementum. Maecenas risus turpis, iaculis id aliquam id, interdum a purus. Donec vel mauris mollis, feugiat nibh quis, vulputate tortor. Quisque nec mi risus. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.

Nulla quis venenatis urna. Integer vitae erat ex. Aenean luctus erat orci, nec congue magna vehicula eu. Quisque eget dignissim nisi. Proin leo enim, sagittis ut dapibus vitae, laoreet sit amet nisl. Vivamus volutpat dictum tristique. Morbi vestibulum justo a neque cursus lacinia. `

	pattern := "venenatis"
	doTesting(s, pattern, t)
}

func doTesting(s, pattern string, t *testing.T) {
	index, err := IndexOf(pattern, s)
	if err != nil || string(s[index:index+len(pattern)]) != pattern {
		t.Fatalf("Finding \"%s\" in text failed, %v", pattern, err)
	}
	fmt.Printf("Found %s in text at index %d", pattern, index)
}
