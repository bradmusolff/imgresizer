package main

import (
    "github.com/nfnt/resize"
    "image/jpeg"
    "log"
    "strconv"
    "os"
)

func main() {

    var width uint = 0
    args := os.Args[1:]
    numargs := len(os.Args)

    if numargs < 2 {
       log.Fatal("Usage: imgresize <inputfile> <outputfile> [width]")
    }

    file, err := os.Open(args[0])
    if err != nil {
        log.Fatal(err)
    } 

    // decode jpeg into image.Image
    img, err := jpeg.Decode(file)
    if err != nil {
	log.Fatal(err)
    }
    file.Close()
    if numargs >=3 {
 
        wd,err := strconv.ParseUint(args[2],0,32)
        if err != nil {
              log.Fatal(err)
        }
        width=uint(wd) 
    } else {
        width=640
    }
    
    m := resize.Resize(width, 0, img, resize.Lanczos3)

    out, err := os.Create(args[1])
    if err != nil {
	log.Fatal(err)
    }
    defer out.Close()

    // write new image to file
    jpeg.Encode(out, m, nil)
}
