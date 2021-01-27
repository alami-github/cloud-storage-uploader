# CS-Uploader | ALAMI Sharia DevOps Toolkit

This toolkit is used by our team to upload object to Google Cloud Storage. Mostly our servers run on GNU/Linux so we built on Golang to avoid any dependencies which needed to install first.

Currently we're just provide GNU/Linux (x64) and Microsoft Windows (x64) binary file.

## Pre-requisites:
* Go ([download here](https://golang.org/dl/))
* Create your own bucket of Google Cloud Storage
* After you created one, please follow [this tutorial](https://cloud.google.com/storage/docs/reference/libraries) to create Service-Account also setup necessary things

## How to run locally
```
[GNU/Linux or macOS]
$ go run cs-uploader.go -filepath "/home/h3rucutu/cloud-storage-uploader/README.md" -bucket cs-bucket -directory testing-1 -object README.md

[Microsoft Windows]
> go run .\cs-uploader.go -filepath "C:\\Users\\h3rucutu\\workspace\\alami\\cloud-storage-uploader\\README.md" -bucket database-archive-bucket -directory testing-2 -object README.md
```

## How to run on your server
```
[GNU/Linux (x64)]
$ curl -o cs-uploader -L https://github.com/alami-github/cloud-storage-uploader/releases/download/v0.1.0/cs-uploader-linux-amd64-0.1.0
$ chmod +x cs-uploader
$ ./cs-uploader -filepath "/home/h3rucutu/cloud-storage-uploader/README.md" -bucket cs-bucket -directory testing-1 -object README.md
```

This code brought to you by [ALAMI](https://www.alamisharia.co.id) Engineering Team.

### Contributors:
* [h3rucutu](https://github.com/h3rucutu)