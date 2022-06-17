/*
 * @Author: your name
 * @Date: 2022-04-23 15:55:42
 * @LastEditTime: 2022-04-23 16:39:56
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_basic/log/logrus.go
 */
package logrus

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LogrusPrint() {
	log.SetLevel(log.TraceLevel)
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}

func LogrusPrintf() {
	log.SetLevel(log.TraceLevel)
	// log.SetFormatter(&log.TextFormatter{
	// 	DisableColors: false,
	// 	FullTimestamp: true,
	// })

	log.Tracef("Something very low level.%s", "I am trancef")
	log.Debugf("Useful debugging information.%s", "I am Debugf")
	log.Infof("Something noteworthy happened!,%s", "I am Infof")
	log.Warnf("You should probably take a look at this.")
	log.Errorf("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatalf("Bye.")
	// Calls panic() after logging
	log.Panicf("I'm bailing.")
}

func LogrusWithFields() {
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.WithFields(log.Fields{
		"field": "info",
		"task":  "test info",
	}).Info("test logrus with field")
	log.SetOutput(os.Stdout)
	// or
	logEntry := log.WithFields(log.Fields{"field": "debug", "task": "test Debug"})
	logEntry.Debug("test logrus With Field")
}
