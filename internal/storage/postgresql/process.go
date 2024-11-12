package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jedyEvgeny/wallet-service/internal/app/service"
)

func (db *DataBase) Write(req *service.Wallet, requestID string) (err error) {
	if req.OperationType == service.Withdrow {
		req.Amount = -req.Amount
	}

	tx, err := db.db.Begin()
	if err != nil {
		return fmt.Errorf(errCreateTx, err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		commitErr := tx.Commit()
		if commitErr != nil {
			err = fmt.Errorf(errCommitTx, err)
		}
	}()

	sqlStmt, err := tx.Prepare(requestInsertWallet())
	if err != nil {
		return fmt.Errorf(errStmt, err)
	}
	defer func() { _ = sqlStmt.Close() }()

	startInsert := time.Now()
	res, err := sqlStmt.Exec(
		req.WalletId,
		req.Amount,
	)
	endInsert := time.Now()
	if err != nil {
		return fmt.Errorf(errExec, err)
	}
	resAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf(errRes, err)
	}
	if resAffected == 0 {
		return fmt.Errorf(errResAffected, err)
	}

	log.Printf(msgTimeInsert, requestID, endInsert.Sub(startInsert))
	return nil
}

func (db *DataBase) Read(id, requestID string) (amount *int, err error) {
	tx, err := db.db.Begin()
	if err != nil {
		return amount, fmt.Errorf(errCreateTx, err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		commitErr := tx.Commit()
		if commitErr != nil {
			err = fmt.Errorf(errCommitTx, err)
		}
	}()

	sqlStmt, err := tx.Prepare(requestFindWallet())
	if err != nil {
		return amount, fmt.Errorf(errStmt, err)
	}
	defer func() { _ = sqlStmt.Close() }()

	startSelect := time.Now()
	err = sqlStmt.QueryRow(id).Scan(&amount)
	endSelect := time.Now()

	if err == sql.ErrNoRows {
		return amount, fmt.Errorf(errIsNoUUID, id, err)
	}
	if err != nil {
		return amount, fmt.Errorf(errExec, err)
	}

	log.Printf(msgTimeSelect, requestID, endSelect.Sub(startSelect))
	return amount, nil
}

func requestInsertWallet() string {
	return `
	INSERT INTO wallets
	(wallet_id, amount)
	VALUES ($1, $2)	
	ON CONFLICT (wallet_id) DO UPDATE 
	SET amount = wallets.amount + EXCLUDED.amount
	`
}

func requestFindWallet() string {
	return `
    SELECT amount 
	FROM wallets
	WHERE wallet_id = $1
    `
}
