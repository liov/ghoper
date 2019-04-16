local _M = { _VERSION = '0.1' }
local mt = { __index = _M}
local tinsert = table.insert
local tconcat = table.concat
local tonumber = tonumber
local function uritoken(uri)
    local ret = {}
    for token in uri:gmatch("[^/.]+") do
        tinsert(ret,token)
    end
    return ret
end

function _M.new(self,rconfig)
    local tab = {}
    local routeMap = require(rconfig)
    local route_data = {}
    local whitelist = routeMap.whitelist
    for i=1,#whitelist do
        local ret = uritoken(whitelist[i])
        if #ret > 0 then
            tinsert(route_data,tconcat({'^',tconcat(ret,'/'),'$'}))
        end
    end
    local rewritelist = routeMap.rewritelist
    local x = 1
    local rewrite_data = {}
    local rewrite_urls = {}
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
    local uri = ngx.var.api_path
    local ret = uritoken(uri)
    uri = tconcat(ret,'/')
    ngx.var.api_path = uri
    local m = ngx.re.match(uri,self.route_pattern)
    if m == nil then
        m = ngx.re.match(uri,self.rewrite_pattern)
        if m == nil then
            ngx.exit(404)
        else
            local locant = ngx.re.match(next(m,#m), "^z(\\d+)z")
            ngx.var.api_path = self.rewrite_urls[tonumber(locant[1])]
        end
    end
end
return _M