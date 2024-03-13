package db

import (
	"context"
	"database/sql"
	"fmt"
)

//store provides all functions to execute db queries
type Store struct{
	*Queries
	db *sql.DB
}

//Newstore creates a new tore 
func Newstore(db *sql.DB) *Store{
	return &Store{
		db: db,
		Queries: New(db),
	}
}

//exeTx executes a funtion within a database transactions
func (store *Store)execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err:%w, rb err:%v",err,rbErr)
		}
		return err
	}
	return tx.Commit()
}
// CreateFileParams contains the input parameters for creating a file
type CreateShareFilesParams struct {
	TaskName string `json:"task_name"`
	TaskTime string `json:"task_time"`
	TaskDate string `json:"task_date"`
}

// CreateFileResult is the result of the create file operation
type CreateShareFileResult struct {
	ShareFile string `json:"task_details"`
	
}

// // Creates a new file in the database
// func (store *Store) CreateShareFile(ctx context.Context, arg CreateShareFilesParams) (CreateShareFileResult, error) {
// 	var result CreateShareFileResult

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error

// 		result.ShareFile, err = q.CreateShareFile(ctx, CreateShareFilesParams{
// 			TaskName: arg.TaskName,
// 			TaskTime: arg.TaskTime,
// 			TaskDate: arg.TaskDate,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// 	return result, err
// }