package status

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/disk"
	"runtime"
	"time"
)

// 获取系统运行信息
func GetSysInfo() (info LSysInfo) {
	unit := uint64(1024 * 1024) // MB

	v, _ := mem.VirtualMemory()

	info.MemAll = v.Total
	info.MemFree = v.Free
	info.MemUsed = info.MemAll - info.MemFree
	// 注：使用SwapMemory或VirtualMemory，在不同系统中使用率不一样，因此直接计算一次
	info.MemUsedPercent = float64(info.MemUsed) / float64(info.MemAll) * 100.0 // v.UsedPercent
	info.MemAll /= unit
	info.MemUsed /= unit
	info.MemFree /= unit

	info.OS = runtime.GOOS
	info.Arch = runtime.GOARCH
	info.CpuCores = runtime.GOMAXPROCS(0)

	// 获取200ms内的CPU信息，太短不准确，也可以获几秒内的，但这样会有延时，因为要等待
	cc, _ := cpu.Percent(time.Millisecond*200, false)
	info.CpuUsedPercent = cc[0]

	// 获取开机时间
	boottime, _ := host.BootTime()
	ntime := time.Now().Unix()
	btime := time.Unix(int64(boottime), 0).Unix()
	deltatime := ntime - btime

	info.Seconds = int64(deltatime)
	info.Minutes = info.Seconds / 60
	info.Seconds -= info.Minutes * 60
	info.Hours = info.Minutes / 60
	info.Minutes -= info.Hours * 60
	info.Days = info.Hours / 24
	info.Hours -= info.Days * 24
	return
}

// 获取硬盘信息
func GetDiskInfo() *disk.UsageStat {
	d, _ := disk.Usage("/")
	return d
}
