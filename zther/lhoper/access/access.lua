local need_moudle = router_filter:route_verify()

local function load_moudle()
    local path = require "path"
    local info = debug.getinfo(1,"S")
    --获取当前路径
    local pathinfo = info.short_src
    --由于获取的路径为反斜杠(\)所以用上面的函数转为正斜杠(/)
    local path = string.match(path.conversion(pathinfo),"^(.*/).*/.*$")
    package.loaded[ngx.var.lua_path] = dofile(path.."/lua/"..ngx.var.lua_path..".lua")
end

local function get_moudle(uri)
    local ret =  string.match(uri,"[^/.]+") 
    return ret
end

if need_moudle then
    if package.loaded[ngx.var.lua_path] == nil then
        load_moudle()
    else
        if ngx.var.arg_reload then
            package.loaded[ngx.var.lua_path] = nil
            load_moudle()
        end
    end
    ngx.var.module = get_moudle(ngx.var.lua_path)
    ngx.var.lua_path = "router"
end