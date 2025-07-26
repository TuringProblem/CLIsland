package tests

import (
	"os"
	"testing"
)

// TestConfig holds configuration for tests
type TestConfig struct {
	TestDataPath string
	Verbose      bool
	RunSlowTests bool
}

// GetTestConfig returns test configuration based on environment
func GetTestConfig() TestConfig {
	return TestConfig{
		TestDataPath: getEnvOrDefault("TEST_DATA_PATH", "data"),
		Verbose:      getEnvOrDefault("TEST_VERBOSE", "false") == "true",
		RunSlowTests: getEnvOrDefault("TEST_RUN_SLOW", "false") == "true",
	}
}

// getEnvOrDefault gets environment variable or returns default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// SkipIfSlowTest skips the test if slow tests are disabled
func SkipIfSlowTest(t *testing.T) {
	config := GetTestConfig()
	if !config.RunSlowTests {
		t.Skip("Skipping slow test. Set TEST_RUN_SLOW=true to run.")
	}
}

// SkipIfNoTestData skips the test if test data is not available
func SkipIfNoTestData(t *testing.T) {
	config := GetTestConfig()
	if _, err := os.Stat(config.TestDataPath); os.IsNotExist(err) {
		t.Skipf("Test data not found at %s. Set TEST_DATA_PATH to specify location.", config.TestDataPath)
	}
}
