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

import axios from 'axios'
// import cookie from '../plugins/utils/cookie'
const cookieparser = process.server ? require('cookieparser') : undefined

export const state = () => ({
  user: null,
  token: ''
})

export const mutations = {
  SET_USER: function(state, user) {
    state.user = user
  },
  SET_TOKEN: function(state, token) {
    state.token = token
  }
}

export const actions = {
  // nuxtServerInit is called by Nuxt.js before server-rendering every page
  async nuxtServerInit({ commit }, { store, req }) {
    // let token =  cookie.getCookie("token",req);
    let user
    let token = null
    if (req.headers.cookie) {
      const parsed = cookieparser.parse(req.headers.cookie)
      try {
        token = parsed.token
        if (token) {
          commit('SET_TOKEN', token)
          // axios.defaults.headers.common['Authorization'] = token;
          axios.defaults.headers.common.Cookie = req.headers.cookie
          await axios.get(`/api/user/get`).then(res => {
            // 跟后端的初始化配合
            if (res.data.msg === '登录成功') {
              user = res.data.data
              commit('SET_USER', user)
            }
          })
        }
      } catch (err) {
        // No valid cookie found
      }
    }
  }
  /*    async login({ commit }, { username, password }) {
        try {
            const { data } = await axios.post('/user/login', { username, password })
            commit('SET_USER', data)
        } catch (error) {
            if (error.response && error.response.status === 401) {
                throw new Error('Bad credentials')
            }
            throw error
        }
    },

    async logout({ commit }) {
        await axios.get('/user/logout')
        commit('SET_USER', null)
    } */
}