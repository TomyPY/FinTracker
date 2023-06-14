package handlers

import (
	"net/http"
	"strconv"

	"github.com/TomyPY/FinTracker/internal/wallet"
	"github.com/TomyPY/FinTracker/pkg/web"
	"github.com/gin-gonic/gin"
)

type Wallet struct {
	s wallet.Service
}

func NewWallet(s wallet.Service) *Wallet {
	return &Wallet{s: s}
}

// GetAll handler
// @Summary Show an array of wallets
// @Schemes
// @Description Get all wallets
// @Tags Wallet
// @Produce json
// @Success 200 {array} wallet.Wallet
// @Failure 500 {object} web.errorResponse
// @Router /wallets [get]
func (w *Wallet) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		wallets, err := w.s.GetAll()
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.Success(ctx, http.StatusAccepted, wallets)
	}
}

// Get handler
// @Summary Show a wallet by id
// @Schemes
// @Description Get just one wallet by his ID
// @Tags Wallet
// @Produce json
// @Param walletId path int true "Wallet ID"
// @Success 200 {object} wallet.Wallet
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /wallets/{id} [get]
func (w *Wallet) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, wallet.ErrInvalidWalletId.Error())
			return
		}

		w, err := w.s.Get(id)
		if err != nil {
			switch err {
			case wallet.ErrNotFound:
				web.Error(ctx, http.StatusNotFound, err.Error())
			default:
				web.Error(ctx, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		web.Success(ctx, http.StatusAccepted, w)
	}
}

// GetReportTransactions handler
// @Summary Show an array of transactions of a wallet
// @Schemes
// @Description Get a list of transactions of a wallet
// @Tags Wallet
// @Produce json
// @Param transactionId query int false "Transaction ID"
// @Param walletId path int true "Wallet ID"
// @Success 200 {array} transaction.Transaction
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /wallets/{id}/transactions [get]
func (w *Wallet) GetReportTransactions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var transactionId *int

		//Get wallet_id
		walletId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, wallet.ErrInvalidWalletId.Error())
			return
		}

		//Get transaction_id
		idTransactionString := ctx.Query("transaction_id")
		if idTransactionString != "" {
			i, err := strconv.Atoi(idTransactionString)
			if err != nil {
				web.Error(ctx, http.StatusBadRequest, wallet.ErrInvalidTransactionId.Error())
				return
			}
			transactionId = &i
		}

		txs, err := w.s.GetReportTransactions(transactionId, walletId)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
		}

		web.Success(ctx, http.StatusAccepted, txs)
	}
}

// AddMoney handler
// @Summary Add money to a wallet by his ID
// @Schemes
// @Description Add the specified money to a wallet by his ID
// @Tags Wallet
// @Produce json
// @Param walletId path int true "Wallet ID"
// @Param transaction body transaction.Transaction true "Transaction"
// @Success 200 {object} wallet.Wallet
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /wallets/{id}/addMoney [patch]
func (w *Wallet) AddMoney() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// SubstractMoney handler
// @Summary Substract money from a wallet
// @Schemes
// @Description Add the specified money to a wallet by his ID
// @Tags Wallet
// @Produce json
// @Param walletId path int true "Wallet ID"
// @Success 200 {object} wallet.Wallet
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /wallets/{id}/substractMoney [patch]
func (w *Wallet) SubstractMoney() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// Delete handler
// @Summary Delete a wallet
// @Schemes
// @Description Delete a wallet from db
// @Tags Wallet
// @Produce json
// @Param walletId path int true "Wallet ID"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /wallets/{id} [delete]
func (w *Wallet) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
