package common

import "math"

type Vec2 struct { X, Y int }

// Evaluate wether two vectors are equal.
func IsEqual(v1, v2 Vec2) bool {
    return v1.X == v2.X && v1.Y == v2.Y
}

// Add two vectors together.
func Add(v1, v2 Vec2) Vec2 {
    return Vec2{v1.X+v2.X, v1.Y+v2.Y}
}

// Subtract two vectors.
func Sub(v1, v2 Vec2) Vec2 {
    return Vec2{v1.X-v2.X, v1.Y-v2.Y}
}

func (v Vec2) ManhattanDist() int {
    return int(math.Abs(float64(v.X)) + math.Abs(float64(v.Y)))
}

