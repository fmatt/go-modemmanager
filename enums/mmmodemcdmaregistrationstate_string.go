// Code generated by "stringer -type=MMModemCdmaRegistrationState"; DO NOT EDIT.

package enums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModemCdmaRegistrationStateUnknown-0]
	_ = x[MmModemCdmaRegistrationStateRegistered-1]
	_ = x[MmModemCdmaRegistrationStateHome-2]
	_ = x[MmModemCdmaRegistrationStateRoaming-3]
}

const _MMModemCdmaRegistrationState_name = "MmModemCdmaRegistrationStateUnknownMmModemCdmaRegistrationStateRegisteredMmModemCdmaRegistrationStateHomeMmModemCdmaRegistrationStateRoaming"

var _MMModemCdmaRegistrationState_index = [...]uint8{0, 35, 73, 105, 140}

func (i MMModemCdmaRegistrationState) String() string {
	if i >= MMModemCdmaRegistrationState(len(_MMModemCdmaRegistrationState_index)-1) {
		return "MMModemCdmaRegistrationState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMModemCdmaRegistrationState_name[_MMModemCdmaRegistrationState_index[i]:_MMModemCdmaRegistrationState_index[i+1]]
}