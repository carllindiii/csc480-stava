package path

import (
   "math"
)

func ClassifyDifficultyWithSimplificationCount(points [][2]float64) float64 {
   var value = CountDRPSimplifications(points)

   if value > 40 {
      return 1.0
   } else if value > 20 {
      return 0.5
   } else {
      return 0.0
   }
}

const epsilonRatio float64 = 0.2

func CountDRPSimplifications(points [][2]float64) int64 {
   var averageSegmentLength float64 = AveragePolylineSegmentLength(points)
   return CountDRPSimplificationsWithEpsilon(points, averageSegmentLength * epsilonRatio)
}

// This is a modification of the Ramer-Douglas-Peucker algorithm that just
// counts the number of times that it would simplify the path.
func CountDRPSimplificationsWithEpsilon(points [][2]float64, epsilon float64) int64 {
   // Base case: 2 or fewer points.
   if len(points) <3 {
      return 1
   } else {
      // First, figure out if this path would be simplified.
      // To do this, start by finding the point with the greatest distance
      // from the overall line segment.
      var maxDistance float64 = 0
      var maxDistanceIndex int64 = -1

      var wholeSegment [2][2]float64
      wholeSegment[0] = points[0]
      wholeSegment[1] = points[len(points) - 1]

      for i := 0; i < len(points); i++ {
         var distance float64 = getDistanceFromSegment(points[i], wholeSegment)
         if distance > maxDistance {
            maxDistance = distance
            maxDistanceIndex = int64(i)
         }
      }

      if maxDistance > epsilon {
         // If the distance is too great, split the problem
         // Note that both subproblems include the split point. This is fine.
         var firstSubproblemPoints [][2]float64 = points[:maxDistanceIndex]
         var secondSubproblemPoints [][2]float64 = points[maxDistanceIndex:]

         // Recurse on the subproblems
         var firstSimplificationCount = CountDRPSimplificationsWithEpsilon(firstSubproblemPoints, epsilon)
         var secondSimplificationCount = CountDRPSimplificationsWithEpsilon(secondSubproblemPoints, epsilon)

         return firstSimplificationCount + secondSimplificationCount
      } else {
         // If the points are all within epsilon of the overall line segment,
         // then just do (count) a single simplification.
         return 1
      }
   }
}

// Return some measure of the distance of 'point' from the line segment
// 'segment', an array of two points.
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
      return math.Sqrt(AtoP[0] * AtoP[0] + AtoP[1] * AtoP[1])
   } else if dotB < 0 {
      // Find the euclidean distance from P to B.
      return math.Sqrt(BtoP[0] * BtoP[0] + BtoP[1] * BtoP[1])
   } else {
      // Otherwise find the point to line distance.
      // from mathworld.wolfram.com/Point-LineDistance2-Dimensional.html
      // equation (14)
      var numerator = math.Abs(AtoP[0] * AtoB[1] - AtoP[1] * AtoB[0])
      var denominator = math.Sqrt(AtoB[0] * AtoB[0] + AtoB[1] * AtoB[1])
      return numerator / denominator
   }
}
