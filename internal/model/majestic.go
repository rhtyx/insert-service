package model

import "strconv"

type Majestic struct {
	ID             int64 `gorm:"primaryKey"`
	GlobalRank     int64
	TLDRank        int64
	Domain         string
	TLD            string
	RefSubNets     int64 `gorm:"column:ref_subnets"`
	RefIPs         int64
	IDNDomain      string `gorm:"column:idn_domain"`
	IDNTLD         string `gorm:"column:idn_tld"`
	PrevGlobalRank int64
	PrevTLDRank    int64
	PrevRefSubNets int64 `gorm:"column:prev_ref_subnets"`
	PrevRefIPs     int64
}

func NewMajestic(data []string) *Majestic {
	globalRank, _ := strconv.ParseInt(data[0], 10, 0)
	TLDRank, _ := strconv.ParseInt(data[1], 10, 0)
	refSubNets, _ := strconv.ParseInt(data[4], 10, 0)
	refIPs, _ := strconv.ParseInt(data[5], 10, 0)
	prevGlobalRank, _ := strconv.ParseInt(data[8], 10, 0)
	prevTLDRank, _ := strconv.ParseInt(data[9], 10, 0)
	prevRefSubNets, _ := strconv.ParseInt(data[10], 10, 0)
	prevRefIPs, _ := strconv.ParseInt(data[11], 10, 0)

	return &Majestic{
		GlobalRank:     globalRank,
		TLDRank:        TLDRank,
		Domain:         data[2],
		TLD:            data[3],
		RefSubNets:     refSubNets,
		RefIPs:         refIPs,
		IDNDomain:      data[6],
		IDNTLD:         data[7],
		PrevGlobalRank: prevGlobalRank,
		PrevTLDRank:    prevTLDRank,
		PrevRefSubNets: prevRefSubNets,
		PrevRefIPs:     prevRefIPs,
	}
}
