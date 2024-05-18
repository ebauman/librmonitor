package librmonitor

import (
	"encoding/csv"
	"strings"
	"testing"
)

var (
	valid_f_values = map[string]string{
		"noflag": `$F,9999,"00:00:00","14:10:03","00:59:59","      "`,
		"green":  `$F,9999,"00:00:00","14:10:05","00:00:00","Green "`,
		"yellow": `$F,9999,"00:00:00","14:10:05","00:00:00","Yellow"`,
		"red":    `$F,9999,"00:00:00","14:10:05","00:00:00","Red   "`,
	}

	invalid_f_values = map[string]string{
		"garbage":      "ljksdfljksdfljsdf",
		"malformat":    `$F`,
		"weird_commas": `$F,,,,,,`,
	}

	valid_a_values = map[string]string{
		"one":   `$A,"21","21",8852130,"Farnbacher /","James","Panoz Esperante",1`,
		"two":   `$A,"28","28",7347646,"Gigliotti / Drissi /","Ruhlman / Goossens","Riley Corvette  C6",1`,
		"three": `$A,"40","40",9367020,"D. Robertson/ Murry/","A.Robertson","Doran GT MK 7",1`,
		"four":  `$A,"44","44",4201897,"van Overbeek /","Law/ Neiman","Porsche 911 GT3  RSR",1`,
		"five":  `$A,"45","45",7729804,"Bergmeister/","Long","Porsche 911 GT3  RSR",1`,
		"six":   `$A,"46","46",1946249,"TBA /","TBA","Porsche 911 GT3  RSR",1`,
	}

	invalid_a_values = map[string]string{
		"garbage":      "lksjdfljsdfljsdfljsf",
		"malformat":    "$A",
		"weird_commas": "$A,,,,",
	}

	valid_comp_values = map[string]string{
		"one":    `$COMP,"21","21",1,"Farnbacher /","James","Panoz Esperante",""`,
		"two":    `$COMP,"28","28",1,"Gigliotti / Drissi /","Ruhlman / Goossens","Riley Corvette  C6",""`,
		"three":  `$COMP,"40","40",1,"D. Robertson/ Murry/","A.Robertson","Doran GT MK 7",""`,
		"four":   `$COMP,"44","44",1,"van Overbeek /","Law/ Neiman","Porsche 911 GT3  RSR",""`,
		"five":   `$COMP,"45","45",1,"Bergmeister/","Long","Porsche 911 GT3  RSR",""`,
		"six":    `$COMP,"46","46",1,"TBA /","TBA","Porsche 911 GT3  RSR",""`,
		"seven":  `$COMP,"87","87",1,"Werner /","Henzler","Porsche 911 GT3  RSR",""`,
		"eight":  `$COMP,"92","92",1,"Auberlen / Hand /","Milner / Mueller","",""`,
		"nine":   `$COMP,"9","9",2,"Brabham /","Sharp","Acura ARX 02a",""`,
		"ten":    `$COMP,"66","66",2,"de Ferran /","Pagenaud","Acura ARX 02a",""`,
		"eleven": `$COMP,"15","15",3,"Fernandez / Diaz /","Jourdain","Acura",""`,
	}

	invalid_comp_values = map[string]string{
		"garbage":      "lkjsdfljsdfljsdf",
		"malformat":    "$COMP",
		"weird_commas": "$COMP,,",
	}

	valid_b_values = map[string]string{
		"one":   `$B,33,"Test Session 5"`,
		"two":   `$B,19,"sdfsdafsadf"`,
		"three": `$B,2,"TESTING session"`,
	}

	invalid_b_values = map[string]string{
		"garbage":        "ljsdflkjsdflkjsdlfjsdf",
		"malformat":      "$B",
		"weird_commas":   "$B,",
		"weird_commas_2": "$B,,,,,,,,",
	}
)

func Test_F(t *testing.T) {
	runValids(t, valid_f_values, ToF)
	runInvalids(t, invalid_f_values, ToF)
}

func Test_A(t *testing.T) {
	runValids(t, valid_a_values, ToA)
	runInvalids(t, invalid_a_values, ToA)
}

func Test_Comp(t *testing.T) {
	runValids(t, valid_comp_values, ToComp)
	runInvalids(t, invalid_comp_values, ToComp)
}

func Test_B(t *testing.T) {
	runValids(t, valid_b_values, ToB)
	runInvalids(t, invalid_b_values, ToB)
}

func runValids[F any](t *testing.T, values map[string]string, f func(raw string, records []string) F) {
	for k, v := range values {
		t.Run(k+" should be valid", func(t *testing.T) {
			records, err := csv.NewReader(strings.NewReader(v)).Read()
			if err != nil {
				t.Error(err)
			}

			defer func() {
				if r := recover(); r != nil {
					t.Error("string index error")
				}
			}()

			f(v, records)
		})
	}
}

func runInvalids[F any](t *testing.T, values map[string]string, f func(raw string, records []string) F) {
	for k, v := range values {
		t.Run(k+" should be invalid", func(t *testing.T) {
			records, _ := csv.NewReader(strings.NewReader(v)).Read()

			defer func() {
				if r := recover(); r != nil {
				}
			}()

			f(v, records)
		})
	}
}
