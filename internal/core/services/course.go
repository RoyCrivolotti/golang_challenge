package services

import (
	"context"
	"golangchallenge/internal/core/domain"
	"golangchallenge/internal/core/ports"
)

type courseService struct {
	courseRepository ports.ICourseRepository
}

func NewCourseService(courseRepository ports.ICourseRepository) ports.ICourseService {
	return &courseService{
		courseRepository: courseRepository,
	}
}

func (c courseService) SortCourses(ctx context.Context, userCourseData domain.UserCourseData) []string {
	courseTree := make(map[string]*Node)

	//Iterate over the input entries, create a node for every course and attach its parent and children, should it have them
	for _, entry := range userCourseData.Courses {
		var parent *Node
		var child *Node

		if node, exists := courseTree[entry.Dependency]; exists == false {
			parent = &Node{CourseName: entry.Dependency}
			courseTree[entry.Dependency] = parent
		} else {
			parent = node
		}

		if node, exists := courseTree[entry.Name]; exists == false {
			child = &Node{CourseName: entry.Name}
			courseTree[entry.Name] = child
		} else {
			child = node
		}

		parent.Children = append(parent.Children, child)
		child.Parent = parent
	}

	var root *Node = nil
	var sortedCourses []string
	var queue []*Node

	//Get first element from map and traverse the parents until it is nil, and set the last one as root
	for _, node := range courseTree {
		root = node
		break
	}

	for root.Parent != nil {
		root = root.Parent
	}

	//Save root in results and append root's children, in order, to queue
	sortedCourses = append(sortedCourses, root.CourseName)
	queue = append(queue, root.Children...)

	//Append the node's children to the end of the queue until there are no more elements in it
	for len(queue) > 0 {
		//Get the first element from the queue, save it to a local variable and delete it from the queue
		currNode, rest := queue[0], queue[1:]
		queue = rest

		//Add the current node's children to the end of the queue
		queue = append(queue, currNode.Children...)
		sortedCourses = append(sortedCourses, currNode.CourseName)
	}

	return sortedCourses
}

type Node struct {
	CourseName string
	Parent     *Node
	Children   []*Node
}
