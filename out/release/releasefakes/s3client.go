// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"
)

type S3Client struct {
	ComputeAWSObjectKeyStub        func(string) (string, string, error)
	computeAWSObjectKeyMutex       sync.RWMutex
	computeAWSObjectKeyArgsForCall []struct {
		arg1 string
	}
	computeAWSObjectKeyReturns struct {
		result1 string
		result2 string
		result3 error
	}
	computeAWSObjectKeyReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	UploadFileStub        func(string) error
	uploadFileMutex       sync.RWMutex
	uploadFileArgsForCall []struct {
		arg1 string
	}
	uploadFileReturns struct {
		result1 error
	}
	uploadFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *S3Client) ComputeAWSObjectKey(arg1 string) (string, string, error) {
	fake.computeAWSObjectKeyMutex.Lock()
	ret, specificReturn := fake.computeAWSObjectKeyReturnsOnCall[len(fake.computeAWSObjectKeyArgsForCall)]
	fake.computeAWSObjectKeyArgsForCall = append(fake.computeAWSObjectKeyArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ComputeAWSObjectKey", []interface{}{arg1})
	fake.computeAWSObjectKeyMutex.Unlock()
	if fake.ComputeAWSObjectKeyStub != nil {
		return fake.ComputeAWSObjectKeyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.computeAWSObjectKeyReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *S3Client) ComputeAWSObjectKeyCallCount() int {
	fake.computeAWSObjectKeyMutex.RLock()
	defer fake.computeAWSObjectKeyMutex.RUnlock()
	return len(fake.computeAWSObjectKeyArgsForCall)
}

func (fake *S3Client) ComputeAWSObjectKeyCalls(stub func(string) (string, string, error)) {
	fake.computeAWSObjectKeyMutex.Lock()
	defer fake.computeAWSObjectKeyMutex.Unlock()
	fake.ComputeAWSObjectKeyStub = stub
}

func (fake *S3Client) ComputeAWSObjectKeyArgsForCall(i int) string {
	fake.computeAWSObjectKeyMutex.RLock()
	defer fake.computeAWSObjectKeyMutex.RUnlock()
	argsForCall := fake.computeAWSObjectKeyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *S3Client) ComputeAWSObjectKeyReturns(result1 string, result2 string, result3 error) {
	fake.computeAWSObjectKeyMutex.Lock()
	defer fake.computeAWSObjectKeyMutex.Unlock()
	fake.ComputeAWSObjectKeyStub = nil
	fake.computeAWSObjectKeyReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *S3Client) ComputeAWSObjectKeyReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.computeAWSObjectKeyMutex.Lock()
	defer fake.computeAWSObjectKeyMutex.Unlock()
	fake.ComputeAWSObjectKeyStub = nil
	if fake.computeAWSObjectKeyReturnsOnCall == nil {
		fake.computeAWSObjectKeyReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.computeAWSObjectKeyReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *S3Client) UploadFile(arg1 string) error {
	fake.uploadFileMutex.Lock()
	ret, specificReturn := fake.uploadFileReturnsOnCall[len(fake.uploadFileArgsForCall)]
	fake.uploadFileArgsForCall = append(fake.uploadFileArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UploadFile", []interface{}{arg1})
	fake.uploadFileMutex.Unlock()
	if fake.UploadFileStub != nil {
		return fake.UploadFileStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uploadFileReturns
	return fakeReturns.result1
}

func (fake *S3Client) UploadFileCallCount() int {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return len(fake.uploadFileArgsForCall)
}

func (fake *S3Client) UploadFileCalls(stub func(string) error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = stub
}

func (fake *S3Client) UploadFileArgsForCall(i int) string {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	argsForCall := fake.uploadFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *S3Client) UploadFileReturns(result1 error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = nil
	fake.uploadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *S3Client) UploadFileReturnsOnCall(i int, result1 error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = nil
	if fake.uploadFileReturnsOnCall == nil {
		fake.uploadFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *S3Client) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.computeAWSObjectKeyMutex.RLock()
	defer fake.computeAWSObjectKeyMutex.RUnlock()
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *S3Client) recordInvocation(key string, args []interface{}) {
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
