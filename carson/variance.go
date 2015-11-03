package main

import (
   "fmt"
   "math"
   "os"
   "flag"
   "github.com/strava/go.strava"
)

func main() {
   var accessToken string

   // Provide an access token, with write permissions.
   // You'll need to complete the oauth flow to get one.
   flag.StringVar(&accessToken, "token", "", "Access Token")

   flag.Parse()

   if accessToken == "" {
      fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")

      flag.PrintDefaults()
      os.Exit(1)
   }

   segmentIds := [...]int64{365235, 6452581, 664647, 1089563, 4956199, 2187, 5732938, 654030, 616554, 3139189}

   client := strava.NewClient(accessToken)
   service := strava.NewSegmentsService(client)
   var times []float64

   for i := 0; i < len(segmentIds); i++ {
      times = nil
      segmentId := segmentIds[i]
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

      fmt.Printf("%s, %s %d times by %d athletes\n\n", segment.Name, verb, segment.EffortCount, segment.AthleteCount)

      fmt.Printf("Fetching leaderboard...\n")

      times = make([]float64, segment.AthleteCount, segment.AthleteCount)

      pageNum := 1
      index := 0

      for index < segment.AthleteCount {
        results, err := service.GetLeaderboard(segmentId).Page(pageNum).PerPage(200).Do()

         if err != nil {
            fmt.Println(err)
            os.Exit(1)
         }

         for _, result := range results.Entries {
            times[index] = float64(result.ElapsedTime)
            index++
         }

         pageNum++
      }

      fmt.Printf("Variance of elapsed times (min): %f\n", variance(times) / 60.0)
      fmt.Printf("StdDev of elapsed times (min): %f\n\n", stdDev(variance(times)) / 60.0)
      fmt.Printf("---------------------------------------------\n\n")
   }
}

func variance(values []float64) float64 {
   mean := mean(values)

   sumOfSquares := 0.0

   for i := 0; i < len(values); i++ {
      sumOfSquares = sumOfSquares + math.Pow(values[i] - mean, 2)
   }

   return sumOfSquares / float64(len(values))
}

func mean(values []float64) float64 {
   sum := 0.0

   for i := 0; i < len(values); i++ {
      sum = sum +  values[i]
   }

   return sum / float64(len(values))
}

func stdDev(variance float64) float64 {
   return math.Sqrt(variance)
}