# ipdf

__convert pdf to image, by `ghostscript` or `ImageMagick`.__


## install [imagick](https://github.com/gographics/imagick)

## pdf 转 image 工具链 for Mac OS

### 1. ghostscript(gs)

 1. 下载链接： http://cactuslab.com/imagemagick/
 

### 2. ImageMagick

 1. 下载链接：http://cactuslab.com/imagemagick/


### 注意

 1. 下载pkg安装目录在 `/opt/`; 

 2. ImageMagick 需要配置环境变量：

```
export PKG_CONFIG_PATH=/opt/ImageMagick/lib/pkgconfig
```

 3. 需要安装X11 ==> https://www.xquartz.org


### port

也可以通过port安装：

```
sudo port install ghostscript
sudo port install ImageMagick
```

_port安装_: 下载 [Macports](https://www.macports.org/)
