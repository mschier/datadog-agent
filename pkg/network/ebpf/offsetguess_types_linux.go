// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- offsetguess_types.go

package ebpf

type Proc struct {
	Comm [16]int8
}

const ProcCommMaxLen = 0x10 - 1

type TracerStatus struct {
	State                  uint64
	Tcp_info_kprobe_status uint64
	Proc                   Proc
	What                   uint64
	Offset_saddr           uint64
	Offset_daddr           uint64
	Offset_sport           uint64
	Offset_dport           uint64
	Offset_netns           uint64
	Offset_ino             uint64
	Offset_family          uint64
	Offset_rtt             uint64
	Offset_rtt_var         uint64
	Offset_daddr_ipv6      uint64
	Offset_saddr_fl4       uint64
	Offset_daddr_fl4       uint64
	Offset_sport_fl4       uint64
	Offset_dport_fl4       uint64
	Offset_saddr_fl6       uint64
	Offset_daddr_fl6       uint64
	Offset_sport_fl6       uint64
	Offset_dport_fl6       uint64
	Offset_socket_sk       uint64
	Err                    uint64
	Daddr_ipv6             [4]uint32
	Netns                  uint32
	Rtt                    uint32
	Rtt_var                uint32
	Saddr                  uint32
	Daddr                  uint32
	Sport                  uint16
	Dport                  uint16
	Sport_via_sk           uint16
	Dport_via_sk           uint16
	Family                 uint16
	Saddr_fl4              uint32
	Daddr_fl4              uint32
	Sport_fl4              uint16
	Dport_fl4              uint16
	Saddr_fl6              [4]uint32
	Daddr_fl6              [4]uint32
	Sport_fl6              uint16
	Dport_fl6              uint16
	Ipv6_enabled           uint8
	Fl4_offsets            uint8
	Fl6_offsets            uint8
	Pad_cgo_0              [5]byte
}

type TracerState uint8

const (
	StateUninitialized TracerState = 0.000000
	StateChecking      TracerState = 1.000000
	StateChecked       TracerState = 2.000000
	StateReady         TracerState = 3.000000
)

type GuessWhat uint64

const (
	GuessSAddr     GuessWhat = 0.000000
	GuessDAddr     GuessWhat = 1.000000
	GuessFamily    GuessWhat = 2.000000
	GuessSPort     GuessWhat = 3.000000
	GuessDPort     GuessWhat = 4.000000
	GuessNetNS     GuessWhat = 5.000000
	GuessRTT       GuessWhat = 6.000000
	GuessDAddrIPv6 GuessWhat = 7.000000

	GuessSAddrFl4 GuessWhat = 8.000000
	GuessDAddrFl4 GuessWhat = 9.000000
	GuessSPortFl4 GuessWhat = 10.000000
	GuessDPortFl4 GuessWhat = 11.000000

	GuessSAddrFl6 GuessWhat = 12.000000
	GuessDAddrFl6 GuessWhat = 13.000000
	GuessSPortFl6 GuessWhat = 14.000000
	GuessDPortFl6 GuessWhat = 15.000000
	GuessSocketSK GuessWhat = 16.000000

	GuessNotApplicable GuessWhat = 99999
)
