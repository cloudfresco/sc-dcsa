package common

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	"go.uber.org/zap"
)

// CurrencyService - For accessing Currency services
type CurrencyService struct {
	log       *zap.Logger
	DBService *DBService
}

// NewCurrencyService - Create Currency Service
func NewCurrencyService(log *zap.Logger, dbOpt *DBService) *CurrencyService {
	return &CurrencyService{
		log:       log,
		DBService: dbOpt,
	}
}

func (cs *CurrencyService) GetCurrency(ctx context.Context, code string) (*commonproto.Currency, error) {
	const query = `SELECT code, numeric_code, currency_name, minor_unit
                   FROM currencies WHERE code = ?`

	row := cs.DBService.DB.QueryRowxContext(ctx, query, code)
	currency := commonproto.Currency{}
	err := row.StructScan(&currency)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("currency not found: %s", code)
		}
		return nil, fmt.Errorf("failed to get currency: %w", err)
	}
	return &currency, nil
}

func ParseAmountString(amountStr string, currency *commonproto.Currency) (int64, error) {
	// Check for valid number format
	if !strings.Contains(amountStr, ".") {
		amountStr += "."
	}

	parts := strings.Split(amountStr, ".")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid amount format")
	}

	// Handle whole number part
	major, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	// Calculate major part in minor units
	amount := major * int64(math.Pow10(int(currency.MinorUnit)))

	// Handle fractional part if it exists
	if len(parts) == 2 {
		// Pad or truncate fractional part to match minor units
		fractionalPart := parts[1]
		if int32(len(fractionalPart)) > currency.MinorUnit {
			fractionalPart = fractionalPart[:currency.MinorUnit]
		} else {
			fractionalPart = fractionalPart + strings.Repeat("0", int(currency.MinorUnit)-len(fractionalPart))
		}

		fractional, err := strconv.ParseInt(fractionalPart, 10, 64)
		if err != nil {
			return 0, err
		}

		// Add fractional part (sign depends on major part)
		if major >= 0 {
			amount += fractional
		} else {
			amount -= fractional
		}
	}

	return amount, nil
}

// FormatAmountString converts minor units back to properly formatted string
func FormatAmountString(amountMinor int64, currency *commonproto.Currency) string {
	var amountStr string
	if currency.MinorUnit == 0 {
		amountStr = strconv.FormatInt(amountMinor, 10)
		return amountStr
	}

	divisor := int64(math.Pow10(int(currency.MinorUnit)))
	major := amountMinor / divisor
	minor := amountMinor % divisor

	// Handle negative amounts correctly
	if amountMinor < 0 {
		minor = -minor
	}

	amountStr = strconv.FormatInt(major, 10) + "." +
		fmt.Sprintf("%0*d", currency.MinorUnit, minor)

	return amountStr
}
