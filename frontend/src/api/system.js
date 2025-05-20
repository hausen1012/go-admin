import axios from 'axios'

export const systemApi = {
  // 加载系统设置
  loadSystemSettings: async () => {
    const response = await axios.get('/api/admin/options')
    return response.data
  },

  // 获取系统信息
  getSysInfo: async () => {
    const response = await axios.get('/api/sysinfo')
    return response.data
  },

  // 更新系统设置
  updateSystemSettings: async (options) => {
    const response = await axios.put('/api/admin/options', { options })
    return response.data
  }
} 