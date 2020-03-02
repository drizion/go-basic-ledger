package action

import (
	"encoding/json"
	"net/http"

	"github.com/gsabadini/go-bank-transfer/domain"
	"github.com/gsabadini/go-bank-transfer/infrastructure/database"
	"github.com/gsabadini/go-bank-transfer/repository"
	"github.com/gsabadini/go-bank-transfer/usecase"

	"github.com/sirupsen/logrus"
)

//Transfer armazena as dependências de uma transferência
type Transfer struct {
	dbHandler database.NoSQLDbHandler
	logger    *logrus.Logger
}

//NewTransfer constrói uma transferência com suas dependências
func NewTransfer(dbHandler database.NoSQLDbHandler, log *logrus.Logger) Transfer {
	return Transfer{dbHandler: dbHandler, logger: log}
}

//Store é um handler para criação de transferência
func (t Transfer) Store(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_transfer"

	var transfer domain.Transfer
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		t.logError(
			logKey,
			"error when decoding json",
			http.StatusBadRequest,
			err,
		)

		ErrorMessage(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	var (
		transferRepository = repository.NewTransfer(t.dbHandler)
		accountRepository  = repository.NewAccount(t.dbHandler)
	)

	result, err := usecase.StoreTransfer(transferRepository, accountRepository, transfer)
	if err != nil {
		switch err {
		case domain.ErrInsufficientBalance:
			t.logError(
				logKey,
				"insufficient balance",
				http.StatusUnprocessableEntity,
				err,
			)

			ErrorMessage(err, http.StatusUnprocessableEntity).Send(w)
			return
		default:
			t.logError(
				logKey,
				"error when creating a new transfer",
				http.StatusInternalServerError,
				err,
			)

			ErrorMessage(err, http.StatusInternalServerError).Send(w)
			return
		}
	}

	t.logSuccess(logKey, "success create transfer", http.StatusCreated)

	Success(result, http.StatusCreated).Send(w)
}

//Index é um handler para retornar a lista de transferências
func (t Transfer) Index(w http.ResponseWriter, _ *http.Request) {
	const logKey = "index_transfer"

	var transferRepository = repository.NewTransfer(t.dbHandler)

	result, err := usecase.FindAllTransfer(transferRepository)
	if err != nil {
		t.logError(
			logKey,
			"error when returning the transfer list",
			http.StatusInternalServerError,
			err,
		)

		ErrorMessage(err, http.StatusInternalServerError).Send(w)
		return
	}

	t.logSuccess(logKey, "success when returning transfer list", http.StatusOK)

	Success(result, http.StatusOK).Send(w)
}

func (t Transfer) logSuccess(key string, message string, httpStatus int) {
	t.logger.WithFields(logrus.Fields{
		"key":         key,
		"http_status": httpStatus,
	}).Info(message)
}

func (t Transfer) logError(key string, message string, httpStatus int, err error) {
	t.logger.WithFields(logrus.Fields{
		"key":         key,
		"http_status": httpStatus,
		"error":       err.Error(),
	}).Error(message)
}