package flags

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/axatol/go-utils/ds"
	"gopkg.in/yaml.v3"
)

type Set struct{ flag.FlagSet }

func (s *Set) UnsetFlags() []*flag.Flag {
	unset := ds.NewSet[*flag.Flag]()
	s.VisitAll(func(f *flag.Flag) { unset.Add(f) })
	s.Visit(func(f *flag.Flag) { unset.Del(f) })
	return unset.Entries()
}

func (s *Set) LoadUnsetFromEnv() error {
	for _, f := range s.UnsetFlags() {
		key := f.Name
		key = strings.ReplaceAll(key, "-", "_")
		key = strings.ToUpper(key)
		val, ok := os.LookupEnv(key)
		if !ok {
			continue
		}

		if err := f.Value.Set(val); err != nil {
			return fmt.Errorf("could not set value from environment variable %s: %s", key, err)
		}
	}

	return nil
}

func (s *Set) LoadUnsetFromMap(inputs map[string]any) error {
	for _, f := range s.UnsetFlags() {
		val, ok := inputs[f.Name]
		if !ok {
			continue
		}

		raw, err := json.Marshal(val)
		if err != nil {
			return fmt.Errorf("failed to serialise value at key %s: %s", f.Name, err)
		}

		if err := f.Value.Set(string(raw)); err != nil {
			return fmt.Errorf("failed to set %s from file: %s", f.Name, err)
		}
	}

	return nil
}

func (s *Set) LoadUnsetFromJSONFile(filename string) error {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return fmt.Errorf("failed to parse file: %s", err)
	}

	return s.LoadUnsetFromMap(parsed)
}

func (s *Set) LoadUnsetFromYAMLFile(filename string) error {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", err)
	}

	var parsed map[string]any
	if err := yaml.Unmarshal(raw, &parsed); err != nil {
		return fmt.Errorf("failed to parse file: %s", err)
	}

	return s.LoadUnsetFromMap(parsed)
}
