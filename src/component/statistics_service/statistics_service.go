package statistics_service

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

type CommandInfo struct {
	SendTime   time.Time //发送时间
	Command    string    //请求的指令
	ReturnTime time.Time //
	IsSuccess  bool      //是否成功
	Flow       int64     //流量(字节)
}

/**
 * 时间记录
 */
type TimeRecorder struct {
	totalCount int64 //总次数
	totalTime  int64 //时间
	errorCount int64 //失败次数
	maxTime    int64 //
	minTime    int64 //
	flow       int64 //流量(字节)
}

func (s *TimeRecorder) recorder(commandInfo *CommandInfo) {
	s.totalCount++
	temp := commandInfo.ReturnTime.UnixNano() - commandInfo.SendTime.UnixNano()
	if temp > s.maxTime {
		s.maxTime = temp
	}
	if s.minTime < 0 {
		s.minTime = temp
	} else if temp < s.minTime {
		s.minTime = temp
	}
	s.totalTime += temp
	s.flow += commandInfo.Flow
	if !commandInfo.IsSuccess {
		s.errorCount++
	}
}

/**
统计服务
*/
type StatisticsService struct {
	timeRecorderMap map[string]*TimeRecorder
	inputChan       chan interface{}
}

func (s *StatisticsService) StartUp() {
	s.timeRecorderMap = make(map[string]*TimeRecorder)
	s.inputChan = make(chan interface{})

	go func() {
		for input := range s.inputChan {
			switch input.(type) {
			case *CommandInfo:
				commandInfo := input.(*CommandInfo)
				timeRecorder := s.timeRecorderMap[commandInfo.Command]
				if timeRecorder == nil {
					timeRecorder = &TimeRecorder{
						minTime: -1,
					}
					s.timeRecorderMap[commandInfo.Command] = timeRecorder
				}
				//记录
				timeRecorder.recorder(commandInfo)
				//fmt.Println("recorder ", commandInfo)
				break
			case string:
				c := input.(string)
				if c == "show" {
					var sb bytes.Buffer
					sb.WriteString("指令统计:")
					for k, v := range s.timeRecorderMap {
						time := float64(v.totalTime) / float64(v.totalCount)
						sb.WriteString("\r\n")
						sb.WriteString("Command=")
						sb.WriteString(k)
						sb.WriteString(" [totalCount=")
						sb.WriteString(strconv.FormatInt(v.totalCount, 10))
						sb.WriteString(",totalTime=")
						sb.WriteString(strconv.FormatInt(v.totalTime/1000000, 10))
						sb.WriteString("ms,errorCount=")
						sb.WriteString(strconv.FormatInt(v.errorCount, 10))
						sb.WriteString(",averageTime=")
						sb.WriteString(strconv.FormatInt(int64(time/1000000), 10))
						sb.WriteString("ms,minTime=")
						sb.WriteString(strconv.FormatInt(v.minTime/1000000, 10))
						sb.WriteString("ms,maxTime=")
						sb.WriteString(strconv.FormatInt(v.maxTime/1000000, 10))
						sb.WriteString("ms,flow=")
						sb.WriteString(strconv.FormatInt(v.flow, 10))
						sb.WriteString("]")
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

func (s *StatisticsService) Recorder(commandInfo *CommandInfo) {
	commandInfo.ReturnTime = time.Now()
	s.inputChan <- commandInfo
}

func (s *StatisticsService) Show() {
	s.inputChan <- "show"
}
