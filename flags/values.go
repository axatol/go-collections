package flags

import (
	"flag"
	"fmt"
	"strings"
)

var (
	_ flag.Getter = (*EnumValue)(nil)
)

type EnumValue struct {
	Valid   []string
	Default *string
	value   *string
}

func (v *EnumValue) String() string {
	if v != nil && len(v.Valid) > 0 {
		return v.Valid[0]
	}

	if v.Default != nil {
		return *v.Default
	}

	return ""
}

func (v *EnumValue) Set(val string) error {
	for _, valid := range v.Valid {
		if val == valid {
			v.value = &val
			return nil
		}
	}

	return fmt.Errorf("value must be one of [%s]", strings.Join(v.Valid, ", "))
}

func (v *EnumValue) Get() any {
	return v.String()
}

// matches zerolog log levels
var ValidLogLevels = []string{"trace", "debug", "info", "warn", "error", "fatal", "panic"}

type LogLevelValue struct{ value *string }

func (v *LogLevelValue) String() string {
	if v == nil || v.value == nil {
		return ""
	}

	return string(*v.value)
}

func (v *LogLevelValue) Set(val string) error {
	for _, valid := range ValidLogLevels {
		if val == valid {
			v.value = &valid
			return nil
		}
	}

	return fmt.Errorf("value must be a valid log level")
}

func (v *LogLevelValue) Get() any {
	return v.String()
}
