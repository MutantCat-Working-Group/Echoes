package status

type LSysInfo struct {
	MemAll         uint64
	MemFree        uint64
	MemUsed        uint64
	MemUsedPercent float64
	Days           int64
	Hours          int64
	Minutes        int64
	Seconds        int64

	CpuUsedPercent float64
	OS             string
	Arch           string
	CpuCores       int
}
