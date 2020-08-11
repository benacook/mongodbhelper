package mongodbhelper_test

import (
	"errors"
	"github.com/benacook/mongodbhelper/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestConnect(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mock_mongodbhelper.NewMockMongoInterface(mockCtrl)
	mdb := mockDoer

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().Connect("mongodb://mongo:27017").Return(nil).Times(1)
	mockDoer.EXPECT().Connect("mongodb://mongo:27010").Return(errors.New(
		"invalid host")).Times(1)
	err := mdb.Connect("mongodb://mongo:27017")
	if err != nil {
		t.Fail()
	}
	err = mdb.Connect("mongodb://mongo:27010")
	if err == nil {
		t.Fail()
	}
}

type dummyRecord struct {
	Name string
	Age int

}

func TestInsertElement(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mock_mongodbhelper.NewMockMongoInterface(mockCtrl)
	mdb := mockDoer
	dr := dummyRecord{"ben", 27}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().InsertElement(dr).Return(nil).Times(1)
	mockDoer.EXPECT().InsertElement(nil).Return(errors.New("no element to insert")).
		Times(1)

	err := mdb.InsertElement(dr)
	if err != nil {
		t.Fail()
	}

	err = mdb.InsertElement(nil)
	if err == nil {
		t.Fail()
	}
}
