package gtop

import (
	"github.com/gizak/termui"
	//"log"
)

type Bar struct {
	Width int
	Height int
	X int
	Y int
	Gauge *termui.Gauge
	Label *termui.Par
	Bufferer []termui.Bufferer
	ValueChannel chan int
}

func NewBar(width int, x int, y int, label string, targetBufferer *[]termui.Bufferer) *Bar {
	e := &Bar{}

	e.Height = 3
	e.Width = width
	e.X = x
	e.Y = y

	e.Gauge = termui.NewGauge()
	e.Gauge.Percent = 0
	e.Gauge.Percent = 0
	e.Gauge.Width = width - 6
	e.Gauge.Height = e.Height
	e.Gauge.Border.Label = ""
	e.Gauge.BgColor = termui.ColorBlack
	e.Gauge.BarColor = termui.ColorGreen
	e.Gauge.Border.FgColor = termui.ColorWhite
	e.Gauge.Border.LabelFgColor = termui.ColorCyan
	e.Gauge.HasBorder = true
	e.Gauge.LabelAlign = termui.AlignRight

	e.Label = termui.NewPar("")
	e.Label.Width = 20
	e.Label.Height = 1
	e.Label.TextFgColor = termui.ColorWhite
	e.Label.HasBorder = false
	e.Label.Text = label

	e.Bufferer = append(e.Bufferer, e.Gauge)
	e.Bufferer = append(e.Bufferer, e.Label)
	*targetBufferer = append(*targetBufferer, e.Bufferer ...)

	e.ValueChannel = make(chan int, 100)
	go e.readValueChannel()

	e.Invalidate()

	return e
}

func (e *Bar) Invalidate()  {
	e.Gauge.Width = e.Width - 6
	e.Gauge.Height = e.Height

	e.Gauge.X = e.X + 2
	e.Label.X = e.X + 0
	e.Gauge.Y = e.Y
	e.Label.Y = e.Y + (e.Height / 2)
}

func (e *Bar) SetLabel(text string)  {
	e.Label.Text = text
}

func (e *Bar) SetValue(value int)  {
	e.Gauge.Percent = value
}

func (e *Bar) readValueChannel() {
	for {
		value, ok := <- e.ValueChannel
		if ok {
			e.SetValue(value)
		}
	}
}