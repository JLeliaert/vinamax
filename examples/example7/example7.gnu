reset

set term postscript eps enhanced monochrome "Helvetica" 18 
set style line 1 lt 1 lc 1 lw 1 pt 9 ps 2
set style line 2 lt 1 lc 2 lw 1 pt 9 ps 2
set style line 3 lt 1 lc 3 lw 1 pt 9 ps 2
set style line 4 lt 2 lc 1 lw 1 pt 9 ps 2
set style line 5 lt 2 lc 2 lw 1 pt 9 ps 2
set style line 6 lt 2 lc 3 lw 1 pt 9 ps 2

set key right top

set rmargin 4.5

set output "example7.eps"
set xlabel "t (ns)" offset 0,0.4
set xrange[0:100]

set ylabel "m_z ()" offset 1.5,0
unset grid

set size 1.15,1

file="./example7.out/table.txt"
filenotemp="./notemp.out/table.txt"

plot file using ($1*1e9):4 ls 3 w l title "<m_z> at 300 K",\
 filenotemp u ($1*1e9):4 ls 6 w l title "<m_z> at 0 K"

