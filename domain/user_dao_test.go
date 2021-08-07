package domain

import "testing"

func TestGetUser_userNotFound(t *testing.T) {
	var userId int64 = 0
	user, err := GetUser(userId)
	if user != nil {
		t.Errorf("GetUser(%d) user: got = %v, want = %v", userId, user, nil)
	}
	if err == nil {
		t.Fatalf("GetUser(%d) err: got = %v", userId, err)
	}
	if err.StatusCode != 404 {
		t.Errorf("GetUser(%d) err.StatusCode: got = %d, want = %d", userId, err.StatusCode, 404)
	}
}

func TestGetUser_userFound(t *testing.T) {
	var userId int64 = 123
	user, err := GetUser(userId)
	if user == nil {
		t.Fatalf("GetUser(%d) user: got = %v", userId, user)
	}
	if user.First != "James" {
		t.Errorf("GetUser(%d) user: got = %q, want = %q", userId, user.First, "James")
	}
	// ... User: assert other fields
	if err != nil {
		t.Fatalf("GetUser(%d) err: got = %v, want = %v", userId, err, nil)
	}
}
