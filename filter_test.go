package goimgrz

import "testing"

func TestFilter_DetectName(t *testing.T) {
	test := []struct {
		name    string
		pattern string
		want    bool
	}{
		{"ab.png", "*.png", true},
		{"/a/b.png", "/*/*.png", true},
		{"a/b.png", "*.png", false},
		{"a.png", "*.jpg", false},
	}

	for _, tt := range test {
		flt := &Filter{pattern: tt.pattern}
		img := &LocImage{tt.name, nil}
		if ok, err := flt.DetectName(img); ok != tt.want {
			t.Errorf("name=%s, pattern=%s, got=%t, want=%t, err:%s", tt.name, tt.pattern, false, tt.want, err)
		}
	}
}

func TestFilter_DetectSize(t *testing.T) {
	test := []struct {
		filename string
		limit    string
		want     bool
	}{
		// IMG_2489.JPG 1.9M
		{"./testdata/IMG_2489.JPG", "+1M", true},
		{"./testdata/IMG_2489.JPG", "+2M", false},
		{"./testdata/IMG_2489.JPG", "+3M", false},
		{"./testdata/IMG_2489.JPG", "-3M", true},
		{"./testdata/IMG_2489.JPG", "-2M", true},
		{"./testdata/IMG_2489.JPG", "-1M", false},
		// gopher2018.png 106k (190717b)
		{"./testdata/gopher2018.png", "-100k", false},
		{"./testdata/gopher2018.png", "-1000k", true},
		{"./testdata/gopher2018.png", "-1M", true},
		{"./testdata/gopher2018.png", "+100k", true},
		{"./testdata/gopher2018.png", "+106k", true},
		{"./testdata/gopher2018.png", "+107k", false},
	}

	for _, tt := range test {
		flt := &Filter{limit: tt.limit}
		img := &LocImage{tt.filename, nil}

		if ok, err := flt.DetectSize(img); ok != tt.want {
			t.Errorf("filename=%s, limit=%s, got=%t, want=%t, err:%s", tt.filename, tt.limit, false, tt.want, err)
		}
	}
}
