package sdhash

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func entr64(buffer []byte, entropyTable [64]uint64) (uint64, [256]byte) {
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
	return entr, ascii
}

func entr64Inc(buffer []byte, prevEntropy *uint64, ascii *[256]byte, entropyTable [64]uint64) (entropy uint64) {
	if buffer[0] == buffer[63] {
		return *prevEntropy
	}

	old_char_cnt := ascii[buffer[0]]
	new_char_cnt := ascii[buffer[63]]

	ascii[buffer[0]]--
	ascii[buffer[63]]++

	if old_char_cnt == new_char_cnt+1 {
		return *prevEntropy
	}

	old_diff := entropyTable[old_char_cnt] - entropyTable[old_char_cnt-1]
	new_diff := entropyTable[new_char_cnt+1] - entropyTable[new_char_cnt]

	entropy = *prevEntropy - old_diff + new_diff
	if entropy < 0 {
		entropy = 0
	} else if entropy > ENTR_SCALE {
		entropy = ENTR_SCALE
	}

	return
}

// Hash returns an sdhash of the input file
func Hash(filename string) (hash string, err error) {
	// Initialization: The entropy score Hnorm, precedence rank Rprec and popularity score Rpop are initialized to zero.
	/*var (
		hNorm int
		rPrec int
		rPop  int
	)*/
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return
	}
	r := bufio.NewReader(f)
	p := make([]byte, featureSize)
	buffer := []byte{}
	_, err = r.Read(p)
	if err != nil {
		return
	}
	buffer = p
	// Hnorm Calculation: The Shannon entropy is first computed for every feature (B-byte sequence)
	entropyTable := entr64Table()
	entropy, ascii := entr64(p, entropyTable)
	for {
		entropy = entr64Inc(p, &entropy, &ascii, entropyTable)
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		p = append(p[1:], b)
		buffer = append(buffer, b)
	}
	fmt.Println(entropy)
	println(len(buffer))
	// Rprec Calculation: The precedence rank Rprec value is obtained by mapping the entropy score Hnorm based on empirical observations.
	// Rpop Calculation: For every sliding window of W consecutive features, the leftmost feature with the lowest precedence rank Rprec is identified. The popularity score Rpop of the identified feature is incremented by one.
	// Feature Selection: Features with popularity rank Rpop >= t, where t is a threshold parameter, are selected.
	// Filtering Weak Features
	// Generating Similarity Digests
	// Comparing Digests
	return "sdbf:", nil
}
