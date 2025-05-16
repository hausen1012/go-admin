import { defineStore } from 'pinia'
import axios from 'axios'
import { userApi } from '../api/user'

export const useUserStore = defineStore('user', {
  state: () => {
    const token = localStorage.getItem('token') || ''
    const user = JSON.parse(localStorage.getItem('user')) || null
    
    // 初始化时设置 axios 默认配置
    if (token) {
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
    }
    
    return {
      token,
      user
    }
  },

  getters: {
    isLoggedIn: (state) => !!state.token,
    isAdmin: (state) => state.user?.is_admin || false
  },

  actions: {
    async login(username, password) {
      try {
        const data = await userApi.login(username, password)
        this.token = data.token
        this.user = data.user
        localStorage.setItem('token', this.token)
        localStorage.setItem('user', JSON.stringify(this.user))
        axios.defaults.headers.common['Authorization'] = `Bearer ${this.token}`
        return true
      } catch (error) {
        throw error.response?.data?.error || '登录失败'
      }
    },

    async register(username, password) {
      try {
        return await userApi.register(username, password)
      } catch (error) {
        throw error.response?.data?.error || '注册失败'
      }
    },

    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      delete axios.defaults.headers.common['Authorization']
    },

    async updatePassword(oldPassword, newPassword) {
      try {
        await userApi.updatePassword(oldPassword, newPassword)
        return true
      } catch (error) {
        throw error.response?.data?.error || '更新密码失败'
      }
    },

    async fetchUsers() {
      try {
        return await userApi.fetchUsers()
      } catch (error) {
        throw error.response?.data?.error || '获取用户列表失败'
      }
    },

    async createUser(username, password) {
      try {
        return await userApi.createUser(username, password)
      } catch (error) {
        throw error.response?.data?.error || '创建用户失败'
      }
    },

    async updateUser(id, data) {
      try {
        return await userApi.updateUser(id, data)
      } catch (error) {
        throw error.response?.data?.error || '更新用户失败'
      }
    },

    async deleteUser(id) {
      try {
        return await userApi.deleteUser(id)
      } catch (error) {
        throw error.response?.data?.error || '删除用户失败'
      }
    },

    async resetUserPassword(id) {
      try {
        return await userApi.resetUserPassword(id)
      } catch (error) {
        throw error.response?.data?.error || '重置密码失败'
      }
    }
  }
}) 