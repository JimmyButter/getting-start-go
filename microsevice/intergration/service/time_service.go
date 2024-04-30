package service

import (
	"time"

	"hertz_demo/request/http"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TimerStart() {

	// var mutex sync.Mutex
	instrumentSyncTimer := time.NewTicker(2 * time.Second)
	shakeRecordSyncTimer := time.NewTicker(7 * time.Second)

	for {
		select {
		case <-instrumentSyncTimer.C:
			InstrumentSync()
		case <-shakeRecordSyncTimer.C:
			ShakeRecordSync()
		}
	}

	// for range instrumentSyncTimer.C {
	// 	instrumentSync()
	// }

	// for range shakeRecordSyncTimer.C {
	// 	shakeRecordSync()
	// }

	// go func() {
	// 	mutex.Lock()
	// 	defer mutex.Unlock()

	// 	if !instrumentSyncTimer.Stop() {
	// 		<-instrumentSyncTimer.C
	// 	}

	// 	instrumentSyncTimer.Reset(2 * time.Second)
	// }()

	// go func() {
	// 	mutex.Lock()
	// 	defer mutex.Unlock()
	// 	if !shakeRecordSyncTimer.Stop() {
	// 		<-shakeRecordSyncTimer.C
	// 	}
	// }()

	// <-instrumentSyncTimer.C
	// instrumentSync()

	// <-shakeRecordSyncTimer.C
	// shakeRecordSync()
}

func InstrumentSync() {
	hlog.Infof("Instrumenting...\n")
	instruments := http.GetInstruments()
	for inst := range instruments {
		hlog.Info(inst)
	}

}

func ShakeRecordSync() {
	hlog.Infof("shakeRecordSync...\n")
}
