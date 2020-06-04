local productSurplus
-- 判断用户是否已抢过
local userKey = tostring(KEYS[1])
local result = "FAIL"

-- 一般将商品数量换存在redis中，这里传商品库存Key
local produceSurplusKey = tostring(KEYS[2])
local hasBuy = redis.pcall("HEXISTS", userKey, "hasBuy")

-- 已经抢购过，返回0
if hasBuy ~= 0 then
    return result
end

-- 准备抢购
productSurplus =  redis.pcall("GET", produceSurplusKey)
if productSurplus == nil then
    return result
end

-- 没有剩余可抢购物品
productSurplus = tonumber(productSurplus)
if productSurplus <= 0 then
    return result
end

-- 更新抢购标识
redis.pcall("HSET", userKey, "hasBuy", "1")
-- 减库存
redis.pcall("DECR", produceSurplusKey)
result = "OK"

return result