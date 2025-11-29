package models

type StructSizes struct{
	NormalSize int `json:"normal-size"`
	NormalAlignment int `json:"normalsalignment"`
	PackedSize int `json:"packed-size"`
	PackedAlignment int `json:"packed-alignment"`
}
