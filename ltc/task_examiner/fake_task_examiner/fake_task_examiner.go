// This file was generated by counterfeiter
package fake_task_examiner

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/task_examiner"
)

type FakeTaskExaminer struct {
	TaskStatusStub        func(taskName string) (task_examiner.TaskInfo, error)
	taskStatusMutex       sync.RWMutex
	taskStatusArgsForCall []struct {
		taskName string
	}
	taskStatusReturns struct {
		result1 task_examiner.TaskInfo
		result2 error
	}
	ListTasksStub        func() ([]task_examiner.TaskInfo, error)
	listTasksMutex       sync.RWMutex
	listTasksArgsForCall []struct{}
	listTasksReturns     struct {
		result1 []task_examiner.TaskInfo
		result2 error
	}
}

func (fake *FakeTaskExaminer) TaskStatus(taskName string) (task_examiner.TaskInfo, error) {
	fake.taskStatusMutex.Lock()
	fake.taskStatusArgsForCall = append(fake.taskStatusArgsForCall, struct {
		taskName string
	}{taskName})
	fake.taskStatusMutex.Unlock()
	if fake.TaskStatusStub != nil {
		return fake.TaskStatusStub(taskName)
	} else {
		return fake.taskStatusReturns.result1, fake.taskStatusReturns.result2
	}
}

func (fake *FakeTaskExaminer) TaskStatusCallCount() int {
	fake.taskStatusMutex.RLock()
	defer fake.taskStatusMutex.RUnlock()
	return len(fake.taskStatusArgsForCall)
}

func (fake *FakeTaskExaminer) TaskStatusArgsForCall(i int) string {
	fake.taskStatusMutex.RLock()
	defer fake.taskStatusMutex.RUnlock()
	return fake.taskStatusArgsForCall[i].taskName
}

func (fake *FakeTaskExaminer) TaskStatusReturns(result1 task_examiner.TaskInfo, result2 error) {
	fake.TaskStatusStub = nil
	fake.taskStatusReturns = struct {
		result1 task_examiner.TaskInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskExaminer) ListTasks() ([]task_examiner.TaskInfo, error) {
	fake.listTasksMutex.Lock()
	fake.listTasksArgsForCall = append(fake.listTasksArgsForCall, struct{}{})
	fake.listTasksMutex.Unlock()
	if fake.ListTasksStub != nil {
		return fake.ListTasksStub()
	} else {
		return fake.listTasksReturns.result1, fake.listTasksReturns.result2
	}
}

func (fake *FakeTaskExaminer) ListTasksCallCount() int {
	fake.listTasksMutex.RLock()
	defer fake.listTasksMutex.RUnlock()
	return len(fake.listTasksArgsForCall)
}

func (fake *FakeTaskExaminer) ListTasksReturns(result1 []task_examiner.TaskInfo, result2 error) {
	fake.ListTasksStub = nil
	fake.listTasksReturns = struct {
		result1 []task_examiner.TaskInfo
		result2 error
	}{result1, result2}
}

var _ task_examiner.TaskExaminer = new(FakeTaskExaminer)
