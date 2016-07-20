// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
	"github.com/pivotal-golang/lager"
)

type FakeFetchContainerCreatorFactory struct {
	NewFetchContainerCreatorStub        func(logger lager.Logger, resourceTypes atc.ResourceTypes, tags atc.Tags, session resource.Session, metadata resource.Metadata, imageFetchingDelegate worker.ImageFetchingDelegate) resource.FetchContainerCreator
	newFetchContainerCreatorMutex       sync.RWMutex
	newFetchContainerCreatorArgsForCall []struct {
		logger                lager.Logger
		resourceTypes         atc.ResourceTypes
		tags                  atc.Tags
		session               resource.Session
		metadata              resource.Metadata
		imageFetchingDelegate worker.ImageFetchingDelegate
	}
	newFetchContainerCreatorReturns struct {
		result1 resource.FetchContainerCreator
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetchContainerCreatorFactory) NewFetchContainerCreator(logger lager.Logger, resourceTypes atc.ResourceTypes, tags atc.Tags, session resource.Session, metadata resource.Metadata, imageFetchingDelegate worker.ImageFetchingDelegate) resource.FetchContainerCreator {
	fake.newFetchContainerCreatorMutex.Lock()
	fake.newFetchContainerCreatorArgsForCall = append(fake.newFetchContainerCreatorArgsForCall, struct {
		logger                lager.Logger
		resourceTypes         atc.ResourceTypes
		tags                  atc.Tags
		session               resource.Session
		metadata              resource.Metadata
		imageFetchingDelegate worker.ImageFetchingDelegate
	}{logger, resourceTypes, tags, session, metadata, imageFetchingDelegate})
	fake.recordInvocation("NewFetchContainerCreator", []interface{}{logger, resourceTypes, tags, session, metadata, imageFetchingDelegate})
	fake.newFetchContainerCreatorMutex.Unlock()
	if fake.NewFetchContainerCreatorStub != nil {
		return fake.NewFetchContainerCreatorStub(logger, resourceTypes, tags, session, metadata, imageFetchingDelegate)
	} else {
		return fake.newFetchContainerCreatorReturns.result1
	}
}

func (fake *FakeFetchContainerCreatorFactory) NewFetchContainerCreatorCallCount() int {
	fake.newFetchContainerCreatorMutex.RLock()
	defer fake.newFetchContainerCreatorMutex.RUnlock()
	return len(fake.newFetchContainerCreatorArgsForCall)
}

func (fake *FakeFetchContainerCreatorFactory) NewFetchContainerCreatorArgsForCall(i int) (lager.Logger, atc.ResourceTypes, atc.Tags, resource.Session, resource.Metadata, worker.ImageFetchingDelegate) {
	fake.newFetchContainerCreatorMutex.RLock()
	defer fake.newFetchContainerCreatorMutex.RUnlock()
	return fake.newFetchContainerCreatorArgsForCall[i].logger, fake.newFetchContainerCreatorArgsForCall[i].resourceTypes, fake.newFetchContainerCreatorArgsForCall[i].tags, fake.newFetchContainerCreatorArgsForCall[i].session, fake.newFetchContainerCreatorArgsForCall[i].metadata, fake.newFetchContainerCreatorArgsForCall[i].imageFetchingDelegate
}

func (fake *FakeFetchContainerCreatorFactory) NewFetchContainerCreatorReturns(result1 resource.FetchContainerCreator) {
	fake.NewFetchContainerCreatorStub = nil
	fake.newFetchContainerCreatorReturns = struct {
		result1 resource.FetchContainerCreator
	}{result1}
}

func (fake *FakeFetchContainerCreatorFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newFetchContainerCreatorMutex.RLock()
	defer fake.newFetchContainerCreatorMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFetchContainerCreatorFactory) recordInvocation(key string, args []interface{}) {
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

var _ resource.FetchContainerCreatorFactory = new(FakeFetchContainerCreatorFactory)
