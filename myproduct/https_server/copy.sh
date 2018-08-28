#!/bin/sh
number=1
while test $number -le 100
do
    cp img/image.png img/image_$number.png
    number=`expr $number + 1`
done
