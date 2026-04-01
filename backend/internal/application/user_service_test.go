package application

import (
	"context"
	"errors"
	"testing"

	"github.com/myorg/myapp/backend/internal/domain"
)

type mockUserRepository struct {
	users map[string]*domain.User
}

func newMockRepo() *mockUserRepository {
	return &mockUserRepository{users: make(map[string]*domain.User)}
}

func (m *mockUserRepository) Create(_ context.Context, user *domain.User) error {
	m.users[user.Email] = user
	return nil
}

func (m *mockUserRepository) GetByEmail(_ context.Context, email string) (*domain.User, error) {
	user, ok := m.users[email]
	if !ok {
		return nil, nil
	}
	return user, nil
}

func (m *mockUserRepository) GetByID(_ context.Context, id domain.UserID) (*domain.User, error) {
	for _, u := range m.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

type mockPasswordHasher struct{}

func (m *mockPasswordHasher) Hash(password string) (string, error) {
	return "hashed_" + password, nil
}

func TestRegister_Success(t *testing.T) {
	svc := NewUserService(newMockRepo(), &mockPasswordHasher{})

	user, err := svc.Register(context.Background(), "test@example.com", "password123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", user.Email)
	}
	if user.Password != "hashed_password123" {
		t.Errorf("expected hashed password, got %s", user.Password)
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	svc := NewUserService(newMockRepo(), &mockPasswordHasher{})

	_, err := svc.Register(context.Background(), "test@example.com", "password123")
	if err != nil {
		t.Fatalf("first registration failed: %v", err)
	}

	_, err = svc.Register(context.Background(), "test@example.com", "password456")
	if err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}

	var alreadyExists *domain.UserAlreadyExistsError
	if !errors.As(err, &alreadyExists) {
		t.Errorf("expected UserAlreadyExistsError, got %T: %v", err, err)
	}
}

func TestRegister_EmptyEmail(t *testing.T) {
	svc := NewUserService(newMockRepo(), &mockPasswordHasher{})

	_, err := svc.Register(context.Background(), "", "password123")
	if err == nil {
		t.Fatal("expected error for empty email, got nil")
	}
	if !errors.Is(err, domain.ErrEmailRequired) {
		t.Errorf("expected ErrEmailRequired, got: %v", err)
	}
}

func TestRegister_EmptyPassword(t *testing.T) {
	svc := NewUserService(newMockRepo(), &mockPasswordHasher{})

	_, err := svc.Register(context.Background(), "test@example.com", "")
	if err == nil {
		t.Fatal("expected error for empty password, got nil")
	}
	if !errors.Is(err, domain.ErrPasswordRequired) {
		t.Errorf("expected ErrPasswordRequired, got: %v", err)
	}
}

func TestRegister_RepositoryError(t *testing.T) {
	repo := &failingRepo{err: errors.New("connection lost")}
	svc := NewUserService(repo, &mockPasswordHasher{})

	_, err := svc.Register(context.Background(), "test@example.com", "password123")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// failingRepo always returns an error from GetByEmail.
type failingRepo struct {
	err error
}

func (f *failingRepo) Create(context.Context, *domain.User) error {
	return f.err
}

func (f *failingRepo) GetByEmail(context.Context, string) (*domain.User, error) {
	return nil, f.err
}

func (f *failingRepo) GetByID(context.Context, domain.UserID) (*domain.User, error) {
	return nil, f.err
}
