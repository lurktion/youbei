package controllers

import (
	"time"

	"github.com/gin-gonic/gin"

	md "youbei/models"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func Sysinfo(c *gin.Context) {
	//disk
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	diskinfo := map[string]interface{}{
		"total":   diskInfo.Total,
		"used":    diskInfo.Used,
		"percent": diskInfo.UsedPercent,
	}
	//mem
	memInfo, _ := mem.VirtualMemory()
	meminfo := map[string]interface{}{
		"total":   memInfo.Total,
		"used":    memInfo.Used,
		"percent": memInfo.UsedPercent,
	}
	//cpu
	percent, _ := cpu.Percent(time.Second, false)
	counts, _ := cpu.Counts(true)
	cpuinfo := map[string]interface{}{
		"cpucounts": counts,
		"percent":   percent[0],
	}
	//net
	//host
	hostInfo, _ := host.Info()
	hostinfo := map[string]interface{}{
		"name":  hostInfo.Hostname,
		"os":    hostInfo.OS,
		"arch":  hostInfo.KernelArch,
		"plate": hostInfo.Platform,
		"id":    hostInfo.HostID,
	}
	APIReturn(c, 200, "查询成功", map[string]interface{}{
		"meminfo":  meminfo,
		"cpuinfo":  cpuinfo,
		"diskinfo": diskinfo,
		"hostinfo": hostinfo,
	})
}

func DashBoardInfo(c *gin.Context) {
	task_ok := md.CountRemoteSendLogByStatus(0)
	task_ing := md.CountRemoteSendLogByStatus(1)
	task_err := md.CountRemoteSendLogByStatus(2)
	APIReturn(c, 200, "查询成功", map[string]interface{}{
		"task_ok":  task_ok,
		"task_ing": task_ing,
		"task_err": task_err,
	})
}
