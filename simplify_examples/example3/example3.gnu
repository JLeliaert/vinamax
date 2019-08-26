reset

set term postscript eps enhanced monochrome "Helvetica" 18 
set style line 1 lt 1 lc 1 lw 1 pt 7 ps 0.8
set style line 2 lt 1 lc 2 lw 1 pt 7 ps 0.8
set style line 3 lt 1 lc 3 lw 1 pt 7 ps 0.8
set key right bottom

set rmargin 4.5

set output "example3a.eps"
set xlabel "t (ns)" offset 0,0.4

set ylabel "m ()" offset 2.3,0
unset grid

set size 1.15,1
set xrange [0.001:0.0015]
file= "./example3a.out/table.txt" 


plot file using 1:($5*333) w l ls 1 title "B/B_{max}",\
file u 1:2 w l ls 2 title "m_x",\
file u 1:8 w l ls 3 title "u_x"
###################################################
set output "example3b.eps"
set xlabel "t (ns)" offset 0,0.4

set ylabel "m ()" offset 2.3,0
unset grid

set size 1.15,1
set xrange [0.001:0.00105]
file= "./example3b.out/table.txt" 


plot file using 1:($5*333) w l ls 1 title "B/B_{max}",\
file u 1:2 w l ls 2 title "m_x",\
file u 1:8 w l ls 3 title "u_x"
###################################################
set output "example3c.eps"
set xlabel "t (ns)" offset 0,0.4

set ylabel "m ()" offset 2.3,0
unset grid

set size 1.15,1
set xrange [0.001:0.00105]
file= "./example3c.out/table.txt" 


plot file using 1:($5*100) w l ls 1 title "B/B_{max}",\
file u 1:2 w l ls 2 title "m_x",\
file u 1:8 w l ls 3 title "u_x"
