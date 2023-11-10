package main

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/wcharczuk/go-chart/v2"
)

func main() {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	buildHeader(m)
	graphGenerator()
	productsTable(m)
	includeGraph(m)
	pieCharGenerator()
	includeChart(m)
	err := m.OutputFileAndClose("pdfs/output.pdf")
	if err != nil {
		fmt.Println("Cannot save output", err)
	}

	fmt.Println("Pdf was successfully generated")

}

func buildHeader(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				m.FileImage("images/logo.png", props.Rect{
					Center:  true,
					Percent: 75,
				})
			})
		})
	})

	m.Row(10, func() {
		m.Row(12, func() {
			m.Text("This is Header Section", props.Text{})
		})
	})
}

func graphGenerator() {
	graph := chart.BarChart{
		Title:      "Production sales",
		TitleStyle: chart.StyleTextDefaults(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 30,
			},
		},
		Height:   256,
		BarWidth: 100,
		Bars: []chart.Value{
			{Value: 40, Label: "Production 1"},
			{Value: 60, Label: "Production 2"},
			{Value: 30, Label: "Production 3"},
			{Value: 80, Label: "Production 4"},
		},
	}

	f, _ := os.Create("graphs/Chart.png")
	defer f.Close()

	graph.Render(chart.PNG, f)
}

func productsTable(m pdf.Maroto) {

	tableHeading := []string{"ProductName", "Version", "Color"}
	content := [][]string{{"AirForce", "1", "White"}, {"Air Jordan", "1", "Black"}, {"Travis Scott", "V3", "Black"}}

	lightPurpleColor := getLightPurpleColor()
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Style:  consts.Bold,
				Family: consts.Courier,
				Color:  color.NewWhite(),
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeading, content, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightPurpleColor,
	})

	m.Row(12, func() {
		m.Col(10, func() {

		})
	})

}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Blue:  166,
		Green: 166,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Blue:  200,
		Green: 230,
	}
}

func includeGraph(m pdf.Maroto) {
	m.Row(50, func() {
		m.Col(12, func() {
			m.FileImage("graphs/Chart.png", props.Rect{
				Center:  true,
				Percent: 100,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("This is Production Sales Graph", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
}

func pieCharGenerator() {
	graph := chart.PieChart{
		Title:      "Most Color Sales",
		TitleStyle: chart.StyleTextDefaults(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 30,
			},
		},
		Height: 512,
		Width:  512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
		},
	}

	f, _ := os.Create("graphs/pie.png")
	defer f.Close()

	graph.Render(chart.PNG, f)
}

func includeChart(m pdf.Maroto) {
	m.Row(50, func() {
		m.Col(12, func() {
			m.FileImage("graphs/pie.png", props.Rect{
				Center:  true,
				Percent: 100,
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("This is simple pie chart", props.Text{})
		})
	})
}
