reset

set term postscript eps enhanced monochrome "Helvetica" 18 
set style line 1 lt 1 lc 1 lw 1 pt 9 ps 2
set style line 2 lt 1 lc 2 lw 1 pt 9 ps 2
set style line 3 lt 1 lc 3 lw 1 pt 9 ps 2
set key right bottom

set rmargin 4.5

set output "example2.eps"
set xlabel "t (ns)" offset 0,0.4
set xrange[0:100]

set ylabel "m ()" offset 2.3,0
unset grid

set size 1.15,1

file="./example2.out/table.txt"
filemumax="./example2_mumax_output.txt"

plot file using ($1*1e9):2 ls 1 w l title "vinamax <m_x>",\
 file u ($1*1e9):3 ls 2 w l title "vinamax <m_y>",\
 file u ($1*1e9):4 ls 3 w l title "vinamax <m_z>",\
 filemumax u ($1*1e9):2 every 10 pt 7 lc 1 w p t "mumax <m_x>",\
 filemumax u ($1*1e9):3 every 10 pt 7 lc 2 w p t "mumax <m_y>",\
 filemumax u ($1*1e9):4 every 10 pt 7 lc 3 w p t "mumax <m_z>"
