package show

import (
	"biback/app/models"
	"biback/app/repository"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type dbShowRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewShowRepository(Conn *sql.DB) repository.Repository {
	return &dbShowRepository{Conn}
}

func (m *dbShowRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Show, error) {

	rows, err := m.Conn.QueryContext(ctx, query)
	//rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		//logrus.Error(err)
		//fmt.Printf(err)
		fmt.Printf("Error 1.2")
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			//logrus.Error(err)
			//fmt.Printf(err)
			fmt.Printf("Error 1.3")
		}
	}()

	result := make([]*models.Show, 0)
	for rows.Next() {
		t := new(models.Show)
		//authorID := int64(0)
		err = rows.Scan(
			&t.Id,
			&t.Name,
		)

		if err != nil {
			//logrus.Error(err)
			//fmt.Printf(err)
			fmt.Printf("Error 1.4")
			return nil, err
		}
		/*t.Author = models.Author{
			ID: authorID,
		}*/
		result = append(result, t)
	}

	return result, nil
}

func (m *dbShowRepository) Fetch(ctx context.Context, cursor string, num int64) ([]*models.Show, string, error) {
	query := `SELECT ShowID, ShowNAME FROM shows`
	fmt.Printf(cursor)
	decodedCursor, err := DecodeCursor(cursor)
	if err != nil && cursor != "" {
		fmt.Printf("Error 1")
		return nil, "", nil
	}

	res, err := m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		fmt.Printf("Error 2")
		return nil, "", err
	}

	nextCursor := ""
	if len(res) == int(num) {
		nextCursor = EncodeCursor(time.Time{})
		//nextCursor = EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return res, nextCursor, err

}

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
