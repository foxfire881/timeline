package util

import (
	"regexp"
	"strconv"
	"time"
)

// ====================================================
func FormatTimeI2S(sec *int64, precision string) string {
	if sec == nil {
		return ""
	}
	t := time.Unix(*sec, 0).UTC()
	switch precision {
	case "y":
		return t.Format("2006")
	case "M":
		return t.Format("2006-01")
	case "h":
		return t.Format("2006-01-02 15:00:00")
	case "m":
		return t.Format("2006-01-02 15:04:00")
	case "s":
		return t.Format("2006-01-02 15:04:05")
	default:
		return t.Format("2006-01-02")
	}
}

func FormatTimeS2I(str string) *int64 {
	// x2006-01-02 15:04:05
	// x2006-01-02 15:04
	// x2006-01-02 15
	// x2006-01-02
	// x2006-01
	// x2006
	if str == "" {
		return nil
	}
	var match []string
	match = regexp.MustCompile(`^(?i)([-]?\d+)-(\d+)-(\d+) (\d+):(\d+):(\d+)$`).FindStringSubmatch(str)
	if match != nil {
		year, _ := strconv.Atoi(match[1])
		month, _ := strconv.Atoi(match[2])
		day, _ := strconv.Atoi(match[3])
		hour, _ := strconv.Atoi(match[4])
		minute, _ := strconv.Atoi(match[5])
		second, _ := strconv.Atoi(match[6])
		if month < 1 || month > 12 || day < 1 || day > 31 || hour < 0 || hour > 23 || minute < 0 || minute > 59 || second < 0 || second > 59 {
			panic("时间格式错误:" + str)
		}
		t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC).UTC().Unix()
		return &t
	}
	match = regexp.MustCompile(`^(?i)([-]?\d+)-(\d+)-(\d+) (\d+):(\d+)$`).FindStringSubmatch(str)
	if match != nil {
		year, _ := strconv.Atoi(match[1])
		month, _ := strconv.Atoi(match[2])
		day, _ := strconv.Atoi(match[3])
		hour, _ := strconv.Atoi(match[4])
		minute, _ := strconv.Atoi(match[5])
		if month < 1 || month > 12 || day < 1 || day > 31 || hour < 0 || hour > 23 || minute < 0 || minute > 59 {
			panic("时间格式错误:" + str)
		}
		t := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC).UTC().Unix()
		return &t
	}
	match = regexp.MustCompile(`^(?i)([-]?\d+)-(\d+)-(\d+) (\d+)$`).FindStringSubmatch(str)
	if match != nil {
		year, _ := strconv.Atoi(match[1])
		month, _ := strconv.Atoi(match[2])
		day, _ := strconv.Atoi(match[3])
		hour, _ := strconv.Atoi(match[4])
		if month < 1 || month > 12 || day < 1 || day > 31 || hour < 0 || hour > 23 {
			panic("时间格式错误:" + str)
		}
		t := time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.UTC).UTC().Unix()
		return &t
	}
	match = regexp.MustCompile(`^(?i)([-]?\d+)-(\d+)-(\d+)$`).FindStringSubmatch(str)
	if match != nil {
		year, _ := strconv.Atoi(match[1])
		month, _ := strconv.Atoi(match[2])
		day, _ := strconv.Atoi(match[3])
		if month < 1 || month > 12 || day < 1 || day > 31 {
			panic("时间格式错误:" + str)
		}
		t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).UTC().Unix()
		return &t
	}
	match = regexp.MustCompile(`^(?i)([-]?\d+)-(\d+)$`).FindStringSubmatch(str)
	if match != nil {
		year, _ := strconv.Atoi(match[1])
		month, _ := strconv.Atoi(match[2])
		if month < 1 || month > 12 {
			panic("时间格式错误:" + str)
		}
		t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).UTC().Unix()
		return &t
	}
	match = regexp.MustCompile(`^(?i)([-]?\d+)$`).FindStringSubmatch(str)
	if match != nil {
		year, _ := strconv.Atoi(match[1])
		t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).UTC().Unix()
		return &t
	}
	panic("时间格式错误:" + str)
}
