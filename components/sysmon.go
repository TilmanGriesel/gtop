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