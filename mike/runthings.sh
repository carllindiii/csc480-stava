#!/bin/sh

idFile=$1
token=$2

cat $idFile | while read id;
do
   go run measureSegmentComplexity.go -token="$2" -id=$id
done
