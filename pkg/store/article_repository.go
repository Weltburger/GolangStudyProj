package store

import (
	"context"
	"database/sql"
	"errors"
	"forRoma/pkg/models"
	"golang.org/x/sync/errgroup"
)

type ArticleRepository struct {
	store *Store
}

func (articleRepository *ArticleRepository) AllArticles(ctx context.Context, take int64, skip int64) ([]*models.Article, int64, error) {
	articles := make([]*models.Article, 0)
	total := int64(0)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		rows, err := articleRepository.store.DB.QueryxContext(ctx, `SELECT uuid, 
       title, 
       text, 
       create_at, 
       updated_at, 
       deleted_at 
FROM "public"."articles" aes offset $1 limit $2`, skip, take)

		if err != nil {
			return err
		}

		defer rows.Close()

		for rows.Next() {
			article := &models.Article{}
			if err := rows.StructScan(article); err != nil {
				return err
			}
			articles = append(articles, article)
		}
		return nil
	})

	eg.Go(func() error {
		if err := articleRepository.store.DB.QueryRowxContext(ctx, `SELECT count(1)
FROM "public"."articles"`).Scan(&total); err != nil {
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func (articleRepository *ArticleRepository) ArticleByUUID(ctx context.Context, article *models.Article) (*models.Article, error) {
	if err := articleRepository.store.DB.QueryRowxContext(ctx, `SELECT uuid, 
       title, 
       text, 
       create_at, 
       updated_at, 
       deleted_at 
FROM "public"."articles" aes 
WHERE aes.uuid = $1`, article.UUID).StructScan(article); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("there is no articles")
		}
		return nil, err
	}

	return article, nil
}

func (articleRepository *ArticleRepository) InsertArticle(ctx context.Context, article *models.Article) (*models.Article, error) {
	if err := articleRepository.store.DB.QueryRowxContext(ctx, `insert into "public"."articles"(user_uuid, title, text) 
VALUES($1, $2, $3) returning uuid`, article.User.UUID, article.Title, article.Text).Scan(&article.UUID); err != nil {
		return nil, err
	}

	return article, nil
}

func (articleRepository *ArticleRepository) DeleteArticle(ctx context.Context, article *models.Article) error {
	if _, err := articleRepository.store.DB.ExecContext(ctx, `update "public"."articles" SET deleted_at = now()
where uuid = $1`, article.UUID); err != nil {
		return err
	}

	return nil
}

func (articleRepository *ArticleRepository) InsertLike(ctx context.Context, like *models.LikeArticle) (*models.LikeArticle, error) {
	if err := articleRepository.store.DB.QueryRowxContext(ctx, `insert into "public"."article_likes"(article_uuid, user_uuid) 
VALUES($1, $2) returning uuid`, like.User.UUID, like.Article.UUID).Scan(&like.UUID); err != nil {
		return nil, err
	}

	return like, nil
}

func (articleRepository *ArticleRepository) DeleteLike(ctx context.Context, like *models.LikeArticle) error {
	if err := articleRepository.store.DB.QueryRowxContext(ctx, `DELETE FROM "public"."article_likes"
WHERE uuid=$1`, like.UUID).Scan(&like.UUID); err != nil {
		return err
	}

	return nil
}

func (articleRepository *ArticleRepository) LikeExist(ctx context.Context, like *models.LikeArticle) (bool, error) {
	var isExist bool
	if err := articleRepository.store.DB.QueryRowxContext(ctx, `select true, uuid from "public"."article_likes"
where user_uuid=$1 and article_uuid=$2`, like.User.UUID, like.Article.UUID).Scan(&isExist, &like.UUID); err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return isExist, nil
}
