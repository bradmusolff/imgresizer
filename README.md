# imgresizer
Image resizer (golang)
==================

 This is a simple CLI tool for resizing jpeg (currently) images
 Written in Go and originally created to build on Joyent Triton/SmartOS for demoing Manta Compute Job capabilities
 Utilizes the nfnt/resize library


Build
-----
# retrieve nfnt library
go get
go build

Usage:
imgresize <inputfile> <outputfile> [width]
