// Code generated by "stringer -type=MMSmsStorage"; DO NOT EDIT.

package enums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSmsStorageUnknown-0]
	_ = x[MmSmsStorageSm-1]
	_ = x[MmSmsStorageMe-2]
	_ = x[MmSmsStorageMt-3]
	_ = x[MmSmsStorageSr-4]
	_ = x[MmSmsStorageBm-5]
	_ = x[MmSmsStorageTa-6]
}

const _MMSmsStorage_name = "MmSmsStorageUnknownMmSmsStorageSmMmSmsStorageMeMmSmsStorageMtMmSmsStorageSrMmSmsStorageBmMmSmsStorageTa"

var _MMSmsStorage_index = [...]uint8{0, 19, 33, 47, 61, 75, 89, 103}

func (i MMSmsStorage) String() string {
	if i >= MMSmsStorage(len(_MMSmsStorage_index)-1) {
		return "MMSmsStorage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMSmsStorage_name[_MMSmsStorage_index[i]:_MMSmsStorage_index[i+1]]
}