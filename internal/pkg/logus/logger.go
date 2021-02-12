package logus

import (
	murlog "github.com/Melenium2/Murlog"
	"log"
)

type LogusExternal interface {
	LogExternal(url string, log Log)
}

type Logus interface {
	Log(key, value interface{})
	LogMany(units ...LUnit)
	LogusExternal
}

type LUnit struct {
	Key interface{}
	Val interface{}
}

func NewLUnit(k, v interface{}) LUnit {
	return LUnit{k, v}
}

type LogusImpl struct {
	l murlog.Logger
}

func (l *LogusImpl) Log(key, value interface{}) {
	l.LogMany(NewLUnit(key, value))
}

func (l *LogusImpl) LogMany(units ...LUnit) {
	vls := make([]interface{}, 0)
	for _, u := range units {
		if u.Key == nil {
			u.Key = "None"
		}
		if u.Val == nil {
			u.Val = "None"
		}
		vls = append(vls, u.Key, u.Val)
	}
	l.l.Log(vls...)
}

func (l *LogusImpl) LogExternal(url string, log Log) {
	l.Log("msg", log.String())
	go func() {
		err := log.SendLog(url)
		l.Log("error", err)
	}()
}


func New(logger murlog.Logger) *LogusImpl {
	return &LogusImpl{
		logger,
	}
}

type EmptyLogusImpl struct {

}

func (l *EmptyLogusImpl) Log(key, value interface{}) {
	log.Print(key, " ", value)
}

func (l *EmptyLogusImpl) LogMany(units ...LUnit) {
	log.Print(units)
}

func (l *EmptyLogusImpl) LogExternal(url string, log Log) {
	l.Log("msg", log.String())
}


func NewEmptyLogger() *EmptyLogusImpl {
	return &EmptyLogusImpl{}
}
