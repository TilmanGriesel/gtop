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
	"time"
	"github.com/shirou/gopsutil/cpu"
)

type Sysmon struct {
	CpuChannel chan []float64
}

func NewSysmon() *Sysmon {
	e := &Sysmon{}
	e.CpuChannel = make(chan []float64, 100)
	return e
}

func (e *Sysmon) CPUCount() (logical int, real int, err error)  {
	logical, err = cpu.CPUCounts(true)
	real, err = cpu.CPUCounts(false)
	return
}

func (e *Sysmon) MonCPU()  {
	for {
		// TODO: shirou/gopsutil is utilizing WMI querys
		// and supports no cpu load for each core on windows
		// ~ BUMMER
		cpuPerc, err := cpu.CPUPercent(1, false)
		if(err == nil) {
			e.CpuChannel <- cpuPerc
		}
		time.Sleep(time.Second)
	}
}