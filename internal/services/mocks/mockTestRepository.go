package mocks

import . "contest/internal/domain"

type MockTestRepository struct{}

func (r MockTestRepository) AddItem(taskID int, input string, expectedResult string, points int) error {
	panic("not implemented")
}
func (r MockTestRepository) DeleteItem(id int) error {
	panic("not implemented")
}
func (r MockTestRepository) UpdateItem(id int, newItem Test) error {
	panic("not implemented")
}
func (r MockTestRepository) GetTable() ([]Test, error) {
	panic("not implemented")
}
func (r MockTestRepository) FindItemByID(id int) (Test, error) {
	panic("not implemented")
}
func (r MockTestRepository) FindItemByCondition(condition func(item Test) bool) (Test, error) {
	panic("not implemented")
}

func (r MockTestRepository) FindItemsByCondition(condition func(item Test) bool) ([]Test, error) {
	return []Test{
		{ID: 1, TaskID: 1, Input: "1", ExpectedResult: "1", Points: 1},
		{ID: 2, TaskID: 1, Input: "2", ExpectedResult: "4", Points: 1},
		{ID: 2, TaskID: 1, Input: "3", ExpectedResult: "9", Points: 1},
	}, nil
}
