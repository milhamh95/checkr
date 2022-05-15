// Code generated by counterfeiter. DO NOT EDIT.
package counterfeiter

import (
	"sync"

	"github.com/milhamh95/checkr/domain"
)

type FakeCartProductStorage struct {
	GetProductStub        func(string) (domain.Product, error)
	getProductMutex       sync.RWMutex
	getProductArgsForCall []struct {
		arg1 string
	}
	getProductReturns struct {
		result1 domain.Product
		result2 error
	}
	getProductReturnsOnCall map[int]struct {
		result1 domain.Product
		result2 error
	}
	ReduceInventoryQuantityStub        func(string, int) error
	reduceInventoryQuantityMutex       sync.RWMutex
	reduceInventoryQuantityArgsForCall []struct {
		arg1 string
		arg2 int
	}
	reduceInventoryQuantityReturns struct {
		result1 error
	}
	reduceInventoryQuantityReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCartProductStorage) GetProduct(arg1 string) (domain.Product, error) {
	fake.getProductMutex.Lock()
	ret, specificReturn := fake.getProductReturnsOnCall[len(fake.getProductArgsForCall)]
	fake.getProductArgsForCall = append(fake.getProductArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetProductStub
	fakeReturns := fake.getProductReturns
	fake.recordInvocation("GetProduct", []interface{}{arg1})
	fake.getProductMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCartProductStorage) GetProductCallCount() int {
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	return len(fake.getProductArgsForCall)
}

func (fake *FakeCartProductStorage) GetProductCalls(stub func(string) (domain.Product, error)) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = stub
}

func (fake *FakeCartProductStorage) GetProductArgsForCall(i int) string {
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	argsForCall := fake.getProductArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCartProductStorage) GetProductReturns(result1 domain.Product, result2 error) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = nil
	fake.getProductReturns = struct {
		result1 domain.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeCartProductStorage) GetProductReturnsOnCall(i int, result1 domain.Product, result2 error) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = nil
	if fake.getProductReturnsOnCall == nil {
		fake.getProductReturnsOnCall = make(map[int]struct {
			result1 domain.Product
			result2 error
		})
	}
	fake.getProductReturnsOnCall[i] = struct {
		result1 domain.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeCartProductStorage) ReduceInventoryQuantity(arg1 string, arg2 int) error {
	fake.reduceInventoryQuantityMutex.Lock()
	ret, specificReturn := fake.reduceInventoryQuantityReturnsOnCall[len(fake.reduceInventoryQuantityArgsForCall)]
	fake.reduceInventoryQuantityArgsForCall = append(fake.reduceInventoryQuantityArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	stub := fake.ReduceInventoryQuantityStub
	fakeReturns := fake.reduceInventoryQuantityReturns
	fake.recordInvocation("ReduceInventoryQuantity", []interface{}{arg1, arg2})
	fake.reduceInventoryQuantityMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCartProductStorage) ReduceInventoryQuantityCallCount() int {
	fake.reduceInventoryQuantityMutex.RLock()
	defer fake.reduceInventoryQuantityMutex.RUnlock()
	return len(fake.reduceInventoryQuantityArgsForCall)
}

func (fake *FakeCartProductStorage) ReduceInventoryQuantityCalls(stub func(string, int) error) {
	fake.reduceInventoryQuantityMutex.Lock()
	defer fake.reduceInventoryQuantityMutex.Unlock()
	fake.ReduceInventoryQuantityStub = stub
}

func (fake *FakeCartProductStorage) ReduceInventoryQuantityArgsForCall(i int) (string, int) {
	fake.reduceInventoryQuantityMutex.RLock()
	defer fake.reduceInventoryQuantityMutex.RUnlock()
	argsForCall := fake.reduceInventoryQuantityArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCartProductStorage) ReduceInventoryQuantityReturns(result1 error) {
	fake.reduceInventoryQuantityMutex.Lock()
	defer fake.reduceInventoryQuantityMutex.Unlock()
	fake.ReduceInventoryQuantityStub = nil
	fake.reduceInventoryQuantityReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCartProductStorage) ReduceInventoryQuantityReturnsOnCall(i int, result1 error) {
	fake.reduceInventoryQuantityMutex.Lock()
	defer fake.reduceInventoryQuantityMutex.Unlock()
	fake.ReduceInventoryQuantityStub = nil
	if fake.reduceInventoryQuantityReturnsOnCall == nil {
		fake.reduceInventoryQuantityReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reduceInventoryQuantityReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCartProductStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	fake.reduceInventoryQuantityMutex.RLock()
	defer fake.reduceInventoryQuantityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCartProductStorage) recordInvocation(key string, args []interface{}) {
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
