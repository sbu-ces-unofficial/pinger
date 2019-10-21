package pflags

import (
	"bufio"
	"os"
	"regexp"
)

var arrayNamePattern = regexp.MustCompile(`\[[-_a-zA-Z]+\]`)

type config struct {
	Arrays map[string][]string
}

func Parse(filename string) (*config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &config{
		Arrays: make(map[string][]string),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if matches := arrayNamePattern.FindAllString(line, -1); len(matches) > 0 && matches[0] == line {
			arrayName := line[1 : len(line)-1]
			cfg.Arrays[arrayName] = make([]string, 0)
			for scanner.Scan() {
				line = scanner.Text()
				if line != "" {
					cfg.Arrays[arrayName] = append(cfg.Arrays[arrayName], line)
				} else {
					break
				}
			}
		}
	}

	return cfg, nil
}
