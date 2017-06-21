package sdhash

import (
	"fmt"
	"math"
)

const (
	featureSize = 64
	BINS        = 1000
	ENTR_POWER  = 10
	ENTR_SCALE  = (BINS * (1 << ENTR_POWER))
)

func entr64Table() [64]uint64 {
	ENTROPY_64_INT := [64]uint64{}
	for _, i := range ENTROPY_64_INT {
		p := float64(i / 64)
		ENTROPY_64_INT[i] = uint64((-p * (math.Log(p) / math.Log(2)) / 6) * ENTR_SCALE)
	}
	return ENTROPY_64_INT
}

func entr64(buffer []byte, entropyTable [64]uint64) uint64 {
	ascii := [256]byte{}
	for _, b := range buffer {
		ascii[b]++
	}
	var entr uint64
	for _, i := range ascii {
		if ascii[i] != 0 {
			entr += entropyTable[ascii[i]]
		}
	}
	return entr
}

func entr64Inc(buffer []byte, prevEntropy int64, ascii [256]byte, entropyTable [64]uint64) (entropy int64) {
	if buffer[0] == buffer[64] {
		return prevEntropy
	}

	old_char_cnt := ascii[buffer[0]]
	new_char_cnt := ascii[buffer[64]]

	ascii[buffer[0]]--
	ascii[buffer[64]]++

	if old_char_cnt == new_char_cnt+1 {
		return prevEntropy
	}

	old_diff := int64(entropyTable[old_char_cnt]) - int64(entropyTable[old_char_cnt-1])
	new_diff := int64(entropyTable[new_char_cnt+1]) - int64(entropyTable[new_char_cnt])

	entropy = int64(prevEntropy) - old_diff + new_diff
	if entropy < 0 {
		entropy = 0
	} else if entropy > ENTR_SCALE {
		entropy = ENTR_SCALE
	}

	return
}

// Hash returns an sdhash of the input file
func Hash(file string) (hash string) {
	// Initialization: The entropy score Hnorm, precedence rank Rprec and popularity score Rpop are initialized to zero.
	/*var (
		hNorm int
		rPrec int
		rPop  int
	)*/
	// Hnorm Calculation: The Shannon entropy is first computed for every feature (B-byte sequence)
	entropyTable := entr64Table()
	entropy := entr64([]byte{1, 3, 7}, entropyTable)
	fmt.Println(entropy)
	// Rprec Calculation: The precedence rank Rprec value is obtained by mapping the entropy score Hnorm based on empirical observations.
	// Rpop Calculation: For every sliding window of W consecutive features, the leftmost feature with the lowest precedence rank Rprec is identified. The popularity score Rpop of the identified feature is incremented by one.
	// Feature Selection: Features with popularity rank Rpop >= t, where t is a threshold parameter, are selected.
	// Filtering Weak Features
	// Generating Similarity Digests
	// Comparing Digests
	return "sdbf:"
}
