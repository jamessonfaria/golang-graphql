package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"math/rand"
	"context"
	"fmt"

	"github.com/jamessonfaria/golang-graphql/graph/generated"
	"github.com/jamessonfaria/golang-graphql/graph/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := model.Category {
		ID: fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
		Description: &input.Description,
	}
	r.Categories = append(r.Categories, &category)
	return &category, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	var category *model.Category
	for _, v := range(r.Categories) {
		if v.ID == input.CategoryID {
			category = v
		}
	}

	course := model.Course {
		ID: fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
		Description: &input.Description,
		Category: category,
	}
	r.Courses = append(r.Courses, &course)
	return &course, nil
}

// CreateChapter is the resolver for the createChapter field.
func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	var course *model.Course

	for _, v := range r.Courses {
		if v.ID == input.CourseID {
			course = v
		}
	}

	chapter := &model.Chapter {
		ID: fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
		Course: course,
	}
	r.Chapters = append(r.Chapters, chapter)

	return chapter, nil

}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

// Chapters is the resolver for the chapters field.
func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
