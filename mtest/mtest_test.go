package mtest

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/vsafonkin/sample-go/mtest/mock_mtest"
)

func TestHelloRunner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedRunner := mock_mtest.NewMockRunner(ctrl)
	mockedRunner.EXPECT().Run(123).Return(nil)

	HelloRunner(123, mockedRunner)
}

func TestIntSum(t *testing.T) {

	data := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"test1", 2, 3, 5},
		{"test2", 3, 4, 7},
		{"test3", -5, 5, 0},
		{"test4", 0, 0, 0},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := IntSum(d.a, d.b)
			if result != d.expected {
				t.Errorf("failed: expected %d, got %d\n", d.expected, result)
			}
		})
	}

	t.Run("<SUPERTEST>", func(t *testing.T) {
		t.Log("<START SUPERTEST>")
		t.Log("<SUPERTEST SUCCESS>")
	})
}

func Test_doSomething(t *testing.T) {
	result := doSomething()
	expected := "do something"
	if result != expected {
		t.Errorf("failed: expected: %s, got %s\n", expected, result)
	}
}

func Test_dataRace(t *testing.T) {
	result := dataRace()
	expected := 0
	if result != expected {
		t.Errorf("failed: expected: %d, got %d\n", expected, result)
	}
}

var output int

func BenchmarkIntSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := IntSum(3, 4)
		output = result
	}
}

func BenchmarkRunSomething(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := RunSomething(200 * time.Millisecond)
		output = result
	}
}
