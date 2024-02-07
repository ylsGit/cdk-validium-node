// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
)

// ZKEVMClientGlobalExitRootGetter is an autogenerated mock type for the ZKEVMClientGlobalExitRootGetter type
type ZKEVMClientGlobalExitRootGetter struct {
	mock.Mock
}

type ZKEVMClientGlobalExitRootGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ZKEVMClientGlobalExitRootGetter) EXPECT() *ZKEVMClientGlobalExitRootGetter_Expecter {
	return &ZKEVMClientGlobalExitRootGetter_Expecter{mock: &_m.Mock}
}

// ExitRootsByGER provides a mock function with given fields: ctx, globalExitRoot
func (_m *ZKEVMClientGlobalExitRootGetter) ExitRootsByGER(ctx context.Context, globalExitRoot common.Hash) (*types.ExitRoots, error) {
	ret := _m.Called(ctx, globalExitRoot)

	var r0 *types.ExitRoots
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.ExitRoots, error)); ok {
		return rf(ctx, globalExitRoot)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.ExitRoots); ok {
		r0 = rf(ctx, globalExitRoot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ExitRoots)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, globalExitRoot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExitRootsByGER'
type ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call struct {
	*mock.Call
}

// ExitRootsByGER is a helper method to define mock.On call
//   - ctx context.Context
//   - globalExitRoot common.Hash
func (_e *ZKEVMClientGlobalExitRootGetter_Expecter) ExitRootsByGER(ctx interface{}, globalExitRoot interface{}) *ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call {
	return &ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call{Call: _e.mock.On("ExitRootsByGER", ctx, globalExitRoot)}
}

func (_c *ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call) Run(run func(ctx context.Context, globalExitRoot common.Hash)) *ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call) Return(_a0 *types.ExitRoots, _a1 error) *ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.ExitRoots, error)) *ZKEVMClientGlobalExitRootGetter_ExitRootsByGER_Call {
	_c.Call.Return(run)
	return _c
}

// NewZKEVMClientGlobalExitRootGetter creates a new instance of ZKEVMClientGlobalExitRootGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewZKEVMClientGlobalExitRootGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ZKEVMClientGlobalExitRootGetter {
	mock := &ZKEVMClientGlobalExitRootGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}