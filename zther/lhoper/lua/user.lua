local _M={_VERSION = 0.1}

function _M.handle()
   ngx.say("测试")
end

function _M.test()
   return "test"
end

return _M