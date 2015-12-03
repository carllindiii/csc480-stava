package main

import (
   "flag"
   "fmt"
   "os"

   "github.com/strava/go.strava"

   "./metrics/path"
   "./metrics/time"
   "./metrics/elevation"
   "./carson"
)

func main() {
   var segmentId int64
   var accessToken string

   // Provide an access token, with write permissions.
   // You'll need to complete the oauth flow to get one.
   flag.Int64Var(&segmentId, "id", 229781, "Strava Segment Id")
   flag.StringVar(&accessToken, "token", "", "Access Token")

   flag.Parse()

   if accessToken == "" {
      fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")

      flag.PrintDefaults()
      os.Exit(1)
   }

   client := strava.NewClient(accessToken)

   segmentService := strava.NewSegmentsService(client)

   segIds := [10]int64{365235, 6452581, 664647, 1089563, 4956199, 2187, 5732938, 654030, 616554, 3139189}


   for _, nextSegId := range segIds {
      fmt.Printf("nextSegId: %d\n", nextSegId);

      segment, err := segmentService.Get(nextSegId).Do()
      var latLngCrds [][2]float64 = segment.Map.Polyline.Decode()

      if err != nil {
         fmt.Println(err)
         os.Exit(1)
      }
      
      numSharpTurns_difficulty, numSharpTurns := path.GetNumSharpTurns(latLngCrds)
      fmt.Printf("\tsharp turns (catnullRom): %d (%f)\n", numSharpTurns, numSharpTurns_difficulty)

      simplificationCount_difficulty := path.ClassifyDifficultyWithSimplificationCount(latLngCrds)
      fmt.Printf("\tsimplificationCount_difficulty: %f\n", simplificationCount_difficulty)

      pace_difficulty, paceTop30, paceAll := time.GetPace(segment, segmentService)
      fmt.Printf("\tpace: %f, %f (%f)\n", paceTop30, paceAll, pace_difficulty)

      uphillyNess := elevation.IsUphill(segment)
      fmt.Printf("\tis uphill: %f\n", uphillyNess)

      stdDevLeaderBoard_neural, stdDevLeaderBoard := carson.StdDevOfLeaderBoard(segment, segmentService)
      fmt.Printf("\tstdDev of leaderboard: %f (%f)\n", stdDevLeaderBoard, stdDevLeaderBoard_neural)


   }

   
}