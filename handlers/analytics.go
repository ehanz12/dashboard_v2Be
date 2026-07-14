package handlers

import (
	"be_dashboard/dto/requests"
	"be_dashboard/services"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// parseAnalyticsFilter mengekstrak dan memvalidasi query parameter analytics.
// Default: type=monthly, month=bulan saat ini, year=tahun saat ini.
func parseAnalyticsFilter(c *fiber.Ctx) (requests.AnalyticsFilter, error) {
	now := time.Now()

	typeParam := c.Query("type", "monthly")
	yearParam := c.Query("year", strconv.Itoa(now.Year()))
	monthParam := c.Query("month", strconv.Itoa(int(now.Month())))

	year, err := strconv.Atoi(yearParam)
	if err != nil {
		return requests.AnalyticsFilter{}, fiber.NewError(fiber.StatusBadRequest, "year must be a valid integer")
	}

	month, err := strconv.Atoi(monthParam)
	if err != nil {
		return requests.AnalyticsFilter{}, fiber.NewError(fiber.StatusBadRequest, "month must be a valid integer")
	}

	filter := requests.AnalyticsFilter{
		Type:  typeParam,
		Month: month,
		Year:  year,
	}

	if err := filter.Validate(); err != nil {
		return requests.AnalyticsFilter{}, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return filter, nil
}

// GetAnalyticsSummaryHandler godoc
// GET /analytics/summary
func GetAnalyticsSummaryHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	filter, err := parseAnalyticsFilter(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generate cache key
	cacheKey, err := services.GenerateCacheKey(userID, "analytics:summary", filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate cache key"})
	}

	// Check cache
	cachedData, found, err := services.GetFromCache(ctx, cacheKey)
	if err == nil && found {
		var data interface{}
		if err := json.Unmarshal(cachedData, &data); err == nil {
			return c.JSON(fiber.Map{
				"message": "Success (from cache)",
				"data":    data,
			})
		}
	}

	// Cache miss - query database
	data, err := services.GetAnalyticsSummary(userID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Set cache
	if err := services.SetCache(ctx, cacheKey, data); err != nil {
		// Log error tapi jangan stop response
		fmt.Printf("Failed to set cache: %v\n", err)
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}

// GetFinanceAnalyticsHandler godoc
// GET /analytics/finance
func GetFinanceAnalyticsHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	filter, err := parseAnalyticsFilter(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generate cache key
	cacheKey, err := services.GenerateCacheKey(userID, "analytics:finance", filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate cache key"})
	}

	// Check cache
	cachedData, found, err := services.GetFromCache(ctx, cacheKey)
	if err == nil && found {
		var data interface{}
		if err := json.Unmarshal(cachedData, &data); err == nil {
			return c.JSON(fiber.Map{
				"message": "Success (from cache)",
				"data":    data,
			})
		}
	}

	// Cache miss - query database
	data, err := services.GetFinanceAnalytics(userID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Set cache
	if err := services.SetCache(ctx, cacheKey, data); err != nil {
		// Log error tapi jangan stop response
		fmt.Printf("Failed to set cache: %v\n", err)
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}

// GetCategoryAnalyticsHandler godoc
// GET /analytics/categories
func GetCategoryAnalyticsHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	filter, err := parseAnalyticsFilter(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generate cache key
	cacheKey, err := services.GenerateCacheKey(userID, "analytics:category", filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate cache key"})
	}

	// Check cache
	cachedData, found, err := services.GetFromCache(ctx, cacheKey)
	if err == nil && found {
		var data interface{}
		if err := json.Unmarshal(cachedData, &data); err == nil {
			return c.JSON(fiber.Map{
				"message": "Success (from cache)",
				"data":    data,
			})
		}
	}

	// Cache miss - query database
	data, err := services.GetCategoryAnalytics(userID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Set cache
	if err := services.SetCache(ctx, cacheKey, data); err != nil {
		// Log error tapi jangan stop response
		fmt.Printf("Failed to set cache: %v\n", err)
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}

// GetHabitAnalyticsHandler godoc
// GET /analytics/habits
func GetHabitAnalyticsHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	filter, err := parseAnalyticsFilter(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generate cache key
	cacheKey, err := services.GenerateCacheKey(userID, "analytics:habit", filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate cache key"})
	}

	// Check cache
	cachedData, found, err := services.GetFromCache(ctx, cacheKey)
	if err == nil && found {
		var data interface{}
		if err := json.Unmarshal(cachedData, &data); err == nil {
			return c.JSON(fiber.Map{
				"message": "Success (from cache)",
				"data":    data,
			})
		}
	}

	// Cache miss - query database
	data, err := services.GetHabitAnalytics(userID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Set cache
	if err := services.SetCache(ctx, cacheKey, data); err != nil {
		// Log error tapi jangan stop response
		fmt.Printf("Failed to set cache: %v\n", err)
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}

// GetTaskAnalyticsHandler godoc
// GET /analytics/tasks
func GetTaskAnalyticsHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	filter, err := parseAnalyticsFilter(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generate cache key
	cacheKey, err := services.GenerateCacheKey(userID, "analytics:task", filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate cache key"})
	}

	// Check cache
	cachedData, found, err := services.GetFromCache(ctx, cacheKey)
	if err == nil && found {
		var data interface{}
		if err := json.Unmarshal(cachedData, &data); err == nil {
			return c.JSON(fiber.Map{
				"message": "Success (from cache)",
				"data":    data,
			})
		}
	}

	// Cache miss - query database
	data, err := services.GetTaskAnalytics(userID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Set cache
	if err := services.SetCache(ctx, cacheKey, data); err != nil {
		// Log error tapi jangan stop response
		fmt.Printf("Failed to set cache: %v\n", err)
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}

// GetStreakAnalyticsHandler godoc
// GET /analytics/streak
func GetStreakAnalyticsHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	filter, err := parseAnalyticsFilter(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	data, err := services.GetStreakAnalytics(userID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}
