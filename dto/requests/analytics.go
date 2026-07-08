package requests

import (
	"errors"
	"time"
)

// AnalyticsFilter adalah query parameter yang digunakan oleh semua endpoint analytics.
// type=monthly  → wajib sertakan month & year
// type=yearly   → cukup sertakan year
type AnalyticsFilter struct {
	Type  string // "monthly" | "yearly"
	Month int    // 1–12, relevan hanya saat type=monthly
	Year  int    // selalu wajib
}

// Validate memastikan kombinasi Type / Month / Year konsisten.
func (f *AnalyticsFilter) Validate() error {
	if f.Type != "monthly" && f.Type != "yearly" {
		return errors.New("type must be 'monthly' or 'yearly'")
	}
	if f.Year < 2000 || f.Year > 2100 {
		return errors.New("year must be between 2000 and 2100")
	}
	if f.Type == "monthly" {
		if f.Month < 1 || f.Month > 12 {
			return errors.New("month must be between 1 and 12")
		}
	}
	return nil
}

// DateRange mengembalikan waktu awal (start) dan akhir (end, eksklusif) dari filter.
// Contoh:
//   type=monthly, month=7, year=2026  →  2026-07-01 00:00:00  s/d  2026-08-01 00:00:00
//   type=yearly,  year=2026           →  2026-01-01 00:00:00  s/d  2027-01-01 00:00:00
func (f *AnalyticsFilter) DateRange() (time.Time, time.Time) {
	if f.Type == "monthly" {
		start := time.Date(f.Year, time.Month(f.Month), 1, 0, 0, 0, 0, time.Local)
		end := start.AddDate(0, 1, 0) // awal bulan berikutnya
		return start, end
	}
	// yearly
	start := time.Date(f.Year, time.January, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(1, 0, 0) // awal tahun berikutnya
	return start, end
}

// DaysInMonth mengembalikan jumlah hari pada bulan filter (hanya valid untuk type=monthly).
func (f *AnalyticsFilter) DaysInMonth() int {
	// Hari terakhir bulan = hari ke-0 bulan berikutnya
	return time.Date(f.Year, time.Month(f.Month+1), 0, 0, 0, 0, 0, time.Local).Day()
}
