package mydata

//"time"
//"golang.org/x/sys/unix"

func getUptime() string /*(time.Duration, error)*/ {
	return "Not ready yet"
	/*	var info unix.Sysinfo_t
		if err := unix.Sysinfo(&info); err != nil {
			return time.Duration(0), err
		}
		return time.Duration(info.Uptime) * time.Second, nil*/
}
