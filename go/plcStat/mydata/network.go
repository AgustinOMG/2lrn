package mydata

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Get network statistics
func GetNet() ([]NetStats, error) {
	// Reference: man 5 proc, Documentation/filesystems/proc.txt in Linux source code
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return collectNetworkStats(file)
}

// Stats represents network statistics for linux
type NetStats struct {
	Name             string
	RxBytes, TxBytes uint64
}

func collectNetworkStats(out io.Reader) ([]NetStats, error) {
	scanner := bufio.NewScanner(out)
	var networks []NetStats
	for scanner.Scan() {
		// Reference: dev_seq_printf_stats in Linux source code
		kv := strings.SplitN(scanner.Text(), ":", 2)
		if len(kv) != 2 {
			continue
		}
		fields := strings.Fields(kv[1])
		if len(fields) < 16 {
			continue
		}
		name := strings.TrimSpace(kv[0])
		if name == "lo" {
			continue
		}
		rxBytes, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse rxBytes of %s", name)
		}
		txBytes, err := strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse txBytes of %s", name)
		}
		networks = append(networks, NetStats{Name: name, RxBytes: rxBytes, TxBytes: txBytes})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for /proc/net/dev: %s", err)
	}
	return networks, nil
}