package logger

import (
	"log"
)

//LdxProbeLogger Object
type LdxProbeLogger struct {
}

//GetLog return native logger
func GetLog() *LdxProbeLogger {
	return &LdxProbeLogger{}
}

//Console print to console
func (BLogger *LdxProbeLogger) Console(str string) {
	log.SetFlags(0)
	log.Print(str)
}

//Table print to console
func (BLogger *LdxProbeLogger) Table(v interface{}) {
	log.SetFlags(0)
	log.Print(v)
}
