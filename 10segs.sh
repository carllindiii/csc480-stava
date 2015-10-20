#!/bin/bash

# this script fetches all the segment efforts for our 10 segments
# writes results to dist/segmentEfforts_*.json
# usage: ./10segs.sh stravaToken

stravaToken="$1"
authStr="Authorization: Bearer $stravaToken"

mkdir dist
mkdir dist/easy
mkdir dist/medium
mkdir dist/hard

if [ "$#" -ne 1 ]; then
	echo 'usage: ./10segs.sh stravaToken'
	exit 1
fi

segments=(365235 6452581 664647 1089563 4956199 2187 5732938 654030 616554 3139189)

hard=(365235 6452581 1089563 4956199)
medium=(664647 2187 5732938)
easy=(654030 616554 3139189)

fetchSegmentEfforts() {
	segId="$1"

	curl -G "https://www.strava.com/api/v3/segments/$segId/all_efforts" -H "$authStr" | python -m json.tool
}

for seg in ${easy[@]}; do
	fetchSegmentEfforts $seg > dist/easy/"segmentEfforts_$seg.json" &
done

for seg in ${medium[@]}; do
	fetchSegmentEfforts $seg > dist/medium/"segmentEfforts_$seg.json" &
done

for seg in ${hard[@]}; do
	fetchSegmentEfforts $seg > dist/hard/"segmentEfforts_$seg.json" &
done

wait

node 10segsHelper.js > dist/stats.txt

echo "writing results to dist/stats.txt"
