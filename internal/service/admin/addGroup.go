package admin

import (
	"context"
	"homeworktodolist/internal/entity"
)

type AddGroup struct {
	Name     string
	Course   int8
	IcalLink string
}

func (s *Service) AddGroup(ctx context.Context, req AddGroup) error {
	group := req.toGroup()

	err := s.manager.Do(ctx, func(ctx context.Context) error {
		var err error
		group.GroupID, err = s.groupService.Create(ctx, group)
		if err != nil {
			return err
		}

		err = s.subjectService.UpdGroupSubjects(ctx, group)
		if err != nil {
			return err
		}

		err = s.classService.UpdGroupClasses(ctx, group)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *AddGroup) toGroup() entity.Group {
	return entity.Group{
		Name:     r.Name,
		Course:   r.Course,
		IcalLink: r.IcalLink,
	}
}
