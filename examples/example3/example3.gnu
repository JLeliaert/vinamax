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

set output "example3.eps"
set xlabel "t (ns)" offset 0,0.4
set xrange[0:150.]

set ylabel "m ()" offset 2.3,0
unset grid

set size 1.15,1

file="./example3.out/table.txt"
filewode="./wode.out/table.txt"
filenofmm="./nofmm.out/table.txt"

plot filenofmm using ($1*1e9):2 ls 1 w l title "brute force <m_x>",\
 filenofmm u ($1*1e9):3 ls 2 w l title "brute force <m_y>",\
 filenofmm u ($1*1e9):4 ls 3 w l title "brute force <m_z>",\
 file u ($1*1e9):2 every 10 pt 7 lc 1 w p t "Dipole approximation <m_x>",\
 file u ($1*1e9):3 every 10 pt 7 lc 2 w p t "Dipole approximation <m_y>",\
 file u ($1*1e9):4 every 10 pt 7 lc 3 w p t "Dipole approximation <m_z>",\
 filewode using ($1*1e9):2 ls 4 w l title "no demag <m_x>",\
 filewode u ($1*1e9):3 ls 5 w l title "no demag <m_y>",\
 filewode u ($1*1e9):4 ls 6 w l title "no demag <m_z>"

