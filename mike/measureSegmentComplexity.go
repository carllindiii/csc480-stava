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

   measureSegmentComplexity(segmentId, accessToken)
}

func getSegment(segmentId int64, accessToken string) SegmentDetailed {
   var client *Client = strava.NewClient(accessToken)
   var segmentService *SegmentsService = strava.NewSegmentsService(client)
   var segment SegmentDetailed
   var err error
   segment, err = segmentService.Get(segmentId).Do()

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   return segment
}

func measureSegmentComplexity(segmentId int64, accessToken string) {
   var segment SegmentDetailed = getSegment(segmentId, accessToken)
   var polyline Polyline = segment.Polyline

   var points [][2]float64 = Polyline.Decode()

   fmt.Println(points)
}
