package statistics

import (
	"go_demo/src/net_demo/service"
	"fmt"
	"reflect"
	"strings"
	"time"
)

/**
统计服务
*/
type StatisticsService struct {
	timeRecorderMap map[string]*TimeRecorder
	inputChan       chan interface{}
}

func NewStatisticsService() *StatisticsService {
	return &StatisticsService{}
}

func (s *StatisticsService) Startup() {
	s.timeRecorderMap = make(map[string]*TimeRecorder)
	s.inputChan = make(chan interface{})

	go func() {
		for input := range s.inputChan {
			switch reflect.TypeOf(input).Name() {
			case "CommandInfo":
				commandInfo := input.(*CommandInfo)
				timeRecorder := s.timeRecorderMap[commandInfo.Command]
				if timeRecorder == nil {
					timeRecorder = new(TimeRecorder)
					s.timeRecorderMap[commandInfo.Command] = timeRecorder
				}
				//记录
				timeRecorder.recorder(commandInfo)
				break
			case "string":
				c := input.(string)
				if c == "show" {
					var sb strings.Builder
					for k, v := range s.timeRecorderMap {
						time := float64(v.totalTime) / float64(v.totalCount)
						fmt.Fprint(&sb, "\r\n")
						fmt.Fprint(&sb, "Command=")
						fmt.Fprint(&sb, k)
						fmt.Fprint(&sb, " [totalCount=")
						fmt.Fprint(&sb, v.totalCount)
						fmt.Fprint(&sb, " totalTime=")
						fmt.Fprint(&sb, v.totalTime)
						fmt.Fprint(&sb, "ms errorCount=")
						fmt.Fprint(&sb, v.errorCount)
						fmt.Fprint(&sb, " averageTime=")
						fmt.Fprint(&sb, time)
						fmt.Fprint(&sb, "ms]")
					}
					fmt.Println(sb.String())
				}
				break
			}
		}
	}()
}

func (s *StatisticsService) Shutdown() {
	close(s.inputChan)
}

func (s *StatisticsService) ServiceType() string {
	return service.ServiceType_Statistics
}

func (s *StatisticsService) ServiceName() string {
	return service.ServiceDefaultName
}

func (s *StatisticsService) Recorder(commandInfo *CommandInfo) {
	s.inputChan <- commandInfo
}

func (s *StatisticsService) Show() {
	s.inputChan <- "show"
}

//----------------------------------------
type CommandInfo struct {
	SendTime  time.Time //发送时间
	Command   string    //请求的指令
	Time      int64     //消耗时间
	IsSuccess bool      //是否成功
}

//----------------------------------------
/**
 * 时间记录
 */
type TimeRecorder struct {
	totalCount int64 //总次数
	totalTime  int64 //时间
	errorCount int64 //失败次数
}

func (s *TimeRecorder) recorder(commandInfo *CommandInfo) {
	s.totalCount++
	s.totalTime += commandInfo.Time
	if !commandInfo.IsSuccess {
		s.errorCount++
	}
}
