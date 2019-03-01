const Koa = require('koa')
const next = require('next')
const Router = require('koa-router')

const port = parseInt(process.env.PORT, 10) || 3000
const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handle = app.getRequestHandler()
const devProxy = {
    '/api': {
        target: 'https://hoper.xyz/api/',
        pathRewrite: { '^/api': '/' },
        changeOrigin: true
    }
}
app.prepare()
    .then(() => {
        const server = new Koa()
        const router = new Router()

        // Set up the proxy.
        if (dev && devProxy) {
            const proxyMiddleware = require('http-proxy-middleware')
            Object.keys(devProxy).forEach(function (context) {
                server.use(proxyMiddleware(context, devProxy[context]))
            })
        }

        router.get('/moment/:id', async ctx => {
            await app.render(ctx.req, ctx.res, '/moment/id', ctx.query)
            ctx.respond = false
        })

        router.get('*', async ctx => {
            await handle(ctx.req, ctx.res)
            ctx.respond = false
        })



        server.use(async (ctx, next) => {
            ctx.res.statusCode = 200
            await next()
        })

        server.use(router.routes())
        server.listen(port, () => {
            console.log(`> Ready on http://localhost:${port}`)
        })
    })
