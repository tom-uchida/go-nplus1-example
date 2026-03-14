package main

import (
	"bytes"
	"fmt"
	"image/color"
	"os/exec"
	"regexp"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var sizes = []int{10, 100, 1000, 10000}

func run(program string, size int) float64 {
	cmd := exec.Command("go", "run", "./"+program, strconv.Itoa(size))

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	output := out.String()

	re := regexp.MustCompile(`([0-9.]+)(ms|s)`)
	match := re.FindStringSubmatch(output)

	value, _ := strconv.ParseFloat(match[1], 64)

	if match[2] == "ms" {
		value = value / 1000
	}

	return value
}

func measure(program string) plotter.XYs {
	points := make(plotter.XYs, len(sizes))

	for i, s := range sizes {
		t := run(program, s)

		fmt.Println(program, s, t)

		points[i].X = float64(s)
		points[i].Y = t
	}

	return points
}

func main() {
	n1 := measure("nplus1")
	in := measure("in_clause")
	join := measure("join")

	p := plot.New()

	p.Title.Text = "N+1 Scaling Problem"
	p.X.Label.Text = "Number of Users"
	p.Y.Label.Text = "Seconds"

	// ライン作成
	l1, _ := plotter.NewLine(n1)
	l2, _ := plotter.NewLine(in)
	l3, _ := plotter.NewLine(join)

	// 色設定
	l1.Color = color.RGBA{R: 255, A: 255} // 赤
	l2.Color = color.RGBA{B: 255, A: 255} // 青
	l3.Color = color.RGBA{G: 180, A: 255} // 緑

	// ポイント追加（見やすくする）
	s1, _ := plotter.NewScatter(n1)
	s2, _ := plotter.NewScatter(in)
	s3, _ := plotter.NewScatter(join)

	s1.Color = l1.Color
	s2.Color = l2.Color
	s3.Color = l3.Color

	p.Add(l1, l2, l3)
	p.Add(s1, s2, s3)

	p.Legend.Add("N+1", l1)
	p.Legend.Add("IN clause", l2)
	p.Legend.Add("JOIN", l3)

	err := p.Save(8*vg.Inch, 4*vg.Inch, "./benchmark/result.png")
	if err != nil {
		panic(err)
	}

	fmt.Println("./benchmark/result.png generated")
}
