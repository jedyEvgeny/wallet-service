package storage

import (
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

	request := `
	INSERT INTO wallets
	(wallet_id, amount)
	VALUES ($1, $2)	
	ON CONFLICT (wallet_id) DO UPDATE 
	SET amount = wallets.amount + EXCLUDED.amount
	`
	sqlStmt, err := db.db.Prepare(request)
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
