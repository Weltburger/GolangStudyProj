package store

import (
	"context"
	"database/sql"
	"errors"
	"forRoma/pkg/models"
)

type CommentRepository struct {
	store *Store
}

func (commentRepository *CommentRepository) CommentByUUID(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	if err := commentRepository.store.DB.QueryRowxContext(ctx, `SELECT uuid, 
       text,
       create_at, 
       updated_at, 
       deleted_at 
FROM "public"."comments" cs 
WHERE cs.uuid = $1`, comment.UUID).StructScan(comment); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("there is no comments")
		}
		return nil, err
	}

	return comment, nil
}

func (commentRepository *CommentRepository) InsertComment(ctx context.Context, comment *models.Comment, article *models.Article) (*models.Comment, error) {
	if err := commentRepository.store.DB.QueryRowxContext(ctx, `insert into "public"."comments"(user_uuid, article_uuid, text) 
VALUES($1, $2, $3) returning uuid`, comment.User.UUID, article.UUID, comment.Text).Scan(&comment.UUID); err != nil {
		return nil, err
	}

	return comment, nil
}

func (commentRepository *CommentRepository) DeleteComment(ctx context.Context, comment *models.Comment) error {
	if err := commentRepository.store.DB.QueryRowxContext(ctx, `DELETE FROM "public"."comments"
WHERE uuid=$1`, comment.UUID).Scan(&comment.UUID); err != nil {
		return err
	}

	return nil
}

func (commentRepository *CommentRepository) InsertLike(ctx context.Context, like *models.LikeComment) (*models.LikeComment, error) {
	if err := commentRepository.store.DB.QueryRowxContext(ctx, `insert into "public"."comment_likes"(comment_uuid, user_uuid) 
VALUES($1, $2) returning uuid`, like.User.UUID, like.Comment.UUID).Scan(&like.UUID); err != nil {
		return nil, err
	}

	return like, nil
}

func (commentRepository *CommentRepository) DeleteLike(ctx context.Context, like *models.LikeComment) error {
	if err := commentRepository.store.DB.QueryRowxContext(ctx, `DELETE FROM "public"."comment_likes"
WHERE uuid=$1`, like.UUID).Scan(&like.UUID); err != nil {
		return err
	}

	return nil
}

func (commentRepository *CommentRepository) LikeExist(ctx context.Context, like *models.LikeComment) (bool, error) {
	var isExist bool
	if err := commentRepository.store.DB.QueryRowxContext(ctx, `select true, uuid from "public"."comment_likes"
where user_uuid=$1 and comment_uuid=$2`, like.User.UUID, like.Comment.UUID).Scan(&isExist, &like.UUID); err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return isExist, nil
}
