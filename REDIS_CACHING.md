# Redis Caching Implementation Guide

## Workflow Caching

Sistem caching yang sudah diimplementasikan mengikuti workflow berikut:

```
Request (Analytics API)
   │
   ▼
Redis Check
   │
   ├── Ada cache?
   │      │
   │      ├── Ya  → Unmarshal & Return (response: "Success (from cache)")
   │      │
   │      └── Tidak → Lanjut ke database
   │
   ▼
MariaDB Query
   │
   ▼
JSON Marshal
   │
   ▼
Redis SET (TTL 5 menit)
   │
   ▼
Return Response (response: "Success")
```

## File yang Dimodifikasi

### 1. **config/config.go**
Menambahkan Redis configuration fields:
- `RedisHost` - Host Redis (default: localhost)
- `RedisPort` - Port Redis (default: 6379)
- `RedisPass` - Password Redis (bisa kosong)

### 2. **database/redis.go** (BARU)
- `InitRedis()` - Inisialisasi Redis client
- `CloseRedis()` - Menutup koneksi Redis
- Global `RedisClient` variable

### 3. **services/cache.go** (BARU)
Cache utility functions:
- `GetFromCache(ctx, cacheKey)` - Ambil data dari cache
- `SetCache(ctx, cacheKey, data)` - Simpan data ke cache (TTL 5 menit)
- `DeleteCache(ctx, cacheKey)` - Hapus cache
- `GenerateCacheKey(userID, prefix, filter)` - Generate cache key unik

### 4. **handlers/analytics.go**
Implemented caching untuk 4 endpoints:
- `GET /analytics/summary` - Cache key: `analytics:summary:{userID}:{filter}`
- `GET /analytics/finance` - Cache key: `analytics:finance:{userID}:{filter}`
- `GET /analytics/categories` - Cache key: `analytics:category:{userID}:{filter}`
- `GET /analytics/habits` - Cache key: `analytics:habit:{userID}:{filter}`
- `GET /analytics/tasks` - Cache key: `analytics:task:{userID}:{filter}`

### 5. **main.go**
Menambahkan `database.InitRedis()` saat startup aplikasi

## Environment Variables

Tambahkan ke `.env`:
```
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
```

## Cara Kerja

### Cache Hit (Data ada di Redis)
```
GET /analytics/summary?type=monthly&month=1&year=2025
↓
Cache Key: analytics:summary:user123:{type:monthly,month:1,year:2025}
↓
Found in Redis!
↓
Response: "Success (from cache)" + data
```

### Cache Miss (Data tidak ada di Redis)
```
GET /analytics/summary?type=monthly&month=1&year=2025
↓
Cache Key not found in Redis
↓
Query MariaDB
↓
Marshal to JSON
↓
Save to Redis (TTL: 5 menit)
↓
Response: "Success" + data
```

## TTL (Time To Live)

Cache akan otomatis expire setelah **5 menit** (300 detik). Kamu bisa ubah di `services/cache.go`:

```go
const CacheTTL = 5 * time.Minute  // Ubah sesuai kebutuhan
```

## Error Handling

- Jika Redis tidak available saat startup, aplikasi akan panic
- Jika cache SET gagal, sistem akan log warning tapi tetap return response
- Fallback: Handler tetap bisa berjalan walau Redis error

## Testing

### Dengan Redis CLI
```bash
# Check cache key
KEYS "analytics:*"

# Get cache value
GET "analytics:summary:user123:{...}"

# Check TTL
TTL "analytics:summary:user123:{...}"

# Clear specific cache
DEL "analytics:summary:user123:{...}"

# Clear all analytics cache
DEL $(KEYS "analytics:*")
```

### Expected Behavior
1. Request pertama → MariaDB query, response: "Success"
2. Request kedua (< 5 menit) → Cache hit, response: "Success (from cache)"
3. Request ketiga (> 5 menit) → Cache expired, MariaDB query lagi, response: "Success"

## Improvement Ideas (Optional)

1. **Cache Invalidation**: Invalidate cache saat user membuat/update data
   ```go
   // Di handler CREATE/UPDATE
   services.DeleteCache(ctx, "analytics:*")
   ```

2. **Selective Caching**: Cache hanya query tertentu yang expensive
   ```go
   if filter.Type == "yearly" {
       // Cache karena expensive
   }
   ```

3. **Different TTL**: Berbeda TTL untuk berbeda endpoint
   ```go
   SetCacheWithTTL(ctx, key, data, 10*time.Minute)
   ```

4. **Cache Warming**: Pre-populate cache saat startup
   ```go
   func WarmCache() {
       // Pre-compute analytics data
   }
   ```
