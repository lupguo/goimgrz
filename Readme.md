## Girls
Girls is not girl, it just short for **Go Image Resize List**, It is used
to resize network or local images

## Install
```
go install github.com/tkstorm/girls
```

### Help
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

### girls resize local dir images
```
$ girls -scan_dir ./mindnode -width 600 -dst /tmp/mindnode_600
2019/06/22 12:36:34 resize ok: /tmp/mindnode_600/golang.org_x_icmp.png (inputW=600,inputH=0)
2019/06/22 12:36:34 resize ok: /tmp/mindnode_600/golang.org_regexp.png (inputW=600,inputH=0)
2019/06/22 12:36:34 resize ok: /tmp/mindnode_600/golang.org_strings.png (inputW=600,inputH=0)
2019/06/22 12:36:34 resize ok: /tmp/mindnode_600/golang.org_bufio.png (inputW=600,inputH=0)
2019/06/22 12:36:35 resize ok: /tmp/mindnode_600/golang.org_x_fmt.png (inputW=600,inputH=0)
2019/06/22 12:36:35 resize ok: /tmp/mindnode_600/golang.org_time.png (inputW=600,inputH=0)
2019/06/22 12:36:36 resize ok: /tmp/mindnode_600/golang.org_flag.png (inputW=600,inputH=0)
2019/06/22 12:36:36 resize ok: /tmp/mindnode_600/golang.org_io.png (inputW=600,inputH=0)
2019/06/22 12:36:36 resize ok: /tmp/mindnode_600/golang.org_image.png (inputW=600,inputH=0)
2019/06/22 12:36:37 resize ok: /tmp/mindnode_600/golang.org_database_sql.png (inputW=600,inputH=0)
2019/06/22 12:36:37 resize ok: /tmp/mindnode_600/golang.org_os(os,file,path).png (inputW=600,inputH=0)

$ girls -scan_dir ./mindnode -width 1200 -dst /tmp
2019/06/22 12:35:43 error(21): dir is not exist ./mindnode, stat ./mindnode: no such file or directory
```

### girls resize http image
```
$ girls -img_url https://physicsworld.com/wp-content/uploads/2006/09/LLOYDblack-hole-635x496.jpg -dir /tmp
2019/06/22 00:08:26 resize ok: /tmp/LLOYDblack-hole-635x496.jpg (inputW=300,inputH=0)
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