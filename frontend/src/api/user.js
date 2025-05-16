import axios from 'axios'

export const userApi = {
  // 登录
  login: async (username, password) => {
    const response = await axios.post('/api/login', { username, password })
    return response.data
  },

  // 注册
  register: async (username, password) => {
    const response = await axios.post('/api/register', { username, password })
    return response.data
  },

  // 更新密码
  updatePassword: async (oldPassword, newPassword) => {
    const response = await axios.put('/api/user/password', { 
      old_password: oldPassword, 
      new_password: newPassword 
    })
    return response.data
  },

  // 获取用户列表
  fetchUsers: async () => {
    const response = await axios.get('/api/admin/users')
    return response.data
  },

  // 创建用户
  createUser: async (username, password) => {
    const response = await axios.post('/api/admin/users', { username, password })
    return response.data
  },

  // 更新用户
  updateUser: async (id, data) => {
    const response = await axios.put(`/api/admin/users/${id}`, data)
    return response.data
  },

  // 删除用户
  deleteUser: async (id) => {
    const response = await axios.delete(`/api/admin/users/${id}`)
    return response.data
  },

  // 重置用户密码
  resetUserPassword: async (id) => {
    const response = await axios.post(`/api/admin/users/${id}/reset-password`)
    return response.data
  }
}