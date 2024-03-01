package cnitypes

import (
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	nadapi "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	"github.com/vishvananda/netlink"
)

const ServerRunDir string = "/var/run/dpu-cni"
const ServerSocketName string = "dpu-cni-server.sock"
const ServerSocketPath string = ServerRunDir + "/" + ServerSocketName

// Request sent to the Server by the DPU CNI plugin
type Request struct {
	// CNI environment variables, like CNI_COMMAND and CNI_NETNS
	Env map[string]string `json:"env,omitempty"`
	// CNI configuration passed via stdin to the CNI plugin
	Config []byte `json:"config,omitempty"`
	// The DeviceInfo struct
	nadapi.DeviceInfo
}

// Response sent to this DPU CNI plugin by the Server
type Response struct {
	Result *current.Result
}

// FIXME: This file is copied from sriov-cni intentionally. We plan to trim this down once
// we know what we want to support from SR-IOV.

const (
	Proto8021q  = "802.1q"
	Proto8021ad = "802.1ad"
)

var VlanProtoInt = map[string]int{Proto8021q: 33024, Proto8021ad: 34984}

// VfState represents the state of the VF
type VfState struct {
	HostIFName   string
	SpoofChk     bool
	Trust        bool
	AdminMAC     string
	EffectiveMAC string
	Vlan         int
	VlanQoS      int
	VlanProto    int
	MinTxRate    int
	MaxTxRate    int
	LinkState    uint32
}

// FillFromVfInfo - Fill attributes according to the provided netlink.VfInfo struct
func (vs *VfState) FillFromVfInfo(info *netlink.VfInfo) {
	vs.AdminMAC = info.Mac.String()
	vs.LinkState = info.LinkState
	vs.MaxTxRate = int(info.MaxTxRate)
	vs.MinTxRate = int(info.MinTxRate)
	vs.Vlan = info.Vlan
	vs.VlanQoS = info.Qos
	vs.VlanProto = info.VlanProto
	vs.SpoofChk = info.Spoofchk
	vs.Trust = info.Trust != 0
}

// NetConf extends types.NetConf for dpu-sriov-cni
type NetConf struct {
	types.NetConf
	OrigVfState   VfState // Stores the original VF state as it was prior to any operations done during cmdAdd flow
	DPDKMode      bool    `json:"-"`
	Master        string
	MAC           string
	Vlan          *int    `json:"vlan"`
	VlanQoS       *int    `json:"vlanQoS"`
	VlanProto     *string `json:"vlanProto"` // 802.1ad|802.1q
	DeviceID      string  `json:"deviceID"`  // PCI address of a VF in valid sysfs format
	VFID          int
	MinTxRate     *int   `json:"min_tx_rate"`          // Mbps, 0 = disable rate limiting
	MaxTxRate     *int   `json:"max_tx_rate"`          // Mbps, 0 = disable rate limiting
	SpoofChk      string `json:"spoofchk,omitempty"`   // on|off
	Trust         string `json:"trust,omitempty"`      // on|off
	LinkState     string `json:"link_state,omitempty"` // auto|enable|disable
	RuntimeConfig struct {
		Mac string `json:"mac,omitempty"`
	} `json:"runtimeConfig,omitempty"`
	LogLevel string `json:"logLevel,omitempty"`
	LogFile  string `json:"logFile,omitempty"`
}