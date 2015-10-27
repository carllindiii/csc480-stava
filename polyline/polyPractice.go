package main

import (
   "flag"
   "fmt"
   "os"
   "math"

   "github.com/strava/go.strava"
   "github.com/go-gl/mathgl/mgl64"
)

func AngleBetween(v1 mgl64.Vec2, v2 mgl64.Vec2) float64 {
   return math.Acos(v1.Dot(v2) / (v1.Len() * v2.Len()))
}

func getCurviness(points [][2]float64) int {
   count := getNumSharpTurns(points)
   if (count > 5) {
      return 2;
   }
   if (count > 3) {
      return 1;
   }
   return 0;
}

func getNumSharpTurns(points [][2]float64) int {
   arrLens := len(points) - 2
   // dists := make([]float64, arrLens)
   angles := make([]float64, arrLens)

   for idx := 0; idx < arrLens; idx++ {
      var p1, p2, p3, v1, v2 mgl64.Vec2

      p1 = points[idx]
      p2 = points[idx + 1]
      p3 = points[idx + 2]

      v1 = p2.Sub(p1)
      v2 = p3.Sub(p2)

      // dists[idx] := v1.Len()
      angles[idx] = AngleBetween(v1, v2) // (radians)
   }

   const maxAngleChange = .4 * math.Pi
   const maxDistChange = 4

   numSharpTurns := 0   // sharp turn is whenn the angles add up to maxAngleChange for 4 consecutive points

   for idx := 0; idx < arrLens; idx++ {
      dAngle := 0.0
      for count := 0; count < maxDistChange && idx < arrLens; count++ {
         dAngle += angles[idx]
         idx++
      }

      if dAngle > maxAngleChange {
         numSharpTurns++
      }
   }

   return numSharpTurns
}

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

   segment, err := segmentService.Get(segmentId).Do()

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   var latLngCrds [][2]float64 = segment.Map.Polyline.Decode()
   fmt.Println(getCurviness(latLngCrds))
}