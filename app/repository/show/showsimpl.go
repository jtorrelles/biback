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

	rows, err := m.Conn.QueryContext(ctx, query, args...)

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
			&t.Active,
			&t.Category1,
			&t.Category2,
			&t.Category3,
			&t.Category4,
			&t.Category5,
			&t.Category6,
			&t.Category7,
			&t.Age,
			&t.WeeklyNut,
			&t.NumberOfCast,
			&t.NumberOfMusicians,
			&t.NumberOfStageHands,
			&t.NumberOfTrucks,
			&t.Notes,
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
	query := `SELECT ShowID, ShowNAME, ShowACTIVE, CategoryID_1, CategoryID_2, CategoryID_3, CategoryID_4, CategoryID_5, CategoryID_6, CategoryID_7, ShowAGE, ShowWEEKLY_NUT, ShowNUMBER_OF_CAST, ShowNUMBER_OF_MUSICIANS, ShowNUMBER_OF_STAGEHANDS, ShowNUMBER_OF_TRUCKS, ShowNOTES FROM shows`
	fmt.Printf(cursor)
	/*decodedCursor, err := DecodeCursor(cursor)
	if err != nil && cursor != "" {
		fmt.Printf("Error 1")
		return nil, "", nil
	}*/

	res, err := m.fetch(ctx, query)
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

func (m *dbShowRepository) GetByID(ctx context.Context, id int64) ([]*models.Show, error) {
	query := `SELECT ShowID, ShowNAME, ShowACTIVE, 
					CategoryID_1, CategoryID_2, CategoryID_3, 
					CategoryID_4, CategoryID_5, CategoryID_6, 
					CategoryID_7, ShowAGE, ShowWEEKLY_NUT, ShowNUMBER_OF_CAST, 
					ShowNUMBER_OF_MUSICIANS, ShowNUMBER_OF_STAGEHANDS, ShowNUMBER_OF_TRUCKS, ShowNOTES 
			FROM shows 
			WHERE ShowID=?;`

	res, err := m.fetch(ctx, query, id)
	if err != nil {
		fmt.Printf("Error 2")
		return nil, err
	}

	return res, err

}

func (m *dbShowRepository) Store(ctx context.Context, a *models.Show) error {
	query := `INSERT INTO shows (ShowNAME, ShowACTIVE, 
					CategoryID_1, CategoryID_2, CategoryID_3, 
					CategoryID_4, CategoryID_5, CategoryID_6, 
					CategoryID_7, ShowAGE, ShowWEEKLY_NUT, ShowNUMBER_OF_CAST, 
					ShowNUMBER_OF_MUSICIANS, ShowNUMBER_OF_STAGEHANDS, ShowNUMBER_OF_TRUCKS, ShowNOTES) 
			VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, a.Name, a.Active, a.Category1, a.Category2, a.Category3, a.Category4, a.Category5, a.Category6, a.Category7, a.Age, a.WeeklyNut, a.NumberOfCast, a.NumberOfMusicians, a.NumberOfStageHands, a.NumberOfTrucks, a.Notes)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	a.Id = int(lastID)
	return nil
}

func (m *dbShowRepository) Update(ctx context.Context, id int64, a *models.Show) error {

	query := `UPDATE shows SET ShowNAME = ?, ShowACTIVE = ?, 
					CategoryID_1 = ?, CategoryID_2 = ?, CategoryID_3 = ?, 
					CategoryID_4 = ?, CategoryID_5 = ?, CategoryID_6 = ?, 
					CategoryID_7 = ?, ShowAGE = ?, ShowWEEKLY_NUT = ?, ShowNUMBER_OF_CAST = ?, 
					ShowNUMBER_OF_MUSICIANS = ?, ShowNUMBER_OF_STAGEHANDS = ?, ShowNUMBER_OF_TRUCKS = ?, ShowNOTES = ? 
				WHERE ShowID = ?;`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, a.Name, a.Active, a.Category1, a.Category2, a.Category3, a.Category4, a.Category5, a.Category6, a.Category7, a.Age, a.WeeklyNut, a.NumberOfCast, a.NumberOfMusicians, a.NumberOfStageHands, a.NumberOfTrucks, a.Notes, id)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		return err
	}

	return nil
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
