## goimgrz
~~goimgrz is not girl, it just short for **Go Image Resize List**, It is used for
resize network or local images to specified size and destination~~

This is a command line tool based on resize Http images or native images written by golang, supporting batch and concurrent processing.

Please note: the toolkit may continue to make major changes until version 0.1

## Install
```
go get -v github.com/tkstorm/goimgrz
```

## Help
```
$ goimgrz -h
Usage of goimgrz:
  -dir string
    	scan the dir where image inside to be resize
  -dst string
    	the output dir where image after resize store (default "/tmp")
  -height uint
    	set resize image's height
  -img string
    	the local image file which to be resize
  -imgs string
    	local image files which to be resize, separated by ','
  -name string
    	using shell pattern to filter image, like *.png (default "*")
  -size string
    	using file size to filter image, like +200k
  -url string
    	the image's http(s) url to be resize, image resource(url|urls|img|imgs|dir) at least need set one
  -urls string
    	image's http(s) urls to be resize, separated by ','
  -verbose
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

### 1. goimgrz resize single local image
```
$ goimgrz -img ./testdata/gopher2018.png
2019/06/22 18:46:45 resize ok: /tmp/gopher2018.png (inputW=300,inputH=0)
```

### 2. goimgrz resize local dir images
```
$ goimgrz -dir ./testdata
2019/06/23 23:02:14 resize ok: /tmp/gopher2018.png (inputW=300,inputH=0)
2019/06/23 23:02:14 resize ok: /tmp/web_bg.jpg (inputW=300,inputH=0)
2019/06/23 23:02:14 resize ok: /tmp/IMG_2489.JPG (inputW=300,inputH=0)
```

### 3. goimgrz resize http url image
```
$ goimgrz -url https://cdn-images-1.medium.com/max/1600/1\*n1kWgo0dPS80uoE430hqSQ.jpeg -width 300
2019/06/22 18:42:17 resize ok: /tmp/1*n1kWgo0dPS80uoE430hqSQ.jpeg (inputW=300,inputH=0)
```

Tips: Be careful of the special characters in the Url and try to use single quotation marks.

### 4. goimgrz resize batch http url image
```
$ goimgrz -urls https://cdn-images-1.medium.com/max/1600/1\*k74qnaAcJd3bzRj7PnLIbg.jpeg,https://cdn-images-1.medium.com/max/2600/0\*jraDH1ztolpSmT9I
2019/06/22 18:49:29 resize ok: /tmp/1*k74qnaAcJd3bzRj7PnLIbg.jpeg (inputW=300,inputH=0)
2019/06/22 18:49:30 resize ok: /tmp/0*jraDH1ztolpSmT9I (inputW=300,inputH=0)
```

## goimgrz filter

### 1. filter by file size
```
$ ll -h ./testdata
total 4480
-rw-r--r--@ 1 Terry  access_bpf   1.9M  6 19 02:26 IMG_2489.JPG
-rw-r--r--@ 1 Terry  staff        106K  9 10  2018 gopher2018.png
-rw-r--r--@ 1 Terry  staff        186K  4  2 15:04 web_bg.jpg

$ goimgrz -dir ./testdata -size +1M
2019/06/23 23:02:20 resize ok: /tmp/IMG_2489.JPG (inputW=300,inputH=0)

$ goimgrz -dir ./testdata -size -200k
2019/06/23 23:14:45 resize ok: /tmp/gopher2018.png (inputW=300,inputH=0)
2019/06/23 23:14:45 resize ok: /tmp/web_bg.jpg (inputW=300,inputH=0)

$ goimgrz -dir ./testdata -size -100k
2019/06/23 23:14:40 resize image task is empty, check the parameters
```

### 2. filter by shell pattern 
```
$ goimgrz -dir ./testdata -size -1M -name '*/*.png'
2019/06/23 23:08:22 resize ok: /tmp/gopher2018.png (inputW=300,inputH=0)
```

## Verbose
output detail message synchronously (include error message)

```
$ goimgrz -dir ./testdata -size -1M -name '*/*.png' -verbose
2019/06/23 23:13:37 resize fail: error(30): task filter detect name pattern not match, pattern:*/*.png, name:testdata/IMG_2489.JPG
2019/06/23 23:13:37 resize fail: error(30): task filter detect name pattern not match, pattern:*/*.png, name:testdata/web_bg.jpg
2019/06/23 23:13:37 resize ok: /tmp/gopher2018.png (inputW=300,inputH=0)
```

### Support image format
- jpeg
- png
- gif

## Bug
1. Now, http image download is no timeout(used default Http Client), this will change soon

## Todo
1. [ ] Add more image format support. (like webp and svg)
2. [ ] Support crawler_url, convenience for crawler crawl the whole html page
   image and resize them to an folder
3. [ ] Update the command flag, support short mode
4. [ ] Update Http Url Image download model, support timeout
5. [ ] Support HttpProxy Setting (Because network block)
