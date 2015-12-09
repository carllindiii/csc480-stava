package main

import (
   "flag"
   "fmt"
   "os"
   "github.com/carllindiii/csc480-stava/stravaAI"
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

   stravaClient := strava.NewClient(accessToken)
   stravaAIClient := stravaAI.NewClient(stravaClient)

   segIds := [10]int64{365235, 6452581, 664647, 1089563, 4956199, 2187, 5732938, 654030, 616554, 3139189}

   for _, nextSegId := range segIds {
      fmt.Println(stravaAIClient.GetSegmentIdDifficulty(nextSegId))
   }
}