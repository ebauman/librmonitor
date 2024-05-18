package librmonitor

import (
	"encoding/csv"
	"strings"
)

func Parse(raw string) any {
	records, err := csv.NewReader(strings.NewReader(raw)).Read()

	if err != nil {
		return msg{raw}
	}

	// F, A, B, C, E, G, H, I, J, Cor, L, T, RMS, RMLT, RMDTL, RMCA, RMHL,
	switch len(records) {
	case 0:
		return msg{raw}
	case 1:
		if records[0] == "$RMDTL" {
			return ToRMDTL(raw, records[0:])
		} else {
			return msg{raw}
		}
	case 2:
		// RMS, RMCA
		switch records[0] {
		case "$RMS":
			return ToRMS(raw, records[1:])
		case "$RMCA":
			return ToRMCA(raw, records[1:])
		default:
			return msg{raw}
		}
	case 3:
		// B, C, E, I, RMLT
		switch records[0] {
		case "$B":
			return ToB(raw, records[1:])
		case "$C":
			return ToC(raw, records[1:])
		case "$E":
			return ToE(raw, records[1:])
		case "$I":
			return ToI(raw, records[1:])
		case "$RMLT":
			return ToRMLT(raw, records[1:])
		default:
			return msg{raw}
		}
	case 4:
		// J
		if records[0] == "$J" {
			return ToJ(raw, records[1:])
		} else {
			return msg{raw}
		}
	case 5:
		// G, H
		switch records[0] {
		case "$G":
			return ToG(raw, records[1:])
		case "$H":
			return ToH(raw, records[1:])
		default:
			return msg{raw}
		}
	case 6:
		// F, Cor
		switch records[0] {
		case "$F":
			return ToF(raw, records[1:])
		case "$COR":
			return ToCor(raw, records[1:])
		default:
			return msg{raw}
		}
	case 7:
		// RMHL
		if records[0] == "$RMHL" {
			return ToRMHL(raw, records[1:])
		} else {
			return msg{raw}
		}
	case 8:
		// A, Comp, L
		switch records[0] {
		case "$A":
			return ToA(raw, records[1:])
		case "$Comp":
			return ToComp(raw, records[1:])
		case "$L":
			return ToL(raw, records[1:])
		default:
			return msg{raw}
		}
	default:
		return msg{raw}
	}
}
