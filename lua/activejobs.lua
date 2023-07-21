local logtable = {};

local function logit(msg)
  logtable[#logtable+1] = msg;
end

-- hgetall "bull:Calculate Security & Hardening Scores:repeat:147ee85d39ad511f33596517b8aa68d7:1669770000000"
-- hget "bull:Calculate Security & Hardening Scores:repeat:ea6709aed9baebb25d9cd18077c7759f:1670029200000" progress
local queue = "bull:Calculate Security & Hardening Scores"
local activeJobs = queue .. ":active"
local result = redis.call("LRANGE", activeJobs, 0 , -1)
for i, val in ipairs(result) do
  local key = queue .. ":" .. val
  local progress = redis.call("HGET", key, "progress");
  local data = redis.call("HGET", key, "data");
  local jsonData = cjson.decode(data)
  logit(jsonData.ip .. ":" .. progress)
end

return logtable;
