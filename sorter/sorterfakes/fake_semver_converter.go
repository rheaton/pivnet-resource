// This file was generated by counterfeiter
package sorterfakes

import (
	"sync"

	"github.com/blang/semver"
	"github.com/pivotal-cf-experimental/pivnet-resource/sorter"
)

type FakeSemverConverter struct {
	ToValidSemverStub        func(string) (semver.Version, error)
	toValidSemverMutex       sync.RWMutex
	toValidSemverArgsForCall []struct {
		arg1 string
	}
	toValidSemverReturns struct {
		result1 semver.Version
		result2 error
	}
	invocations map[string][][]interface{}
}

func (fake *FakeSemverConverter) ToValidSemver(arg1 string) (semver.Version, error) {
	fake.toValidSemverMutex.Lock()
	fake.toValidSemverArgsForCall = append(fake.toValidSemverArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.guard("ToValidSemver")
	fake.invocations["ToValidSemver"] = append(fake.invocations["ToValidSemver"], []interface{}{arg1})
	fake.toValidSemverMutex.Unlock()
	if fake.ToValidSemverStub != nil {
		return fake.ToValidSemverStub(arg1)
	} else {
		return fake.toValidSemverReturns.result1, fake.toValidSemverReturns.result2
	}
}

func (fake *FakeSemverConverter) ToValidSemverCallCount() int {
	fake.toValidSemverMutex.RLock()
	defer fake.toValidSemverMutex.RUnlock()
	return len(fake.toValidSemverArgsForCall)
}

func (fake *FakeSemverConverter) ToValidSemverArgsForCall(i int) string {
	fake.toValidSemverMutex.RLock()
	defer fake.toValidSemverMutex.RUnlock()
	return fake.toValidSemverArgsForCall[i].arg1
}

func (fake *FakeSemverConverter) ToValidSemverReturns(result1 semver.Version, result2 error) {
	fake.ToValidSemverStub = nil
	fake.toValidSemverReturns = struct {
		result1 semver.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeSemverConverter) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeSemverConverter) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ sorter.SemverConverter = new(FakeSemverConverter)