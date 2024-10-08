package ephem

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

var deviceGlobs = []string{
	"/dev/stratis/*/*",
	"/dev/disk/by-uuid/*",
	"/dev/mapper/*",
	"/dev/disk/by-label/*",
}

func StableDevicePath(device string) (string, error) {
	l := log.WithPrefix("stableDevicePath")

	if !strings.HasPrefix("/", device) {
		return device, nil
	}
	stat, err := os.Stat(device)
	if err != nil {
		return "", err
	}

	for idx := range deviceGlobs {
		glob := deviceGlobs[idx]
		l.Debugf("searching glob: %s", glob)

		matches, err := filepath.Glob(glob)
		if err != nil {
			// the only possible error is ErrBadPattern
			return "", err
		}
		for _, match := range matches {
			matchStat, err := os.Stat(match)
			if err != nil {
				l.Debugf("failed to stat match %s: %s", match, err)
				continue
			}
			if os.SameFile(stat, matchStat) {
				l.Debugf("match %s found for device %s", match, device)
				return match, nil
			}
		}
	}

	l.Debugf("no match found for device %s", device)
	// if not match was found, we return the original device path
	return device, nil
}
