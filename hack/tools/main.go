//go:build tools

package tools

/* for ensuring that the go mod tidy will not removed the deps for codegen */

import (
	_ "go.uber.org/mock/mockgen"
)
