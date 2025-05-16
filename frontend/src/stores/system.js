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
    }
  }
}) 