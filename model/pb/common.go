package pb

func (f *StoreFilter) Empty() bool {
	return len(f.Province) == 0 &&
		len(f.District) == 0 &&
		len(f.Ward) == 0 &&
		len(f.Status) == 0
}
