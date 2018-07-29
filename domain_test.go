package opsutilsgo

import (
	"fmt"
	"testing"
)

func TestGetDomain(t *testing.T) {
	result := GetDomain("現代醫藥衛生.cn")
	if result != "現代醫藥衛生.cn" {
		t.Fatal("error!")
	}
	fmt.Println(result)
}
