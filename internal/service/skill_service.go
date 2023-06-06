package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type SkillServiceContract interface {
	FindAll(ctx context.Context) ([]*entity.Skill, error)
	FindAllByProjectIDs(ctx context.Context, projectIDs []int64) ([]*entity.SkillBelongsTo, error)
	FindAllByWorkExperienceIDs(ctx context.Context, workExperienceIDs []int64) ([]*entity.SkillBelongsTo, error)
	FindAllByCategorySkillIDs(ctx context.Context, categorySkillIDs []int64) ([]*entity.Skill, error)
	FindByID(ctx context.Context, id int64) (*entity.Skill, error)
	Create(ctx context.Context, skill *entity.Skill) (*entity.Skill, error)
	Update(ctx context.Context, skill *entity.Skill) (*entity.Skill, error)
	Delete(ctx context.Context, id int64) error
}

type SkillService struct {
	SkillRepository repository.SkillRepositoryContract
}

func NewSkillService(skillRepository repository.SkillRepositoryContract) SkillServiceContract {
	return &SkillService{SkillRepository: skillRepository}
}

func (service *SkillService) FindAll(ctx context.Context) ([]*entity.Skill, error) {
	skills, err := service.SkillRepository.FindAll(ctx)
	if err != nil {
		log.Println("[SkillService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return skills, nil
}

func (service *SkillService) FindAllByProjectIDs(ctx context.Context, projectIDs []int64) ([]*entity.SkillBelongsTo, error) {
	skills, err := service.SkillRepository.FindAllByProjectIDs(ctx, projectIDs)
	if err != nil {
		log.Println("[SkillService][FindAllByProjectIDs] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return skills, nil
}

func (service *SkillService) FindAllByWorkExperienceIDs(ctx context.Context, workExperienceIDs []int64) ([]*entity.SkillBelongsTo, error) {
	skills, err := service.SkillRepository.FindAllByWorkExperienceIDs(ctx, workExperienceIDs)
	if err != nil {
		log.Println("[SkillService][FindAllByWorkExperienceIDs] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return skills, nil
}

func (service *SkillService) FindAllByCategorySkillIDs(ctx context.Context, categorySkillIDs []int64) ([]*entity.Skill, error) {
	skills, err := service.SkillRepository.FindAllByCategorySkillIDs(ctx, categorySkillIDs)
	if err != nil {
		log.Println("[SkillService][FindAllByCategorySkillIDs] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return skills, nil
}

func (service *SkillService) FindByID(ctx context.Context, id int64) (*entity.Skill, error) {
	skill, err := service.SkillRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[SkillService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return skill, nil
}

func (service *SkillService) Create(ctx context.Context, skill *entity.Skill) (*entity.Skill, error) {
	skill, err := service.SkillRepository.Create(ctx, skill)
	if err != nil {
		log.Println("[SkillService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return skill, nil
}

func (service *SkillService) Update(ctx context.Context, skill *entity.Skill) (*entity.Skill, error) {
	_, err := service.SkillRepository.FindByID(ctx, skill.ID)
	if err != nil {
		log.Println("[SkillService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.SkillRepository.Update(ctx, skill)
	if err != nil {
		log.Println("[SkillService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	skillUpdated, err := service.SkillRepository.FindByID(ctx, skill.ID)
	if err != nil {
		log.Println("[SkillService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return skillUpdated, nil
}

func (service *SkillService) Delete(ctx context.Context, id int64) error {
	_, err := service.SkillRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[SkillService][FindByID] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.SkillRepository.Delete(ctx, id)
	if err != nil {
		log.Println("[SkillService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
