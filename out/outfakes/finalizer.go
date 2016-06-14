// This file was generated by counterfeiter
package outfakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/pivnet-resource/concourse"
	"github.com/pivotal-cf-experimental/pivnet-resource/pivnet"
)

type Finalizer struct {
	FinalizeStub        func(release pivnet.Release) (concourse.OutResponse, error)
	finalizeMutex       sync.RWMutex
	finalizeArgsForCall []struct {
		release pivnet.Release
	}
	finalizeReturns struct {
		result1 concourse.OutResponse
		result2 error
	}
}

func (fake *Finalizer) Finalize(release pivnet.Release) (concourse.OutResponse, error) {
	fake.finalizeMutex.Lock()
	fake.finalizeArgsForCall = append(fake.finalizeArgsForCall, struct {
		release pivnet.Release
	}{release})
	fake.finalizeMutex.Unlock()
	if fake.FinalizeStub != nil {
		return fake.FinalizeStub(release)
	} else {
		return fake.finalizeReturns.result1, fake.finalizeReturns.result2
	}
}

func (fake *Finalizer) FinalizeCallCount() int {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return len(fake.finalizeArgsForCall)
}

func (fake *Finalizer) FinalizeArgsForCall(i int) pivnet.Release {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return fake.finalizeArgsForCall[i].release
}

func (fake *Finalizer) FinalizeReturns(result1 concourse.OutResponse, result2 error) {
	fake.FinalizeStub = nil
	fake.finalizeReturns = struct {
		result1 concourse.OutResponse
		result2 error
	}{result1, result2}
}
