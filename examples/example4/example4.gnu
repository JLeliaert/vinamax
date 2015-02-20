reset

set term postscript eps enhanced monochrome "Helvetica" 18 
set style line 1 lt 1 lc 1 lw 1 pt 9 ps 2
set style line 2 lt 1 lc 2 lw 1 pt 9 ps 2
set style line 3 lt 1 lc 3 lw 1 pt 9 ps 2
set key right bottom

set rmargin 4.5

set output "example4.eps"
set xlabel "B_{ext} along the x-axis (T)" offset 0,0.4

set ylabel "(B_x^2+B_y^2)^{0.5} ,1 mm above the sample (T)" offset 2,0
unset grid

set size 1.15,1

file="./example4.out/table.txt"

plot file using 7:(sqrt($9*$9+$10*$10)) ls 1 w l notitle
