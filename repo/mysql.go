package content

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

	"project-art-museum/business"
	"project-art-museum/business/content"
)

//MySQLRepository The implementation of content.Repository object
type MySQLRepository struct {
	db *sql.DB
}

//NewMySQLRepository Generate mongo DB content repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db,
	}
}

//FindContentByID Find content based on given ID. Its return nil if not found
func (repo *MySQLRepository) FindContentByID(ID string) (*content.Content, error) {
	var content content.Content

	selectQuery := `SELECT id, name, description, biography, birth_year, deadth_year, version, 
	id_artwork, accession_number, title_artworks, tombstone, url
	FROM creator LEFT JOIN artworks GROUP BY id ON id.creator = artworks.id_artwork`

	//var tags string
	err := repo.db.
		QueryRow(selectQuery, ID).
		Scan(
			&content.ID, &content.Name, &content.Nationality, &content.Description,
			&content.Biography, &content.Birth_year, &content.Death_year, &content.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	//content.Tags = constructTagArray(tags)

	return &content, nil
}

//InsertContent Insert new content into database. Its return content id if success
func (repo *MySQLRepository) InsertContent(content content.Content) (ID string, err error) {
	var contentID int
	ctx := context.Background()

	tx, err := repo.db.BeginTx(ctx, nil)

	if err != nil {
		return ID, err
	}

	contentQuery := `INSERT INTO content (id, name, description, biography, birth_year, deadth_year, version, id_artwork, accession_number, title_artworks, tombstone, url)`

	stmt, err := tx.Prepare(contentQuery)

	if err != nil {
		log.Fatal(err)
		return ID, err
	}

	err = stmt.QueryRow(content.Name,
		content.Nationality,
		content.Description,
		content.Biography,
		content.Birth_year,
		content.Death_year).Scan(&contentID)

	if err != nil {
		return ID, err
	}

	if err != nil {
		tx.Rollback()
		return ID, err
	}

	ID = strconv.Itoa(contentID)

	err = tx.Commit()

	if err != nil {
		return ID, err
	}

	return ID, err
}

//UpdateContent Update existing content in database
func (repo *MySQLRepository) UpdateContent(content content.Content, currentVersion int) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	contentInsertQuery := `UPDATE content 
		SET
			name = $1,
			description = $2,
			modified_at = NOW(),
			modified_by = $3,
			version = $4
		WHERE id = $5 AND version = $6`

	res, err := tx.Exec(contentInsertQuery,
		content.Name,
		content.Nationality,
		content.Description,
		content.Biography,
		content.Birth_year,
		content.Death_year,
		content.ID,
		currentVersion,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		tx.Rollback()
		return err
	}

	if affected == 0 {
		tx.Rollback()
		return business.ErrZeroAffected
	}

	//TODO: maybe better if we only delete the record that we need to delete
	//add logic slice to find which deleted and which want to added
	tagDeleteQuery := "DELETE FROM content_tag WHERE content_id = $1"
	_, err = tx.Exec(tagDeleteQuery, content.ID)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func constructTagArray(tags string) []string {
	if tags == "" {
		return make([]string, 0)
	}

	return strings.Split(tags, ",")
}
