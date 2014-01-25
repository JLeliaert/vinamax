#!/bin/bash
go run example1.go
gnuplot example1.gnu
epstopdf example1.eps
rm example1.eps
cp example1.pdf ../../vinamax/images/
