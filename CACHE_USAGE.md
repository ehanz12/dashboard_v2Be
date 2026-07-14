# Redis Caching Usage & Testing Guide

## Quick Start

### 1. Setup Redis
Pastikan Redis sudah running di local:
```bash
# Windows (dengan Redis installer)
redis-server

# atau menggunakan Docker
docker run -d -p 6379:6379 redis:latest
```

### 2. Update .env
Sesuaikan konfigurasi Redis di `.env`:
```env
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=  # kosong jika tidak ada password
```

### 3. Run Application
```bash
go run main.go
```

---

## Testing Workflow dengan Postman

### Scenario: Get Analytics Summary

#### **REQUEST 1 - Cache Miss** (First Request)
```
GET http://localhost:3000/analytics/summary?type=monthly&month=1&year=2025
Headers:
  Authorization: Bearer <your_jwt_token>
```

**Expected Response (200)**:
```json
{
  "message": "Success",
  "data": {
    "finance": {
      "income": 5000000,
      "expense": 2000000,
      "balance": 3000000
    },
    "habit": {
      "completed_rate": 75,
      "current_streak": 12,
      "longest_streak": 45
    },
    "task": {
      "completed": 23,
      "total": 30,
      "completed_rate": 77
    }
  }
}
```

**Response Time**: ~500-1000ms (database query + JSON marshal + Redis SET)

---

#### **REQUEST 2 - Cache Hit** (Second Request, < 5 menit)
```
GET http://localhost:3000/analytics/summary?type=monthly&month=1&year=2025
Headers:
  Authorization: Bearer <your_jwt_token>
```

**Expected Response (200)**:
```json
{
  "message": "Success (from cache)",
  "data": {
    "finance": { ... },
    "habit": { ... },
    "task": { ... }
  }
}
```

**Response Time**: ~10-50ms (Redis GET only, no database query)

---

#### **REQUEST 3 - After TTL Expires** (> 5 menit)
```
GET http://localhost:3000/analytics/summary?type=monthly&month=1&year=2025
Headers:
  Authorization: Bearer <your_jwt_token>
```

**Expected Response (200)**:
```json
{
  "message": "Success",
  "data": { ... }
}
```

**Response Time**: ~500-1000ms (cache expired, database query again)

---

## Testing dengan Redis CLI

### 1. Connect ke Redis
```bash
redis-cli
```

### 2. Monitor Cache Keys
```bash
# Lihat semua cache keys
KEYS "analytics:*"

# Contoh output:
# 1) "analytics:summary:user123:{\"Type\":\"monthly\",\"Month\":1,\"Year\":2025}"
# 2) "analytics:finance:user123:{\"Type\":\"monthly\",\"Month\":1,\"Year\":2025}"
```

### 3. Get Specific Cache Value
```bash
# Format: GET "analytics:<endpoint>:<userID>:{filter}"
GET "analytics:summary:user123:{\"Type\":\"monthly\",\"Month\":1,\"Year\":2025}"

# Output: JSON string dari analytics data
```

### 4. Check TTL
```bash
TTL "analytics:summary:user123:{\"Type\":\"monthly\",\"Month\":1,\"Year\":2025}"

# Output:
# 299  -> 4 menit 59 detik remaining
# -1   -> key tidak ada TTL (shouldn't happen)
# -2   -> key expired/tidak ada
```

### 5. Manual Cache Invalidation
```bash
# Hapus 1 specific cache
DEL "analytics:summary:user123:{\"Type\":\"monthly\",\"Month\":1,\"Year\":2025}"

# Hapus semua analytics cache
DEL $(KEYS "analytics:*")

# Clear specific user's cache
EVAL "return redis.call('del',unpack(redis.call('keys','analytics:*:user123:*')))" 0
```

---

## Cache Keys Structure

**Pattern**: `<prefix>:<userID>:<filter_json>`

### Examples

| Endpoint | Cache Key |
|----------|-----------|
| `/analytics/summary` | `analytics:summary:user123:{"Type":"monthly","Month":1,"Year":2025}` |
| `/analytics/finance` | `analytics:finance:user123:{"Type":"monthly","Month":1,"Year":2025}` |
| `/analytics/categories` | `analytics:category:user123:{"Type":"yearly","Month":0,"Year":2025}` |
| `/analytics/habits` | `analytics:habit:user123:{"Type":"monthly","Month":3,"Year":2025}` |
| `/analytics/tasks` | `analytics:task:user123:{"Type":"monthly","Month":12,"Year":2025}` |

---

## Performance Metrics

### Before Caching (Database Only)
```
Request Time: 500-1500ms
Database Load: High (aggregation queries)
Network Traffic: Full payload each request
```

### After Caching (Cache Hit)
```
Request Time: 10-50ms (50-100x faster)
Database Load: Only on cache miss
Network Traffic: Reduced (cached responses)
```

### Improvement
- **10-50x faster** response time pada cache hit
- **Database queries reduced** secara signifikan
- **Better user experience** dengan faster API response

---

## Monitoring Cache Performance

### Redis Command untuk Monitoring
```bash
# Real-time monitoring
MONITOR

# Cache statistics
INFO stats

# Memory usage
INFO memory

# Key count
DBSIZE
```

### Log Output
Check application logs untuk warning saat cache set failed:
```
Failed to set cache: connection refused
Failed to set cache: timeout
```

---

## Advanced: Cache Invalidation Strategies

### 1. Manual Invalidation (After Data Update)
**Pada handler CREATE/UPDATE/DELETE**, tambahkan:
```go
// Delete related cache saat data berubah
ctx := context.Background()
services.DeleteCache(ctx, "analytics:summary:"+userID+":*")
services.DeleteCache(ctx, "analytics:finance:"+userID+":*")
services.DeleteCache(ctx, "analytics:category:"+userID+":*")
```

### 2. Pattern-based Invalidation
```go
// Delete semua cache user
func InvalidateUserCache(userID string) error {
    return services.DeleteCache(ctx, "analytics:*:"+userID+":*")
}
```

### 3. Scheduled Cache Cleanup
```go
// Clear old cache every hour
scheduler.Every(1).Hour().Do(func() {
    // Run cache cleanup
    services.DeleteCache(ctx, "analytics:*")
})
```

---

## Troubleshooting

### Issue: "redis: nil" Error
**Cause**: Redis connection failed
**Solution**:
1. Check Redis is running: `redis-cli ping` (should return PONG)
2. Check .env configuration
3. Check firewall/network connectivity

### Issue: Cache Not Working (Always Database Query)
**Cause**: Cache SET failing silently
**Solution**:
1. Check application logs for "Failed to set cache" messages
2. Verify Redis memory not full: `INFO memory`
3. Check Redis permissions

### Issue: Stale Data in Cache
**Cause**: TTL too long or no invalidation
**Solution**:
1. Reduce TTL: Change `CacheTTL` in `services/cache.go`
2. Implement cache invalidation after updates
3. Manual cache clear via Redis CLI

---

## Best Practices

✅ **DO:**
- Cache expensive queries (analytics, aggregations)
- Monitor cache hit/miss ratio
- Implement cache invalidation for data changes
- Set reasonable TTL based on data freshness requirements
- Log cache errors for debugging

❌ **DON'T:**
- Cache sensitive user data without encryption
- Set TTL too high (stale data)
- Cache POST/PUT/DELETE requests
- Forget to invalidate cache after updates
- Ignore cache errors in production

---

## Next Steps

1. **Monitor Cache Performance**: Track hit/miss ratio
2. **Implement Cache Invalidation**: Add on data mutations
3. **Scale Redis**: Use Redis cluster for production
4. **Add Cache Warming**: Pre-populate frequently accessed data
5. **Implement Different TTLs**: Different endpoints, different TTL
