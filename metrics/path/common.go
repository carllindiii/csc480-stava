package path

import (
   "math"
)

func PolylineLength(points [][2]float64) float64 {
   var total float64 = 0

   for i := 0; i < len(points) - 1; i++ {
      var xDistance = points[i + 1][0] - points[i][0]
      var yDistance = points[i + 1][1] - points[i][1]
      total += math.Sqrt(xDistance * xDistance + yDistance * yDistance)
   }

   return total
}

func AveragePolylineSegmentLength(points [][2]float64) float64 {
   if len(points) <= 1 {
      return 0.0
   } else {
      return PolylineLength(points) / float64(len(points) - 1)
   }
}
