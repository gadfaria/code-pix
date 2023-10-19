package model_test

import (
	"testing"

	"github.com/gadfaria/codepix/domain/model"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Nubank"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Gabriel"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "hello@gadfaria.com"
	pixKey, err := model.NewPixKey(key, kind, account)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, model.PixKeyStatusActive)

	kind = "cpf"
	_, err = model.NewPixKey(key, kind, account)
	require.Nil(t, err)

	_, err = model.NewPixKey(key, "nome", account)
	require.NotNil(t, err)
}
