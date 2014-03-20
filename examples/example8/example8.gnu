reset

set term postscript eps enhanced monochrome "Helvetica" 18 
set style line 1 lt 1 lc 1 lw 1 pt 9 ps 2
set style line 2 lt 1 lc 2 lw 1 pt 9 ps 2
set style line 3 lt 1 lc 3 lw 1 pt 9 ps 2
set style line 4 lt 2 lc 1 lw 1 pt 9 ps 2
set style line 5 lt 2 lc 2 lw 1 pt 9 ps 2
set style line 6 lt 2 lc 3 lw 1 pt 9 ps 2

unset key
set rmargin 4.5

set output "example8.eps"
set xlabel "t (s)" offset 0,0.4
set xrange[0:1.]
set yrange[-1.5:1.5]
set ytics 1

set ylabel "m_z ()" offset 1.5,0
unset grid

set size 1.15,0.5

file="./example8.out/table.txt"

plot file using ($1):4 ls 3 w l notitle 

