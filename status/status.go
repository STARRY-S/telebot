package status

// #include <sys/sysinfo.h>
import "C"

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/STARRY-S/telebot/utils"
)

const (
	BYTE    = 1
	KB      = 1024 * BYTE
	MB      = 1024 * KB
	GB      = 1024 * MB
	SECOND  = 1
	MINUTES = 60 * SECOND
	HOUR    = 60 * MINUTES
	DAY     = 24 * HOUR
)

type sysinfo struct {
	uptime    float64
	totalram  float64
	freeram   float64
	totalswap float64
	freeswap  float64
}

func GetStatus() (string, error) {
	cpuTemp, err := cpuTemp()
	if err != nil {
		return "", err
	}
	var info C.struct_sysinfo
	C.sysinfo(&info)
	sysinfo := sysinfo{
		uptime:    float64(info.uptime),
		totalram:  float64(info.totalram),
		freeram:   float64(info.freeram),
		totalswap: float64(info.totalswap),
		freeswap:  float64(info.freeswap),
	}

	infoBuff := &bytes.Buffer{}
	fmt.Fprintf(infoBuff, "CPU:        %v\n", cpuTemp)
	fmt.Fprintf(infoBuff, "Uptime:     %.2f hour\n", sysinfo.uptime/HOUR)
	fmt.Fprintf(infoBuff, "Total RAM:  %.2fG\n", sysinfo.totalram/GB)
	fmt.Fprintf(infoBuff, "Free RAM:   %.2fG\n", sysinfo.freeram/GB)
	fmt.Fprintf(infoBuff, "Total Swap: %.2fG\n", sysinfo.totalswap/GB)
	fmt.Fprintf(infoBuff, "Free Swap:  %.2fG\n", sysinfo.freeswap/GB)

	return infoBuff.String(), nil
}

func sensors() (string, error) {
	path, err := exec.LookPath("sensors")
	if err != nil {
		return "", errors.New("sensors not found")
	}

	return utils.RunCommandFunc(path)
}

func cpuTemp() (string, error) {
	sensors, err := sensors()
	if err != nil {
		return "", fmt.Errorf("cpuTemp: exec command: %w", err)
	}

	sensors = strings.TrimSpace(sensors)
	scanner := bufio.NewScanner(strings.NewReader(sensors))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Tctl") {
			var param []string
			for _, v := range strings.Split(line, " ") {
				if len(v) > 0 {
					param = append(param, v)
				}
			}
			if len(param) == 2 {
				return param[1], nil
			}
		}
		if strings.Contains(line, "Tccd") {
			var param []string
			for _, v := range strings.Split(line, " ") {
				if len(v) > 0 {
					param = append(param, v)
				}
			}
			if len(param) == 2 {
				return param[1], nil
			}
		}
	}

	return "", fmt.Errorf("cpuTemp: failed to get CPU temp")
}
