package helper

import (
	"database/sql"
	"log"
)

// <-- digunakan untuk mengecek jika ada error maka transaksi akan dirollback

func CommitOrRollback(tx *sql.Tx, err error) {
	func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			log.Printf("Rollback failed: %v", errRollback)
			IfErrNotNul(errRollback)

		} else {
			errCommit := tx.Commit()
			log.Printf("Rollback failed: %v", errCommit)
			IfErrNotNul(errCommit)
		}
	}()
}
