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
go run without_demag.go
gnuplot example2.gnu
epstopdf example2.eps
rm example2.eps

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example2.pdf -quality 100 -sharpen 0x1.0 example2.png

cd ../example3
go run example3.go
go run fmm.go
gnuplot example3.gnu
epstopdf example3.eps
rm example3.eps

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example3.pdf -quality 100 -sharpen 0x1.0 example3.png
