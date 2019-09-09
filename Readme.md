
## 1. goimgrz - Go Image Resize

Go Image Resize is a command line tool resize Http Url images or native images.

Go Image Resize written by golang, based on `github.com/nfnt/resize` library, support interpolation functions setting.

### 1.1. Feature
1. Local file, url image resource file resize with special width and height.
2. Scanning local file directory images and resize them.
3. Image file size, name filtering 
4. Resize interpolation function specification
5. Jpeg quality setting
6. Support batch and concurrent image resize processing

### 1.2. Todo
1. Water Image Support
2. Http Request timeout setting
3. Http Proxy Setting (when download special network image)
4. Crawler url Support (when one want one html page's image)
5. Short Param Support
6. Concurrent resource setting (now no limit concurrent resource setting)
7. More image format support (now only support jpeg,png,gif)

Please note: the toolkit may continue to make major changes until version 0.1

## 2. Install
```
go get -u github.com/tkstorm/goimgrz
```

## 3. Quick Start

### 3.1. Image File Type

#### 3.1.1. resize single local image with special width
```
$ goimgrz -img ./testdata/gopher2018.png -w 400
2019/06/25 16:10:16 resize ok: /tmp/goimgrz/gopher2018.png (inputW=400,inputH=0)
```

#### 3.1.2. resize local dir images to dst dir
```
$ goimgrz -scdir ./testdata -w 500  -dst /tmp/goimgrz/500
2019/06/25 16:11:22 resize ok: /tmp/goimgrz/500/gopher2018.png (inputW=500,inputH=0)
2019/06/25 16:11:22 resize ok: /tmp/goimgrz/500/web_bg.jpg (inputW=500,inputH=0)
2019/06/25 16:11:23 resize ok: /tmp/goimgrz/500/IMG_2489.JPG (inputW=500,inputH=0)
```

#### 3.1.3. resize http url image
```
$ goimgrz -url https://cdn-images-1.medium.com/max/1600/1\*n1kWgo0dPS80uoE430hqSQ.jpeg -w 300
2019/06/22 18:42:17 resize ok: /tmp/1*n1kWgo0dPS80uoE430hqSQ.jpeg (inputW=300,inputH=0)
```

Tips: Be careful of the special characters in the Url and try to use single quotation marks.

#### 3.1.4. resize batch http url image
```
$ goimgrz -urls https://cdn-images-1.medium.com/max/1600/1\*k74qnaAcJd3bzRj7PnLIbg.jpeg,https://cdn-images-1.medium.com/max/2600/0\*jraDH1ztolpSmT9I
2019/06/22 18:49:29 resize ok: /tmp/1*k74qnaAcJd3bzRj7PnLIbg.jpeg (inputW=300,inputH=0)
2019/06/22 18:49:30 resize ok: /tmp/0*jraDH1ztolpSmT9I (inputW=300,inputH=0)
```

### 3.2. Resize Image Param

#### 3.2.1. resize setting quality (only jpeg has effect)
```
$ goimgrz -img ./testdata/IMG_2489.JPG -w 500 -qty 85
2019/06/25 16:14:28 resize ok: /tmp/goimgrz/IMG_2489.JPG (inputW=500,inputH=0)
```

#### 3.2.2. resize setting interpolation functions support
The provided interpolation functions support (from fast to slow execution time)

- 0: NearestNeighbor
- 1: Bilinear
- 2: Bicubic
- 3: MitchellNetravali
- 4: Lanczos2
- 5: Lanczos3

```
$ goimgrz -img ./testdata/IMG_2489.JPG -w 500 -itp 2
2019/06/25 16:16:40 resize ok: /tmp/goimgrz/IMG_2489.JPG (inputW=500,inputH=0)
```

### 3.3. Filter

#### 3.3.1. filter by file size
```
$ ll -h ./testdata
total 4480
-rw-r--r--@ 1 Terry  access_bpf   1.9M  6 19 02:26 IMG_2489.JPG
-rw-r--r--@ 1 Terry  staff        106K  9 10  2018 gopher2018.png
-rw-r--r--@ 1 Terry  staff        186K  4  2 15:04 web_bg.jpg

$ goimgrz -scdir ./testdata -size +1M
2019/06/23 23:02:20 resize ok: /tmp/IMG_2489.JPG (inputW=300,inputH=0)

$ goimgrz -scdir ./testdata -size +1M -v
2019/06/25 16:20:44 resize fail: error(31): not satisfy, limit size:+1M, file size:108873
2019/06/25 16:20:44 resize fail: error(31): not satisfy, limit size:+1M, file size:190717
2019/06/25 16:20:45 resize ok: /tmp/goimgrz/IMG_2489.JPG (inputW=0,inputH=0)

$ goimgrz -scdir ./testdata -size -200k
2019/06/23 23:14:45 resize ok: /tmp/gopher2018.png (inputW=300,inputH=0)
2019/06/23 23:14:45 resize ok: /tmp/web_bg.jpg (inputW=300,inputH=0)
```

#### 3.3.2. filter by name (using shell pattern)
```
$ goimgrz -scdir ./testdata -size -1M -name '*/*.png'
2019/06/23 23:08:22 resize ok: /tmp/gopher2018.png (inputW=300,inputH=0)
```

## 4. Command Helper
```
$ goimgrz --help
Usage: goimgrz [options...]

Goimgrz version 0.0.1, a simple image resizing tool that supports web or local images.
https://github.com/tkstorm/goimgrz

Options:
Resize image source:
-url	Web images to be resize,
	image source(url|urls|img|imgs|scdir) at least need set one
-urls	Multiple web images to be resize, separated by ','
-img	Local images to be resize
-imgs	Multiple local images to be resize, separated by ','
-scdir	Scanned file image file directory

Filter:
-name	Using shell pattern to filter image, like *.png
-size	Using file size to filter image, like +200k

Resize Setting:
-w	Set resize image's width, default width is 0 represent origin image
-h	Set resize image's height, default height is 0 represent origin image
-cfmt	Convert image output format(jpg|png|gif)
-itp	The provided interpolation functions support (from fast to slow execution time).
	0:NearestNeighbor,1:Bilinear,2:Bicubic,3:MitchellNetravali,4:Lanczos2,5:Lanczos3
-qty 	Set resize image's quality percent

Image Saving:
-dst	The output dir

Other:
-v	Verbose message
```