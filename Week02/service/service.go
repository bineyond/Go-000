package service

import (
	"errors"
	"strconv"

	"github.com/bineyond/Go-000/Week02/dao"

	xerrors "github.com/pkg/errors"
)

var (
	ErrorParam = errors.New("Param error")
	ErrorConv  = errors.New("conversion faile")
)

// Get User
func GetUserByID(id string) (string, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return "", xerrors.Wrap(ErrorConv, "userId faild")
	}
	if userId <= 0 {
		return "", xerrors.Wrap(ErrorParam, "userId invalid")
	}
	return dao.GetUserByID(userId), nil
}
