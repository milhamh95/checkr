// Code generated by counterfeiter. DO NOT EDIT.
package counterfeiter

import (
	"sync"

	"github.com/milhamh95/checkr/domain"
)

type FakePromoStorageSource struct {
	GetPromoStub        func(string) (domain.Promo, error)
	getPromoMutex       sync.RWMutex
	getPromoArgsForCall []struct {
		arg1 string
	}
	getPromoReturns struct {
		result1 domain.Promo
		result2 error
	}
	getPromoReturnsOnCall map[int]struct {
		result1 domain.Promo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePromoStorageSource) GetPromo(arg1 string) (domain.Promo, error) {
	fake.getPromoMutex.Lock()
	ret, specificReturn := fake.getPromoReturnsOnCall[len(fake.getPromoArgsForCall)]
	fake.getPromoArgsForCall = append(fake.getPromoArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetPromoStub
	fakeReturns := fake.getPromoReturns
	fake.recordInvocation("GetPromo", []interface{}{arg1})
	fake.getPromoMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePromoStorageSource) GetPromoCallCount() int {
	fake.getPromoMutex.RLock()
	defer fake.getPromoMutex.RUnlock()
	return len(fake.getPromoArgsForCall)
}

func (fake *FakePromoStorageSource) GetPromoCalls(stub func(string) (domain.Promo, error)) {
	fake.getPromoMutex.Lock()
	defer fake.getPromoMutex.Unlock()
	fake.GetPromoStub = stub
}

func (fake *FakePromoStorageSource) GetPromoArgsForCall(i int) string {
	fake.getPromoMutex.RLock()
	defer fake.getPromoMutex.RUnlock()
	argsForCall := fake.getPromoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePromoStorageSource) GetPromoReturns(result1 domain.Promo, result2 error) {
	fake.getPromoMutex.Lock()
	defer fake.getPromoMutex.Unlock()
	fake.GetPromoStub = nil
	fake.getPromoReturns = struct {
		result1 domain.Promo
		result2 error
	}{result1, result2}
}

func (fake *FakePromoStorageSource) GetPromoReturnsOnCall(i int, result1 domain.Promo, result2 error) {
	fake.getPromoMutex.Lock()
	defer fake.getPromoMutex.Unlock()
	fake.GetPromoStub = nil
	if fake.getPromoReturnsOnCall == nil {
		fake.getPromoReturnsOnCall = make(map[int]struct {
			result1 domain.Promo
			result2 error
		})
	}
	fake.getPromoReturnsOnCall[i] = struct {
		result1 domain.Promo
		result2 error
	}{result1, result2}
}

func (fake *FakePromoStorageSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPromoMutex.RLock()
	defer fake.getPromoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePromoStorageSource) recordInvocation(key string, args []interface{}) {
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