local m = package.loaded[ngx.var.module]
local cjson = require "cjson"
ngx.say(cjson.encode(test_var_exec_every_time).."<br>")