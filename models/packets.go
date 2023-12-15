package models

import "time"

type Packet struct {
	Timestamp    time.Time `json:"timestamp"`
	Length       int       `json:"length"`
	SrcMAC       string    `json:"srcMAC"`
	DstMAC       string    `json:"dstMAC"`
	SrcIP        string    `json:"srcIP"`
	DstIP        string    `json:"dstIP"`
	SrcPort      uint16    `json:"srcPort"`
	DstPort      uint16    `json:"dstPort"`
	MatchedRules []string  `json:"matchedRules"`
}
