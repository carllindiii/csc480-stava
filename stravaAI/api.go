package stravaAI

import (
   "./metrics/path"
   "./metrics/time"
   "./metrics/elevation"

   "github.com/strava/go.strava"
)

// "github.com/carllindiii/csc480-stava/stravaAI

type Client struct {
   stravaClient *strava.Client
   segmentSrvc *strava.SegmentsService
}

type DIFFICULTY_LABEL int

const (
   EASY DIFFICULTY_LABEL = iota
   MEDIUM = iota
   HARD = iota
)

func NewClient(stravaClient *strava.Client) Client {
   var stravaAI Client
   stravaAI.stravaClient = stravaClient
   stravaAI.segmentSrvc = strava.NewSegmentsService(stravaClient)

   return stravaAI
}

func neuralToDifficultyLabel(neuralVal float64) DIFFICULTY_LABEL {
   if (neuralVal <= 0.4) {
      return EASY
   } else if (neuralVal <= 0.7) {
      return MEDIUM
   }
   return HARD
}

func (client *Client) GetSegmentIdDifficulty(segId int64) (DIFFICULTY_LABEL, error) {
   segment, err := client.segmentSrvc.Get(segId).Do()
   if (err != nil) {
      return -1, err
   }
   return client.GetSegmentDifficulty(segment), nil
}

func (client *Client) GetSegmentDifficulty(segment *strava.SegmentDetailed) DIFFICULTY_LABEL {
   var latLngCrds [][2]float64 = segment.Map.Polyline.Decode()
   
   numSharpTurns_difficulty, _ := path.GetNumSharpTurns(latLngCrds)
   simplificationCount_difficulty := path.ClassifyDifficultyWithSimplificationCount(latLngCrds)
   pace_difficulty, _, _ := time.GetPace(segment, client.segmentSrvc)
   elevationDifficulty := elevation.IsUphill(segment)
   stdDevLeaderBoard_neural, _ := time.StdDevOfLeaderBoard(segment, client.segmentSrvc)

   difficulty := numSharpTurns_difficulty * .2 + simplificationCount_difficulty * .2 + pace_difficulty * .2 + elevationDifficulty * .2 + stdDevLeaderBoard_neural * .2

   return neuralToDifficultyLabel(difficulty)
}