package lifecycle

import (
	"com.mutantcat.echoes/scheduler"
	"com.mutantcat.echoes/speaker"
	"com.mutantcat.echoes/status"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"github.com/shirou/gopsutil/load"
	"runtime"
)

func RegisterWarning(server_name string, interval_time int, loadavg_max_percent float64, mem_used_percent float64, cpu_used_percent float64, notice_mod string, args ...string) {
	switch notice_mod {
	case "dingbot":
		bot := dingtalk.InitDingTalkWithSecret(args[0], args[1])
		f := DingBotWarning(server_name, bot, loadavg_max_percent, mem_used_percent, cpu_used_percent)
		scheduler.IntervalDo(interval_time, f)
	default:
		fmt.Println("未知的通知方式")
		return
	}
}

func RegisterNotice(server_name string, daily_time, notice_mod string, args ...string) {
	switch notice_mod {
	case "dingbot":
		bot := dingtalk.InitDingTalkWithSecret(args[0], args[1])
		bot.SendTextMessage("您的服务器\"" + server_name + "\"已开启通知和告警服务")
		f := DingBotNotice(server_name, bot)
		scheduler.DayDo(daily_time, f)
	default:
		fmt.Println("未知的通知方式")
	}
}

func DingBotWarning(server_name string, bot *dingtalk.DingTalk, loadavg_max_percent, mem_used_percent, cpu_used_percent float64) func() {
	return func() {
		info := status.GetSysInfo()
		avg, _ := load.Avg()
		cpuNum := runtime.NumCPU()
		loadavg_max := float64(cpuNum) * loadavg_max_percent / 100
		if avg.Load1 > loadavg_max || avg.Load5 > loadavg_max || avg.Load15 > loadavg_max {
			bot.SendMarkDownMessage("服务器告警", "⚠️<font color=\"#d30c0c\">【警告】</font>您的云服务器\""+server_name+"\"当前负载过高，当前负载为<font color=\"#d30c0c\">"+fmt.Sprintf("%.2f", avg.Load1)+"</font>，请及时检查系统是否存在问题。")
		}
		if info.MemUsedPercent > mem_used_percent {
			bot.SendMarkDownMessage("服务器告警", "⚠️<font color=\"#d30c0c\">【警告】</font>您的云服务器\""+server_name+"\"当前内存使用率为<font color=\"#d30c0c\">"+fmt.Sprintf("%.2f", info.MemUsedPercent)+"%</font>，请及时检查系统是否存在问题。")
		}
		if info.CpuUsedPercent > cpu_used_percent {
			bot.SendMarkDownMessage("服务器告警", "⚠️<font color=\"#d30c0c\">【警告】</font>您的云服务器\""+server_name+"\"当前CPU使用率为<font color=\"#d30c0c\">"+fmt.Sprintf("%.2f", info.CpuUsedPercent)+"%</font>，请及时检查系统是否存在问题。")
		}
		return
	}
}

func DingBotNotice(server_name string, bot *dingtalk.DingTalk) func() {
	return func() {
		i := status.GetSysInfo()
		d := status.GetDiskInfo()
		message := speaker.GetMarkDownInfoBySysInfoAndDiskInfo(i, d, server_name)
		bot.SendMarkDownMessage(server_name, message)
	}
}
