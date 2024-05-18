package librmonitor

import (
	"strconv"
	"strings"
)

type Message interface {
	Raw() string
}

type msg struct {
	msg string
}

func (m msg) Raw() string {
	return m.msg
}

type F struct {
	msg
	LapsToGo   string `json:"lapsToGo"`
	TimeToGo   string `json:"timeToGo"`
	TimeOfDay  string `json:"timeOfDay"`
	RaceTime   string `json:"raceTime"`
	FlagStatus string `json:"flagStatus"`
}

func ToF(raw string, parts []string) F {
	return F{
		msg{raw},
		NoQuote(parts[0]),
		NoQuote(parts[1]),
		NoQuote(parts[2]),
		NoQuote(parts[3]),
		NoQuote(parts[4]),
	}
}

type A struct {
	msg
	RegistrationNumber string `json:"registrationNumber"`
	Number             string `json:"number"`
	TransponderNumber  string `json:"transponderNumber"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Nationality        string `json:"nationality"`
	ClassNumber        string `json:"classNumber"`
}

func ToA(raw string, parts []string) A {
	return A{
		msg{raw},
		parts[0],
		parts[1],
		parts[2],
		parts[3],
		parts[4],
		parts[5],
		parts[6],
	}
}

type Comp struct {
	msg
	RegistrationNumber string `json:"registrationNumber"`
	Number             string `json:"number"`
	ClassNumber        string `json:"classNumber"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Nationality        string `json:"nationality"`
	AdditionalData     string `json:"additionalData"`
}

func ToComp(raw string, parts []string) Comp {
	return Comp{
		msg{raw},
		parts[0],
		parts[1],
		parts[2],
		parts[3],
		parts[4],
		parts[5],
		parts[6],
	}
}

type B struct {
	msg
	UniqueNumber string `json:"uniqueNumber"`
	Description  string `json:"description"`
}

func ToB(raw string, parts []string) B {
	return B{
		msg{raw},
		parts[0],
		parts[1],
	}
}

type C struct {
	msg
	UniqueNumber string `json:"uniqueNumber"`
	Description  string `json:"description"`
}

func ToC(raw string, parts []string) C {
	return C{
		msg{raw},
		parts[0],
		parts[1],
	}
}

type E struct {
	msg
	Description string `json:"description"`
	Value       string `json:"value"`
}

func ToE(raw string, parts []string) E {
	return E{
		msg{raw},
		parts[0],
		parts[1],
	}
}

type G struct {
	msg
	Position           string `json:"position"`
	RegistrationNumber string `json:"registrationNumber"`
	Laps               string `json:"laps"`
	TotalTime          string `json:"totalTime"`
}

func ToG(raw string, parts []string) G {
	return G{
		msg{raw},
		parts[0],
		parts[1],
		parts[2],
		parts[3],
	}
}

type H struct {
	msg
	Position           string `json:"position"`
	RegistrationNumber string `json:"registrationNumber"`
	BestLap            string `json:"bestLap"`
	BestLaptime        string `json:"bestLaptime"`
}

func ToH(raw string, parts []string) H {
	return H{
		msg{raw},
		parts[0],
		parts[1],
		parts[2],
		parts[3],
	}
}

type I struct {
	msg
	TimeOfDay string `json:"timeOfDay"`
	Date      string `json:"date"`
}

func ToI(raw string, parts []string) I {
	return I{
		msg{raw},
		parts[0],
		parts[1],
	}
}

type J struct {
	msg
	RegistrationNumber string `json:"registrationNumber"`
	Laptime            string `json:"laptime"`
	TotalTime          string `json:"totalTime"`
}

func ToJ(raw string, parts []string) J {
	return J{
		msg{raw},
		parts[0],
		parts[1],
		parts[2],
	}
}

type Cor struct {
	msg
	RegistrationNumber string `json:"registrationNumber"`
	Number             string `json:"number"`
	Laps               string `json:"laps"`
	TotalTime          string `json:"totalTime"`
	Correction         string `json:"correction"`
}

func ToCor(raw string, parts []string) Cor {
	return Cor{
		msg{raw},
		parts[0],
		parts[1],
		parts[2],
		parts[3],
		parts[4],
	}
}

type L struct {
	msg
	CarNumber      string `json:"carNumber"`
	TimeLineNumber string `json:"timeLineNumber"`
	TimeLineName   string `json:"timeLineName"`
	DateOfCrossing string `json:"dateOfCrossing"`
	TimeOfCrossing string `json:"timeOfCrossing"`
	DriverId       string `json:"driverId"`
	ClassName      string `json:"className"`
}

func ToL(raw string, parts []string) L {
	return L{
		msg{raw},
		parts[0],
		parts[1],
		parts[2],
		parts[3],
		parts[4],
		parts[5],
		parts[6],
	}
}

type T struct {
	msg
	TrackName      string `json:"trackName"`
	TrackShortName string `json:"trackShortName"`
	TrackDistance  string `json:"trackDistance"`

	SectionCount int `json:"sectionCount"`

	Sections []TSection `json:"sections"`
}

func ToT(raw string, parts []string) T {
	t := T{
		msg:            msg{raw},
		TrackName:      parts[0],
		TrackShortName: parts[1],
		TrackDistance:  parts[2],
	}

	count, err := strconv.ParseInt(parts[3], 10, 32)
	if err != nil {
		return t
	}

	t.Sections = make([]TSection, count)

	pos := 3
	for i := range count {
		sec := TSection{
			msg:                   msg{},
			SectionName:           parts[pos],
			SectionStart:          parts[pos+1],
			SectionEnd:            parts[pos+2],
			SectionDistanceInches: parts[pos+3],
		}
		t.Sections[i] = sec

		pos = pos + 4
	}

	return t
}

type TSection struct {
	msg
	SectionName           string `json:"sectionName"`
	SectionStart          string `json:"sectionStart"`
	SectionEnd            string `json:"sectionEnd"`
	SectionDistanceInches string `json:"sectionDistanceInches"`
}

type RMS struct {
	msg
	SortMode string `json:"sortMode"`
}

func ToRMS(raw string, parts []string) RMS {
	return RMS{
		msg{raw},
		parts[0],
	}
}

type RMLT struct {
	msg
	RacerId     string `json:"racerId"`
	LastPassing string `json:"lastPassing"`
}

func ToRMLT(raw string, parts []string) RMLT {
	return RMLT{
		msg{raw},
		parts[0],
		parts[1],
	}
}

type RMDTL struct {
	msg
}

func ToRMDTL(raw string, parts []string) RMDTL {
	return RMDTL{
		msg{raw},
	}
}

type RMCA struct {
	msg
	RelayServerTime string `json:"relayServerTime"`
}

func ToRMCA(raw string, parts []string) RMCA {
	return RMCA{
		msg{raw},
		parts[0],
	}
}

type RMHL struct {
	msg
	RacerId      string `json:"racerId"`
	LapNumber    string `json:"lapNumber"`
	RacePosition string `json:"racePosition"`
	LapTime      string `json:"lapTime"`
	FlagStatus   string `json:"flagStatus"`
	TotalTime    string `json:"totalTime"`
}

func ToRMHL(raw string, parts []string) RMHL {
	return RMHL{
		msg{raw},
		parts[0],
		parts[1],
		parts[2],
		parts[3],
		parts[4],
		parts[5],
	}
}

func NoQuote(s string) string {
	return strings.Trim(s, "\"")
}
