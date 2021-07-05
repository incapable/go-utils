//go:build !release

package logger

func initLogger() {
	setupLogger(true)
}
