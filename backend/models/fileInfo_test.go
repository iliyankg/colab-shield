package models

import (
	"testing"

	pb "github.com/iliyankg/colab-shield/protos"
)

func TestNewFileInfo(t *testing.T) {
	// Setup
	fi := NewFileInfo("fileId", "fileHash", "branchName")

	// Assert
	if fi.FileId != "fileId" {
		t.Errorf("Expected file ID to be 'fileId', got %s", fi.FileId)
	}
	if fi.FileHash != "fileHash" {
		t.Errorf("Expected file hash to be 'fileHash', got %s", fi.FileHash)
	}
	if len(fi.UserIds) != 0 {
		t.Errorf("Expected user IDs to be empty, got %v", fi.UserIds)
	}
	if fi.BranchName != "branchName" {
		t.Errorf("Expected branch name to be 'branchName', got %s", fi.BranchName)
	}
	if fi.ClaimMode != pb.ClaimMode_UNCLAIMED {
		t.Errorf("Expected claim mode to be 0, got %d", fi.ClaimMode)
	}
}

func TestFileInfo_Claim(t *testing.T) {
	t.Run("Claim NewFileInfo", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", pb.ClaimMode_EXCLUSIVE); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if len(fi.UserIds) != 1 {
			t.Errorf("Expected user IDs to have length 1, got %v", fi.UserIds)
		}
		if fi.UserIds[0] != "userId" {
			t.Errorf("Expected user ID to be 'userId', got %s", fi.UserIds[0])
		}
		if fi.ClaimMode != pb.ClaimMode_EXCLUSIVE {
			t.Errorf("Expected claim mode to be 1, got %d", fi.ClaimMode)
		}
	})

	t.Run("Claim Out of Date", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHashTwo", pb.ClaimMode_EXCLUSIVE); err != ErrFileOutOfDate {
			t.Errorf("Expected error to be ErrFileOutOfDate, got %v", err)
		}
	})

	t.Run("Claim Exclusive Exclusive", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", pb.ClaimMode_EXCLUSIVE); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if err := fi.Claim("userIdTwo", "fileHash", pb.ClaimMode_EXCLUSIVE); err != ErrFileAlreadyClaimed {
			t.Errorf("Expected error to be ErrFileAlreadyClaimed, got %v", err)
		}
	})

	t.Run("Claim Shared Shared", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", pb.ClaimMode_SHARED); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if err := fi.Claim("userIdTwo", "fileHash", pb.ClaimMode_SHARED); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(fi.UserIds) != 2 {
			t.Errorf("Expected user IDs to have length 2, got %v", fi.UserIds)
		}
	})

	t.Run("Invalid Claim Mode", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", pb.ClaimMode_UNCLAIMED); err != ErrInvalidClaimMode {
			t.Errorf("Expected error to be ErrInvalidClaimMode, got %v", err)
		}
	})

	t.Run("Claim Exclusive Shared", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", pb.ClaimMode_EXCLUSIVE); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if err := fi.Claim("userIdTwo", "fileHash", pb.ClaimMode_SHARED); err != ErrFileAlreadyClaimed {
			t.Errorf("Expected error to be ErrInvalidClaimMode, got %v", err)
		}
	})

	t.Run("Claim Shared Exclusive", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", pb.ClaimMode_SHARED); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if err := fi.Claim("userIdTwo", "fileHash", pb.ClaimMode_EXCLUSIVE); err != ErrInvalidClaimMode {
			t.Errorf("Expected error to be ErrInvalidClaimMode, got %v", err)
		}
	})
}

func TestFileInfo_Update(t *testing.T) {
	t.Run("Update Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")
		fi.UserIds = []string{"userId"}

		// Action
		if err := fi.Update("userId", "fileHashTwo", "branchNameTwo"); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if fi.FileHash != "fileHashTwo" {
			t.Errorf("Expected file hash to be 'fileHashTwo', got %s", fi.FileHash)
		}
		if fi.BranchName != "branchNameTwo" {
			t.Errorf("Expected branch name to be 'branchNameTwo', got %s", fi.BranchName)
		}
	})

	t.Run("Update Not Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Update("userId", "fileHashTwo", "branchNameTwo"); err != ErrUserNotOwner {
			t.Errorf("Expected error to be ErrUserNotOwner, got %v", err)
		}
	})
}

func TestFileInfo_Release(t *testing.T) {
	t.Run("Release Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")
		fi.UserIds = []string{"userId"}

		// Action
		if err := fi.Release("userId"); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if len(fi.UserIds) != 0 {
			t.Errorf("Expected user IDs to be empty, got %v", fi.UserIds)
		}
		if fi.ClaimMode != pb.ClaimMode_UNCLAIMED {
			t.Errorf("Expected claim mode to be 0, got %d", fi.ClaimMode)
		}
	})

	t.Run("Release Not Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Release("userId"); err != ErrUserNotOwner {
			t.Errorf("Expected error to be ErrUserNotOwner, got %v", err)
		}
	})
}

func TestFileInfo_CheckOwner(t *testing.T) {
	t.Run("Check Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")
		fi.UserIds = []string{"userId"}

		// Action
		if !fi.CheckOwner("userId") {
			t.Errorf("Expected user to be owner, got false")
		}
	})

	t.Run("Check Not Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if fi.CheckOwner("userId") {
			t.Errorf("Expected user to not be owner, got true")
		}
	})
}
