package main

import (
	"sync"
)

type Doc struct {
	Title string
	Val   string
}

var poolDoc = &sync.Pool{
	New: func() any {
		return &Doc{}
	},
}
