// Package mytask
// @Author: zhangdi
// @File: manager
// @Version: 1.0.0
// @Date: 2023/11/16 14:24
package mytask

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
)

// TaskManager 定时任务管理器
type TaskManager struct {
	//运行中的任务列表
	RunTasks []*TaskInfo
	//cron 自动任务实例
	c *cron.Cron
}

// NewTaskManager 创建定时任务管理器
func NewTaskManager() *TaskManager {
	//创建定时任务对象
	c := cron.New(cron.WithSeconds())
	//开启定时任务
	c.Start()

	return &TaskManager{c: c}
}

// AddTask 添加一个定时任务
// spec可以使用spec.go中的函数进行快速组装
// taskName 任务名，不可重复，重复会导致覆盖
func (s *TaskManager) AddTask(taskName string, spec string, cmd func()) (cron.EntryID, error) {
	if s.CheckTaskExists(taskName) != nil {
		return 0, errors.New(fmt.Sprintf("任务名%s，已存在，请更换名称", taskName))
	}
	id, err := s.c.AddFunc(spec, cmd)
	if err != nil {
		return id, err
	}
	s.runTaskAdd(taskName, int(id))
	return id, nil
}

// CheckTaskExists 根据任务名检测人物是否存在，不存在则返回nil,存在则返回info
func (s *TaskManager) CheckTaskExists(taskName string) *TaskInfo {
	for _, task := range s.RunTasks {
		if task.TaskName == taskName { //已存在任务
			return task
		}
	}
	return nil
}

// 向正在运行中的任务加入该任务
func (s *TaskManager) runTaskAdd(taskName string, taskId int) {
	s.RunTasks = append(s.RunTasks, &TaskInfo{
		EntryID:  taskId,
		TaskName: taskName,
		RunCount: 0,
	})
}

// RunTaskRemove 删除一个正在运行的任务
func (s *TaskManager) RunTaskRemove(taskName string) error {
	taskInfo := s.CheckTaskExists(taskName)
	if taskInfo == nil {
		return errors.New(fmt.Sprintf("任务%s不存在，无法删除！", taskName))
	}
	//移除一个正在运行的任务
	s.c.Remove(cron.EntryID(taskInfo.EntryID))

	//更新任务列表
	newRunTasks := make([]*TaskInfo, 0)
	for i, task := range s.RunTasks {
		if task.TaskName != taskName { //已存在任务
			newRunTasks = append(newRunTasks, s.RunTasks[i])
		}
	}
	s.RunTasks = newRunTasks

	return nil
}
