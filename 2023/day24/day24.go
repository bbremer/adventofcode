package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	/*
		file, _ := os.Open("../inputs/test.txt")
		min := 7.
		max := 27.
	*/
	file, _ := os.Open("../inputs/day24.txt")
	min := 200000000000000.
	max := 400000000000000.
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trajectories := []Trajectory{}
	for scanner.Scan() {
		trajectories = append(trajectories, NewTrajectory(scanner.Text()))
	}
	fmt.Println("Part 1:", Part1(min, max, trajectories))

	p1 := trajectories[0].p
	p2 := trajectories[1].p
	p3 := trajectories[2].p
	v1 := trajectories[0].d
	v2 := trajectories[1].d
	v3 := trajectories[2].d

	A1 := []float64{0., -v1.z + v2.z, v1.y - v2.y, 0., p1.z - p2.z, -p1.y + p2.y}
	A2 := []float64{v1.z - v2.z, 0., -v1.x + v2.x, -p1.z + p2.z, 0., p1.x - p2.x}
	A3 := []float64{-v1.y + v2.y, v1.x - v2.x, 0., p1.y - p2.y, -p1.x + p2.x, 0.}
	A4 := []float64{0., -v1.z + v3.z, v1.y - v3.y, 0., p1.z - p3.z, -p1.y + p3.y}
	A5 := []float64{v1.z - v3.z, 0., -v1.x + v3.x, -p1.z + p3.z, 0., p1.x - p3.x}
	A6 := []float64{-v1.y + v3.y, v1.x - v3.x, 0., p1.y - p3.y, -p1.x + p3.x, 0.}

	var ASlice []float64
	ASlice = append(ASlice, A1...)
	ASlice = append(ASlice, A2...)
	ASlice = append(ASlice, A3...)
	ASlice = append(ASlice, A4...)
	ASlice = append(ASlice, A5...)
	ASlice = append(ASlice, A6...)
	A := mat.NewDense(6, 6, ASlice)

	b1 := p1.z*v1.y - p1.y*v1.z - p2.z*v2.y + p2.y*v2.z
	b2 := p1.x*v1.z - p1.z*v1.x - p2.x*v2.z + p2.z*v2.x
	b3 := p1.y*v1.x - p1.x*v1.y - p2.y*v2.x + p2.x*v2.y
	b4 := p1.z*v1.y - p1.y*v1.z - p3.z*v3.y + p3.y*v3.z
	b5 := p1.x*v1.z - p1.z*v1.x - p3.x*v3.z + p3.z*v3.x
	b6 := p1.y*v1.x - p1.x*v1.y - p3.y*v3.x + p3.x*v3.y
	b := mat.NewVecDense(6, []float64{b1, b2, b3, b4, b5, b6})

	var x mat.VecDense
	x.SolveVec(A, b)
	X := x.At(0, 0)
	Y := x.At(1, 0)
	Z := x.At(2, 0)
	fmt.Printf("Part 2: %.0f\n", math.Round(X)+math.Round(Y)+math.Round(Z))
}

func Part1(min, max float64, trajectories []Trajectory) (sum int) {
	for i, t1 := range trajectories {
		for _, t2 := range trajectories[i+1:] {
			if m, ok := t1.Intersection(t2, 2); ok {
				time1 := m.At(0, 0)
				time2 := m.At(1, 0)
				v := t1.TimeToVec(time1)
				x, y := v.At(0, 0), v.At(1, 0)
				if time1 >= 0. && time2 >= 0. && min <= x && x <= max && min <= y && y <= max {
					sum++
				}
			}
		}
	}
	return
}

func NewMatrix(c1, c2 *mat.Dense) *mat.Dense {
	var m mat.Dense
	m.Augment(c1, c2)
	return &m
}

type Trajectory struct {
	p Vec
	d Vec
}

func NewTrajectory(line string) Trajectory {
	splitLine := strings.Split(line, " @ ")
	return Trajectory{NewVec(splitLine[0]), NewVec(splitLine[1])}
}

func (t1 Trajectory) Intersection(t2 Trajectory, l int) (mat.VecDense, bool) {
	var A mat.Dense
	var d2 mat.Dense
	d2.Scale(-1, t2.d.ToMatrix(l))
	A.Augment(t1.d.ToMatrix(l), &d2)

	var sub mat.VecDense
	sub.SubVec(t2.p.ToVec(l), t1.p.ToVec(l))

	var ret mat.VecDense
	err := ret.SolveVec(&A, &sub)
	if err != nil {
		return ret, false
	}
	return ret, true
}

func (t Trajectory) TimeToVec(time float64) (v mat.VecDense) {
	p := t.p.ToVec(2)
	d := t.d.ToVec(2)
	d.ScaleVec(time, d)
	v.AddVec(p, d)
	return
}

type Vec struct {
	x float64
	y float64
	z float64
}

func NewVec(line string) Vec {
	splitLine := strings.Split(strings.ReplaceAll(line, " ", ""), ",")
	v1, _ := strconv.Atoi(splitLine[0])
	v2, _ := strconv.Atoi(splitLine[1])
	v3, _ := strconv.Atoi(splitLine[2])
	return Vec{float64(v1), float64(v2), float64(v3)}
}

func (v Vec) ToMatrix(l int) *mat.Dense {
	fs := []float64{v.x, v.y, v.z}[:l]
	return mat.NewDense(l, 1, fs)
}

func (v Vec) ToVec(l int) *mat.VecDense {
	fs := []float64{v.x, v.y, v.z}[:l]
	return mat.NewVecDense(l, fs)
}
