package store

import (
	"forRoma/pkg/config"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	DB                    *sqlx.DB
	Redis                 *redis.Client
	userRepository        *UserRepository
	commentRepository     *CommentRepository
	articleRepository     *ArticleRepository
}

func (store *Store) UserRepository() *UserRepository {
	if store.userRepository != nil {
		return store.userRepository
	}

	store.userRepository = &UserRepository{store: store}

	return store.userRepository
}

func (store *Store) ArticleRepository() *ArticleRepository {
	if store.articleRepository != nil {
		return store.articleRepository
	}

	store.articleRepository = &ArticleRepository{store: store}

	return store.articleRepository
}

func (store *Store) CommentRepository() *CommentRepository {
	if store.commentRepository != nil {
		return store.commentRepository
	}

	store.commentRepository = &CommentRepository{store: store}

	return store.commentRepository
}

func Open(config *config.Config) (*Store, error) {
	sqlDB, err := sqlx.Open("postgres", config.DBConnection)
	if err != nil {
		return nil, err
	}

	r := redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Addr,
		Password: "",
		DB:       config.RedisConfig.DB,
	})

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return &Store{DB: sqlDB, Redis: r}, nil
}
