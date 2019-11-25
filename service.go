package goproc_rpc

import (
	"fmt"
	"github.com/c9s/goprocinfo/linux"
)

type Service int

func (*Service) Stat(input *int, reply *linux.Stat) error {
	stat, err := linux.ReadStat("/proc/stat")
	if err != nil {
		return err
	}
	*reply = *stat
	return nil
}

func (*Service) ListPID(input *int, reply *[]uint64) error {
	list, err := linux.ListPID("/proc", 2048)
	if err != nil {
		return err
	}
	*reply = list
	return nil
}

func (*Service) ReadProcessCmdline(input *uint64, reply *string) error {
	r, err := linux.ReadProcessCmdline(fmt.Sprintf("/proc/%d/cmdline", *input))
	if err != nil {
		return err
	}
	*reply = r
	return nil
}

func (*Service) ReadCPUInfo(input *int, reply *linux.CPUInfo) error {
	info, err := linux.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		return err
	}
	*reply = *info
	return nil
}

func (*Service) ReadReadDisk(input *int, reply *linux.Disk) error {
	dist, err := linux.ReadDisk("/")
	if err != nil {
		return err
	}
	*reply = *dist
	return nil
}

func (*Service) ReadReadDiskStats(input *int, reply *[]linux.DiskStat) error {
	diskStat, err := linux.ReadDiskStats("/proc/diskstats")
	if err != nil {
		return err
	}
	*reply = diskStat
	return nil
}

func (*Service) ReadInterrupts(input *int, reply *linux.Interrupts) error {
	interrupts, err := linux.ReadInterrupts("/proc/interrupts")
	if err != nil {
		return err
	}
	*reply = *interrupts
	return nil
}

func (*Service) ReadLoadAvg(input *int, reply *linux.LoadAvg) error {
	loadAvg, err := linux.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		return err
	}
	*reply = *loadAvg
	return nil
}

func (*Service) ReadMemInfo(input *int, reply *linux.MemInfo) error {
	memInfo, err := linux.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return err
	}
	*reply = *memInfo
	return nil
}

func (*Service) ReadMounts(input *int, reply *linux.Mounts) error {
	mounts, err := linux.ReadMounts("/proc/mounts")
	if err != nil {
		return err
	}
	*reply = *mounts
	return nil
}

func (*Service) ReadNetTCPSocketsV4(input *int, reply *linux.NetTCPSockets) error {
	tcp, err := linux.ReadNetTCPSockets("/proc/net/tcp", linux.NetIPv4Decoder)
	if err != nil {
		return err
	}
	*reply = *tcp
	return nil
}

func (*Service) ReadNetTCPSocketsV6(input *int, reply *linux.NetTCPSockets) error {
	tcp, err := linux.ReadNetTCPSockets("/proc/net/tcp6", linux.NetIPv6Decoder)
	if err != nil {
		return err
	}
	*reply = *tcp
	return nil
}
