// Code generated by "stringer -type=MMCallState -trimprefix=MmCallState"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmCallStateUnknown-0]
	_ = x[MmCallStateDialing-1]
	_ = x[MmCallStateRingingOut-2]
	_ = x[MmCallStateRingingIn-3]
	_ = x[MmCallStateActive-4]
	_ = x[MmCallStateHeld-5]
	_ = x[MmCallStateWaiting-6]
	_ = x[MmCallStateTerminated-7]
}

const _MMCallState_name = "UnknownDialingRingingOutRingingInActiveHeldWaitingTerminated"

var _MMCallState_index = [...]uint8{0, 7, 14, 24, 33, 39, 43, 50, 60}

func (i MMCallState) String() string {
	if i >= MMCallState(len(_MMCallState_index)-1) {
		return "MMCallState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMCallState_name[_MMCallState_index[i]:_MMCallState_index[i+1]]
}