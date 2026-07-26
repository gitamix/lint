package config

import (
	"os"
)

// Load reads content from provided filepath
// and unmarshals YAML file to Config
// or returns error if file reading or unmarshaling failed.
func Load(fpath string) (*Config, error) {
	bb, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return Unmarshal(bb)
}
