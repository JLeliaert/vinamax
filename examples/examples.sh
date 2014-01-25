#!/bin/bash
go run example1.go
gnuplot example1.gnu
epstopdf example1.eps
rm example1.eps

#only to update the figure for the webpage
#convert -verbose -density 125 -trim example1.pdf -quality 100 -sharpen 0x1.0 example1.png
