// Package comments provides comment repository abstractions.
package comments

import (
	"context"

	"commentTree/internal/domain"
)

// Repository defines the persistence operations required by the comments service.
type Repository interface {
	Create(ctx context.Context, comment *domain.Comment) (int, error)
	Delete(ctx context.Context, id int) error
	GetRootComments(ctx context.Context, search *string, limit, offset int) ([]*domain.Comment, error)
	GetChildComments(ctx context.Context, parentID int) ([]*domain.Comment, error)
	GetAllComments(ctx context.Context) ([]domain.CommentNode, error)
}
