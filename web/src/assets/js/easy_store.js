import Vue from 'vue'
import Vuex from 'vuex'
import Cookie from './easy_cookie'

//使用vuex
Vue.use(Vuex);

const USER = 'user';//用户
const TOKEN = 'token';//口令

const store = new Vuex.Store({
  state: {
    user: (Cookie.getVal(USER) == undefined ? '' : JSON.parse(Cookie.getVal(USER))) || '',
    token: Cookie.getVal(TOKEN) || '',
  },
  getters: {
    getLoginUser(state) {
      return state.user
    },
    getToken(state) {
      return state.token
    }
  },
  mutations: {
    setLoginUser(state, user) {
      state.user = user;
    },
    setToken(state, token) {
      state.token = token;
    }
  },
  actions: {
    // 登录用户信息
    saveLoginUser({commit}, user) {
      return new Promise((resolve => {
        Cookie.setVal(USER, JSON.stringify(user));
        commit('setLoginUser', user)
        resolve();
      }))
    },
    // Token
    saveToken({commit}, token) {
      return new Promise((resolve => {
        Cookie.setVal(TOKEN, token);
        commit('setToken', token)
        resolve();
      }))
    }
  }
});

export default store
