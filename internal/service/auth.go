package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-user/internal/repository/entity"
	"github.com/smhdhsn/restaurant-user/internal/service/dto"
	"github.com/smhdhsn/restaurant-user/pkg/encryption"

	repositoryContract "github.com/smhdhsn/restaurant-user/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-user/internal/service/contract"
)

// AuthService contains repositories that will be used within this service.
type AuthService struct {
	uRepo repositoryContract.UserRepository
}

// NewAuthService creates a user's auth service with it's dependencies.
func NewAuthService(uRepo repositoryContract.UserRepository) serviceContract.UserAuthService {
	return &AuthService{
		uRepo: uRepo,
	}
}

// FindBy is responsible for fetching user's details from database.
func (s *AuthService) FindBy(uDTO *dto.User) (*dto.User, error) {
	uEntity := singleUserDTOToEntity(uDTO)

	uEntity, err := s.uRepo.FindBy(uEntity)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrRecordNotFound) {
			return nil, serviceContract.ErrRecordNotFound
		}

		return nil, errors.Wrap(err, "error on calling findby on user repository")
	}

	uDTO = singleUserEntityToDTO(uEntity)

	return uDTO, nil
}

// Store is responsible for storing user data inside database.
func (s *AuthService) Store(uDTO *dto.User) (*dto.User, error) {
	uEntity := singleUserDTOToEntity(uDTO)

	uEntity, err := s.uRepo.Store(uEntity)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrDuplicateEntry) {
			return nil, serviceContract.ErrDuplicateEntry
		}

		return nil, errors.Wrap(err, "error on calling store on user repository")
	}

	uDTO = singleUserEntityToDTO(uEntity)

	return uDTO, nil
}

// singleUserDTOToEntity is responsible for transforming a user dto into user entity struct.
func singleUserDTOToEntity(uDTO *dto.User) *entity.User {
	return &entity.User{
		ID:        uDTO.ID,
		FirstName: uDTO.FirstName,
		LastName:  uDTO.LastName,
		Email:     uDTO.Email,
		Password:  encryption.EncodeMD5(uDTO.Password),
		CreatedAt: uDTO.CreatedAt,
		UpdatedAt: uDTO.UpdatedAt,
	}
}

// singleUserEntityToDTO is responsible for transforming a user entity into user dto struct.
func singleUserEntityToDTO(uEntity *entity.User) *dto.User {
	return &dto.User{
		ID:        uEntity.ID,
		FirstName: uEntity.FirstName,
		LastName:  uEntity.LastName,
		Email:     uEntity.Email,
		Password:  uEntity.Password,
		CreatedAt: uEntity.CreatedAt,
		UpdatedAt: uEntity.UpdatedAt,
	}
}
