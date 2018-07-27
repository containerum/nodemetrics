package models

import "time"

type Report struct {
	Name          string         `json:"name"`
	Subcontainers []Subcontainer `json:"subcontainers"`
	Spec          Spec           `json:"spec"`
	Stats         []Stat         `json:"stats"`
}

type Subcontainer struct {
	Name string `json:"name"`
}

type Stat struct {
	Timestamp  time.Time    `json:"timestamp"`
	CPU        CPU          `json:"cpu"`
	DiskIO     DiskIO       `json:"diskio"`
	Memory     Memory       `json:"memory"`
	Network    Network      `json:"network"`
	Filesystem []Filesystem `json:"filesystem"`
	TaskStats  TaskStats    `json:"task_stats"`
}

type Memory struct {
	Usage            int64            `json:"usage"`
	MaxUsage         int64            `json:"max_usage"`
	Cache            int64            `json:"cache"`
	Rss              int64            `json:"rss"`
	Swap             int              `json:"swap"`
	WorkingSet       int64            `json:"working_set"`
	Failcnt          int              `json:"failcnt"`
	ContainerData    ContainerData    `json:"container_data"`
	HierarchicalData HierarchicalData `json:"hierarchical_data"`
}

type HierarchicalData struct {
	Pgfault    int `json:"pgfault"`
	Pgmajfault int `json:"pgmajfault"`
}

type ContainerData struct {
	Pgfault    int `json:"pgfault"`
	Pgmajfault int `json:"pgmajfault"`
}

type DiskIO struct {
	IoServiceBytes []struct {
		Device string `json:"device"`
		Major  int    `json:"major"`
		Minor  int    `json:"minor"`
		Stats  struct {
			Async int `json:"Async"`
			Read  int `json:"Read"`
			Sync  int `json:"Sync"`
			Total int `json:"Total"`
			Write int `json:"Write"`
		} `json:"stats"`
	} `json:"io_service_bytes"`
	IoServiced []struct {
		Device string `json:"device"`
		Major  int    `json:"major"`
		Minor  int    `json:"minor"`
		Stats  struct {
			Async int `json:"Async"`
			Read  int `json:"Read"`
			Sync  int `json:"Sync"`
			Total int `json:"Total"`
			Write int `json:"Write"`
		} `json:"stats"`
	} `json:"io_serviced"`
}

type Spec struct {
	CreationTime time.Time `json:"creation_time"`
	HasCPU       bool      `json:"has_cpu"`
	CPU          struct {
		Limit    int    `json:"limit"`
		MaxLimit int    `json:"max_limit"`
		Mask     string `json:"mask"`
		Period   int    `json:"period"`
	} `json:"cpu"`
	HasMemory bool `json:"has_memory"`
	Memory    struct {
		Limit       int64 `json:"limit"`
		Reservation int64 `json:"reservation"`
	} `json:"memory"`
	HasNetwork       bool `json:"has_network"`
	HasFilesystem    bool `json:"has_filesystem"`
	HasDiskio        bool `json:"has_diskio"`
	HasCustomMetrics bool `json:"has_custom_metrics"`
}

type CPU struct {
	Usage struct {
		Total       int64   `json:"total"`
		PerCPUUsage []int64 `json:"per_cpu_usage"`
		User        int64   `json:"user"`
		System      int64   `json:"system"`
	} `json:"usage"`
	Cfs struct {
		Periods          int `json:"periods"`
		ThrottledPeriods int `json:"throttled_periods"`
		ThrottledTime    int `json:"throttled_time"`
	} `json:"cfs"`
	Schedstat struct {
		RunTime      int `json:"run_time"`
		RunqueueTime int `json:"runqueue_time"`
		RunPeriods   int `json:"run_periods"`
	} `json:"schedstat"`
	LoadAverage int `json:"load_average"`
}

type Network struct {
	Name       string      `json:"name"`
	RxBytes    int         `json:"rx_bytes"`
	RxPackets  int         `json:"rx_packets"`
	RxErrors   int         `json:"rx_errors"`
	RxDropped  int         `json:"rx_dropped"`
	TxBytes    int         `json:"tx_bytes"`
	TxPackets  int         `json:"tx_packets"`
	TxErrors   int         `json:"tx_errors"`
	TxDropped  int         `json:"tx_dropped"`
	Interfaces []Interface `json:"interfaces"`
	TCP        TCP         `json:"tcp"`
	TCP6       TCP         `json:"tcp6"`
	UDP        UDP         `json:"udp"`
	UDP6       UDP         `json:"udp6"`
}

type Interface struct {
	Name      string `json:"name"`
	RxBytes   int    `json:"rx_bytes"`
	RxPackets int    `json:"rx_packets"`
	RxErrors  int    `json:"rx_errors"`
	RxDropped int    `json:"rx_dropped"`
	TxBytes   int    `json:"tx_bytes"`
	TxPackets int    `json:"tx_packets"`
	TxErrors  int    `json:"tx_errors"`
	TxDropped int    `json:"tx_dropped"`
}

type TCP struct {
	Established int `json:"Established"`
	SynSent     int `json:"SynSent"`
	SynRecv     int `json:"SynRecv"`
	FinWait1    int `json:"FinWait1"`
	FinWait2    int `json:"FinWait2"`
	TimeWait    int `json:"TimeWait"`
	Close       int `json:"Close"`
	CloseWait   int `json:"CloseWait"`
	LastAck     int `json:"LastAck"`
	Listen      int `json:"Listen"`
	Closing     int `json:"Closing"`
}

type UDP struct {
	Listen   int `json:"Listen"`
	Dropped  int `json:"Dropped"`
	RxQueued int `json:"RxQueued"`
	TxQueued int `json:"TxQueued"`
}

type TaskStats struct {
	NrSleeping        int `json:"nr_sleeping"`
	NrRunning         int `json:"nr_running"`
	NrStopped         int `json:"nr_stopped"`
	NrUninterruptible int `json:"nr_uninterruptible"`
	NrIoWait          int `json:"nr_io_wait"`
}

type Filesystem struct {
	Device          string `json:"device"`
	Type            string `json:"type"`
	Capacity        int64  `json:"capacity"`
	Usage           int64  `json:"usage"`
	BaseUsage       int    `json:"base_usage"`
	Available       int64  `json:"available"`
	HasInodes       bool   `json:"has_inodes"`
	Inodes          int    `json:"inodes"`
	InodesFree      int    `json:"inodes_free"`
	ReadsCompleted  int    `json:"reads_completed"`
	ReadsMerged     int    `json:"reads_merged"`
	SectorsRead     int    `json:"sectors_read"`
	ReadTime        int    `json:"read_time"`
	WritesCompleted int    `json:"writes_completed"`
	WritesMerged    int    `json:"writes_merged"`
	SectorsWritten  int    `json:"sectors_written"`
	WriteTime       int    `json:"write_time"`
	IoInProgress    int    `json:"io_in_progress"`
	IoTime          int    `json:"io_time"`
	WeightedIoTime  int    `json:"weighted_io_time"`
}
