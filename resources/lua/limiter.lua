local userKey = KEYS[1]
local rate = tonumber(ARGV[1])
local nowTime = tonumber(ARGV[2])
local lastTime = nowTime

local storedPermits = 1

if (redis.pcall("exists", userKey) == 0)
then
    storedPermits = rate
    redis.pcall("HSET", userKey, "lastTime", lastTime)
else
    lastTime = tonumber(redis.pcall("HGET", userKey, "lastTime"))
    if ((nowTime - lastTime) > 1000)
    then
        storedPermits = rate
        redis.pcall("HSET", userKey, "lastMS", nowTime)
    else
        storedPermits = tonumber(redis.pcall("HGET", userKey, "storedPermits"))
    end
end

storedPermits = storedPermits - 1

redis.pcall("HSET", userKey, "storedPermits", storedPermits)
redis.pcall("EXPIRE", userKey, 600)

local result = storedPermits >= 0 and "OK" or "FAIL"

return result