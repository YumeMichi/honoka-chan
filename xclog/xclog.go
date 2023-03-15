// Copyright (C) 2022 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package xclog

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Logg 日志
type Logg struct {
	fileDir  string
	fileName string
	filePre  string
	saveFile bool
	level    int
	date     string
	log      *log.Logger
}

// LogConf config struct of log
type LogConf struct {
	FileDir  string `json:"file_dir"`
	FilePre  string `json:"file_pre"`
	Level    int    `json:"level"`
	SaveFile bool   `json:"save_file"`
}

var (
	lgg          *Logg
	levelLinePre = map[int]string{
		1: "[F]",
		2: "[E]",
		3: "[W]",
		4: "[I]",
		5: "[D]",
	}
)

// 为防止未初始化调用奔溃，使用init默认参数初始化
func init() {
	lgg, _ = New("./", "", "", 5, false)
}

// Init 使用参数初始化
func Init(fileDir, filePre string, level int, saveFile bool) (*Logg, error) {
	var fileName string

	if level > 5 || level < 1 {
		level = 5
	}

	if !strings.HasSuffix(fileDir, "/") {
		fileDir = fileDir + "/"
	} else if fileDir == "" {
		fileDir = "./"
	}

	date := getNowDate()

	if filePre != "" {
		filePre = filePre + "_"
	}
	stat, err := os.Stat(fileDir)
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(fileDir, 0777)
	} else if err != nil {
		return nil, err
	} else if !stat.IsDir() {
		return nil, errors.New("log_dir is not dir:" + fileDir)
	}
	fileName = fileDir + filePre + date + ".log"
	newLogg, err := New(fileDir, filePre, fileName, level, saveFile)
	if err != nil {
		fmt.Println("Init log error:", err.Error())
		return nil, err
	}

	lgg = newLogg
	return lgg, nil
}

// InitByLogConf
func InitByLogConf(lc LogConf) (*Logg, error) {
	return Init(lc.FileDir, lc.FilePre, lc.Level, lc.SaveFile)
}

// New 返回一个日志实体，与默认不同的独立的日志，通过返回的Logg调用方法
func New(fileDir, filePre, fileName string, level int, saveFile bool) (*Logg, error) {
	// file, err := os.Create(fileName)
	var newLogg = &Logg{
		fileDir:  fileDir,
		filePre:  filePre,
		date:     getNowDate(),
		level:    level,
		fileName: fileName,
		saveFile: saveFile,
	}

	if saveFile {
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, err
		}
		os.Chmod(fileName, 0777)
		newLogg.log = log.New(file, "", 0)
	}

	return newLogg, nil
}

func getNowDate() string {
	return time.Now().Format("2006-01-02")
}

// Debug Debug
func Debug(args ...interface{}) {
	lgg.Debug(args...)
}

// Debugf Debugf
func Debugf(format string, args ...interface{}) {
	lgg.Debugf(format, args...)
}

// Debug Debug
func (lg *Logg) Debug(args ...interface{}) {
	lg.writeLine(5, args...)
}

// Debugf format debug
func (lg *Logg) Debugf(format string, args ...interface{}) {
	lg.writeLine(5, fmt.Sprintf(format, args...))
}

// Info Info
func Info(args ...interface{}) {
	lgg.Info(args...)
}

// Infof Infof
func Infof(format string, args ...interface{}) {
	lgg.Infof(format, args...)
}

// Info Info
func (lg *Logg) Info(args ...interface{}) {
	lg.writeLine(4, args...)
}

// Infof format debug
func (lg *Logg) Infof(format string, args ...interface{}) {
	lg.writeLine(4, fmt.Sprintf(format, args...))
}

// Warn Warn
func Warn(args ...interface{}) {
	lgg.Warn(args...)
}

// Warnf Warnf
func Warnf(format string, args ...interface{}) {
	lgg.Warnf(format, args...)
}

// Warn Warn
func (lg *Logg) Warn(args ...interface{}) {
	lg.writeLine(3, args...)
}

// Warnf format debug
func (lg *Logg) Warnf(format string, args ...interface{}) {
	lg.writeLine(3, fmt.Sprintf(format, args...))
}

// Error Error
func Error(args ...interface{}) {
	lgg.Error(args...)
}

// Errorf Errorf
func Errorf(format string, args ...interface{}) {
	lgg.Errorf(format, args...)
}

// Error Error
func (lg *Logg) Error(args ...interface{}) {
	lg.writeLine(2, args...)
}

// Errorf format debug
func (lg *Logg) Errorf(format string, args ...interface{}) {
	lg.writeLine(2, fmt.Sprintf(format, args...))
}

// Fatal Fatal
func Fatal(args ...interface{}) {
	lgg.Fatal(args...)
}

// Fatalf Fatalf
func Fatalf(format string, args ...interface{}) {
	lgg.Fatalf(format, args...)
}

// Fatal Fatal
func (lg *Logg) Fatal(args ...interface{}) {
	lg.writeLine(1, args...)
}

// Fatalf format debug
func (lg *Logg) Fatalf(format string, args ...interface{}) {
	lg.writeLine(1, fmt.Sprintf(format, args...))
}

// writeLine ...
func (lg *Logg) writeLine(level int, args ...interface{}) {
	_, file, line, _ := runtime.Caller(3)
	fileArr := strings.Split(file, "/")
	a := []interface{}{time.Now().Format("2006/01/02 15:04:05"), fileArr[len(fileArr)-1] + ":" + strconv.Itoa(line) + ":", levelLinePre[level]}
	a = append(a, args...)
	if !lg.saveFile {
		fmt.Println(a...)
		if level == 1 {
			os.Exit(1)
		}
		return
	}
	nowDate := getNowDate()
	if nowDate != lg.date {
		// 切割日志文件
		newFile := lg.fileDir + lg.filePre + nowDate + ".log"
		f, err := os.OpenFile(newFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			lg.log.Println("[E]", "create new log file:", err.Error())
		} else {
			lg.log.SetOutput(f)
			lg.date = nowDate
			lg.fileName = newFile
		}
	}

	if level == 1 {
		lg.log.Fatal(a...)
	}

	if level <= lg.level {
		lg.log.Println(a...)
	}
}
