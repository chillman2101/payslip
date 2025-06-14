// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/payslip/models"
	mock "github.com/stretchr/testify/mock"
)

// PayrollRepository is an autogenerated mock type for the PayrollRepository type
type PayrollRepository struct {
	mock.Mock
}

// CreatePayroll provides a mock function with given fields: ctx, payroll
func (_m *PayrollRepository) CreatePayroll(ctx context.Context, payroll *models.Payroll) error {
	ret := _m.Called(ctx, payroll)

	if len(ret) == 0 {
		panic("no return value specified for CreatePayroll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Payroll) error); ok {
		r0 = rf(ctx, payroll)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindPayrollByDate provides a mock function with given fields: ctx, payroll
func (_m *PayrollRepository) FindPayrollByDate(ctx context.Context, payroll *models.Payroll) (*models.Payroll, error) {
	ret := _m.Called(ctx, payroll)

	if len(ret) == 0 {
		panic("no return value specified for FindPayrollByDate")
	}

	var r0 *models.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Payroll) (*models.Payroll, error)); ok {
		return rf(ctx, payroll)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Payroll) *models.Payroll); ok {
		r0 = rf(ctx, payroll)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Payroll) error); ok {
		r1 = rf(ctx, payroll)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSummaryPayrollByPayrollId provides a mock function with given fields: ctx, payroll_id
func (_m *PayrollRepository) GetSummaryPayrollByPayrollId(ctx context.Context, payroll_id int) (*models.Payroll, error) {
	ret := _m.Called(ctx, payroll_id)

	if len(ret) == 0 {
		panic("no return value specified for GetSummaryPayrollByPayrollId")
	}

	var r0 *models.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*models.Payroll, error)); ok {
		return rf(ctx, payroll_id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Payroll); ok {
		r0 = rf(ctx, payroll_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, payroll_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSummaryPayrollByPayrollIdAndEmployeeId provides a mock function with given fields: ctx, payroll_id, employee_id
func (_m *PayrollRepository) GetSummaryPayrollByPayrollIdAndEmployeeId(ctx context.Context, payroll_id int, employee_id int) (*models.Payroll, error) {
	ret := _m.Called(ctx, payroll_id, employee_id)

	if len(ret) == 0 {
		panic("no return value specified for GetSummaryPayrollByPayrollIdAndEmployeeId")
	}

	var r0 *models.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (*models.Payroll, error)); ok {
		return rf(ctx, payroll_id, employee_id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *models.Payroll); ok {
		r0 = rf(ctx, payroll_id, employee_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, payroll_id, employee_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPayrollUnprocessed provides a mock function with given fields: ctx
func (_m *PayrollRepository) ListPayrollUnprocessed(ctx context.Context) ([]models.Payroll, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListPayrollUnprocessed")
	}

	var r0 []models.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Payroll, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Payroll); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessPayroll provides a mock function with given fields: ctx, payroll_id
func (_m *PayrollRepository) ProcessPayroll(ctx context.Context, payroll_id int) error {
	ret := _m.Called(ctx, payroll_id)

	if len(ret) == 0 {
		panic("no return value specified for ProcessPayroll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, payroll_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPayrollRepository creates a new instance of PayrollRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPayrollRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PayrollRepository {
	mock := &PayrollRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
