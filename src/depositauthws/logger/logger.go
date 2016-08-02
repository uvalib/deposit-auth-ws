package logger

import (
    "log"
)

func Log( msg string ) {
    log.Printf( "DEPOSITAUTH: %s", msg )
}