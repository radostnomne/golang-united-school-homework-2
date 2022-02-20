package square

import "math"

const (
	SidesTriangle int64 = 3
	SidesSquare   int64 = 4
	SidesCircle   int64 = 0
)

type Sides int64

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	}

	if sidesNum == SidesTriangle {
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	}

	if sidesNum == SidesSquare {
		return math.Pow(sideLen, sideLen)
	}

	return 0
}
