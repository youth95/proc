package proc

import (
	"github.com/c9s/goprocinfo/linux"
	"github.com/stretchr/testify/assert"
	"net/rpc"
	"testing"
)

var client *rpc.Client

func TestMain(m *testing.M) {
	err := <-StartRpcService(":3399")
	if err != nil {
		panic(err)
	}
	c, err := rpc.Dial("tcp", ":3399")
	if err != nil {
		panic(err)
	}
	client = c
	m.Run()
}

func TestService_Stat(t *testing.T) {
	a := assert.New(t)
	var result linux.Stat
	err := client.Call("linux.Stat", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ListPID(t *testing.T) {
	a := assert.New(t)
	var result []uint64
	err := client.Call("linux.ListPID", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadProcessCmdline(t *testing.T) {
	a := assert.New(t)
	var result string
	err := client.Call("linux.ReadProcessCmdline", uint64(1), &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadCPUInfo(t *testing.T) {
	a := assert.New(t)
	var result linux.CPUInfo
	err := client.Call("linux.ReadCPUInfo", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadReadDisk(t *testing.T) {
	a := assert.New(t)
	var result linux.Disk
	err := client.Call("linux.ReadReadDisk", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadReadDiskStats(t *testing.T) {
	a := assert.New(t)
	var result []linux.DiskStat
	err := client.Call("linux.ReadReadDiskStats", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadInterrupts(t *testing.T) {
	a := assert.New(t)
	var result linux.Interrupts
	err := client.Call("linux.ReadInterrupts", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadLoadAvg(t *testing.T) {
	a := assert.New(t)
	var result linux.LoadAvg
	err := client.Call("linux.ReadLoadAvg", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadMemInfo(t *testing.T) {
	a := assert.New(t)
	var result linux.MemInfo
	err := client.Call("linux.ReadMemInfo", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadMounts(t *testing.T) {
	a := assert.New(t)
	var result linux.Mounts
	err := client.Call("linux.ReadMounts", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadNetTCPSocketsV4(t *testing.T) {
	a := assert.New(t)
	var result linux.NetTCPSockets

	err := client.Call("linux.ReadNetTCPSocketsV4", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadNetTCPSocketsV6(t *testing.T) {
	a := assert.New(t)
	var result linux.NetTCPSockets
	err := client.Call("linux.ReadNetTCPSocketsV6", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}

func TestService_ReadAllProcessInfo(t *testing.T) {
	a := assert.New(t)
	var result []ProcessInfo
	err := client.Call("linux.ReadAllProcessInfo", 0, &result)
	a.Nil(err)
	a.NotEmpty(result)
	t.Log(result)
}
