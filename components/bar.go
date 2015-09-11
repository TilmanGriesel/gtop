/*The MIT License (MIT)

Copyright (c) 2015 Tilman Griesel

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.*/

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