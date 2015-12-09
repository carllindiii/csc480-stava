// segment_hill.go fetches the given segments and
// labels whether they are an uphill or downhill segment
// This will be used to assist us when looking at average
// speed of a segment
package elevation

import (
	"github.com/strava/go.strava"
)

// returns value between 0, 1 for how uphill/downhill segment is
// 0 is maximum downhilly-ness, 1 is max uphilly-ness
func IsUphill(segment *strava.SegmentDetailed) float64 {
	if segment.TotalElevationGain >= -5 && segment.TotalElevationGain <= 5 {
      return 0.2
   } else if segment.TotalElevationGain >= -10 && segment.TotalElevationGain < -5 {
      return 0.4
   } else if segment.TotalElevationGain > -10 {
      return 0.6
   } else if segment.TotalElevationGain > 5 && segment.TotalElevationGain <= 10 {
      return 0.8
   } else {
      return 1.0
   }
}
