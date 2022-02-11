package dms_to_dd_go

import (
	"strconv"
	"strings"
	"errors"
)

type dms struct {
	Degrees string
	Minutes  string
	Seconds string
}

//Pass degrees, minutes, seconds.
//Examples:
//NewDMS("44/1","44/1","39/1")
//NewDMS("47/1","23/1","110762/10000")
func NewDMS(d, m, s string) *dms {
	return &dms{
		Degrees: d,
		Minutes:  m,
		Seconds: s,
	}
}

//Returns coordinates in the format DD
func (d *dms) ConverterDMSToDD() (float64, error) {
	var de, m, s float64
	deSplit := split(d.Degrees)
	if len(deSplit) == 2 {
		deRes, err := calculateTheDifference(deSplit[0], deSplit[1])
		if err != nil {
			return 0, err
		}
		de = deRes
	} else {
		return 0, errors.New(`wrong degrees format`)
	}
	mSplit := split(d.Minutes)
	if len(mSplit) == 2 {
		mRes, err := calculateTheDifference(mSplit[0], mSplit[1])
		if err != nil {
			return 0, err
		}
		m = mRes
	} else {
		return 0, errors.New(`wrong minutes format`)
	}
	sSplit := split(d.Seconds)
	if len(sSplit) == 2 {
		sRes, err := calculateTheDifference(sSplit[0], sSplit[1])
		if err != nil {
			return 0, err
		}
		s = sRes
	} else {
		return 0, errors.New(`wrong seconds format`)
	}
	return calculateDMSToDD(de, m, s), nil
}

func calculateDMSToDD(d, m, s float64) float64 {
	return d + m/60 + s/3600
}

func calculateTheDifference(d1, d2 string) (float64, error) {
	s1, err := strconv.ParseFloat(d1, 64)
	if err != nil {
		return 0, err
	}
	s2, err := strconv.ParseFloat(d2, 64)
	if err != nil {
		return 0, err
	}
	res := s1 / s2
	return res, nil

}

func split(s string) []string {
	res := make([]string, 2)
	strings.Split(s, `/`)
	res = strings.Split(s, `/`)
	return res
}
