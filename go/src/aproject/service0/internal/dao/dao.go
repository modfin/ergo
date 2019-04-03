package dao

import (
	"aproject/service0"
)

func GetModel(a, b int) service0.ModelService0 {
	return service0.ModelService0{
		A: a,
		B: b,
	}
}
