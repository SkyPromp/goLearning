package services

import (
	"unsafe"
	"github.com/SkyPromp/goLearning/models"
)

func GetByteAlignment() (models.StructSizes){
	type Rectangle struct{
		X uint8
		Y uint8
		Color uint32
		W uint8
		H uint8
	}

	type RectangleOptimized struct{
		X uint8
		Y uint8
		W uint8
		H uint8
		Color uint32
	}

	return models.StructSizes{
		NormalSize: int(unsafe.Sizeof(Rectangle{})),
		NormalAlignment: int(unsafe.Alignof(Rectangle{})),
		PackedSize: int(unsafe.Sizeof(RectangleOptimized{})),
		PackedAlignment: int(unsafe.Alignof(RectangleOptimized{})),
	}
}

func GetUnsafeExample() (int32){
	var x int32 = 10

	base := uintptr(unsafe.Pointer(&x))
	ptr := (*byte)(unsafe.Pointer(base + 1))

	*ptr = 1

	return x
}
