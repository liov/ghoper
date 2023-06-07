# antd

> hoper

## Build Setup

``` bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn run dev

# build for production and launch server
$ yarn run build
$ yarn start

# generate static project
$ yarn run generate
```

For detailed explanation on how things work, checkout [Nuxt.js docs](https://nuxtjs.org).

为了消除30版本的bug
VueApolloMutationOptions (in types/options.d.ts) works around the error.
optimisticResponse?: ((this: ApolloVueThisType<V>) => any) | Object;=>
optimisticResponse?: ((this: ApolloVueThisType<V>) => any) | any;
