package util

import (
	"log"
	"testing"
)

type StudentData struct {
	Name        string         `json:"name"`
	Host        string         `json:"host"`
	Personality map[string]any `json:"personality"`
	Count       int            `json:"count"`
}

func (student *StudentData) LessThan(student2 Comparable) bool {
	return student.Count < student2.GetInteger()
}

func (student *StudentData) MoreThan(student2 Comparable) bool {
	return student.Count > student2.GetInteger()
}

func (student *StudentData) EqualTo(student2 Comparable) bool {
	return student.Count == student2.GetInteger()
}

func (student *StudentData) GetInteger() int {
	return student.Count
}

func TestAVLTree(t *testing.T) {
	vals := []StudentData{{
		Name:        "a",
		Host:        "a",
		Personality: nil,
		Count:       3,
	}, {
		Name:        "b",
		Host:        "b",
		Personality: nil,
		Count:       1,
	}, {
		Name:        "c",
		Host:        "c",
		Personality: nil,
		Count:       32,
	}, {
		Name:        "d",
		Host:        "d",
		Personality: nil,
		Count:       2,
	}, {
		Name:        "e",
		Host:        "e",
		Personality: nil,
		Count:       10,
	}}

	tree := AVLTree[*StudentData]{}
	for _, val := range vals {
		tree.Insert(&tree.Root, &val)
	}
	log.Println(tree.Size)
}
