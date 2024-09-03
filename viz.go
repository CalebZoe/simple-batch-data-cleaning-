import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Create bar chart
	createBarChart()

	// Create line chart
	createLineChart()
}

func createBarChart() {
	// Create a new plot
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// Set the title of the plot
	p.Title.Text = "Bar Chart Example"

	// Create the bars
	bars := plotter.Values{20, 35, 30, 45, 25}

	// Add the bars to the plot
	bar, err := plotter.NewBarChart(bars, vg.Points(20))
	if err != nil {
		panic(err)
	}
	p.Add(bar)

	// Save the plot to a PNG file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
}

func createLineChart() {
	// Create a new plot
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// Set the title of the plot
	p.Title.Text = "Line Chart Example"

	// Create a plotter.XYs to hold our data
	pts := plotter.XYs{
		{X: 1, Y: 1},
		{X: 2, Y: 4},
		{X: 3, Y: 9},
		{X: 4, Y: 16},
		{X: 5, Y: 25},
	}

	// Add the line to the plot
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Save the plot to a PNG file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "linechart.png"); err != nil {
		panic(err)
	}
}
