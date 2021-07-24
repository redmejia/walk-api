package logs

import "log"

type Logers struct {
	Error *log.Logger
	Info  *log.Logger
}
