package sdhash

import "testing"

var (
	testCases = []struct {
		filename string
		hash     string
	}{
		{"LICENSE", "sdbf:03:0::11359:sha1:256:5:7ff:160:1:160:IoFBClIQqFAxCa4JCEns8ACBIAQ1UEwAAkUiSoDIEiyNm5QQCJQDhEGISPghTIDWVVaATIMjJChQK4CkgSAgtGCEbIacfGUQgxygkgBEgaRBigAhCoCQO4ZGCEtuB8RgLuQKaAk2AgKA6SAQGCirEEa1doFBwTwyKiAxLEhRKHAYArAUgAkICheDgGY0QVtLKByAwQSQ4CoFAwBWeQHyCIqy4IiACikBBKsAAjXoGAhgFEgCpAzEjYYAFoZT0AAB4QEQCDQC0EoiCkpCUVII33eqdIAJGioMmBXseEq9Wgg4MxhVNCIRPFMLH6pJyZgRDJDRKAIkcaBC4AEgjIjqAQ=="},
	}
)

func TestHash(t *testing.T) {
	for _, tc := range testCases {
		hash := Hash(tc.filename)
		if hash != tc.hash {
			t.Log(tc.hash)
			t.Error(hash)
		}
	}
}
