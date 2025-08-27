package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
)

func UploadProblem(problem *problem_model.Problem) error {
	err := global.DB.Create(problem).Error
	if err != nil {
		return err
	}

	cppConstraint := problem_model.ProblemConstraint{
		ProblemID:   problem.ID,
		TimeLimit:   problem.Constraints.Cpp.TimeLimit,
		MemoryLimit: problem.Constraints.Cpp.MemoryLimit,
		Language:    "cpp",
	}

	if err = global.DB.Create(&cppConstraint).Error; err != nil {
		return err
	}

	pyConstraint := problem_model.ProblemConstraint{
		ProblemID:   problem.ID,
		TimeLimit:   problem.Constraints.Python.TimeLimit,
		MemoryLimit: problem.Constraints.Python.MemoryLimit,
		Language:    "python",
	}
	if err = global.DB.Create(&pyConstraint).Error; err != nil {
		return err
	}

	goConstraint := problem_model.ProblemConstraint{
		ProblemID:   problem.ID,
		TimeLimit:   problem.Constraints.Go.TimeLimit,
		MemoryLimit: problem.Constraints.Go.MemoryLimit,
		Language:    "go",
	}
	if err = global.DB.Create(&goConstraint).Error; err != nil {
		return err
	}

	for _, tag := range problem.Tags {
		tagData := problem_model.ProblemTag{
			ProblemID: problem.ID,
			Tag:       tag,
		}
		if err = global.DB.Create(&tagData).Error; err != nil {
			return nil
		}
	}

	for _, example := range problem.Examples {
		exampleData := problem_model.ProblemExample{
			ProblemID: problem.ID,
			Input:     example.Input,
			Output:    example.Output,
			IsPublic:  example.IsPublic,
		}
		if err = global.DB.Create(&exampleData).Error; err != nil {
			return nil
		}
	}

	return nil
}
