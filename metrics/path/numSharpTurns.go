package path

import (
   "math"
   "github.com/go-gl/mathgl/mgl64"
)

func NumSharpTurns(points [][2]float64) int {
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
