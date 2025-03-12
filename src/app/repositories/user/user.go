package user

import (
	"errors"
	"log"
	dto "todo_list/src/app/dto/user"
	"todo_list/src/infra/helper"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository mendefinisikan metode yang harus diimplementasikan
// untuk operasi pengguna seperti registrasi, login, dan refresh token.
type UserRepository interface {
	RegisterUser(data *dto.RegisterUserReqDTO) (*dto.RegisterUserRespDTO, error)
	SignIn(data *dto.SignInReqDTO) (*dto.RegisterUserRespDTO, error)
	RefreshToken(data *dto.RefreshTokenReq) (*dto.RefreshTokenResp, error)
}

// Query SQL untuk berbagai operasi database
const (
	RegisterUser = `INSERT INTO public.users (name, email, password)
		VALUES ($1, $2, $3) RETURNING id, name, email;`

	SaveRefreshToken = `UPDATE public.users SET refresh_token = $1 WHERE id = $2;`

	SignIn = `SELECT id, name, email, password FROM public.users WHERE email = $1;`

	RefreshToken = `SELECT id, name, email FROM public.users WHERE refresh_token = $1;`
)

// Struct untuk menyimpan statement yang telah diprepare
var statement PreparedStatement

type PreparedStatement struct {
	refreshToken     *sqlx.Stmt
	saveRefreshToken *sqlx.Stmt
	signIn           *sqlx.Stmt
}

// UserRepo menyimpan koneksi database
// dan menyediakan metode untuk mengelola pengguna.
type UserRepo struct {
	Connection *sqlx.DB
}

// NewUserRepository menginisialisasi UserRepo dan menyiapkan prepared statement
func NewUserRepository(db *sqlx.DB) UserRepository {
	repo := &UserRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

// Preparex menyiapkan statement SQL yang telah diprepare
func (p *UserRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

// InitPreparedStatement menginisialisasi prepared statement untuk query tertentu
func InitPreparedStatement(m *UserRepo) {
	statement = PreparedStatement{
		refreshToken:     m.Preparex(RefreshToken),
		signIn:           m.Preparex(SignIn),
		saveRefreshToken: m.Preparex(SaveRefreshToken),
	}
}

// RegisterUser menangani proses registrasi pengguna baru
func (repo *UserRepo) RegisterUser(data *dto.RegisterUserReqDTO) (*dto.RegisterUserRespDTO, error) {
	// Hash password sebelum disimpan ke database
	hashedPassword, err := hashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	// Mulai transaksi database
	tx, err := repo.Connection.Beginx()
	if err != nil {
		log.Println("Failed to start transaction:", err)
		return nil, err
	}

	// Pastikan transaksi rollback jika terjadi error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Siapkan statement untuk insert user
	stmtRegisterUser, err := tx.Preparex(RegisterUser)
	if err != nil {
		log.Println("Failed to prepare registerUser statement:", err)
		return nil, err
	}
	defer stmtRegisterUser.Close()

	// Siapkan statement untuk menyimpan refresh token
	stmtSaveRefreshToken, err := tx.Preparex(SaveRefreshToken)
	if err != nil {
		log.Println("Failed to prepare saveRefreshToken statement:", err)
		return nil, err
	}
	defer stmtSaveRefreshToken.Close()

	// Insert user ke database
	var resp dto.RegisterUserRespDTO
	err = stmtRegisterUser.QueryRowx(data.Name, data.Email, hashedPassword).Scan(
		&resp.ID, &resp.Name, &resp.Email)
	if err != nil {
		log.Println("Failed to insert user:", err)
		return nil, err
	}

	// Generate token
	resp.RefreshToken, err = helper.GenerateToken(resp.ID, resp.Email, true)
	if err != nil {
		return nil, err
	}

	resp.Token, err = helper.GenerateToken(resp.ID, resp.Email, false)
	if err != nil {
		return nil, err
	}

	// Simpan refresh token ke database
	_, err = stmtSaveRefreshToken.Exec(resp.RefreshToken, resp.ID)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// SignIn menangani autentikasi user berdasarkan email dan password
func (repo *UserRepo) SignIn(data *dto.SignInReqDTO) (*dto.RegisterUserRespDTO, error) {
	var user dto.SignInModelDTO

	// Cek user berdasarkan email
	err := statement.signIn.Get(&user, data.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid account")
	}

	// Verifikasi password
	err = verifyPassword(user.Password, data.Password)
	if err != nil {
		return nil, errors.New("invalid account")
	}

	// Generate token
	resp := &dto.RegisterUserRespDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	resp.RefreshToken, err = helper.GenerateToken(resp.ID, resp.Email, true)
	if err != nil {
		return nil, err
	}

	_, err = statement.saveRefreshToken.Exec(resp.RefreshToken, resp.ID)

	if err != nil {
		return nil, err
	}

	resp.Token, err = helper.GenerateToken(resp.ID, resp.Email, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RefreshToken menangani pembuatan token baru berdasarkan refresh token
func (repo *UserRepo) RefreshToken(data *dto.RefreshTokenReq) (*dto.RefreshTokenResp, error) {
	var user dto.SignInModelDTO

	// Cek user berdasarkan refresh token
	err := statement.refreshToken.Get(&user, data.RefreshToken)
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid account")
	}

	// Generate token baru
	resp := dto.RefreshTokenResp{}
	resp.Token, err = helper.GenerateToken(user.ID, user.Email, false)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// hashPassword mengenkripsi password sebelum disimpan ke database
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// verifyPassword membandingkan password yang dimasukkan dengan hash di database
func verifyPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
