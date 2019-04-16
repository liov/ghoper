--白名单列表
local whitelist = {
    'lua/test',
    'lua/user/login',
    'lua/user/register'
}
--路由重写列表
local rewritelist = {
    ['lua/user/([-_a-zA-Z0-9]+)/login'] = 'user/login',
    ['lua/user/([a-zA-Z0-9]+)/register'] = 'user/register'
}
return {
    whitelist = whitelist,
    rewritelist = rewritelist
}