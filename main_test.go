package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shivamarora1/gomock/example-1/foo/mocks"
)

func TestMain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockFoo(ctrl)
	m.EXPECT().Do(12).Return(101)
	ProcessFoo(m)
}
