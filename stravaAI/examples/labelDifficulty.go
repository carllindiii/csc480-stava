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
   flag.Int64Var(&segmentId, "id", -1, "Strava Segment Id")
   flag.StringVar(&accessToken, "token", "", "Access Token")

   flag.Parse()

   if accessToken == "" {
      fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")

      flag.PrintDefaults()
      os.Exit(1)
   }

   if segmentId <0 {
      fmt.Println("\nPlease provide a segment id")

      flag.PrintDefaults()
      os.Exit(1)
   }

   stravaClient := strava.NewClient(accessToken)
   stravaAIClient := stravaAI.NewClient(stravaClient)

   fmt.Println("please wait...")

   difficulty, err := stravaAIClient.GetSegmentIdDifficulty(segmentId)

   if (err != nil) {
      fmt.Println("something went wrong");
      fmt.Println(err)
   }

   if difficulty == stravaAI.EASY {
      fmt.Println("that's an easy segment")
   } else if difficulty == stravaAI.MEDIUM {
      fmt.Println("watch out. That's a medium segment")
   } else if difficulty == stravaAI.HARD {
      fmt.Println("that's a hard segment. Don't go on it unless you are pro")
   } else {
      fmt.Println("something went wrong")
   }
}