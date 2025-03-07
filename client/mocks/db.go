// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	blockstore "github.com/ipfs/boxo/blockstore"
	client "github.com/sourcenetwork/defradb/client"

	context "context"

	datastore "github.com/sourcenetwork/defradb/datastore"

	events "github.com/sourcenetwork/defradb/events"

	mock "github.com/stretchr/testify/mock"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

type DB_Expecter struct {
	mock *mock.Mock
}

func (_m *DB) EXPECT() *DB_Expecter {
	return &DB_Expecter{mock: &_m.Mock}
}

// AddP2PCollection provides a mock function with given fields: ctx, collectionID
func (_m *DB) AddP2PCollection(ctx context.Context, collectionID string) error {
	ret := _m.Called(ctx, collectionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, collectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_AddP2PCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddP2PCollection'
type DB_AddP2PCollection_Call struct {
	*mock.Call
}

// AddP2PCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID string
func (_e *DB_Expecter) AddP2PCollection(ctx interface{}, collectionID interface{}) *DB_AddP2PCollection_Call {
	return &DB_AddP2PCollection_Call{Call: _e.mock.On("AddP2PCollection", ctx, collectionID)}
}

func (_c *DB_AddP2PCollection_Call) Run(run func(ctx context.Context, collectionID string)) *DB_AddP2PCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_AddP2PCollection_Call) Return(_a0 error) *DB_AddP2PCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_AddP2PCollection_Call) RunAndReturn(run func(context.Context, string) error) *DB_AddP2PCollection_Call {
	_c.Call.Return(run)
	return _c
}

// AddSchema provides a mock function with given fields: _a0, _a1
func (_m *DB) AddSchema(_a0 context.Context, _a1 string) ([]client.CollectionDescription, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []client.CollectionDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]client.CollectionDescription, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []client.CollectionDescription); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.CollectionDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_AddSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSchema'
type DB_AddSchema_Call struct {
	*mock.Call
}

// AddSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) AddSchema(_a0 interface{}, _a1 interface{}) *DB_AddSchema_Call {
	return &DB_AddSchema_Call{Call: _e.mock.On("AddSchema", _a0, _a1)}
}

func (_c *DB_AddSchema_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_AddSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_AddSchema_Call) Return(_a0 []client.CollectionDescription, _a1 error) *DB_AddSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_AddSchema_Call) RunAndReturn(run func(context.Context, string) ([]client.CollectionDescription, error)) *DB_AddSchema_Call {
	_c.Call.Return(run)
	return _c
}

// BasicExport provides a mock function with given fields: ctx, config
func (_m *DB) BasicExport(ctx context.Context, config *client.BackupConfig) error {
	ret := _m.Called(ctx, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.BackupConfig) error); ok {
		r0 = rf(ctx, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_BasicExport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BasicExport'
type DB_BasicExport_Call struct {
	*mock.Call
}

// BasicExport is a helper method to define mock.On call
//   - ctx context.Context
//   - config *client.BackupConfig
func (_e *DB_Expecter) BasicExport(ctx interface{}, config interface{}) *DB_BasicExport_Call {
	return &DB_BasicExport_Call{Call: _e.mock.On("BasicExport", ctx, config)}
}

func (_c *DB_BasicExport_Call) Run(run func(ctx context.Context, config *client.BackupConfig)) *DB_BasicExport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*client.BackupConfig))
	})
	return _c
}

func (_c *DB_BasicExport_Call) Return(_a0 error) *DB_BasicExport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_BasicExport_Call) RunAndReturn(run func(context.Context, *client.BackupConfig) error) *DB_BasicExport_Call {
	_c.Call.Return(run)
	return _c
}

// BasicImport provides a mock function with given fields: ctx, filepath
func (_m *DB) BasicImport(ctx context.Context, filepath string) error {
	ret := _m.Called(ctx, filepath)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, filepath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_BasicImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BasicImport'
type DB_BasicImport_Call struct {
	*mock.Call
}

// BasicImport is a helper method to define mock.On call
//   - ctx context.Context
//   - filepath string
func (_e *DB_Expecter) BasicImport(ctx interface{}, filepath interface{}) *DB_BasicImport_Call {
	return &DB_BasicImport_Call{Call: _e.mock.On("BasicImport", ctx, filepath)}
}

func (_c *DB_BasicImport_Call) Run(run func(ctx context.Context, filepath string)) *DB_BasicImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_BasicImport_Call) Return(_a0 error) *DB_BasicImport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_BasicImport_Call) RunAndReturn(run func(context.Context, string) error) *DB_BasicImport_Call {
	_c.Call.Return(run)
	return _c
}

// Blockstore provides a mock function with given fields:
func (_m *DB) Blockstore() blockstore.Blockstore {
	ret := _m.Called()

	var r0 blockstore.Blockstore
	if rf, ok := ret.Get(0).(func() blockstore.Blockstore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blockstore.Blockstore)
		}
	}

	return r0
}

// DB_Blockstore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Blockstore'
type DB_Blockstore_Call struct {
	*mock.Call
}

// Blockstore is a helper method to define mock.On call
func (_e *DB_Expecter) Blockstore() *DB_Blockstore_Call {
	return &DB_Blockstore_Call{Call: _e.mock.On("Blockstore")}
}

func (_c *DB_Blockstore_Call) Run(run func()) *DB_Blockstore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Blockstore_Call) Return(_a0 blockstore.Blockstore) *DB_Blockstore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Blockstore_Call) RunAndReturn(run func() blockstore.Blockstore) *DB_Blockstore_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: _a0
func (_m *DB) Close(_a0 context.Context) {
	_m.Called(_a0)
}

// DB_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type DB_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *DB_Expecter) Close(_a0 interface{}) *DB_Close_Call {
	return &DB_Close_Call{Call: _e.mock.On("Close", _a0)}
}

func (_c *DB_Close_Call) Run(run func(_a0 context.Context)) *DB_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_Close_Call) Return() *DB_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *DB_Close_Call) RunAndReturn(run func(context.Context)) *DB_Close_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteReplicator provides a mock function with given fields: ctx, rep
func (_m *DB) DeleteReplicator(ctx context.Context, rep client.Replicator) error {
	ret := _m.Called(ctx, rep)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Replicator) error); ok {
		r0 = rf(ctx, rep)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_DeleteReplicator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteReplicator'
type DB_DeleteReplicator_Call struct {
	*mock.Call
}

// DeleteReplicator is a helper method to define mock.On call
//   - ctx context.Context
//   - rep client.Replicator
func (_e *DB_Expecter) DeleteReplicator(ctx interface{}, rep interface{}) *DB_DeleteReplicator_Call {
	return &DB_DeleteReplicator_Call{Call: _e.mock.On("DeleteReplicator", ctx, rep)}
}

func (_c *DB_DeleteReplicator_Call) Run(run func(ctx context.Context, rep client.Replicator)) *DB_DeleteReplicator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.Replicator))
	})
	return _c
}

func (_c *DB_DeleteReplicator_Call) Return(_a0 error) *DB_DeleteReplicator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_DeleteReplicator_Call) RunAndReturn(run func(context.Context, client.Replicator) error) *DB_DeleteReplicator_Call {
	_c.Call.Return(run)
	return _c
}

// Events provides a mock function with given fields:
func (_m *DB) Events() events.Events {
	ret := _m.Called()

	var r0 events.Events
	if rf, ok := ret.Get(0).(func() events.Events); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(events.Events)
	}

	return r0
}

// DB_Events_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Events'
type DB_Events_Call struct {
	*mock.Call
}

// Events is a helper method to define mock.On call
func (_e *DB_Expecter) Events() *DB_Events_Call {
	return &DB_Events_Call{Call: _e.mock.On("Events")}
}

func (_c *DB_Events_Call) Run(run func()) *DB_Events_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Events_Call) Return(_a0 events.Events) *DB_Events_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Events_Call) RunAndReturn(run func() events.Events) *DB_Events_Call {
	_c.Call.Return(run)
	return _c
}

// ExecRequest provides a mock function with given fields: _a0, _a1
func (_m *DB) ExecRequest(_a0 context.Context, _a1 string) *client.RequestResult {
	ret := _m.Called(_a0, _a1)

	var r0 *client.RequestResult
	if rf, ok := ret.Get(0).(func(context.Context, string) *client.RequestResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.RequestResult)
		}
	}

	return r0
}

// DB_ExecRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecRequest'
type DB_ExecRequest_Call struct {
	*mock.Call
}

// ExecRequest is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) ExecRequest(_a0 interface{}, _a1 interface{}) *DB_ExecRequest_Call {
	return &DB_ExecRequest_Call{Call: _e.mock.On("ExecRequest", _a0, _a1)}
}

func (_c *DB_ExecRequest_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_ExecRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_ExecRequest_Call) Return(_a0 *client.RequestResult) *DB_ExecRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_ExecRequest_Call) RunAndReturn(run func(context.Context, string) *client.RequestResult) *DB_ExecRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllCollections provides a mock function with given fields: _a0
func (_m *DB) GetAllCollections(_a0 context.Context) ([]client.Collection, error) {
	ret := _m.Called(_a0)

	var r0 []client.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]client.Collection, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []client.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetAllCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllCollections'
type DB_GetAllCollections_Call struct {
	*mock.Call
}

// GetAllCollections is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *DB_Expecter) GetAllCollections(_a0 interface{}) *DB_GetAllCollections_Call {
	return &DB_GetAllCollections_Call{Call: _e.mock.On("GetAllCollections", _a0)}
}

func (_c *DB_GetAllCollections_Call) Run(run func(_a0 context.Context)) *DB_GetAllCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_GetAllCollections_Call) Return(_a0 []client.Collection, _a1 error) *DB_GetAllCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetAllCollections_Call) RunAndReturn(run func(context.Context) ([]client.Collection, error)) *DB_GetAllCollections_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllIndexes provides a mock function with given fields: _a0
func (_m *DB) GetAllIndexes(_a0 context.Context) (map[string][]client.IndexDescription, error) {
	ret := _m.Called(_a0)

	var r0 map[string][]client.IndexDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string][]client.IndexDescription, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string][]client.IndexDescription); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]client.IndexDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetAllIndexes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllIndexes'
type DB_GetAllIndexes_Call struct {
	*mock.Call
}

// GetAllIndexes is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *DB_Expecter) GetAllIndexes(_a0 interface{}) *DB_GetAllIndexes_Call {
	return &DB_GetAllIndexes_Call{Call: _e.mock.On("GetAllIndexes", _a0)}
}

func (_c *DB_GetAllIndexes_Call) Run(run func(_a0 context.Context)) *DB_GetAllIndexes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_GetAllIndexes_Call) Return(_a0 map[string][]client.IndexDescription, _a1 error) *DB_GetAllIndexes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetAllIndexes_Call) RunAndReturn(run func(context.Context) (map[string][]client.IndexDescription, error)) *DB_GetAllIndexes_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllP2PCollections provides a mock function with given fields: ctx
func (_m *DB) GetAllP2PCollections(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetAllP2PCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllP2PCollections'
type DB_GetAllP2PCollections_Call struct {
	*mock.Call
}

// GetAllP2PCollections is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DB_Expecter) GetAllP2PCollections(ctx interface{}) *DB_GetAllP2PCollections_Call {
	return &DB_GetAllP2PCollections_Call{Call: _e.mock.On("GetAllP2PCollections", ctx)}
}

func (_c *DB_GetAllP2PCollections_Call) Run(run func(ctx context.Context)) *DB_GetAllP2PCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_GetAllP2PCollections_Call) Return(_a0 []string, _a1 error) *DB_GetAllP2PCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetAllP2PCollections_Call) RunAndReturn(run func(context.Context) ([]string, error)) *DB_GetAllP2PCollections_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllReplicators provides a mock function with given fields: ctx
func (_m *DB) GetAllReplicators(ctx context.Context) ([]client.Replicator, error) {
	ret := _m.Called(ctx)

	var r0 []client.Replicator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]client.Replicator, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []client.Replicator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Replicator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetAllReplicators_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllReplicators'
type DB_GetAllReplicators_Call struct {
	*mock.Call
}

// GetAllReplicators is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DB_Expecter) GetAllReplicators(ctx interface{}) *DB_GetAllReplicators_Call {
	return &DB_GetAllReplicators_Call{Call: _e.mock.On("GetAllReplicators", ctx)}
}

func (_c *DB_GetAllReplicators_Call) Run(run func(ctx context.Context)) *DB_GetAllReplicators_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_GetAllReplicators_Call) Return(_a0 []client.Replicator, _a1 error) *DB_GetAllReplicators_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetAllReplicators_Call) RunAndReturn(run func(context.Context) ([]client.Replicator, error)) *DB_GetAllReplicators_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionByName provides a mock function with given fields: _a0, _a1
func (_m *DB) GetCollectionByName(_a0 context.Context, _a1 string) (client.Collection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (client.Collection, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) client.Collection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetCollectionByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionByName'
type DB_GetCollectionByName_Call struct {
	*mock.Call
}

// GetCollectionByName is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) GetCollectionByName(_a0 interface{}, _a1 interface{}) *DB_GetCollectionByName_Call {
	return &DB_GetCollectionByName_Call{Call: _e.mock.On("GetCollectionByName", _a0, _a1)}
}

func (_c *DB_GetCollectionByName_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_GetCollectionByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_GetCollectionByName_Call) Return(_a0 client.Collection, _a1 error) *DB_GetCollectionByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetCollectionByName_Call) RunAndReturn(run func(context.Context, string) (client.Collection, error)) *DB_GetCollectionByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionBySchemaID provides a mock function with given fields: _a0, _a1
func (_m *DB) GetCollectionBySchemaID(_a0 context.Context, _a1 string) (client.Collection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (client.Collection, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) client.Collection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetCollectionBySchemaID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionBySchemaID'
type DB_GetCollectionBySchemaID_Call struct {
	*mock.Call
}

// GetCollectionBySchemaID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) GetCollectionBySchemaID(_a0 interface{}, _a1 interface{}) *DB_GetCollectionBySchemaID_Call {
	return &DB_GetCollectionBySchemaID_Call{Call: _e.mock.On("GetCollectionBySchemaID", _a0, _a1)}
}

func (_c *DB_GetCollectionBySchemaID_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_GetCollectionBySchemaID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_GetCollectionBySchemaID_Call) Return(_a0 client.Collection, _a1 error) *DB_GetCollectionBySchemaID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetCollectionBySchemaID_Call) RunAndReturn(run func(context.Context, string) (client.Collection, error)) *DB_GetCollectionBySchemaID_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionByVersionID provides a mock function with given fields: _a0, _a1
func (_m *DB) GetCollectionByVersionID(_a0 context.Context, _a1 string) (client.Collection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (client.Collection, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) client.Collection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetCollectionByVersionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionByVersionID'
type DB_GetCollectionByVersionID_Call struct {
	*mock.Call
}

// GetCollectionByVersionID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) GetCollectionByVersionID(_a0 interface{}, _a1 interface{}) *DB_GetCollectionByVersionID_Call {
	return &DB_GetCollectionByVersionID_Call{Call: _e.mock.On("GetCollectionByVersionID", _a0, _a1)}
}

func (_c *DB_GetCollectionByVersionID_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_GetCollectionByVersionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_GetCollectionByVersionID_Call) Return(_a0 client.Collection, _a1 error) *DB_GetCollectionByVersionID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetCollectionByVersionID_Call) RunAndReturn(run func(context.Context, string) (client.Collection, error)) *DB_GetCollectionByVersionID_Call {
	_c.Call.Return(run)
	return _c
}

// LensRegistry provides a mock function with given fields:
func (_m *DB) LensRegistry() client.LensRegistry {
	ret := _m.Called()

	var r0 client.LensRegistry
	if rf, ok := ret.Get(0).(func() client.LensRegistry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.LensRegistry)
		}
	}

	return r0
}

// DB_LensRegistry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LensRegistry'
type DB_LensRegistry_Call struct {
	*mock.Call
}

// LensRegistry is a helper method to define mock.On call
func (_e *DB_Expecter) LensRegistry() *DB_LensRegistry_Call {
	return &DB_LensRegistry_Call{Call: _e.mock.On("LensRegistry")}
}

func (_c *DB_LensRegistry_Call) Run(run func()) *DB_LensRegistry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_LensRegistry_Call) Return(_a0 client.LensRegistry) *DB_LensRegistry_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_LensRegistry_Call) RunAndReturn(run func() client.LensRegistry) *DB_LensRegistry_Call {
	_c.Call.Return(run)
	return _c
}

// MaxTxnRetries provides a mock function with given fields:
func (_m *DB) MaxTxnRetries() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// DB_MaxTxnRetries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MaxTxnRetries'
type DB_MaxTxnRetries_Call struct {
	*mock.Call
}

// MaxTxnRetries is a helper method to define mock.On call
func (_e *DB_Expecter) MaxTxnRetries() *DB_MaxTxnRetries_Call {
	return &DB_MaxTxnRetries_Call{Call: _e.mock.On("MaxTxnRetries")}
}

func (_c *DB_MaxTxnRetries_Call) Run(run func()) *DB_MaxTxnRetries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_MaxTxnRetries_Call) Return(_a0 int) *DB_MaxTxnRetries_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_MaxTxnRetries_Call) RunAndReturn(run func() int) *DB_MaxTxnRetries_Call {
	_c.Call.Return(run)
	return _c
}

// NewConcurrentTxn provides a mock function with given fields: _a0, _a1
func (_m *DB) NewConcurrentTxn(_a0 context.Context, _a1 bool) (datastore.Txn, error) {
	ret := _m.Called(_a0, _a1)

	var r0 datastore.Txn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool) (datastore.Txn, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool) datastore.Txn); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Txn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_NewConcurrentTxn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewConcurrentTxn'
type DB_NewConcurrentTxn_Call struct {
	*mock.Call
}

// NewConcurrentTxn is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 bool
func (_e *DB_Expecter) NewConcurrentTxn(_a0 interface{}, _a1 interface{}) *DB_NewConcurrentTxn_Call {
	return &DB_NewConcurrentTxn_Call{Call: _e.mock.On("NewConcurrentTxn", _a0, _a1)}
}

func (_c *DB_NewConcurrentTxn_Call) Run(run func(_a0 context.Context, _a1 bool)) *DB_NewConcurrentTxn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool))
	})
	return _c
}

func (_c *DB_NewConcurrentTxn_Call) Return(_a0 datastore.Txn, _a1 error) *DB_NewConcurrentTxn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_NewConcurrentTxn_Call) RunAndReturn(run func(context.Context, bool) (datastore.Txn, error)) *DB_NewConcurrentTxn_Call {
	_c.Call.Return(run)
	return _c
}

// NewTxn provides a mock function with given fields: _a0, _a1
func (_m *DB) NewTxn(_a0 context.Context, _a1 bool) (datastore.Txn, error) {
	ret := _m.Called(_a0, _a1)

	var r0 datastore.Txn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool) (datastore.Txn, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool) datastore.Txn); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Txn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_NewTxn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewTxn'
type DB_NewTxn_Call struct {
	*mock.Call
}

// NewTxn is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 bool
func (_e *DB_Expecter) NewTxn(_a0 interface{}, _a1 interface{}) *DB_NewTxn_Call {
	return &DB_NewTxn_Call{Call: _e.mock.On("NewTxn", _a0, _a1)}
}

func (_c *DB_NewTxn_Call) Run(run func(_a0 context.Context, _a1 bool)) *DB_NewTxn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool))
	})
	return _c
}

func (_c *DB_NewTxn_Call) Return(_a0 datastore.Txn, _a1 error) *DB_NewTxn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_NewTxn_Call) RunAndReturn(run func(context.Context, bool) (datastore.Txn, error)) *DB_NewTxn_Call {
	_c.Call.Return(run)
	return _c
}

// PatchSchema provides a mock function with given fields: _a0, _a1
func (_m *DB) PatchSchema(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_PatchSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchSchema'
type DB_PatchSchema_Call struct {
	*mock.Call
}

// PatchSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) PatchSchema(_a0 interface{}, _a1 interface{}) *DB_PatchSchema_Call {
	return &DB_PatchSchema_Call{Call: _e.mock.On("PatchSchema", _a0, _a1)}
}

func (_c *DB_PatchSchema_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_PatchSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_PatchSchema_Call) Return(_a0 error) *DB_PatchSchema_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_PatchSchema_Call) RunAndReturn(run func(context.Context, string) error) *DB_PatchSchema_Call {
	_c.Call.Return(run)
	return _c
}

// PrintDump provides a mock function with given fields: ctx
func (_m *DB) PrintDump(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_PrintDump_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PrintDump'
type DB_PrintDump_Call struct {
	*mock.Call
}

// PrintDump is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DB_Expecter) PrintDump(ctx interface{}) *DB_PrintDump_Call {
	return &DB_PrintDump_Call{Call: _e.mock.On("PrintDump", ctx)}
}

func (_c *DB_PrintDump_Call) Run(run func(ctx context.Context)) *DB_PrintDump_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_PrintDump_Call) Return(_a0 error) *DB_PrintDump_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_PrintDump_Call) RunAndReturn(run func(context.Context) error) *DB_PrintDump_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveP2PCollection provides a mock function with given fields: ctx, collectionID
func (_m *DB) RemoveP2PCollection(ctx context.Context, collectionID string) error {
	ret := _m.Called(ctx, collectionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, collectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_RemoveP2PCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveP2PCollection'
type DB_RemoveP2PCollection_Call struct {
	*mock.Call
}

// RemoveP2PCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID string
func (_e *DB_Expecter) RemoveP2PCollection(ctx interface{}, collectionID interface{}) *DB_RemoveP2PCollection_Call {
	return &DB_RemoveP2PCollection_Call{Call: _e.mock.On("RemoveP2PCollection", ctx, collectionID)}
}

func (_c *DB_RemoveP2PCollection_Call) Run(run func(ctx context.Context, collectionID string)) *DB_RemoveP2PCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_RemoveP2PCollection_Call) Return(_a0 error) *DB_RemoveP2PCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_RemoveP2PCollection_Call) RunAndReturn(run func(context.Context, string) error) *DB_RemoveP2PCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Root provides a mock function with given fields:
func (_m *DB) Root() datastore.RootStore {
	ret := _m.Called()

	var r0 datastore.RootStore
	if rf, ok := ret.Get(0).(func() datastore.RootStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.RootStore)
		}
	}

	return r0
}

// DB_Root_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Root'
type DB_Root_Call struct {
	*mock.Call
}

// Root is a helper method to define mock.On call
func (_e *DB_Expecter) Root() *DB_Root_Call {
	return &DB_Root_Call{Call: _e.mock.On("Root")}
}

func (_c *DB_Root_Call) Run(run func()) *DB_Root_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Root_Call) Return(_a0 datastore.RootStore) *DB_Root_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Root_Call) RunAndReturn(run func() datastore.RootStore) *DB_Root_Call {
	_c.Call.Return(run)
	return _c
}

// SetMigration provides a mock function with given fields: _a0, _a1
func (_m *DB) SetMigration(_a0 context.Context, _a1 client.LensConfig) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.LensConfig) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_SetMigration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetMigration'
type DB_SetMigration_Call struct {
	*mock.Call
}

// SetMigration is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 client.LensConfig
func (_e *DB_Expecter) SetMigration(_a0 interface{}, _a1 interface{}) *DB_SetMigration_Call {
	return &DB_SetMigration_Call{Call: _e.mock.On("SetMigration", _a0, _a1)}
}

func (_c *DB_SetMigration_Call) Run(run func(_a0 context.Context, _a1 client.LensConfig)) *DB_SetMigration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.LensConfig))
	})
	return _c
}

func (_c *DB_SetMigration_Call) Return(_a0 error) *DB_SetMigration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_SetMigration_Call) RunAndReturn(run func(context.Context, client.LensConfig) error) *DB_SetMigration_Call {
	_c.Call.Return(run)
	return _c
}

// SetReplicator provides a mock function with given fields: ctx, rep
func (_m *DB) SetReplicator(ctx context.Context, rep client.Replicator) error {
	ret := _m.Called(ctx, rep)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Replicator) error); ok {
		r0 = rf(ctx, rep)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_SetReplicator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetReplicator'
type DB_SetReplicator_Call struct {
	*mock.Call
}

// SetReplicator is a helper method to define mock.On call
//   - ctx context.Context
//   - rep client.Replicator
func (_e *DB_Expecter) SetReplicator(ctx interface{}, rep interface{}) *DB_SetReplicator_Call {
	return &DB_SetReplicator_Call{Call: _e.mock.On("SetReplicator", ctx, rep)}
}

func (_c *DB_SetReplicator_Call) Run(run func(ctx context.Context, rep client.Replicator)) *DB_SetReplicator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.Replicator))
	})
	return _c
}

func (_c *DB_SetReplicator_Call) Return(_a0 error) *DB_SetReplicator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_SetReplicator_Call) RunAndReturn(run func(context.Context, client.Replicator) error) *DB_SetReplicator_Call {
	_c.Call.Return(run)
	return _c
}

// WithTxn provides a mock function with given fields: _a0
func (_m *DB) WithTxn(_a0 datastore.Txn) client.Store {
	ret := _m.Called(_a0)

	var r0 client.Store
	if rf, ok := ret.Get(0).(func(datastore.Txn) client.Store); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Store)
		}
	}

	return r0
}

// DB_WithTxn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTxn'
type DB_WithTxn_Call struct {
	*mock.Call
}

// WithTxn is a helper method to define mock.On call
//   - _a0 datastore.Txn
func (_e *DB_Expecter) WithTxn(_a0 interface{}) *DB_WithTxn_Call {
	return &DB_WithTxn_Call{Call: _e.mock.On("WithTxn", _a0)}
}

func (_c *DB_WithTxn_Call) Run(run func(_a0 datastore.Txn)) *DB_WithTxn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(datastore.Txn))
	})
	return _c
}

func (_c *DB_WithTxn_Call) Return(_a0 client.Store) *DB_WithTxn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_WithTxn_Call) RunAndReturn(run func(datastore.Txn) client.Store) *DB_WithTxn_Call {
	_c.Call.Return(run)
	return _c
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
