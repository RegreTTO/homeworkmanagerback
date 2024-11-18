package subject

import "homeworktodolist/internal/entity"

type subject struct {
	SubjectId   entity.SubjectID `json:"subject_id"`
	GroupId     entity.GroupID   `json:"group_id"`
	SubjectName string           `json:"subject_name"`
}

func (s subject) toSubject() entity.Subject {
	return entity.Subject{
		SubjectId:   s.SubjectId,
		GroupId:     s.GroupId,
		SubjectName: s.SubjectName,
	}
}

func toSubjects(subjects []subject) []entity.Subject {
	res := make([]entity.Subject, len(subjects))
	for i, g := range subjects {
		res[i] = g.toSubject()
	}
	return res
}
