import { defineStore } from 'pinia'
import { systemApi } from '../api/system'

export const useSystemStore = defineStore('system', {
  state: () => ({
    systemName: '后台管理系统', // 默认值
    systemDescription: '',
    allowRegistration: true,
    loading: false,
    sysInfoLoaded: false // 标记是否已加载过系统信息
  }),

  actions: {
    async loadSystemSettings() {
      try {
        this.loading = true
        const options = await systemApi.loadSystemSettings()
        
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
    },

    async getSysInfo() {
      // 如果已经加载过系统信息，直接返回缓存的值
      if (this.sysInfoLoaded) {
        return {
          allowRegistration: this.allowRegistration,
          systemName: this.systemName
        }
      }

      try {
        const data = await systemApi.getSysInfo()
        // 更新 store 中的值
        this.systemName = data.systemName || '后台管理系统'
        this.allowRegistration = data.allowRegistration
        this.sysInfoLoaded = true
        
        return {
          allowRegistration: this.allowRegistration,
          systemName: this.systemName
        }
      } catch (error) {
        console.error('获取系统信息失败:', error)
        return {
          allowRegistration: false,
          systemName: '后台管理系统'
        }
      }
    }
  }
}) 