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
  }
} 