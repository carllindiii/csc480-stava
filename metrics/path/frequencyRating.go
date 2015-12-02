package path

import (
   "math"
   "github.com/mjibson/go-dsp/fft"
)

func ClassifyDifficultyWithFrequencyRating(points [][2]float64) string {
   var value = ComputeFrequencyRating(points)

   if value > 400 {
      return "Hard"
   } else if value > 200 {
      return "Medium"
   } else {
      return "Easy"
   }
}

// This basically takes the inner product of the FFT of the path treated as a
// complex signal and a linear function with positive values for
// high frequencies and negative values for low frequencies.
func ComputeFrequencyRating(points [][2]float64) float64 {
   // First, find the "center" of the path
   var centerX, centerY float64

   for i := 0; i < len(points); i++ {
      centerX += points[i][0]
      centerY += points[i][1]
   }

   centerX /= float64(len(points))
   centerY /= float64(len(points))

   // Estimate the scale of the path using the mean distance from the center
   var scale float64 = 0
   for i := 0; i < len(points); i++ {
      var xDistance = points[i][0] - centerX
      var yDistance = points[i][1] - centerY
      scale += math.Sqrt(xDistance * xDistance + yDistance * yDistance)
   }
   scale /= float64(len(points))

   // Now turn the path into a complex signal
   // This is distance-naive
   var signal []complex128
   for i := 0; i < len(points); i++ {
      signal = append(signal, complex((points[i][0] - centerX) / scale, (points[i][0] - centerY) / scale))
   }

   // Now take the FFT
   var signalFFT = fft.FFT(signal)

   // Now add up the frequency magnitudes, weighted by frequency
   var frequencyRating float64 = 0
   for i := 0; i < len(signalFFT); i++ {
      var frequencyMagnitude float64
      var frequencyWeight float64

      // Use a simple linear function ranging from -1 to 1
      frequencyWeight = 2 * float64(i) / (float64(len(signalFFT) - 1) - 1)

      // The frequency magnitude is the complex magnitude
      var realAmplitude = real(signalFFT[i])
      var imaginaryAmplitude = imag(signalFFT[i])
      frequencyMagnitude = math.Sqrt(realAmplitude * realAmplitude + imaginaryAmplitude * imaginaryAmplitude)

      // Add it to the total
      frequencyRating += frequencyMagnitude * frequencyWeight
   }

   return frequencyRating
}
