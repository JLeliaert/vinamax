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
go run nofmm.go
go run wode.go
gnuplot example3.gnu
epstopdf example3.eps
rm example3.eps

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example3.pdf -quality 100 -sharpen 0x1.0 example3.png

cd ../example4
go run example4.go
gnuplot example4.gnu
epstopdf example4.eps
rm example4.eps

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example4.pdf -quality 100 -sharpen 0x1.0 example4.png

cd ../example5
go build example5.go
./example5 -cpuprofile=run.prof
go tool pprof --pdf ./example5 run.prof >profile.pdf
rm example5
rm -r example5.out
rm run.prof
#only to update the figures for the webpage
#convert -verbose -density 125 -trim profile.pdf -quality 100 -sharpen 0x1.0 profile.png

