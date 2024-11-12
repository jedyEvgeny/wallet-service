package service

import "github.com/google/uuid"

type Wallet struct {
	WalletId      uuid.UUID `json:"valletId"`
	OperationType string    `json:"operationType"`
	Amount        int       `json:"amount"`
}

type ResponsePost struct {
	CodeStatus       int        `json:"codeStatus"`
	DescritionStatus string     `json:"descriptionStatus"`
	ID               *uuid.UUID `json:"resourceID,omitempty"`
}

type ResponseGet struct {
	CodeStatus       int    `json:"codeStatus"`
	DescritionStatus string `json:"descriptionStatus"`
	Amount           *int   `json:"walletAmount,omitempty"`
}

const (
	Deposit  = "DEPOSIT"
	Withdrow = "WITHDRAW"
)

const (
	msg200 = "Ресурс существует"
	msg201 = "Ресурс создан"
)

const (
	errMarshalJson    = "ошибка создания json-объекта"
	errDecodeJson     = "ошибка декодирования json-объекта"
	errMethod         = "ошибка метода. Ожидался: %s, имеется: %s"
	errIsNotUUID      = "поле json valletID ожидалось с уникальным UUID. Имеется: %v"
	errOperation      = "поле json operationType: %s. Ожидалось 'DEPOSIT' или 'WITHDRAW'"
	errAmount         = "поле json amount должно быть больше нуля. Имеется: %d"
	errIsNotUUIDInURL = "в URL ожидалось UUID. Имеется: %s"
)
