// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/lock"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/atc/worker"
)

type FakeClient struct {
	CreateVolumeStub        func(lager.Logger, worker.VolumeSpec, worker.WorkerSpec, db.VolumeType) (worker.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 worker.WorkerSpec
		arg4 db.VolumeType
	}
	createVolumeReturns struct {
		result1 worker.Volume
		result2 error
	}
	createVolumeReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 error
	}
	FindContainerStub        func(lager.Logger, int, string) (worker.Container, bool, error)
	findContainerMutex       sync.RWMutex
	findContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}
	findContainerReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	findContainerReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	FindVolumeStub        func(lager.Logger, int, string) (worker.Volume, bool, error)
	findVolumeMutex       sync.RWMutex
	findVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}
	findVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findVolumeReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	RunTaskStepStub        func(context.Context, lager.Logger, lock.LockFactory, db.ContainerOwner, worker.ContainerSpec, worker.WorkerSpec, worker.ContainerPlacementStrategy, db.ContainerMetadata, worker.ImageFetcherSpec, worker.TaskProcessSpec, chan runtime.Event) worker.TaskResult
	runTaskStepMutex       sync.RWMutex
	runTaskStepArgsForCall []struct {
		arg1  context.Context
		arg2  lager.Logger
		arg3  lock.LockFactory
		arg4  db.ContainerOwner
		arg5  worker.ContainerSpec
		arg6  worker.WorkerSpec
		arg7  worker.ContainerPlacementStrategy
		arg8  db.ContainerMetadata
		arg9  worker.ImageFetcherSpec
		arg10 worker.TaskProcessSpec
		arg11 chan runtime.Event
	}
	runTaskStepReturns struct {
		result1 worker.TaskResult
	}
	runTaskStepReturnsOnCall map[int]struct {
		result1 worker.TaskResult
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) CreateVolume(arg1 lager.Logger, arg2 worker.VolumeSpec, arg3 worker.WorkerSpec, arg4 db.VolumeType) (worker.Volume, error) {
	fake.createVolumeMutex.Lock()
	ret, specificReturn := fake.createVolumeReturnsOnCall[len(fake.createVolumeArgsForCall)]
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 worker.WorkerSpec
		arg4 db.VolumeType
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateVolume", []interface{}{arg1, arg2, arg3, arg4})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createVolumeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeClient) CreateVolumeCalls(stub func(lager.Logger, worker.VolumeSpec, worker.WorkerSpec, db.VolumeType) (worker.Volume, error)) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = stub
}

func (fake *FakeClient) CreateVolumeArgsForCall(i int) (lager.Logger, worker.VolumeSpec, worker.WorkerSpec, db.VolumeType) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	argsForCall := fake.createVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) CreateVolumeReturns(result1 worker.Volume, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateVolumeReturnsOnCall(i int, result1 worker.Volume, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	if fake.createVolumeReturnsOnCall == nil {
		fake.createVolumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 error
		})
	}
	fake.createVolumeReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindContainer(arg1 lager.Logger, arg2 int, arg3 string) (worker.Container, bool, error) {
	fake.findContainerMutex.Lock()
	ret, specificReturn := fake.findContainerReturnsOnCall[len(fake.findContainerArgsForCall)]
	fake.findContainerArgsForCall = append(fake.findContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("FindContainer", []interface{}{arg1, arg2, arg3})
	fake.findContainerMutex.Unlock()
	if fake.FindContainerStub != nil {
		return fake.FindContainerStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findContainerReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) FindContainerCallCount() int {
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	return len(fake.findContainerArgsForCall)
}

func (fake *FakeClient) FindContainerCalls(stub func(lager.Logger, int, string) (worker.Container, bool, error)) {
	fake.findContainerMutex.Lock()
	defer fake.findContainerMutex.Unlock()
	fake.FindContainerStub = stub
}

func (fake *FakeClient) FindContainerArgsForCall(i int) (lager.Logger, int, string) {
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	argsForCall := fake.findContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) FindContainerReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.findContainerMutex.Lock()
	defer fake.findContainerMutex.Unlock()
	fake.FindContainerStub = nil
	fake.findContainerReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindContainerReturnsOnCall(i int, result1 worker.Container, result2 bool, result3 error) {
	fake.findContainerMutex.Lock()
	defer fake.findContainerMutex.Unlock()
	fake.FindContainerStub = nil
	if fake.findContainerReturnsOnCall == nil {
		fake.findContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 bool
			result3 error
		})
	}
	fake.findContainerReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindVolume(arg1 lager.Logger, arg2 int, arg3 string) (worker.Volume, bool, error) {
	fake.findVolumeMutex.Lock()
	ret, specificReturn := fake.findVolumeReturnsOnCall[len(fake.findVolumeArgsForCall)]
	fake.findVolumeArgsForCall = append(fake.findVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("FindVolume", []interface{}{arg1, arg2, arg3})
	fake.findVolumeMutex.Unlock()
	if fake.FindVolumeStub != nil {
		return fake.FindVolumeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findVolumeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) FindVolumeCallCount() int {
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	return len(fake.findVolumeArgsForCall)
}

func (fake *FakeClient) FindVolumeCalls(stub func(lager.Logger, int, string) (worker.Volume, bool, error)) {
	fake.findVolumeMutex.Lock()
	defer fake.findVolumeMutex.Unlock()
	fake.FindVolumeStub = stub
}

func (fake *FakeClient) FindVolumeArgsForCall(i int) (lager.Logger, int, string) {
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	argsForCall := fake.findVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) FindVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.findVolumeMutex.Lock()
	defer fake.findVolumeMutex.Unlock()
	fake.FindVolumeStub = nil
	fake.findVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindVolumeReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.findVolumeMutex.Lock()
	defer fake.findVolumeMutex.Unlock()
	fake.FindVolumeStub = nil
	if fake.findVolumeReturnsOnCall == nil {
		fake.findVolumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findVolumeReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) RunTaskStep(arg1 context.Context, arg2 lager.Logger, arg3 lock.LockFactory, arg4 db.ContainerOwner, arg5 worker.ContainerSpec, arg6 worker.WorkerSpec, arg7 worker.ContainerPlacementStrategy, arg8 db.ContainerMetadata, arg9 worker.ImageFetcherSpec, arg10 worker.TaskProcessSpec, arg11 chan runtime.Event) worker.TaskResult {
	fake.runTaskStepMutex.Lock()
	ret, specificReturn := fake.runTaskStepReturnsOnCall[len(fake.runTaskStepArgsForCall)]
	fake.runTaskStepArgsForCall = append(fake.runTaskStepArgsForCall, struct {
		arg1  context.Context
		arg2  lager.Logger
		arg3  lock.LockFactory
		arg4  db.ContainerOwner
		arg5  worker.ContainerSpec
		arg6  worker.WorkerSpec
		arg7  worker.ContainerPlacementStrategy
		arg8  db.ContainerMetadata
		arg9  worker.ImageFetcherSpec
		arg10 worker.TaskProcessSpec
		arg11 chan runtime.Event
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11})
	fake.recordInvocation("RunTaskStep", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11})
	fake.runTaskStepMutex.Unlock()
	if fake.RunTaskStepStub != nil {
		return fake.RunTaskStepStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.runTaskStepReturns
	return fakeReturns.result1
}

func (fake *FakeClient) RunTaskStepCallCount() int {
	fake.runTaskStepMutex.RLock()
	defer fake.runTaskStepMutex.RUnlock()
	return len(fake.runTaskStepArgsForCall)
}

func (fake *FakeClient) RunTaskStepCalls(stub func(context.Context, lager.Logger, lock.LockFactory, db.ContainerOwner, worker.ContainerSpec, worker.WorkerSpec, worker.ContainerPlacementStrategy, db.ContainerMetadata, worker.ImageFetcherSpec, worker.TaskProcessSpec, chan runtime.Event) worker.TaskResult) {
	fake.runTaskStepMutex.Lock()
	defer fake.runTaskStepMutex.Unlock()
	fake.RunTaskStepStub = stub
}

func (fake *FakeClient) RunTaskStepArgsForCall(i int) (context.Context, lager.Logger, lock.LockFactory, db.ContainerOwner, worker.ContainerSpec, worker.WorkerSpec, worker.ContainerPlacementStrategy, db.ContainerMetadata, worker.ImageFetcherSpec, worker.TaskProcessSpec, chan runtime.Event) {
	fake.runTaskStepMutex.RLock()
	defer fake.runTaskStepMutex.RUnlock()
	argsForCall := fake.runTaskStepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9, argsForCall.arg10, argsForCall.arg11
}

func (fake *FakeClient) RunTaskStepReturns(result1 worker.TaskResult) {
	fake.runTaskStepMutex.Lock()
	defer fake.runTaskStepMutex.Unlock()
	fake.RunTaskStepStub = nil
	fake.runTaskStepReturns = struct {
		result1 worker.TaskResult
	}{result1}
}

func (fake *FakeClient) RunTaskStepReturnsOnCall(i int, result1 worker.TaskResult) {
	fake.runTaskStepMutex.Lock()
	defer fake.runTaskStepMutex.Unlock()
	fake.RunTaskStepStub = nil
	if fake.runTaskStepReturnsOnCall == nil {
		fake.runTaskStepReturnsOnCall = make(map[int]struct {
			result1 worker.TaskResult
		})
	}
	fake.runTaskStepReturnsOnCall[i] = struct {
		result1 worker.TaskResult
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	fake.runTaskStepMutex.RLock()
	defer fake.runTaskStepMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ worker.Client = new(FakeClient)
