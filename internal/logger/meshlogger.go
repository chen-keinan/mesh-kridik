package logger

import (
	"log"
)

//MeshKridikLogger Object
type MeshKridikLogger struct {
}

//GetLog return native logger
func GetLog() *MeshKridikLogger {
	return &MeshKridikLogger{}
}

//Console print to console
func (BLogger *MeshKridikLogger) Console(str string) {
	log.SetFlags(0)
	log.Print(str)
}

//Table print to console
func (BLogger *MeshKridikLogger) Table(v interface{}) {
	log.SetFlags(0)
	log.Print(v)
}
