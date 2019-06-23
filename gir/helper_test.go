package gir

import "testing"

func TestGetLocalImages(t *testing.T) {
	imgs, err := GetImagesFromDir("./testdata")
	if err != nil {
		t.Error(err)
	}
	t.Log(imgs)
}

func TestParseHumanDataSize(t *testing.T) {
	test := []struct {
		size string
		want uint64
	}{
		{"1024k", 1024 * 1024},
		{"100k", 100 * 1024},
		{"100M", 100 * 1024 * 1024},
		{"100", 100},
	}

	for _, tt := range test {
		nsize, err := ParseHumanDataSize(tt.size)
		if err != nil {
			t.Errorf("ParseHumanDataSize(%s), %s", tt.size, err)
		}
		if nsize != tt.want {
			t.Errorf("human size: %s, want %d, but got %d", tt.size, tt.want, nsize)
		}
	}
}

func TestSatisfyHumanSize(t *testing.T) {
	test := []struct {
		size  string
		limit string
		want  bool
	}{
		{"1024", "1000", false},
		{"1024", "1024", true},
		{"1024", "+1000", true},
		{"1024", "+1024", true},
		{"1024", "+1025", false},
		{"10240", "+10k", true},
		{"10240", "+11k", false},
		{"10230", "+10k", false},
		{"10240", "-10k", false},
		{"10230", "-10k", true},
		{"10241", "-10k", false},
		{"99", "-1k", true},
		{"100", "1M", false},
		{"100", "+1M", false},
		{"100", "-1M", true},
		{"100", "", true},
	}

	for _, tt := range test {
		satisfy, err := SatisfyHumanSize(tt.size, tt.limit)
		if err != nil {
			t.Errorf("SatisfyHumanSize(%s, %s), %s", tt.size, tt.limit, err)
		}
		if satisfy != tt.want {
			t.Errorf("file size %s, compare %s,  want %t, but got %t", tt.size, tt.limit, tt.want, satisfy)
		}
	}

}
