package base

type DiskInfo struct {
	Path  string `json:"p"`
	Total uint64 `json:"t"`
	Free  uint64 `json:"f"`
}

type CpuInfo struct {
	Cores  int     `json:"c"`
	Family string  `json:"f"`
	Load1  float64 `json:"l1"`
	Load5  float64 `json:"l5"`
	Load15 float64 `json:"l15"`
}

type MemInfo struct {
	Total     uint64 `json:"t"`
	Free      uint64 `json:"f"`
	Available uint64 `json:"v"`
}

type UserInfo struct {
	Host    string `json:"h"`
	User    string `json:"u"`
	Started int    `json:"s"`
}

type MachineInfo struct {
	Hostname string     `json:"n"`
	OS       string     `json:"o"`
	Cpu      CpuInfo    `json:"c"`
	Mem      MemInfo    `json:"m"`
	Disk     []DiskInfo `json:"d"`
	User     []UserInfo `json:"u"`
	Time     int64      `json:"t"`
}
