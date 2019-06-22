## Girls
Girls is not girl, it just short for **Go Image Resize List**, It is used for
resize network or local images to specified size and destination

## Install
```
go get -v github.com/tkstorm/girls
```

## Help
```
$ girls -h
Usage of girls:
  -crawler_url string
    	the crawler url used by girls download the http images and resize only matched image files
  -dst string
    	the output dir where image after resize store (default "/tmp")
  -height uint
    	set resize image's height
  -img string
    	the local image which to be resize
  -match string
    	only matched image will to be resize (default "*")
  -quality uint
    	set resize image's quality percent (default 75)
  -scan_dir string
    	scan the dir where image inside to be resize
  -url string
    	the http(s) url of image which to be resize, image resource(url|urls|dirname|filename) at least need set one
  -urls string
    	the http(s) urls of image which to be resize, separated by ','
  -water_img string
    	append water image
  -width uint
    	set resize image's width (default 300)
```

The `-crawler_url`、`quality`、`water_img` is developing, it will finished 
soon（now is unavailable）

Support image type:

- jpeg
- png
- gif

## Example

### 1. girls resize single local image
```
$ girls -img ./testdata/gopher2018.png
2019/06/22 18:46:45 resize ok: /tmp/gopher2018.png (inputW=300,inputH=0)
```

### 2. girls resize local dir images
```
$ girls -scan_dir ./testdata -width 500 -dst /tmp/resize_500
2019/06/22 18:47:45 resize ok: /tmp/resize_500/gopher2018.png (inputW=500,inputH=0)
2019/06/22 18:47:45 resize ok: /tmp/resize_500/web_bg.jpg (inputW=500,inputH=0)
2019/06/22 18:47:45 resize ok: /tmp/resize_500/IMG_2489.JPG (inputW=500,inputH=0)

$ ll -h ./testdata
total 4480
-rw-r--r--@ 1 Terry  access_bpf   1.9M  6 19 02:26 IMG_2489.JPG
-rw-r--r--@ 1 Terry  staff        106K  9 10  2018 gopher2018.png
-rw-r--r--@ 1 Terry  staff        186K  4  2 15:04 web_bg.jpg

$ ll -h /tmp/resize_500
total 416
-rw-r--r--  1 Terry  wheel    36K  6 22 18:47 IMG_2489.JPG
-rw-r--r--  1 Terry  wheel    50K  6 22 18:47 gopher2018.png
-rw-r--r--  1 Terry  wheel   116K  6 22 18:47 web_bg.jpg
```

### 3. Girls resize http url image
```
$ girls -url https://cdn-images-1.medium.com/max/1600/1\*n1kWgo0dPS80uoE430hqSQ.jpeg -width 300
2019/06/22 18:42:17 resize ok: /tmp/1*n1kWgo0dPS80uoE430hqSQ.jpeg (inputW=300,inputH=0)
```

### 4. Girls resize batch http url image
```
$ girls -urls https://cdn-images-1.medium.com/max/1600/1\*k74qnaAcJd3bzRj7PnLIbg.jpeg,https://cdn-images-1.medium.com/max/2600/0\*jraDH1ztolpSmT9I
2019/06/22 18:49:29 resize ok: /tmp/1*k74qnaAcJd3bzRj7PnLIbg.jpeg (inputW=300,inputH=0)
2019/06/22 18:49:30 resize ok: /tmp/0*jraDH1ztolpSmT9I (inputW=300,inputH=0)
```

### Support 
- jpeg
- png
- gif

## TestCase

### Test resize local image
```
=== RUN   TestResizeLocalImage
--- PASS: TestResizeLocalImage (0.47s)
    gir_resize_test.go:11: ./testdata/IMG_2489.JPG
    gir_resize_test.go:18: /var/folders/pk/2mwxkhlx5g7ckfwyks8vn3n40000gn/T/ /Users/Terry/Library/Caches
    gir_resize_test.go:27: /var/folders/pk/2mwxkhlx5g7ckfwyks8vn3n40000gn/T/resizeCache/IMG_2489.JPG
PASS
```

### Test resize http image
```
=== RUN   TestResizeHttpImage
--- PASS: TestResizeHttpImage (2.76s)
    gir_resize_test.go:36: [https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false https://uidesign.gbtcdn.com/GB/image/2019/20190612_10650/1190x420.gif?imbypass=false]
    gir_resize_test.go:43: /var/folders/pk/2mwxkhlx5g7ckfwyks8vn3n40000gn/T/ /Users/Terry/Library/Caches
    gir_resize_test.go:54: resize ok, /var/folders/pk/2mwxkhlx5g7ckfwyks8vn3n40000gn/T/resizeCache/1*pV0ZUbW1dURx-_YOWu1mzQ.png
    gir_resize_test.go:54: resize ok, /var/folders/pk/2mwxkhlx5g7ckfwyks8vn3n40000gn/T/resizeCache/New_B.jpg
    gir_resize_test.go:54: resize ok, /var/folders/pk/2mwxkhlx5g7ckfwyks8vn3n40000gn/T/resizeCache/1190x420.gif
PASS
```

### Task resize images (include local image and http image)

```
=== RUN   TestGirTask_ResizeImages
2019/06/21 15:06:20 resize ok: /tmp/girls/IMG_2489.JPG
2019/06/21 15:06:22 resize ok: /tmp/girls/1*pV0ZUbW1dURx-_YOWu1mzQ.png
2019/06/21 15:06:22 resize ok: /tmp/girls/New_B.jpg
2019/06/21 15:06:22 resize ok: /tmp/girls/1190x420.gif
--- PASS: TestGirTask_ResizeImages (2.48s)
PASS
```

## Bug
1. Now, http image download is no timeout(used default Http Client), this will change soon