package utils

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// SetupValidation registers custom validation rules
func SetupValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Regex patterns from the OpenAPI specification
		retailerPattern := regexp.MustCompile(`^[\w\s\-\&]+$`)
		purchaseDatePattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
		purchaseTimePattern := regexp.MustCompile(`^\d{2}:\d{2}$`)
		totalPattern := regexp.MustCompile(`^\d+\.\d{2}$`)
		shortDescriptionPattern := regexp.MustCompile(`^[\w\s\-]+$`)
		pricePattern := regexp.MustCompile(`^\d+\.\d{2}$`)

		// Register custom validation functions
		_ = v.RegisterValidation(
			"retailerPattern", func(fl validator.FieldLevel) bool {
				return retailerPattern.MatchString(fl.Field().String())
			},
		)
		_ = v.RegisterValidation(
			"purchaseDatePattern", func(fl validator.FieldLevel) bool {
				return purchaseDatePattern.MatchString(fl.Field().String())
			},
		)
		_ = v.RegisterValidation(
			"purchaseTimePattern", func(fl validator.FieldLevel) bool {
				return purchaseTimePattern.MatchString(fl.Field().String())
			},
		)
		_ = v.RegisterValidation(
			"totalPattern", func(fl validator.FieldLevel) bool {
				return totalPattern.MatchString(fl.Field().String())
			},
		)
		_ = v.RegisterValidation(
			"shortDescriptionPattern", func(fl validator.FieldLevel) bool {
				return shortDescriptionPattern.MatchString(fl.Field().String())
			},
		)
		_ = v.RegisterValidation(
			"pricePattern", func(fl validator.FieldLevel) bool {
				return pricePattern.MatchString(fl.Field().String())
			},
		)
	}
}
