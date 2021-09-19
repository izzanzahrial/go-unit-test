package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// naming convention in go for test function is begin with "Test"
// for unit testing use parameter (t *testing.T) and return nothing
// to run the test run "go test" in your package directory
// "go test -v" verbose, to look at the function that currently testing
// "go test -v -run [name of the function]" to test spesific function
// "go test ./..." to run all of the unit test from the top folder

// how to fail the test in go
// *testing.T already have built in function for it : Fail(), FailNow(), Error(), Fatal()
// Fail(), will fail the test, but keep continuing the test, but at the end the test will still failed
// FailNow(), will fail the test at that moment without continuing the test
// Error(), will log the error, then run like Fail()
// Fatal(), will log the error, then run like FailNow()
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Izzan")
	if result != "Hello Izzan" {
		// unit test failed
		t.Fail()
	}
	fmt.Println("Test hello world Fail()")
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("zan")
	if result != "Hello Izzan" {
		// unit test failed
		t.Fail()
	}
	fmt.Println("Test hello world Fail()") // <-- still run because using Fail()
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("zan")
	if result != "Hello Izzan" {
		// unit test failed
		t.FailNow()
	}
	fmt.Println("Test hello world FailNow()") // <-- won't run because using FailNow()
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("zan")
	if result != "Hello Izzan" {
		// unit test failed
		t.Error("The result is not the same")
	}
	fmt.Println("Test hello world Error()") // <-- still run because using Error() that like Fail()
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("zan")
	if result != "Hello Izzan" {
		// unit test failed
		t.Fatal("The result is not the same")
	}
	fmt.Println("Test hello world Fatal()") // <-- won't run because using Fatal() that like FailNow()
}

// Assertion using testify library
// https://github.com/stretchr/testify
// in assertion if the test fail, it will continue the test like Fail()

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("zan")
	assert.Equal(t, "Hello Izzan", result, "The result is not the same")
	fmt.Println("TestHelloWorld with Assertion") // <-- will run
}

// Require
// in require if the test fail, it won't continue the test like FailNow()

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("zan")
	require.Equal(t, "Hello Izzan", result, "The result is not the same")
	fmt.Println("TestHelloWorld with Assertion") // <-- won't run
}

// Skip Test
// if dont wanna run spesific test function
// example : dont run mobile function in desktop

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run the test on Windows OS") // <-- will skip the test if you run the test on windows OS
	}

	result := HelloWorld("Izzan")
	require.Equal(t, "Hello Izzan", result, "The result is not the same")
}

// Test Main = (m *testing.M)
// test main can be used to arrange how to execute unit test
// test main will run once per Go-lang package
// so you need to make it per package
// "go test -v"

func TestMain(m *testing.M) {
	fmt.Println("Start the hello world test") // <-- will run even tho the test failed

	m.Run() // execute all the test

	fmt.Println("The test is done") // <-- will run even tho the test failed
}

// Sub test = nested test
// t.Run()
// to run spesific subtest use "go test -v -run [NameOfTheTest/NameOfTheSubTest]"
// example : go test -v -run TestSubTest/SubTestName

func TestSubTest(t *testing.T) {
	t.Run("SubTestName", func(t *testing.T) {
		result := HelloWorld("Izzan")
		require.Equal(t, "Hello Izzan", result, "The result is not the same")
	})
	// you can create more than one sub test
	t.Run("SubTest2", func(t *testing.T) {
		result := HelloWorld("Izzan")
		require.Equal(t, "Hello Izzan", result, "The result is not the same")
	})
}

// Table test
// Table test is used for testing a range of data at once

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string // name of the sub test
		request  string // request data
		expected string // expected result
	}{
		// the range of data
		{
			name:     "HelloWorld(Ahmad)",
			request:  "Ahmad",
			expected: "Hello Ahmad",
		},
		{
			name:     "HelloWorld(Izzan)",
			request:  "Izzan",
			expected: "Hello Izzan",
		},
		{
			name:     "HelloWorld(Zahrial)",
			request:  "Zahrial",
			expected: "Hello Zahrial",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

// After this check repository/category_repository_mock

// Benchmark
// (b *testing.B)
// "go test -v -bench=." to run all benchmark and unit test in current directory/package
// "go test -v -run=[Name of non existed unit test] -bench=." to run all benchmark in current directory/package
// "go test -v -run=[Name of non existed unit test] -bench=[name of the benchmark function]"
// example "go test -v -run=[Name of non existed unit test] -bench=BenchmarkHelloWorld"
// to run spesific benchmark in current direcotry/package
// "go test -v -bench=. ./..." to run all benchmark from parent directory
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ { // <-- the value of the N will be calculated by benchmark itself
		HelloWorld("Izzan")
	}
}

// Sub benchmark
// "b.Run()" = sub benchmark
// "go test -v -bench=[Name of the benchmark sub function]"
// all the commands almost the same like sub unit test
func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("BenchmarkSub1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Izzan")
		}
	})

	b.Run("BenchmarkSub2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Izzan")
		}
	})
}

// Benchmark Table
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string // name of the sub test
		request string // request data
	}{
		// the range of data
		{
			name:    "HelloWorld(Ahmad)",
			request: "Ahmad",
		},
		{
			name:    "HelloWorld(Izzan)",
			request: "Izzan",
		},
		{
			name:    "HelloWorld(Zahrial)",
			request: "Zahrial",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
