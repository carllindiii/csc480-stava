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

   flag.Int64Var(&segmentId, "id", 229781, "Strava Segment Id")
   flag.StringVar(&accessToken, "token", "", "Access Token")

   flag.Parse()

   if accessToken == "" {
      fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")
      flag.PrintDefaults()
      os.Exit(1)
   }

   client := strava.NewClient(accessToken)
   segment, err := strava.NewSegmentsService(client).Get(segmentId).Do()

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   verb := "ridden"
   if segment.ActivityType == strava.ActivityTypes.Run {
      verb = "run"
   }

   fmt.Printf("%s, %s %d times by %d athletes\n\n", segment.Name, verb, segment.EffortCount, segment.AthleteCount)

   fmt.Printf("Fetching leaderboard...\n")
   results, err := strava.NewSegmentsService(client).GetLeaderboard(segmentId).Do()
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   for _, e := range results.Entries {
      fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
   }
}
