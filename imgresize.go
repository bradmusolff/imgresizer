package main

import (
    "github.com/nfnt/resize"
    "github.com/pborman/getopt"
    "image/jpeg"
    "log"
    "os"
)

func main() {


    numargs := len(os.Args)
    var file *os.File
    var out *os.File

    if numargs < 2 {
            log.Fatal("Usage: imgresize [ -i inputfile -o outputfile ] -w width")
    }

  //  helpFlag := getopt.Bool('?', "display help")
    outputfile := getopt.StringLong("output", 'o', "", "output file name")
    inputfile := getopt.StringLong("input",'i',"","input file name")
    width := getopt.UintLong("width",'w',640,"new width")
    var opts = getopt.CommandLine

    opts.Parse(os.Args)
    if (((outputfile == nil) || (inputfile == nil)) && !(outputfile == nil && inputfile == nil) ) {
       log.Fatal("Usage: imgresize [ -i inputfile -o outputfile ] -w width ")
    }

    if (*inputfile != "") {
       var err error
       file, err = os.Open(*inputfile)
       if err != nil {
           log.Fatal(err)
       }
    } else {
       file = os.Stdin
    }


    // decode jpeg into image.Image
    img, err := jpeg.Decode(file)
    if err != nil {
	      log.Fatal(err)
    }
    file.Close()


    m := resize.Resize(*width, 0, img, resize.Lanczos3)
    if (*outputfile != "") {
      var err error
      out, err = os.Create(*outputfile)
      if err != nil {
  	      log.Fatal(err)
      }
      defer out.Close()
    } else {
      out = os.Stdout
    }
    // write new image to file
    jpeg.Encode(out, m, nil)
}
