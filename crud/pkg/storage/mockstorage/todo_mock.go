// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockstorage

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	crudapi "sadlil.com/samples/crud/apis/go/crudapi"
)

// TodoQuery is an autogenerated mock type for the TodoQuery type
type TodoQuery struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TodoQuery) GetByID(ctx context.Context, id string) (*crudapi.Todo, error) {
	ret := _m.Called(ctx, id)

	var r0 *crudapi.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*crudapi.Todo, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *crudapi.Todo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*crudapi.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTodoQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoQuery creates a new instance of TodoQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoQuery(t mockConstructorTestingTNewTodoQuery) *TodoQuery {
	mock := &TodoQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
