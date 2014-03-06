#!/bin/bash
cd example1
go run example1.go
gnuplot example1.gnu
epstopdf example1.eps
rm example1.eps
rm -r example1.out

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example1.pdf -quality 100 -sharpen 0x1.0 example1.png

cd ../example2
go run example2.go
go run without_demag.go
gnuplot example2.gnu
epstopdf example2.eps
rm example2.eps
rm -r example2.out

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example2.pdf -quality 100 -sharpen 0x1.0 example2.png

#commented out because long examples
cd ../example3
#go run example3.go
#go run nofmm.go
#go run wode.go
gnuplot example3.gnu
epstopdf example3.eps
rm example3.eps
#rm -r example3.out
#rm -r wode.out
#rm -r nofmm.out

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example3.pdf -quality 100 -sharpen 0x1.0 example3.png

cd ../example4
go run example4.go
gnuplot example4.gnu
epstopdf example4.eps
rm example4.eps
rm -r example4.out

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

cd ../example6
go run example6.go
gnuplot example6.gnu
epstopdf example6.eps
epstopdf lognormal.eps
rm example6.eps
rm lognormal.eps
rm -r example6.out
#only to update the figures for the webpage
#convert -verbose -density 125 -trim example6.pdf -quality 100 -sharpen 0x1.0 example6.png
#convert -verbose -density 125 -trim lognormal.pdf -quality 100 -sharpen 0x1.0 lognormal.png

cd ../example7
go run example7.go
go run notemp.go
gnuplot example7.gnu
epstopdf example7.eps
rm example7.eps
rm -r example7.out
rm -r notemp.out

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example7.pdf -quality 100 -sharpen 0x1.0 example7.png

#commented out because long simulation
cd ../example8
go run example8.go
gnuplot example8.gnu
epstopdf example8.eps
rm example7.eps
#rm -r example7.out

#only to update the figures for the webpage
#convert -verbose -density 125 -trim example7.pdf -quality 100 -sharpen 0x1.0 example7.png


