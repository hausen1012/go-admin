<template>
  <div class="page-container">
    <div class="auth-container">
      <div class="left-panel" v-if="showLeftPanel">
        <div class="welcome-content">
          <img src="@/assets/logo.svg" alt="Logo" class="logo" />
          <h1>
            <span>欢迎使用</span>
            <span>{{ systemName }}</span>
          </h1>
          <p class="description">简洁、高效、安全的管理平台</p>
        </div>
      </div>
      <div class="right-panel">
        <div class="form-container">
          <slot></slot>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const systemName = ref('后台管理系统')
const windowWidth = ref(window.innerWidth)
const breakpointMedium = 992
const breakpointSmall = 768

const showLeftPanel = computed(() => {
  return windowWidth.value > breakpointMedium
})

const handleResize = () => {
  windowWidth.value = window.innerWidth
}

onMounted(async () => {
  window.addEventListener('resize', handleResize)
  try {
    const sysInfo = await userStore.getSysInfo()
    systemName.value = sysInfo.systemName
  } catch (error) {
    console.error('获取系统信息失败:', error)
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>

<style scoped>
.page-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f0f2f5;
  padding: 20px;
  box-sizing: border-box;
}

.auth-container {
  width: 90%;
  max-width: 1200px;
  min-height: 600px;
  display: flex;
  background-color: #fff;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}

.left-panel {
  flex: 1;
  background: linear-gradient(135deg, #1976d2 0%, #64b5f6 100%);
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px;
  position: relative;
  overflow: hidden;
}

.welcome-content {
  text-align: center;
  position: relative;
  z-index: 1;
}

.logo {
  width: 120px;
  height: 120px;
  margin-bottom: 24px;
}

.left-panel h1 {
  color: #ffffff;
  font-size: 36px;
  font-weight: 300;
  line-height: 1.4;
  margin: 0 0 16px;
}

.left-panel h1 span {
  display: block;
}

.description {
  color: rgba(255, 255, 255, 0.8);
  font-size: 16px;
  margin: 0;
}

.right-panel {
  width: 400px;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px;
}

.form-container {
  width: 100%;
  max-width: 400px;
}

/* 中等屏幕适配 */
@media screen and (max-width: 992px) {
  .right-panel {
    width: 100%;
  }
}

/* 移动端适配 */
@media screen and (max-width: 768px) {
  .page-container {
    padding: 0;
    background-color: #fff;
  }

  .auth-container {
    width: 100%;
    min-height: 100vh;
    box-shadow: none;
    border-radius: 0;
  }

  .right-panel {
    padding: 20px;
  }

  .form-container {
    max-width: 100%;
  }
}
</style> 