package path

func ClassifyDifficultyWithFractalDimension(points [][2]float64) string {
   var value = EstimateFractalDimension(points)

   if value < 0.7 {
      return "Hard"
   } else if value < 0.8 {
      return "Medium"
   } else {
      return "Easy"
   }
}

// TODO: tweak this so that it always keeps the first and last points,
// and so that it always goes down to just those 2 points at the end.
func EstimateFractalDimension(points [][2]float64) float64 {
   var pathLengths []float64
   var steps int64 = 0
   var currentPath [][2]float64 = points
   
   for ; len(currentPath) >= 4; steps++ {
      // Find the current length and put it in the pathLengths slice
      pathLengths = append(pathLengths, PolylineLength(currentPath))

      // Make the new path by skipping every other point in the current path
      var newPath [][2]float64
      for i := 0; 2 * i < len(currentPath); i++ {
         newPath = append(newPath, currentPath[2 * i])
      }

      // Keep the end point in the path
      if len(currentPath) % 2 != 0 {
         newPath = append(newPath, currentPath[len(currentPath) - 1])
      }

      currentPath = newPath
   }

   // Now that we have the list of lengths, average the ratios
   var totalRatio float64
   var step int64

   for step = 0; step < steps - 1; step++ {
      totalRatio += pathLengths[step + 1] / pathLengths[step]
   }

   return totalRatio / float64(steps)
}
