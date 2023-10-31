package models

import (
	"errors"
	"strings"
	"time"

	"github.com/upper/db/v4"
)

var (
	ErrDuplicateTitle = errors.New("title already exists in database")
	ErrDuplicateVotes = errors.New("you have already voted")

	queryTemplate = `
	SELECT COUNT(*) OVER() AS total_records, pq.*, u.name as uname FROM (
	    SELECT p.id, p.title, p.url, p.created_at, p.user_id as uid, COUNT(c.post_id) as comment_count, count(v.post_id) as votes
		FROM posts p 
		LEFT JOIN comments c ON p.id = c.post_id 
	    LEFT JOIN votes v ON p.id = v.post_id
	 	#where#
		GROUP BY p.id
		#orderby#
		) AS pq
	LEFT JOIN users u ON u.id = uid
	#limit#
	`
)

// Post holds properties for a post
type Post struct {
	ID           int       `db:"id,omitempty"`
	Title        string    `db:"title"`
	Url          string    `db:"url"`
	Username     string    `db:"username,omitempty"`
	CommentCount int       `db:"comment_count,omitempty"`
	TotalRecords int       `db:"total_records,omitempty"`
	UserID       int       `db:"user_id"`
	CreatedAt    time.Time `db:"created_at"`
	Votes        int       `db:"votes"`
}

type PostsModel struct {
	db db.Session
}

func (m PostsModel) Table() string {
	return "posts"
}

// Get gets a post by id from the database
func (pm PostsModel) Get(id int) (*Post, error) {
	var post Post
	query := strings.Replace(queryTemplate, "#where#", "WHERE p.id = $1", 1)
	query = strings.Replace(query, "#orderby#", "", 1)
	query = strings.Replace(query, "#limit#", "", 1)

	rows, err := pm.db.SQL().Query(query, id)
	if err != nil {
		return nil, err
	}

	iter := pm.db.SQL().NewIterator(rows)
	err = iter.One(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
