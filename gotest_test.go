package main

import (
	"testing"

	"github.com/book-library-seat-system/go-server/server"
)

// UserHandle 测试部分
// func TestUserHandle(t *testing.T) {
// 	server.TesttestGET(t)
// 	server.TesttestPost(t)
// }

// UserHandle 测试部分
func TestSeatHandle(t *testing.T) {
	server.TestShowTimeIntervalInfoHandle(t)
	server.TestShowSeatInfoHandle1(t)
	server.TestShowSeatInfoHandle2(t)
}