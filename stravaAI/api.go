package stravaAI

import (
   "github.com/strava/go.strava"

   "github.com/carllindiii/csc480-stava/stravaAI/metrics/path"
   "github.com/carllindiii/csc480-stava/stravaAI/metrics/time"
   "github.com/carllindiii/csc480-stava/stravaAI/metrics/elevation"
)

type Client struct {
   stravaClient *strava.Client
   service *strava.SegmentsService
}

type DIFFICULTY_LABEL int

const (
   EASY DIFFICULTY_LABEL = iota
   MEDIUM = iota
   HARD = iota
)

func NewClient(stravaClient strava.Client) *Client {
   var stravaAI Client
   stravaAI.stravaClient = stravaClient
}

func (client *Client) getSegmentDifficulty(segment *strava.SegmentDetailed) DIFFICULTY_LABEL {
   // TODO
   return EASY
}