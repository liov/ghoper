--白名单列表
local whitelist = {
    'test',
    'log',
    'user/login',
    'user/register'
}
--路由重写列表
local rewritelist = {
    ['user/([-_a-zA-Z0-9]+)/login'] = 'user',
    ['user/([a-zA-Z0-9]+)/register'] = 'user/register',
    ['user/([a-zA-Z0-9]+)/logout'] = 'user/logout'
}
return {
    whitelist = whitelist,
    rewritelist = rewritelist
}