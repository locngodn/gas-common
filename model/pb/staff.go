package pb

import validation "github.com/go-ozzo/ozzo-validation"

func (staffs *StaffStationWarehouses) Validate() error {
	err := validation.ValidateStruct(staffs,
		validation.Field(&staffs.StaffId, validation.Required),
	)
	if err != nil {
		return err
	}

	if len(staffs.Data) > 0 {
		err = validation.ValidateStruct(staffs,
			validation.Field(&staffs.Data, validation.Required),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (staff *StaffStationWarehouse) Validate() error {
	return validation.ValidateStruct(staff,
		validation.Field(&staff.StaffId, validation.Required),
		validation.Field(&staff.StationWarehouseId, validation.Required),
	)
}
