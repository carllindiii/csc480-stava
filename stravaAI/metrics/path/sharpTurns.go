package path

import (
   "math"
   
   "github.com/go-gl/mathgl/mgl64"
)

// func AngleBetween(v1 mgl64.Vec2, v2 mgl64.Vec2) float64 {
//    return math.Acos(v1.Dot(v2) / (v1.Len() * v2.Len()))
// }

// func getCurviness(points [][2]float64) int {
//    count := GetNumSharpTurns(points)
//    if (count > 5) {
//       return 2
//    }
//    if (count > 3) {
//       return 1
//    }
//    return 0
// }

func catmullRom(p1, p2, p3, p4 mgl64.Vec2, u float64) mgl64.Vec2 {
   var B mgl64.Mat4
   
   B.Set(0, 0, 1)
   B.Set(0, 1, -3)
   B.Set(0, 2, 3)
   B.Set(0, 3, -1)
   B.Set(1, 0, 4)
   B.Set(1, 1, 0)
   B.Set(1, 2, -6)
   B.Set(1, 3, 3)
   B.Set(2, 0, 1)
   B.Set(2, 1, 3)
   B.Set(2, 2, 3)
   B.Set(2, 3, -3)
   B.Set(3, 0, 0)
   B.Set(3, 1, 0)
   B.Set(3, 2, 0)
   B.Set(3, 3, 1)

   B = B.Mul(1.0/6.0)

   var pMat mgl64.Mat2x4

   pMat.Set(0, 0, p1.X())
   pMat.Set(1, 0, p1.Y())
   pMat.Set(0, 1, p2.X())
   pMat.Set(1, 1, p2.Y())
   pMat.Set(0, 2, p3.X())
   pMat.Set(1, 2, p3.Y())
   pMat.Set(0, 3, p4.X())
   pMat.Set(1, 3, p4.Y())

   return pMat.Mul4(B).Mul4x1(mgl64.Vec4{0, 1, 2 * u, 2 * u * u})
}

func angleBetween(v1, v2 mgl64.Vec2) float64 {
   return math.Acos(v1.Dot(v2)/(v1.Len() * v2.Len()));
}

func GetNumSharpTurns(points [][2]float64) (float64, int) {
   arrLens := len(points) - 2

   tangents := make([][2]float64, arrLens)
   dists := make([]float64, arrLens - 1)
   angles := make([]float64, arrLens - 1)

   // arrLens[0] = catmullRom(points[0], points[1], points[2], points[3], 0);

   for idx := 0; idx < arrLens - 1; idx++ {
      tangents[idx] = catmullRom(points[idx], points[idx + 1], points[idx + 2], points[idx + 3], 0.0);
      dists[idx] = 1.0
   }

   for idx := 0; idx < arrLens - 1; idx++ {
      angles[idx] = angleBetween(tangents[idx], tangents[idx + 1])
   }

   const maxAngleChange = .4 * math.Pi
   const maxDistChange = 4

   numSharpTurns := 0   // sharp turn is when the angles add up to maxAngleChange for 4 consecutive points

   for idx := 0; idx < arrLens - 1; idx++ {
      dAngle := 0.0
      for count := 0; count < maxDistChange && idx < arrLens - 1; count++ {
         dAngle += angles[idx]
         idx++
      }

      if dAngle > maxAngleChange {
         numSharpTurns++
      }
   }

   var difficulty float64
   if (numSharpTurns < 5) {
      difficulty =  0
   } else if (numSharpTurns < 10) {
      difficulty = 0.5
   } else {
      difficulty = 1
   }

   return difficulty, numSharpTurns
}

func GetNumSharpTurnsSecant(points [][2]float64) int {
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

// func main() {
//    var segmentId int64
//    var accessToken string

//    // Provide an access token, with write permissions.
//    // You'll need to complete the oauth flow to get one.
//    flag.Int64Var(&segmentId, "id", 229781, "Strava Segment Id")
//    flag.StringVar(&accessToken, "token", "", "Access Token")

//    flag.Parse()

//    if accessToken == "" {
//       fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")

//       flag.PrintDefaults()
//       os.Exit(1)
//    }

//    client := strava.NewClient(accessToken)

//    segmentService := strava.NewSegmentsService(client)

//    segIds := [10]int64{365235, 6452581, 664647, 1089563, 4956199, 2187, 5732938, 654030, 616554, 3139189}


//    for _, nextSegId := range segIds {
//       fmt.Printf("nextSegId: %d\n", nextSegId);
//       segment, err := segmentService.Get(nextSegId).Do()

//       if err != nil {
//          fmt.Println(err)
//          os.Exit(1)
//       }

//       var latLngCrds [][2]float64 = segment.Map.Polyline.Decode()
//       fmt.Printf("\tsecant: %d\n", GetNumSharpTurnsSecant(latLngCrds))
//       fmt.Printf("\tcatnullRom: %d\n", GetNumSharpTurns(latLngCrds))
//    }

   
// }