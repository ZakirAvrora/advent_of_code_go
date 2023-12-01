#!/bin/bash
y=2023
for i in `seq -w 1 25`; do
    mkdir $y/day-$i
    touch $y/day-$i/README.md
    mkdir $y/day-$i/part1
    mkdir $y/day-$i/part2
    touch $y/day-$i/part1/main.go
    touch $y/day-$i/part2/main.go
    touch input.txt
done