package pack1

import(
	"log/slog"
)

func Nop(){
	slog.Info("I am into test-module", slog.String("package", "pack1"))
}
