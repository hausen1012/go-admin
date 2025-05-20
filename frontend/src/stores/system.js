import { defineStore } from 'pinia'
import { systemApi } from '../api/system'

export const useSystemStore = defineStore('system', {
  state: () => ({
    systemName: '后台管理系统', // 默认值
    allowRegistration: true,
    loading: false,
    sysInfoLoaded: false // 标记是否已加载过系统信息
  }),

  actions: {
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
    },

    // 加载系统设置
    async loadSystemSettings() {
      try {
        const options = await systemApi.loadSystemSettings()
        options.forEach(option => {
          if (option.OptionName === 'system_name') {
            this.systemName = option.OptionValue
          } else if (option.OptionName === 'allow_registration') {
            this.allowRegistration = option.OptionValue === 'true'
          }
        })
        return options
      } catch (error) {
        console.error('加载系统设置失败:', error)
        throw error
      }
    },

    // 更新系统设置
    async updateSystemSettings(settings) {
      try {
        const settingsData = settings.map(setting => ({
          option_name: setting.option_name,
          option_value: setting.option_value
        }))
        
        await systemApi.updateSystemSettings(settingsData)
        
        // 更新 store 中的状态
        settings.forEach(setting => {
          if (setting.option_name === 'system_name') {
            this.systemName = setting.option_value
          } else if (setting.option_name === 'allow_registration') {
            this.allowRegistration = setting.option_value === 'true'
          }
        })
      } catch (error) {
        console.error('更新系统设置失败:', error)
        throw error
      }
    }
  }
}) 