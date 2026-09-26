package utils

import "testing"

func Test_ByteSize(t *testing.T) {
	tests := make([]float64, 0)
	tests = append(tests, 0)
	tests = append(tests, uByte)
	tests = append(tests, uKilobyte)
	tests = append(tests, uMegabyte)
	tests = append(tests, uGigabyte)
	tests = append(tests, uTerabyte)
	tests = append(tests, uPetabyte)
	tests = append(tests, uExabyte)
	tests = append(tests, uKilobyte+(uKilobyte/2))
	tests = append(tests, uMegabyte+(uMegabyte/2))
	tests = append(tests, uGigabyte+(uGigabyte/2))
	tests = append(tests, uTerabyte+(uTerabyte/2))
	tests = append(tests, uPetabyte+(uPetabyte/2))
	tests = append(tests, uExabyte+(uExabyte/2))
	
	expectedResults := make([]string, 0)
	expectedResults = append(expectedResults, "0B")
	expectedResults = append(expectedResults, "1B")
	expectedResults = append(expectedResults, "1KB")
	expectedResults = append(expectedResults, "1MB")
	expectedResults = append(expectedResults, "1GB")
	expectedResults = append(expectedResults, "1TB")
	expectedResults = append(expectedResults, "1PB")
	expectedResults = append(expectedResults, "1EB")
	expectedResults = append(expectedResults, "1.5KB")
	expectedResults = append(expectedResults, "1.5MB")
	expectedResults = append(expectedResults, "1.5GB")
	expectedResults = append(expectedResults, "1.5TB")
	expectedResults = append(expectedResults, "1.5PB")
	expectedResults = append(expectedResults, "1.5EB")
	
	for i := range len(tests) {
		et := expectedResults[i]
		r := ByteSize(tests[i])
		if r != et {
			t.Errorf("Expected %s, got %s", et, r)
		}
	}
}
