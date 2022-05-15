// Code generated by counterfeiter. DO NOT EDIT.
package counterfeiter

import (
	"sync"

	"github.com/milhamh95/checkr/domain"
)

type FakeBundlingPromoCartStorage struct {
	AddFreeItemStub        func(domain.CartItem)
	addFreeItemMutex       sync.RWMutex
	addFreeItemArgsForCall []struct {
		arg1 domain.CartItem
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBundlingPromoCartStorage) AddFreeItem(arg1 domain.CartItem) {
	fake.addFreeItemMutex.Lock()
	fake.addFreeItemArgsForCall = append(fake.addFreeItemArgsForCall, struct {
		arg1 domain.CartItem
	}{arg1})
	stub := fake.AddFreeItemStub
	fake.recordInvocation("AddFreeItem", []interface{}{arg1})
	fake.addFreeItemMutex.Unlock()
	if stub != nil {
		fake.AddFreeItemStub(arg1)
	}
}

func (fake *FakeBundlingPromoCartStorage) AddFreeItemCallCount() int {
	fake.addFreeItemMutex.RLock()
	defer fake.addFreeItemMutex.RUnlock()
	return len(fake.addFreeItemArgsForCall)
}

func (fake *FakeBundlingPromoCartStorage) AddFreeItemCalls(stub func(domain.CartItem)) {
	fake.addFreeItemMutex.Lock()
	defer fake.addFreeItemMutex.Unlock()
	fake.AddFreeItemStub = stub
}

func (fake *FakeBundlingPromoCartStorage) AddFreeItemArgsForCall(i int) domain.CartItem {
	fake.addFreeItemMutex.RLock()
	defer fake.addFreeItemMutex.RUnlock()
	argsForCall := fake.addFreeItemArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBundlingPromoCartStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addFreeItemMutex.RLock()
	defer fake.addFreeItemMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBundlingPromoCartStorage) recordInvocation(key string, args []interface{}) {
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