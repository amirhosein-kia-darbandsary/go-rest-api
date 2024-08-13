package db

import (
	"context"
	"database/sql"
	"fmt"

	Icomment "github.com/amirhosein-kia-darbandsary/project/internal/comment"
	"github.com/docker/distribution/uuid"
)

type commentRow struct {
	id     string
	slug   sql.NullString
	body   sql.NullString
	author sql.NullString
}

func convertcmtRowTocomment(cmtRow commentRow) Icomment.Comment {
	return Icomment.Comment{
		ID:     cmtRow.id,
		Slug:   cmtRow.slug.String,
		Body:   cmtRow.body.String,
		Author: cmtRow.author.String,
	}

}
func (d *Database) GetComment(ctx context.Context, uuid string) (Icomment.Comment, error) {
	var cmtRow commentRow
	row := d.Client.QueryRowContext(ctx,
		`SELECT id, slug, body, author
		FROM comments
		WHERE id = $1`,
		uuid)

	err := row.Scan(&cmtRow.id, &cmtRow.slug, &cmtRow.body, &cmtRow.author)

	if err != nil {
		fmt.Println(err)
		return Icomment.Comment{}, err
	}
	comment := convertcmtRowTocomment(cmtRow)
	return comment, nil
}

func (d *Database) PostComment(ctx context.Context, cmt Icomment.Comment) error {
	cmt.ID = uuid.Generate().String()
	postRow := commentRow{
		id:     cmt.ID,
		slug:   sql.NullString{String: cmt.Slug, Valid: true},
		body:   sql.NullString{String: cmt.Body, Valid: true},
		author: sql.NullString{String: cmt.Author, Valid: true},
	}

	query := `INSERT INTO comments (id, slug, body, author)
    VALUES ($1, $2, $3, $4)`

	// Execute the query using Exec
	_, err := d.Client.ExecContext(ctx, query,
		postRow.id,
		postRow.slug,
		postRow.body,
		postRow.author,
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UpdateComment(ctx context.Context, id string, cmt Icomment.Comment) error {

	postRow := commentRow{
		id:     id,
		slug:   sql.NullString{String: cmt.Slug, Valid: true},
		body:   sql.NullString{String: cmt.Body, Valid: true},
		author: sql.NullString{String: cmt.Author, Valid: true},
	}
	query := `UPDATE comments SET
				slug = $1,
				author = $2,
				body = $3
				WHERE id = $4;`

	_, err := d.Client.Exec(query,
		postRow.slug,
		postRow.author,
		postRow.body,
		postRow.id)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	query := `
        DELETE FROM comments
        WHERE id = $1`

	_, err := d.Client.ExecContext(ctx, query, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
