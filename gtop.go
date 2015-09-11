package main

import (
	"github.com/gizak/termui"
	"github.com/TilmanGriesel/gtop/components"
	"strconv"
	"time"
)

var sysmon *gtop.Sysmon
var cpu_bars []*gtop.Bar

func update_cpu_bars() {
	for {
		data, ok := <- sysmon.CpuChannel
		if(ok) {
			cpu_bars[0].ValueChannel <- int(data[0])
			// TODO [TG] Upgrade or change psutil lib
			// Possible if gopsutil is kind enough to
			// support multi cpu on windows or I find a different lib
			//for i := 0; i < len(cpu_bars); i++ {
			//	cpu_bars[i].ValueChannel <- data[i]
			//}

		}
	}
}

func render_loop(bufferer []termui.Bufferer) {
	for {
		termui.Render(bufferer ...)
		time.Sleep(time.Second)
	}
}

func main() {

	sysmon = gtop.NewSysmon()
	go sysmon.MonCPU()

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	bufferer := []termui.Bufferer{}

	lcpu_c, _, err := sysmon.CPUCount()
	if err == nil {
		cpu_bars = []*gtop.Bar{}
		bar_offset := 0
		for i := 0; i < lcpu_c; i++ {
			b := gtop.NewBar(35, 2, 1 + bar_offset, strconv.Itoa(i + 1), &bufferer)
			cpu_bars = append(cpu_bars, b)
			bar_offset += b.Height
		}
	}

	go render_loop(bufferer)
	go update_cpu_bars()

	evt := termui.EventCh()
	for {
		select {
			case e := <-evt:
				if e.Type == termui.EventKey && e.Ch == 0 {
					return
				}
		}
	}
}