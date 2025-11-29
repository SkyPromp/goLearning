package models

type MemStats struct{
	BeforeAlloc string `json:"before-alloc"`
	AfterAlloc string `json:"after-alloc"`
	AfterGC string `json:"after-gc"`
}
