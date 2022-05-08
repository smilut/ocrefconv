package main

// example: 9a25c95e-36df-11e9-80ed-0050569f2e9f
// example: 80ed0050569f2e9f11e936df9a25c95e

import "testing"

func TestUUID2Ref(t *testing.T) {
	ref, err := uuid2Ref("9a25c95e-36df-11e9-80ed-0050569f2e9f")

	if ref != "80ed0050569f2e9f11e936df9a25c95e" || err != nil {
		t.Error("Ожидается '80ed0050569f2e9f11e936df9a25c95e', получено ", ref, " с ошибкой ", err)
	}
}

func TestRef2UUID(t *testing.T) {
	uuid, err := ref2UUID("80ed0050569f2e9f11e936df9a25c95e")

	if uuid != "9a25c95e-36df-11e9-80ed-0050569f2e9f" || err != nil {
		t.Error("Ожидается '9a25c95e-36df-11e9-80ed-0050569f2e9f', получено ", uuid, " с ошибкой ", err)
	}
}
