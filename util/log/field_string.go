// Code generated by "stringer -type Field"; DO NOT EDIT

package log

import "fmt"

const _Field_name = "NodeIDStoreIDRangeIDKeymaxField"

var _Field_index = [...]uint8{0, 6, 13, 20, 23, 31}

func (i Field) String() string {
	if i < 0 || i >= Field(len(_Field_index)-1) {
		return fmt.Sprintf("Field(%d)", i)
	}
	return _Field_name[_Field_index[i]:_Field_index[i+1]]
}
