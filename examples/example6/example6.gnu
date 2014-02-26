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

set output "example6.eps"
set xlabel "t (ns)" offset 0,0.4
set xrange[0:8]

set ylabel "<m> ()" offset 1.5,0
unset grid

set size 1.15,1

file="./example6.out/table.txt"
g(x)=0.5

plot file using ($1*1e9):(sqrt($2*$2+$3*$3+$4*$4)) ls 1 w l title " <m>", g(x) ls 2 w l title "0.5}"
########################################################################
datafile="./example6.out/m000000.txt"

s=1
m=20
f(x)= 20000./(sqrt(2*pi)*s*x*1e9)*exp(-1./(2*s*s)*log(x*1e9/m)*log(x*1e9/m))
set xlabel "diameter (m)"
set ylabel "#particles"

set output "lognormal.eps"
binwidth=1e-9
width=1e-9
set xrange[0:1e-7]
set boxwidth binwidth
set style fill solid 0.5 #fillstyle
bin(x,width)=width*floor(x/width) + binwidth/2.0
plot datafile using (bin(2*$4,binwidth)):(1.0) smooth freq with boxes lc rgb"blue" title "diameter distribution", f(x) ls 1 w l title "theoretical"
