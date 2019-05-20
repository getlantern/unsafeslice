// package unsafeslice contains unsafe convenience operations on slices. Use with caution.
package unsafeslice

import (
	"reflect"
	"unsafe"
)

// ShiftHeadLeft shifts the head of a slice left. This assumes that the slice was previously
// resliced such that there's actually data available to the left in the underlying array. If
// you use ShiftHeadLeft to shift past this point, you risk data corruption.
//
// s is a pointer to the slice that is to be modified.
func ShiftHeadLeft(s interface{}, offset int) {
	v := reflect.ValueOf(s)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(v.Pointer()))
	elemType := v.Elem().Type().Elem()
	hdr.Data -= uintptr(offset) * elemType.Size()
	hdr.Len += offset
	hdr.Cap += offset
}
