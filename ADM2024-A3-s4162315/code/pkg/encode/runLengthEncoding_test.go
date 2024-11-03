package encode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestRunLengthEncoding(t *testing.T) {
	input := "egular courts above the\nly final dependencies: slyly bold \nriously. regular, express dep\nlites. fluffily even de\n pending foxes. slyly re\narefully slyly ex\nven requests. deposits breach a\nongside of the furiously brave acco\n unusual accounts. eve\nnal foxes wake. \ny. fluffily pending d\nages nag slyly pending\nges sleep after the caref\n- quickly regular packages sleep. idly\nts wake furiously \nsts use slyly quickly special instruc\neodolites. fluffily unusual\np furiously special foxes\nss pinto beans wake against th\nes. instructions\n unusual reques\n. slyly special requests haggl\nns haggle carefully ironic deposits. bl\njole. excuses wake carefully alongside of \nithely regula\nsleep quickly. req\nlithely regular deposits. fluffily \n express accounts wake according to the\ne slyly final pac\nsymptotes nag according to the ironic depo"
	var encoded strings.Builder
	n := len(input)

	// Iterate over the input string
	for i := 0; i < n; i++ {
		count := 1

		// Count consecutive characters
		for i+1 < n && input[i] == input[i+1] {
			count++
			i++
		}

		// Append the character and its count to the encoded string
		encoded.WriteString(string(input[i]))
		if count > 1 {
			encoded.WriteString(strconv.Itoa(count))
		}
	}
	fmt.Println("Encoded:", encoded.String())
}
