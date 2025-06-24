package service

import "github.com/labbs/zotion/pkg/models"

type groupService struct {
	groupRepository models.GroupRepository
}

// NewGroupService creates a new group service
func NewGroupService(groupRepository models.GroupRepository) *groupService {
	return &groupService{groupRepository: groupRepository}
}

// CreateGroup creates a new group
func (s *groupService) CreateGroup(group models.Group) (models.Group, error) {
	return s.groupRepository.Create(group)
}

// GetGroupById returns a group by id
func (s *groupService) GetGroupById(id string) (models.Group, error) {
	return s.groupRepository.GetById(id)
}

// GetGroupWithUsers returns a group with users
func (s *groupService) GetGroupWithUsers(id string) (models.Group, error) {
	return s.groupRepository.GetGroupWithUsers(id)
}

// GetAllGroups returns all groups
func (s *groupService) GetAllGroups() ([]models.Group, error) {
	return s.groupRepository.GetAll()
}

// UpdateGroup updates a group
func (s *groupService) UpdateGroup(group models.Group) (models.Group, error) {
	return s.groupRepository.Update(group)
}

// DeleteGroup deletes a group
func (s *groupService) DeleteGroup(id string) error {
	return s.groupRepository.Delete(id)
}

// GetAllGroupsWithUsers returns all groups with users
func (s *groupService) GetAllGroupsWithUsers() ([]models.Group, error) {
	return s.groupRepository.GetAllGroupsWithUsers()
}
