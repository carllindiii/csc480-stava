package main

import (
   "flag"
   "fmt"
   "os"
   "math"
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

   measureSegmentComplexity(segmentId, accessToken)
}

func getSegment(segmentId int64, accessToken string) strava.SegmentDetailed {
   var client *strava.Client = strava.NewClient(accessToken)
   var segmentService *strava.SegmentsService = strava.NewSegmentsService(client)
   var segment *strava.SegmentDetailed
   var err error
   segment, err = segmentService.Get(segmentId).Do()

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   return *segment
}

func measureSegmentComplexity(segmentId int64, accessToken string) {
   var segment strava.SegmentDetailed = getSegment(segmentId, accessToken)
   var polyline strava.Polyline = segment.Map.Polyline

   var points [][2]float64 = polyline.Decode()

   var complexityScore = countDRPSimplifications(points)

   fmt.Println("Segment ", segmentId, " has a complexity score of ", complexityScore)
}

func countDRPSimplifications(points [][2]float64) int64 {
   var length int64 = points.length

   // Base case: 2 or fewer points.
   if points.length <3 {
      return 0
   } else {
      // First, figure out if this path would be simplified.
      // To do this, start by finding the point with the greatest distance
      // from the overall line segment.
      var maxDistance float64 = 0
      var maxDistanceIndex int64 = -1

      var wholeSegment [2][2]float64
      wholeSegment[0] = points[0]
      wholeSegment[1] = points[wholeSegment.length - 1]

      for var i int64 = 0; i < points.length; i++ {
         var distance float64 = getDistanceFromSegment(points[i], wholeSegment)
         if distance > maxDistance {
            maxDistance = distance
            maxDistanceIndex = i
         }
      }

      // TODO
   }
}

func getDistanceFromSegment(point [2]float64, segment [2][2]float64) float64 {
   // segment[0] is point A, segment[1] is point B, point is point P

   var AtoP [2]float64
   AtoP[0] = point[0] - segment[0][0]
   AtoP[1] = point[1] - segment[0][1]

   var AtoB [2]float64
   AtoB[0] = segment[1][0] - segment[0][0]
   AtoB[1] = segment[1][1] - segment[0][1]

   var BtoP [2]float64
   BtoP[0] = point[0] - segment[1][0]
   BtoP[1] = point[1] - segment[1][1]

   var dotA float64 = AtoB[0] * AtoP[0] + AtoB[1] * AtoP[1]
   var dotB float64 = -AtoB[0] * BtoP[0] + -AtoB[1] * BtoP[1]

   // If either of these dot products is negative, it's a special case
   if dotA < 0 {
      // Find the euclidean distance from P to A.
      return math.sqrt(AtoP[0] * AtoP[0] + AtoP[1] * AtoP[1])
   } else if dotB < 0 {
      // Find the euclidean distance form P to B.
      return math.sqrt(BtoP[0] * BtoP[0] + BtoP[1] * BtoP[1])
   } else {
      // Otherwise find the point to line distance.
      // from mathworld.wolfram.com/Point-LineDistance2-Dimensional.html
      // equation (14)
      var numerator = math.abs(AtoP[0] * AtoB[1] - AtoP[1] * AtoB[0])
      var denominator = math.sqrt(AtoB[0] * AtoB[0] + AtoB[1] * AtoB[1])
      return numerator / denominator
   }
}