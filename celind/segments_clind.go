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
	const SEGMENT_SIZE = 10
	const KM_PER_MILE = 1.60934

	segmentList := [SEGMENT_SIZE]int64{365235, 6452581, 664647, 1089563, 4956199,
					 		2187, 5732938, 654030, 616554, 3139189};

	// Provide an access token, with write permissions.
	flag.StringVar(&accessToken, "token", "", "Access Token")
	flag.Parse()

	if accessToken == "" {
		fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := strava.NewClient(accessToken)
	service := strava.NewSegmentsService(client)


	for i := 0; i < SEGMENT_SIZE; i++{

		segmentId = segmentList[i]

		fmt.Printf("Fetching segment %d info...\n", segmentId)
		segment, err := service.Get(segmentId).Do()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		verb := "ridden"
		if segment.ActivityType == strava.ActivityTypes.Run {
			verb = "run"
		}
		fmt.Printf("%s, %s %d times by %d athletes\n",
				 segment.Name, verb, segment.EffortCount, segment.AthleteCount)

		//for all the segments, grab and calculate all the segment efforts
		var totalTime int
		// return list of segment efforts
		segmentEfforts, err := service.ListEfforts(segmentId).Do();

		for _, segmentEffortSummary := range segmentEfforts {
   			totalTime += segmentEffortSummary.MovingTime;
		}   

		avgPace := float64(totalTime / segment.EffortCount)
		fmt.Printf("Avg Pace %.2f\n", avgPace);

		kmDistance := segment.Distance / 1000

		avgPacePerKm := avgPace / kmDistance
		avgPacePerMile := avgPacePerKm * KM_PER_MILE


		fmt.Printf("Total distance of %.2f km %s with an average pace of %f per km (%.2f per mile)\n\n",
			kmDistance, verb, avgPacePerMile, avgPacePerMile)

	}
}
