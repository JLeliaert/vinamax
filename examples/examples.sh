#!/bin/bash
cd example1
go run example1.go
gnuplot example1.gnu
epstopdf example1.eps
rm example1.eps

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example1.pdf -quality 100 -sharpen 0x1.0 example1.png

cd ../example2
go run example2.go
gnuplot example2.gnu
epstopdf example2.eps
rm example2.eps

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example2.pdf -quality 100 -sharpen 0x1.0 example2.png
