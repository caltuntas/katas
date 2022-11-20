local keys = {};
local done = false;
local cursor = "0";
local logtable = {};

local function dump(o)
   if type(o) == 'table' then
      local s = '{ '
      for k,v in pairs(o) do
         if type(k) ~= 'number' then k = '"'..k..'"' end
         s = s .. '['..k..'] = ' .. dump(v) .. ','
      end
      return s .. '} '
   else
      return tostring(o)
   end
end

local function logit(msg)
  logtable[#logtable+1] = msg;
end

repeat
    local result = redis.call("SCAN", cursor, "match", ARGV[1], "count", ARGV[2])
    local d = dump(result);
    logit(d);
    cursor = result[1];
    keys = result[2];
    logit(unpack(result));
    logit(cursor);
    for i, key in ipairs(keys) do
        redis.call("DEL", key);
    end
    if cursor == "0" then
        done = true;
    end
until done
return logtable;

