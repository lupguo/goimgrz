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
