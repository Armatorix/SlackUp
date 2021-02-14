package xos

import (
	"fmt"
	"os"
	"strings"
)

func ResolveEnvs(txt string) string {
	envs := os.Environ()
	var replacePairs []string
	for _, env := range envs {
		splitEnv := strings.SplitN(env, "=", 2)
		if len(splitEnv) != 2 {
			continue
		}
		key, value := splitEnv[0], splitEnv[1]
		replacePairs = append(replacePairs,
			fmt.Sprintf("$%s", key), value,
			fmt.Sprintf("${%s}", key), value)
	}
	replacer := strings.NewReplacer(replacePairs...)
	return replacer.Replace(txt)
}
