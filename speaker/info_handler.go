package speaker

import (
	"com.mutantcat.echoes/status"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
)

// 通过系统信息和硬盘信息获取markdown信息
func GetMarkDownInfoBySysInfoAndDiskInfo(info status.LSysInfo, d *disk.UsageStat, name string) string {
	message := "### 服务器基本信息"
	message += "  \n> 服务器 - " + name + " - 运行状态"
	message += "  \n> 您的服务器已运行-" + fmt.Sprintf("%d", info.Days) + "天" + fmt.Sprintf("%d", info.Hours) + "小时" + fmt.Sprintf("%d", info.Minutes) + "分钟" + fmt.Sprintf("%d", info.Seconds) + "秒"
	message += "  \n- 系统运行内存使用率为：" + fmt.Sprintf("%.2f", info.MemUsedPercent) + "%"
	message += "  \n- 系统运行CPU使用率为：" + fmt.Sprintf("%.2f", info.CpuUsedPercent) + "%"
	message += "  \n- 系统物理磁盘使用率为：" + fmt.Sprintf("%.2f", d.UsedPercent) + "%"
	message += "  \n- 系统运行内存已用量为：" + fmt.Sprintf("%d", info.MemUsed) + "MB/" + fmt.Sprintf("%d", info.MemAll) + "MB"
	message += "  \n- 系统运行内存空闲量为：" + fmt.Sprintf("%d", info.MemFree) + "MB"
	message += "  \n- 系统运行CPU核心数为：" + fmt.Sprintf("%d", info.CpuCores) + "核"
	message += "  \n- 系统物理磁盘已用大小为：" + fmt.Sprintf("%d", d.Used/1024/1024/1024) + "GB/" + fmt.Sprintf("%d", d.Total/1024/1024/1024) + "GB"
	message += "  \n- 系统物理磁盘空闲大小为：" + fmt.Sprintf("%d", d.Free/1024/1024/1024) + "GB"
	return message
}
