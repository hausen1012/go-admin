import { defineStore } from 'pinia'
import axios from 'axios'

export const useSystemStore = defineStore('system', {
  state: () => ({
    systemName: '后台管理系统', // 默认值
    systemDescription: '',
    allowRegistration: true,
    loading: false
  }),

  actions: {
    async loadSystemSettings() {
      try {
        this.loading = true
        const response = await axios.get('/api/admin/options')
        const options = response.data
        
        options.forEach(option => {
          if (option.option_name === 'system_name') {
            this.systemName = option.option_value
          } else if (option.option_name === 'system_description') {
            this.systemDescription = option.option_value
          } else if (option.option_name === 'allow_registration') {
            this.allowRegistration = option.option_value === 'true'
          }
        })
      } catch (error) {
        console.error('加载系统配置失败:', error)
      } finally {
        this.loading = false
      }
    }
  }
}) 