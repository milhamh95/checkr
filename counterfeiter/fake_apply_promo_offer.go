// Code generated by counterfeiter. DO NOT EDIT.
package counterfeiter

import (
	"sync"

	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/service"
)

type FakeApplyPromoOffer struct {
	ApplyStub        func(domain.CartItem, domain.Product, domain.Promo) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 domain.CartItem
		arg2 domain.Product
		arg3 domain.Promo
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApplyPromoOffer) Apply(arg1 domain.CartItem, arg2 domain.Product, arg3 domain.Promo) error {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 domain.CartItem
		arg2 domain.Product
		arg3 domain.Promo
	}{arg1, arg2, arg3})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1, arg2, arg3})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApplyPromoOffer) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeApplyPromoOffer) ApplyCalls(stub func(domain.CartItem, domain.Product, domain.Promo) error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeApplyPromoOffer) ApplyArgsForCall(i int) (domain.CartItem, domain.Product, domain.Promo) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeApplyPromoOffer) ApplyReturns(result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApplyPromoOffer) ApplyReturnsOnCall(i int, result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApplyPromoOffer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeApplyPromoOffer) recordInvocation(key string, args []interface{}) {
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

var _ service.ApplyPromoOffer = new(FakeApplyPromoOffer)
