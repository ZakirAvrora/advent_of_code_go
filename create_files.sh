#!/bin/bash
y=2023
for i in `seq -w 1 25`; do
    mkdir $y/day-$i
    touch $y/day-$i/README.md
    touch $y/day-$i/day-$i-part-1.go
    touch $y/day-$i/day-$i-part-2.go
    touch $y/day-$i/day-$i-input.txt
done