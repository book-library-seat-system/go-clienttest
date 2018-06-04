package main

import (
	"testing"

	"github.com/book-library-seat-system/go-server/server"
)

// UserHandle 测试部分
func TestUserHandle(t *testing.T) {
	server.TesttestGET(t)
	server.TesttestPost(t)
}
