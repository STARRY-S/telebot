package status

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	Uptime       float32
	MemTotal     float32
	MemFree      float32
	MemAvailable float32
	SwapTotal    float32
	SwapFree     float32
}

func GetStatus() (string, error) {
	cpuTemp, err := cpuTemp()
	if err != nil {
		return "", err
	}

	meminfo, err := os.Open("/proc/meminfo")
	if err != nil {
		return "", err
	}
	defer meminfo.Close()
	scanner := bufio.NewScanner(meminfo)
	scanner.Split(bufio.ScanLines)
	info := sysinfo{}
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "MemTotal"):
			for _, s := range strings.Split(line, " ") {
				if i, err := strconv.Atoi(s); err == nil {
					info.MemTotal = float32(i) / GB * KB
					break
				}
			}
		case strings.HasPrefix(line, "MemFree"):
			for _, s := range strings.Split(line, " ") {
				if i, err := strconv.Atoi(s); err == nil {
					info.MemFree = float32(i) / GB * KB
					break
				}
			}
		case strings.HasPrefix(line, "MemAvailable"):
			for _, s := range strings.Split(line, " ") {
				if i, err := strconv.Atoi(s); err == nil {
					info.MemAvailable = float32(i) / GB * KB
					break
				}
			}
		case strings.HasPrefix(line, "SwapTotal"):
			for _, s := range strings.Split(line, " ") {
				if i, err := strconv.Atoi(s); err == nil {
					info.SwapTotal = float32(i) / GB * KB
					break
				}
			}
		case strings.HasPrefix(line, "SwapFree"):
			for _, s := range strings.Split(line, " ") {
				if i, err := strconv.Atoi(s); err == nil {
					info.SwapFree = float32(i) / GB * KB
					break
				}
			}
		}
	}
	uptimeInfo, err := os.Open("/proc/uptime")
	if err != nil {
		return "", err
	}
	defer uptimeInfo.Close()
	uptimeBytes := make([]byte, 64)
	if _, err := uptimeInfo.Read(uptimeBytes); err != nil {
		return "", err
	}
	uptimeSlice := strings.Split(string(uptimeBytes), " ")
	if len(uptimeSlice) == 0 {
		return "", fmt.Errorf("failed to get uptime")
	}
	uptime, err := strconv.ParseFloat(uptimeSlice[0], 32)
	if err != nil {
		return "", err
	}
	info.Uptime = float32(uptime) / HOUR * SECOND

	infoBuff := &bytes.Buffer{}
	// Actual Free Memory = Free + Buffers + Cached
	fmt.Fprintf(infoBuff, "CPU: %v\n", cpuTemp)
	fmt.Fprintf(infoBuff, "Uptime: %.2f Hour\n", info.Uptime)
	fmt.Fprintf(infoBuff, "TotalRAM: %.2fG\n", info.MemTotal)
	fmt.Fprintf(infoBuff, "FreeRAM: %.2fG\n", info.MemFree)
	fmt.Fprintf(infoBuff, "AvailableRAM: %.2fG\n", info.MemAvailable)
	fmt.Fprintf(infoBuff, "TotalSwap: %.2fG\n", info.SwapTotal)
	fmt.Fprintf(infoBuff, "FreeSwap: %.2fG\n", info.SwapFree)

	return infoBuff.String(), nil
}

func sensors() (string, error) {
	path, err := exec.LookPath("sensors")
	if err != nil {
		return "", errors.New("sensors not found")
	}

	var stdout bytes.Buffer
	cmd := exec.Command(path)
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("sensors: %w", err)
	}
	return stdout.String(), nil
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
