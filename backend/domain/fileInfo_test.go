package domain

import (
	"testing"
)

func TestNewFileInfo(t *testing.T) {
	// Setup
	fi := NewFileInfo("fileId", "fileHash", "branchName")

	// Assert
	if fi.fileId != "fileId" {
		t.Errorf("Expected file ID to be 'fileId', got %s", fi.fileId)
	}
	if fi.fileHash != "fileHash" {
		t.Errorf("Expected file hash to be 'fileHash', got %s", fi.fileHash)
	}
	if len(fi.userIds) != 0 {
		t.Errorf("Expected user IDs to be empty, got %v", fi.userIds)
	}
	if fi.branchName != "branchName" {
		t.Errorf("Expected branch name to be 'branchName', got %s", fi.branchName)
	}
	if fi.claimMode != ClaimMode_Unclaimed {
		t.Errorf("Expected claim mode to be 0, got %d", fi.claimMode)
	}
}

func TestFileInfo_Claim(t *testing.T) {
	t.Run("Claim NewFileInfo", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", ClaimMode_Exclusive); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if len(fi.userIds) != 1 {
			t.Errorf("Expected user IDs to have length 1, got %v", fi.userIds)
		}
		if fi.userIds[0] != "userId" {
			t.Errorf("Expected user ID to be 'userId', got %s", fi.userIds[0])
		}
		if fi.claimMode != ClaimMode_Exclusive {
			t.Errorf("Expected claim mode to be 1, got %d", fi.claimMode)
		}
	})

	t.Run("Claim Out of Date", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHashTwo", ClaimMode_Exclusive); err != ErrFileOutOfDate {
			t.Errorf("Expected error to be ErrFileOutOfDate, got %v", err)
		}
	})

	t.Run("Claim Exclusive Exclusive", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", ClaimMode_Exclusive); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if err := fi.Claim("userIdTwo", "fileHash", ClaimMode_Exclusive); err != ErrFileAlreadyClaimed {
			t.Errorf("Expected error to be ErrFileAlreadyClaimed, got %v", err)
		}
	})

	t.Run("Claim Shared Shared", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", ClaimMode_Shared); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if err := fi.Claim("userIdTwo", "fileHash", ClaimMode_Shared); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(fi.userIds) != 2 {
			t.Errorf("Expected user IDs to have length 2, got %v", fi.userIds)
		}
	})

	t.Run("Invalid Claim Mode", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", ClaimMode_Unclaimed); err != ErrInvalidClaimMode {
			t.Errorf("Expected error to be ErrInvalidClaimMode, got %v", err)
		}
	})

	t.Run("Claim Exclusive Shared", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", ClaimMode_Exclusive); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if err := fi.Claim("userIdTwo", "fileHash", ClaimMode_Shared); err != ErrFileAlreadyClaimed {
			t.Errorf("Expected error to be ErrInvalidClaimMode, got %v", err)
		}
	})

	t.Run("Claim Shared Exclusive", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Claim("userId", "fileHash", ClaimMode_Shared); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if err := fi.Claim("userIdTwo", "fileHash", ClaimMode_Exclusive); err != ErrInvalidClaimMode {
			t.Errorf("Expected error to be ErrInvalidClaimMode, got %v", err)
		}
	})
}

func TestFileInfo_Update(t *testing.T) {
	t.Run("Update Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")
		fi.userIds = []string{"userId"}

		// Action
		if err := fi.Update("userId", "branchNameTwo", "fileHash", "fileHashTwo"); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if fi.fileHash != "fileHashTwo" {
			t.Errorf("Expected file hash to be 'fileHashTwo', got %s", fi.fileHash)
		}
		if fi.branchName != "branchNameTwo" {
			t.Errorf("Expected branch name to be 'branchNameTwo', got %s", fi.branchName)
		}
	})

	t.Run("Update Not Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.Update("userId", "branchNameTwo", "fileHash", "fileHashTwo"); err != ErrUserNotOwner {
			t.Errorf("Expected error to be ErrUserNotOwner, got %v", err)
		}
	})

	t.Run("Update Out of Date", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")
		fi.userIds = []string{"userId"}

		// Action
		if err := fi.Update("userId", "branchNameTwo", "fileHashTwo", "fileHash"); err != ErrFileOutOfDate {
			t.Errorf("Expected error to be ErrFileOutOfDate, got %v", err)
		}
	})
}

func TestFileInfo_Release(t *testing.T) {
	t.Run("Release Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")
		fi.userIds = []string{"userId"}

		// Action
		if err := fi.Release("userId"); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if len(fi.userIds) != 0 {
			t.Errorf("Expected user IDs to be empty, got %v", fi.userIds)
		}
		if fi.claimMode != ClaimMode_Unclaimed {
			t.Errorf("Expected claim mode to be 0, got %d", fi.claimMode)
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
		fi.userIds = []string{"userId"}

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

func TestFileInfo_addOwner(t *testing.T) {
	t.Run("Add Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		fi.addOwner("userId")

		// Assert
		if len(fi.userIds) != 1 {
			t.Errorf("Expected user IDs to have length 1, got %v", fi.userIds)
		}
		if fi.userIds[0] != "userId" {
			t.Errorf("Expected user ID to be 'userId', got %s", fi.userIds[0])
		}
	})

	t.Run("Add Owner Twice", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		fi.addOwner("userId")
		fi.addOwner("userId")

		// Assert
		if len(fi.userIds) != 1 {
			t.Errorf("Expected user IDs to have length 1, got %v", fi.userIds)
		}
	})
}

func TestFileInfo_removeOwner(t *testing.T) {
	t.Run("Remove Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")
		fi.userIds = []string{"userId"}

		// Action
		if err := fi.removeOwner("userId"); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Assert
		if len(fi.userIds) != 0 {
			t.Errorf("Expected user IDs to be empty, got %v", fi.userIds)
		}
	})

	t.Run("Remove Not Owner", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		if err := fi.removeOwner("userId"); err != ErrUserNotOwner {
			t.Errorf("Expected error to be ErrUserNotOwner, got %v", err)
		}
	})
}

func TestFileInfo_NewMissingFileInfo(t *testing.T) {
	// Setup
	fi := NewMissingFileInfo("fileId")

	// Assert
	if fi.fileId != "fileId" {
		t.Errorf("Expected file ID to be 'fileId', got %s", fi.fileId)
	}
	if len(fi.userIds) != 0 {
		t.Errorf("Expected user IDs to be empty, got %v", fi.userIds)
	}
	if fi.branchName != "" {
		t.Errorf("Expected branch name to be empty, got %s", fi.branchName)
	}
	if fi.claimMode != ClaimMode_Unclaimed {
		t.Errorf("Expected claim mode to be 0, got %d", fi.claimMode)
	}
	if fi.rejectReason != RejectReason_Missing {
		t.Errorf("Expected reject reason to be 5, got %d", fi.rejectReason)
	}
}

func TestFileInfo_UpgradeMissingToNew(t *testing.T) {
	t.Run("Upgrade Missing to New", func(t *testing.T) {
		// Setup
		fi := NewMissingFileInfo("fileId")

		// Action
		err := fi.UpgradeMissingToNew("fileHash", "branchName")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if fi.fileHash != "fileHash" {
			t.Errorf("Expected file hash to be 'fileHash', got %s", fi.fileHash)
		}

		if fi.branchName != "branchName" {
			t.Errorf("Expected branch name to be 'branchName', got %s", fi.branchName)
		}

		if len(fi.userIds) != 0 {
			t.Errorf("Expected user IDs to be empty, got %v", fi.userIds)
		}

		if fi.claimMode != ClaimMode_Unclaimed {
			t.Errorf("Expected claim mode to be 0, got %d", fi.claimMode)
		}

		if fi.rejectReason != RejectReason_None {
			t.Errorf("Expected reject reason to be 0, got %d", fi.rejectReason)
		}
	})

	t.Run("Upgrade New to New", func(t *testing.T) {
		// Setup
		fi := NewFileInfo("fileId", "fileHash", "branchName")

		// Action
		err := fi.UpgradeMissingToNew("fileHash", "branchName")
		if err != ErrorFileNotMissing {
			t.Errorf("Expected error to be ErrorFileNotMissing, got %v", err)
		}
	})
}
