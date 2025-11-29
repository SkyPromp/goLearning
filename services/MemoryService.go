package services

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"runtime"
	"github.com/SkyPromp/goLearning/models"
)

func MemoryManagement() (models.MemStats){
	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)

	beforeAlloc := fmt.Sprintf("%v kB", mem.Alloc / 1024)

	references := make([] * [1_000_000]int, 100)

	for i := range 100{
		references[i] = new([1_000_000]int)
	}

	runtime.ReadMemStats(&mem)

	afterAlloc := fmt.Sprintf("%v kB", mem.Alloc / 1024)

	references = nil
	runtime.GC()

	runtime.ReadMemStats(&mem)

	afterGC := fmt.Sprintf("%v kB", mem.Alloc / 1024)

	return models.MemStats{BeforeAlloc: beforeAlloc, AfterAlloc: afterAlloc, AfterGC: afterGC}
}

func UnsafeMemoryManagement() (models.MemStats) {
    var mem runtime.MemStats

    runtime.ReadMemStats(&mem)
    beforeAlloc := fmt.Sprintf("%v kB", mem.Alloc/1024)

    size := 1_000_000
    ptr := C.malloc(C.size_t(size))

    // Touch memory so OS commits pages
    data := (*[1 << 30]byte)(ptr)[:size:size]
    for i := range size{
        data[i] = 1
    }

    runtime.ReadMemStats(&mem)
    afterAlloc := fmt.Sprintf("%v kB", mem.Alloc/1024)

    C.free(ptr)

    runtime.ReadMemStats(&mem)
    afterFree := fmt.Sprintf("%v kB", mem.Alloc/1024)

    return models.MemStats{
        BeforeAlloc: beforeAlloc,
        AfterAlloc:  afterAlloc,
        AfterGC:     afterFree,
    }
}
