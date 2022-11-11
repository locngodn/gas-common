package model

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/locngodn/gas-common/util"
	"github.com/pkg/errors"
)

type (
	ValidationError struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}
	UserError struct {
		Message    string `json:"message"`
		Code       string `json:"code"`
		StatusCode int    `json:"-"`
	}
	Status struct {
		StatusCode int
		StatusText string
	}
	User struct {
		Id                string
		AccountId         string
		KongId            string
		StationWarehouses []string
	}
)

const (
	store    = "CH"
	staff    = "S"
	customer = "C"
	agency   = "A"
	branch   = "B"
)

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *UserError) Error() string {
	return e.Message
}

func (u *User) IncludeStationWarehouse(stationWarehouseId string) bool {
	if len(u.StationWarehouses) == 0 {
		return true
	}
	if len(stationWarehouseId) == 0 {
		return false
	}
	return util.Contains(u.StationWarehouses, stationWarehouseId)
}

func (u *User) IncludeStationWarehouses(stationWarehouses []string) bool {
	if len(u.StationWarehouses) == 0 {
		return true
	}
	if len(stationWarehouses) == 0 {
		return false
	}
	return util.Include(u.StationWarehouses, stationWarehouses)
}

func (u *User) GetValidStationWarehouses(stationWarehouses []string) ([]string, error) {
	if len(u.StationWarehouses) == 0 {
		return stationWarehouses, nil
	}
	if len(stationWarehouses) == 0 {
		return u.StationWarehouses, nil
	}
	if !util.Include(u.StationWarehouses, stationWarehouses) {
		return nil, &UserError{Message: "Permission denied: Invalid User Station or Warehouse", StatusCode: http.StatusForbidden}
	}
	return stationWarehouses, nil
}

func (u *User) IsStaff() bool {
	return strings.HasPrefix(u.Id, staff)
}

func (u *User) IsAgency() bool {
	return strings.HasPrefix(u.Id, agency)
}

func (u *User) IsBranch() bool {
	return strings.HasPrefix(u.Id, branch)
}

func (u *User) IsCustomer() bool {
	return strings.HasPrefix(u.Id, customer)
}

func (u *User) IsStore() bool {
	return strings.HasPrefix(u.Id, store)
}

func (u *User) IsStoreMember() bool {
	return strings.HasPrefix(u.Id, "CHNV")
}

func (u *User) IsAnonymous() bool {
	return u == nil
}

func (u *User) GetType() string {
	if u.IsStore() {
		return store
	}
	if u.IsStaff() {
		return staff
	}
	if u.IsCustomer() {
		return customer
	}
	if u.IsAgency() {
		return agency
	}
	if u.IsBranch() {
		return branch
	}
	return ""
}

func NewUser(r *http.Request) User {
	userId := r.Header.Get("x-consumer-username")
	kongId := r.Header.Get("x-consumer-id")
	id := r.Header.Get("x-consumer-custom-id")
	return User{Id: userId, AccountId: id, KongId: kongId}
}

func NewStatus(err error, isNoContent bool) Status {
	if err != nil {
		if verr, ok := err.(*ValidationError); ok {
			return Status{
				StatusCode: http.StatusBadRequest,
				StatusText: verr.Error(),
			}
		} else if uerr, ok := err.(*UserError); ok {
			return Status{
				StatusCode: uerr.StatusCode,
				StatusText: uerr.Message,
			}
		} else {
			return Status{
				StatusCode: http.StatusInternalServerError,
				StatusText: err.Error(),
			}
		}
	} else {
		if isNoContent {
			return Status{
				StatusCode: http.StatusNoContent,
				StatusText: "Ok",
			}
		} else {
			return Status{
				StatusCode: http.StatusOK,
				StatusText: "Ok",
			}
		}
	}
}

func (s Status) ToError() error {
	switch s.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return &ValidationError{Message: s.StatusText}
	case http.StatusInternalServerError:
		return errors.New(s.StatusText)
	default:
		return &UserError{Message: s.StatusText, StatusCode: s.StatusCode}
	}
}

func IsInternalError(err error) bool {
	if _, ok := err.(*ValidationError); ok {
		return false
	} else if _, ok := err.(*UserError); ok {
		return false
	} else {
		return true
	}
}

func WrapError(err error, prefix string) error {
	if verr, ok := err.(*ValidationError); ok {
		verr.Message = fmt.Sprintf("%s: %s", prefix, verr.Message)
		return verr
	} else if uerr, ok := err.(*UserError); ok {
		uerr.Message = fmt.Sprintf("%s: %s", prefix, uerr.Message)
		return uerr
	} else {
		return errors.Wrap(err, prefix)
	}
}

func NewError(err error, code, msg string) error {
	if verr, ok := err.(*ValidationError); ok {
		verr.Message = msg
		verr.Code = code
		return verr
	} else if uerr, ok := err.(*UserError); ok {
		uerr.Message = msg
		uerr.Code = code
		return uerr
	} else {
		return errors.New(msg)
	}
}
