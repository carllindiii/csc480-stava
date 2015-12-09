package main

import (
   "https://github.com/carllindiii/csc480-stava"
   "github.com/strava/go.strava"
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

   csvF, err := os.Create("results.csv")

   defer csvF.Close()

   if (err != nil) {
      fmt.Println(err)
      os.Exit(1)
   }


   writer := csv.NewWriter(csvF)

   writer.Write([]string{"segmentId, shapTurns, simplication, pace, uphill, leaderboard"})
   writer.Flush()


   for _, nextSegId := range segIds {
      fmt.Printf("nextSegId: %d\n", nextSegId);

      segment, err := segmentService
      stravaAI.getSegmentDifficulty(segment)
   }
}