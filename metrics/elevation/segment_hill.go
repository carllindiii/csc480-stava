// segment_hill.go fetches the given segments and
// labels whether they are an uphill or downhill segment
// This will be used to assist us when looking at average
// speed of a segment
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/strava/go.strava"
)

func main() {
	var segmentId int64
	var accessToken string
	const SIZE = 10

	segmentList := [SIZE]int64{365235, 6452581, 664647, 1089563, 4956199,
					 		2187, 5732938, 654030, 616554, 3139189};

	// Provide an access token, with write permissions.
	// You'll need to complete the oauth flow to get one.
	//flag.Int64Var(&segmentId, "id", 229781, "Strava Segment Id")
	flag.StringVar(&accessToken, "token", "1db87db0b3dd8fa569bc8bb45c5cdbe544ddd4b5", "Access Token")

	flag.Parse()

	if accessToken == "" {
		fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")

		flag.PrintDefaults()
		os.Exit(1)
	}

	client := strava.NewClient(accessToken)


	for i := 0; i < SIZE; i++{

		segmentId = segmentList[i]

		fmt.Printf("Fetching segment %d info...\n", segmentId)
		segment, err := strava.NewSegmentsService(client).Get(segmentId).Do()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		verb := "ridden"
		if segment.ActivityType == strava.ActivityTypes.Run {
			verb = "run"
		}
		fmt.Printf("%s, %s %d times by %d athletes\n\n",
				 segment.Name, verb, segment.EffortCount, segment.AthleteCount)

		if segment.TotalElevationGain > 0{
			fmt.Printf("Uphill\n\n")
		}else{
			fmt.Printf("Downhill\n\n")
		}
	}
}
