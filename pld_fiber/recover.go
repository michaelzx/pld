package pld_fiber

import (
	"github.com/gofiber/fiber/v2"
)

func Recover() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		// Catch panics
		defer func() {
			// pld_logger.Debug("Recover-defer")
			if r := recover(); r != nil {
				// pld_logger.Debug("Recover-recover")
				if recoverErr, ok := r.(error); ok {
					// pld_logger.Debug("Recover-recoverErr", recoverErr)
					err = recoverErr
				}
			}
		}()
		return c.Next()
	}
}
