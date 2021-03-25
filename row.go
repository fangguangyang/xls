package xls

type rowInfo struct {
	Index    uint16
	Fcell    uint16
	Lcell    uint16
	Height   uint16
	Notused  uint16
	Notused2 uint16
	Flags    uint32
}

// Row handle.
type Row struct {
	wb   *WorkBook
	info *rowInfo
	Cols map[uint16]contentHandler
}

// Col gets the n'th column (zero-based). If not found it will return empty string.
// Merged cells will be reported more then once if iterating.
func (r *Row) Col(n int) string {
	serial := uint16(n)
	if ch, ok := r.Cols[serial]; ok {
		strs := ch.String(r.wb)
		return strs[0]
	}
	for _, v := range r.Cols {
		if v.FirstCol() <= serial && v.LastCol() >= serial {
			strs := v.String(r.wb)
			return strs[serial-v.FirstCol()]
		}
	}
	return ""
}

// ColExact gets the n'th column (zero-based). If not found it will return empty string.
// Merged cells will only show the value at the first cell.
func (r *Row) ColExact(n int) string {
	serial := uint16(n)
	if ch, ok := r.Cols[serial]; ok {
		strs := ch.String(r.wb)
		for _, s := range strs {
			if len(s) == 0 {
				continue
			}
			return s
		}
		return ""
	}
	return ""
}

// Value of the cell.
func (r *Row) Value(n int) CellValue {
	serial := uint16(n)
	if ch, ok := r.Cols[serial]; ok {
		return ch.Value(r.wb)
	}
	return CellValue{}
}

// LastCol gets the index of the last column.
func (r *Row) LastCol() int {
	return int(r.info.Lcell)
}

// FirstCol gets the index of the first column.
func (r *Row) FirstCol() int {
	return int(r.info.Fcell)
}
