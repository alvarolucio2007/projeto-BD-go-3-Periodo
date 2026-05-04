package grpcclient

import (
	"context"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserConexao) DoCreateUser(user *models.Usuario) (uint32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := u.User.Create(ctx, &proto.UsuarioCreateRequest{Username: user.Username, Password: user.Password, Role: user.Role})
	if err != nil {
		return 0, err
	}
	return res.Id, nil
}

func (u *UserConexao) DoReadUser(username string) ([]*models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := u.User.Read(ctx, &proto.UsuarioReadRequest{Nome: username})
	if err != nil {
		return nil, err
	}
	result := make([]*models.Usuario, 0, len(res.Usuarios))
	for _, u := range res.Usuarios {
		user := models.Usuario{
			ID:       u.Id,
			Username: u.Username,
			Password: u.Password,
			Role:     u.Role,
		}
		result = append(result, &user)
	}
	return result, nil
}

func (u *UserConexao) DoReadAllUser(ctx context.Context) ([]*models.Usuario, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res, err := u.User.ReadAll(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	result := make([]*models.Usuario, 0, len(res.Usuarios))
	for _, u := range res.Usuarios {
		user := models.Usuario{
			ID:       u.Id,
			Username: u.Username,
			Password: u.Password,
			Role:     u.Role,
		}
		result = append(result, &user)
	}
	return result, nil
}

func (u *UserConexao) DoUpdateUser(user *models.Usuario) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := u.User.Update(ctx, &proto.UsuarioUpdateRequest{Id: user.ID, Username: user.Username, Password: user.Password, Role: user.Role})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserConexao) DoDeleteUser(id uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := u.User.Delete(ctx, &proto.UsuarioDeleteRequest{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserConexao) DoAuth(username string, password string) (*models.AuthResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := u.User.Auth(ctx, &proto.UsuarioLoginRequest{Username: username, Password: password})
	if err != nil {
		return nil, err
	}
	return &models.AuthResult{Status: response.Status, Mensagem: response.Mensagem, Role: response.Role}, nil
}
