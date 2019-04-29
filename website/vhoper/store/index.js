/*
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

// window.fetch() 的 Polyfill
/!*require('whatwg-fetch');*!/

const store = () => new Vuex.Store({

    state: {
        authUser: null,
        token:''
    },

    mutations: {
        SET_USER: function (state, user) {
            state.authUser = user
        },
        SET_TOKEN:function (state,token) {
            state.token = token
        }
    },

    actions: {
        // ...
    }

});
*/

import cookieUtils from '../plugins/utils/cookie'

export const state = () => ({
  user: null,
  token: ''
})

export const mutations = {
  SET_USER: function (state, user) {
    state.user = user
  },
  SET_TOKEN: function (state, token) {
    state.token = token
  }
}

export const actions = {
  // nuxtServerInit is called by Nuxt.js before server-rendering every page
  async nuxtServerInit({ commit }, { $axios, store, req, route, redirect }) {
    if(store.state.user) return

    const token = cookieUtils.getCookie('token', req)
    if (req.headers.cookie) {
      if (token) {
        commit('SET_TOKEN', token)
        $axios.defaults.headers.common.Authorization = token
        $axios.defaults.headers.common.Cookie = req.headers.cookie
        await $axios.$get(`/api/user/get`).then((res) => {
          // 跟后端的初始化配合
          if (res.code === 200) commit('SET_USER', res.data)
        })
      }
    }

  }
  /*  async login({ commit }, params) {
    try {
      const { data } = await $axios.$post('/user/login', params)
      commit('SET_USER', data)
    } catch (error) {
      if (error.response && error.response.status === 401) {
        throw new Error('Bad credentials')
      }
      throw error
    }
  },

  async logout({ commit }) {
    await $axios.$get('/user/logout')
    commit('SET_USER', null)
  } */
}
