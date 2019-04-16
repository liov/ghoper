local _M = { _VERSION = '0.1' }
local mt = { __index = _M}
local tinsert = table.insert
local tconcat = table.concat
local cjson = require "cjson"
function _M.new(self,config_path)
    local tab = {}
    local routeMap = require(config_path)
    local route_data = {}
    local whitelist = routeMap.whitelist
    for i=1,#whitelist do
        tinsert(route_data,tconcat({'^',whitelist[i],'$'}))
    end
    local rewritelist = routeMap.rewritelist
    local rewrite_data = {}
    local rewrite_urls = {}
    local x = 1
    for k,v in pairs(rewritelist) do
        tinsert(rewrite_data,tconcat({'^(?<z',x,'z>',k,')$'}))
        tinsert(rewrite_urls,v)
         x = x + 1
    end
    tab.rewrite_urls = rewrite_urls
    tab.rewrite_pattern = tconcat(rewrite_data,'|')
    tab.route_pattern = tconcat(route_data,'|')
    return setmetatable(tab, mt)
end

function _M.route_verify(self)
    local lua_path = ngx.var.lua_path
    local m = ngx.re.match(lua_path,self.route_pattern)
    if m == nil then
        m = ngx.re.match(lua_path,self.rewrite_pattern)
        if m == nil then
            ngx.var.lua_path = "error"
        else
            local locant = ngx.re.match(next(m,#m), "^z(\\d+)z")
            test = cjson.encode(m)
            ngx.var.lua_path = self.rewrite_urls[tonumber(locant[1])]
        end
    end
    ngx.var.lua_path = "error"
end
return _M

--莫名其妙的bug
-- set $lua_path $1;
-- ngx.var.lua_path = nil
--ngx.var.1 error
--ngx.var[1]
--原来是windows结束进程，其实后台还有无数进程