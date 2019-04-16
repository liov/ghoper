ngx.say("URL错了哦<br>")
ngx.say(test.."<br>")
local cjson = require "cjson"
ngx.say(cjson.encode(route).."<br>")
ngx.say(exec_every_time)