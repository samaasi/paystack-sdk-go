package paystackapi

import (
	"strings"
	"testing"
)

func TestGenerateUUIDv4_Format(t *testing.T) {
	uuid, err := GenerateUUIDv4()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		t.Errorf("expected 5 hyphen-separated groups, got %d in %q", len(parts), uuid)
	}
	if len(uuid) != 36 {
		t.Errorf("expected 36 characters, got %d in %q", len(uuid), uuid)
	}
}

func TestGenerateUUIDv4_Version4(t *testing.T) {
	uuid, err := GenerateUUIDv4()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// 15th character (index 14) must be '4' for version 4
	if uuid[14] != '4' {
		t.Errorf("UUID version indicator at index 14 should be '4', got %q in %q", uuid[14], uuid)
	}
}

func TestGenerateUUIDv4_Uniqueness(t *testing.T) {
	seen := make(map[string]bool, 100)
	for i := 0; i < 100; i++ {
		uuid, err := GenerateUUIDv4()
		if err != nil {
			t.Fatalf("unexpected error at iteration %d: %v", i, err)
		}
		if seen[uuid] {
			t.Fatalf("duplicate UUID generated: %q", uuid)
		}
		seen[uuid] = true
	}
}
