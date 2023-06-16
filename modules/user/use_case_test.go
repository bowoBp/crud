package user

import (
	"crud/entity"
	"crud/repository"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_useCaseUser_GetUserById(t *testing.T) {
	type fields struct {
		userRepo repository.UserInterfaceRepo
	}
	type args struct {
		id uint
	}
	mockUserRepo := repository.NewMockUserInterface(t)

	defaultArg := args{
		id: uint(1),
	}

	//gagal get user
	mockUserRepo.On("GetUserById", defaultArg.id).
		Return(&entity.User{}, errors.New("failed get user")).Times(1)

	//berhasil get
	user := entity.User{ID: 1}
	mockUserRepo.On("GetUserById", defaultArg.id).
		Return(user, nil).Times(1)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name: "berhasil get",
			fields: fields{
				userRepo: mockUserRepo,
			},
			args: defaultArg,
			want: user,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Nil(t, err, i)
			},
		},
		{
			name: "error get",
			fields: fields{
				userRepo: mockUserRepo,
			},
			args: defaultArg,
			want: entity.User{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Nil(t, err, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseUser{
				userRepo: tt.fields.userRepo,
			}
			got, err := uc.GetUserById(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetUserById(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetUserById(%v)", tt.args.id)

		})
	}
}
