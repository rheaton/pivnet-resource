// This file was generated by counterfeiter
package uploaderfakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/pivnet-resource/uploader"
)

type FakeTransport struct {
	UploadStub        func(fileGlob string, filepathPrefix string, sourcesDir string) error
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		fileGlob       string
		filepathPrefix string
		sourcesDir     string
	}
	uploadReturns struct {
		result1 error
	}
	invocations map[string][][]interface{}
}

func (fake *FakeTransport) Upload(fileGlob string, filepathPrefix string, sourcesDir string) error {
	fake.uploadMutex.Lock()
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		fileGlob       string
		filepathPrefix string
		sourcesDir     string
	}{fileGlob, filepathPrefix, sourcesDir})
	fake.guard("Upload")
	fake.invocations["Upload"] = append(fake.invocations["Upload"], []interface{}{fileGlob, filepathPrefix, sourcesDir})
	fake.uploadMutex.Unlock()
	if fake.UploadStub != nil {
		return fake.UploadStub(fileGlob, filepathPrefix, sourcesDir)
	} else {
		return fake.uploadReturns.result1
	}
}

func (fake *FakeTransport) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *FakeTransport) UploadArgsForCall(i int) (string, string, string) {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return fake.uploadArgsForCall[i].fileGlob, fake.uploadArgsForCall[i].filepathPrefix, fake.uploadArgsForCall[i].sourcesDir
}

func (fake *FakeTransport) UploadReturns(result1 error) {
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTransport) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeTransport) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ uploader.Transport = new(FakeTransport)