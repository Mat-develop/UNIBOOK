package repository

import (
	"database/sql"
	"errors"
	"v1/monorepo/post/model"
)

const (
	CreateQuery            = "INSERT INTO post (title, body, community_id, user_id ) values (?, ?, ?, ?)"
	FindByCommunityIdQuery = "SELECT p.*, u.nick from posts p inner join users u on u.id = p.user_id where p.community_id = ?"
	FindByUserIdQuery      = "SELECT p.*, u.nick from posts p inner join users u on u.id = p.user_id where p.user_id = ?"
)

/*
	-- TABLE CREATION SO I REMEMBER ---

CREATE TABLE post(

	    id uuid not null,
	    title varchar(50) not null,
	    body varchar(300) not null,

	    user_id uuid not null,
	    foreign key (user_id)
	    references users(id)
		ON DELETE CASCADE,

	    community_id uuid,
	    foreign key (community_id)
	    references communities(id)
	    ON DELETE CASCADE,

	    likes int default,
	    created_at timestamp default current_timestamp()
	    )
*/
type PostRepository interface {
	FindCommunityPosts(communityId uint64) ([]model.Post, error)
	FindUserPosts(userId uint64) ([]model.Post, error)
	FindPostByName() ([]model.Post, error)
	Create(userId uint64, postBody model.PostDTO) error
	Update() error
	Delete() error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) FindCommunityPosts(communityId uint64) ([]model.Post, error) {
	rows, err := p.db.Query(FindByCommunityIdQuery, communityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []model.Post{}

	if rows.Next() {
		post := model.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.UserId,
			&post.CommunityId,
			&post.Likes,
			&post.CreatedAt,
		)
		if err != nil {
			return []model.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (p *postRepository) FindUserPosts(userId uint64) ([]model.Post, error) {
	rows, err := p.db.Query(FindByUserIdQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []model.Post{}

	if rows.Next() {
		post := model.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.UserId,
			&post.CommunityId,
			&post.Likes,
			&post.CreatedAt,
		)
		if err != nil {
			return []model.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
func (p *postRepository) FindPostByName() ([]model.Post, error) {
	return []model.Post{}, errors.New("hello")
}

func (p *postRepository) Create(userId uint64, postBody model.PostDTO) error {
	statement, err := p.db.Prepare(CreateQuery)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(postBody.Title, postBody.Body, postBody.CommunityId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (p *postRepository) Update() error {
	return nil
}

func (p *postRepository) Delete() error {
	return nil
}
