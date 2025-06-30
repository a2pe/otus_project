package repository_test

import (
	"fmt"
	"otus_project/internal/model"
	"otus_project/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSaveTask(t *testing.T) {
	tests := []struct {
		name     string
		item     interface{}
		wantErr  bool
		errCheck func(error) bool
	}{
		{
			name: "valid task",
			item: &model.Task{
				Title:   "Test task",
				Status:  "new",
				DueDate: time.Now().Add(24 * time.Hour),
			},
			wantErr: false,
		},
		{
			name:    "empty struct",
			item:    &model.Task{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.SaveItem(tt.item.(*model.Task))
			if tt.wantErr {
				fmt.Println(err)
				require.Error(t, err)
				if tt.errCheck != nil {
					require.True(t, tt.errCheck(err), "unexpected error: %v", err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSaveUser(t *testing.T) {
	tests := []struct {
		name     string
		item     interface{}
		wantErr  bool
		errCheck func(error) bool
	}{
		{
			name: "valid task",
			item: &model.User{
				Name: "Test user",
			},
			wantErr: false,
		},
		{
			name:    "empty struct",
			item:    &model.User{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.SaveItem(tt.item.(*model.User))
			if tt.wantErr {
				fmt.Println(err)
				require.Error(t, err)
				if tt.errCheck != nil {
					require.True(t, tt.errCheck(err), "unexpected error: %v", err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSaveTag(t *testing.T) {
	tests := []struct {
		name     string
		item     interface{}
		wantErr  bool
		errCheck func(error) bool
	}{
		{
			name: "valid tag",
			item: &model.Tag{
				Name: "Test tag",
			},
			wantErr: false,
		},
		{
			name:    "empty struct",
			item:    &model.Tag{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.SaveItem(tt.item.(*model.Tag))
			if tt.wantErr {
				fmt.Println(err)
				require.Error(t, err)
				if tt.errCheck != nil {
					require.True(t, tt.errCheck(err), "unexpected error: %v", err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSaveProject(t *testing.T) {
	tests := []struct {
		name     string
		item     interface{}
		wantErr  bool
		errCheck func(error) bool
	}{
		{
			name: "valid project",
			item: &model.Project{
				Name: "Test tag",
			},
			wantErr: false,
		},
		{
			name:    "empty struct",
			item:    &model.Project{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.SaveItem(tt.item.(*model.Project))
			if tt.wantErr {
				fmt.Println(err)
				require.Error(t, err)
				if tt.errCheck != nil {
					require.True(t, tt.errCheck(err), "unexpected error: %v", err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSaveReminder(t *testing.T) {
	tests := []struct {
		name     string
		item     interface{}
		wantErr  bool
		errCheck func(error) bool
	}{
		{
			name: "valid project",
			item: &model.Reminder{
				RemindAt: time.Now().Add(24 * time.Hour),
			},
			wantErr: false,
		},
		{
			name:    "empty struct",
			item:    &model.Reminder{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.SaveItem(tt.item.(*model.Reminder))
			if tt.wantErr {
				fmt.Println(err)
				require.Error(t, err)
				if tt.errCheck != nil {
					require.True(t, tt.errCheck(err), "unexpected error: %v", err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}
