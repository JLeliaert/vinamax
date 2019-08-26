reset

set term postscript eps enhanced monochrome "Helvetica" 18 
set style line 1 lt 1 lc 1 lw 1 pt 7 ps 0.8
set style line 2 lt 1 lc 2 lw 1 pt 7 ps 0.8
set style line 3 lt 1 lc 3 lw 1 pt 7 ps 0.8
set key right top

set rmargin 4.5

set output "example5.eps"
set xlabel "t (ns)" offset 0,0.4

set ylabel "u ()" offset 2.3,0
unset grid

set size 1.15,1

file="./example5.out/table.txt"
tauB=6.553286282752431e-07
f(x)=exp(-x/tauB)

plot file using 1:7 ls 1 w l title "vinamax",\
f(x) ls 2 w l title "theory"
