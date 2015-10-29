// segment_example.go provides a simple example to fetch a segment details
// and list the top 10 on the leaderboard.
//
// usage:
//   > go get github.com/strava/go.strava
//   > cd $GOPATH/github.com/strava/go.strava/examples
//   > go run segment_example.go -id=segment_id -token=access_token
//
//   You can find an access_token for your app at https://www.strava.com/settings/api
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


	//get distance, average_grade, elevation_high, elevation_low 
	//climb_category, total_elevation_gain, hazardous
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
