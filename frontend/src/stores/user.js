import { defineStore } from 'pinia'
import axios from 'axios'

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
        const response = await axios.post('/api/login', { username, password })
        this.token = response.data.token
        this.user = response.data.user
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
        const response = await axios.post('/api/register', { username, password })
        return response.data
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
        await axios.put('/api/user/password', { old_password: oldPassword, new_password: newPassword })
        return true
      } catch (error) {
        throw error.response?.data?.error || '更新密码失败'
      }
    },

    async fetchUsers() {
      try {
        const response = await axios.get('/api/admin/users')
        return response.data
      } catch (error) {
        throw error.response?.data?.error || '获取用户列表失败'
      }
    },

    async createUser(username, password) {
      try {
        const response = await axios.post('/api/admin/users', { username, password })
        return response.data
      } catch (error) {
        throw error.response?.data?.error || '创建用户失败'
      }
    },

    async updateUser(id, data) {
      try {
        const response = await axios.put(`/api/admin/users/${id}`, data)
        return response.data
      } catch (error) {
        throw error.response?.data?.error || '更新用户失败'
      }
    },

    async deleteUser(id) {
      try {
        const response = await axios.delete(`/api/admin/users/${id}`)
        return response.data
      } catch (error) {
        throw error.response?.data?.error || '删除用户失败'
      }
    },

    async resetUserPassword(id) {
      try {
        const response = await axios.post(`/api/admin/users/${id}/reset-password`)
        return response.data
      } catch (error) {
        throw error.response?.data?.error || '重置密码失败'
      }
    },

    async getRegistrationStatus() {
      try {
        const response = await axios.get('/api/registration-status')
        return response.data.allowRegistration
      } catch (error) {
        console.error('获取注册状态失败:', error)
        return false // 默认不允许注册
      }
    }
  }
}) 