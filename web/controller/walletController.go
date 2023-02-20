package controller

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/q2-payer-challenge/web/business"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IWalletController interface {
	GetWalletByUserCredentials(ctx *fiber.Ctx) error
	CreateWallet(ctx *fiber.Ctx) error
	DeleteWallet(ctx *fiber.Ctx) error
	DepositToWalletBalance(ctx *fiber.Ctx) error
	PayWalletTransaction(ctx *fiber.Ctx) error
}

type walletController struct {
	walletBusiness business.IWalletBusiness
}

func NewWalletController(walletBusiness business.IWalletBusiness) IWalletController {
	return &walletController{
		walletBusiness: walletBusiness,
	}
}

func (c *walletController) GetWalletByUserCredentials(ctx *fiber.Ctx) error {
	credentialsRequest := new(dto.UserCredentials)
	if err := ctx.BodyParser(credentialsRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	wallet, err := c.walletBusiness.GetWalletByUserCredentials(credentialsRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	walletResponse := dto.ParseWalletToResponse(wallet)
	return ctx.Status(http.StatusOK).JSON(walletResponse)
}

func (c *walletController) CreateWallet(ctx *fiber.Ctx) error {
	walletRequest := new(dto.WalletRequest)
	if err := ctx.BodyParser(walletRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	wallet := dto.ParseRequestToWallet(walletRequest)
	if err := c.walletBusiness.CreateWallet(wallet); err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "wallet created successfully!",
		"wallet":  walletRequest,
	})
}

func (c *walletController) DeleteWallet(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.walletBusiness.DeleteWallet(uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "wallet deleted successfully!",
	})
}

func (c *walletController) DepositToWalletBalance(ctx *fiber.Ctx) error {
	walletRequest := new(dto.WalletRequest)
	if err := ctx.BodyParser(walletRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	wallet, err := c.walletBusiness.GetWalletByUserId(walletRequest.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	oldBalance := wallet.Balance
	if err := c.walletBusiness.DepositToWalletBalance(walletRequest.UserID, walletRequest.Balance); err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	newWallet, err := c.walletBusiness.GetWalletByUserId(walletRequest.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "deposit to wallet successfully!",
		"status": fiber.Map{
			"old_balance":     oldBalance,
			"current_balance": newWallet.Balance,
			"incoming_value":  walletRequest.Balance,
		},
	})
}

func (c *walletController) PayWalletTransaction(ctx *fiber.Ctx) error {
	paymentTransactionRequest := new(dto.TransactionRequest)
	if err := ctx.BodyParser(paymentTransactionRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	updatedReceiverWallet, err := c.walletBusiness.PaymentWalletTransaction(paymentTransactionRequest)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	updatedReceiverWalletResponse := dto.ParseWalletToResponse(updatedReceiverWallet)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "payment succeded!",
		"receiver": updatedReceiverWalletResponse,
	})
}
