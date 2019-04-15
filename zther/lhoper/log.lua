local function close_redis(red)
    if not red then
        return
    end
    -- 释放连接(连接池实现)，毫秒
    local pool_max_idle_time = 10000 
    -- 连接池大小
    local pool_size = 100 
    local ok, err = red:set_keepalive(pool_max_idle_time, pool_size)
    local log = ngx_log
    if not ok then
        log(ngx_ERR, "set redis keepalive error : ", err)
    end
end

local redis = require "resty.redis"
local red = redis:new()
local cjson = require "cjson"

red:set_timeout(1000)

local ok, err = red:connect("127.0.0.1", 6379)
if not ok then
        ngx.say("failed to connect: ", err,"<br>")
        return
end

local headers=ngx.req.get_headers()
local ip=headers["X-REAL-IP"] or headers["X_FORWARDED_FOR"] or ngx.var.remote_addr or "0.0.0.0"

local exist, err = red:hexists("ip_list",ip)
if not exist then
	local data = {}
	data[ip] = 1
    red:hmset("ip_list",data)
else	
	red:hincrby("ip_list",ip,1)
end
local res, err = red:hgetall("ip_list")
if not res then
    ngx.say("failed to get ip_list: ", err)
    return
end
ngx.say("ip_list:[IP地址:访问次数]<br>")
ngx.say(cjson.encode(res))

close_redis(red)