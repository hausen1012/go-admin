<template>
  <el-container class="layout-container">
    <!-- 移动端抽屉菜单 -->
    <el-drawer
      v-model="menuVisible"
      direction="ltr"
      size="240px"
      :with-header="false"
      v-if="isMobile"
    >
      <div class="drawer-content">
        <div class="logo">
          <router-link to="/" class="logo-link" @click="closeMenu">
            <img src="/logo.svg" alt="系统logo" class="logo-img" />
            <span>{{ systemStore.systemName.length > 6 ? systemStore.systemName.substring(0, 6) : systemStore.systemName }}</span>
          </router-link>
        </div>
        <el-menu
          :router="true"
          :default-active="$route.path"
          class="el-menu-vertical"
          :collapse-transition="false"
          @select="closeMenu"
        >
          <el-menu-item index="/">
            <el-icon><HomeFilled /></el-icon>
            <span>首页</span>
          </el-menu-item>
          <el-menu-item v-if="userStore.isAdmin" index="/users">
            <el-icon><UserFilled /></el-icon>
            <span>用户管理</span>
          </el-menu-item>
          <el-menu-item v-if="userStore.isAdmin" index="/settings">
            <el-icon><Setting /></el-icon>
            <span>系统设置</span>
          </el-menu-item>
        </el-menu>
      </div>
    </el-drawer>

    <!-- 桌面端侧边栏 -->
    <el-aside width="240px" class="aside" v-if="!isMobile">
      <div class="logo">
        <router-link to="/" class="logo-link">
          <img src="/logo.svg" alt="系统logo" class="logo-img" />
          <span>{{ systemStore.systemName.length > 6 ? systemStore.systemName.substring(0, 6) : systemStore.systemName }}</span>
        </router-link>
      </div>
      <el-menu
        :router="true"
        :default-active="$route.path"
        class="el-menu-vertical"
        :collapse-transition="false"
      >
        <el-menu-item index="/">
          <el-icon><HomeFilled /></el-icon>
          <span>首页</span>
        </el-menu-item>
        <el-menu-item v-if="userStore.isAdmin" index="/users">
          <el-icon><UserFilled /></el-icon>
          <span>用户管理</span>
        </el-menu-item>
        <el-menu-item v-if="userStore.isAdmin" index="/settings">
          <el-icon><Setting /></el-icon>
          <span>系统设置</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="header">
        <div class="header-content">
          <div class="header-left">
            <div class="menu-btn" @click="toggleMenu" v-if="isMobile">
              <el-button type="primary" circle>
                <el-icon><Menu /></el-icon>
              </el-button>
            </div>
            <el-breadcrumb separator="/" v-if="!isMobile">
              <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
              <el-breadcrumb-item>{{ currentRoute }}</el-breadcrumb-item>
            </el-breadcrumb>
            <span class="current-page" v-else>{{ currentRoute }}</span>
          </div>
          <div class="header-right">
            <el-dropdown @command="handleCommand" trigger="click">
              <span class="user-dropdown">
                <el-avatar :size="isMobile ? 28 : 32" class="user-avatar">
                  {{ userStore.user?.username?.charAt(0)?.toUpperCase() }}
                </el-avatar>
                <span class="username" v-if="!isMobile">{{ userStore.user?.username }}</span>
                <el-icon v-if="!isMobile"><ArrowDown /></el-icon>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item command="changePassword">
                    <el-icon><Key /></el-icon>修改密码
                  </el-dropdown-item>
                  <el-dropdown-item command="logout">
                    <el-icon><SwitchButton /></el-icon>退出登录
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>
      </el-header>
      <el-main>
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </el-main>
    </el-container>

    <!-- 修改密码弹窗 -->
    <el-dialog
      v-model="passwordDialogVisible"
      title="修改密码"
      width="400px"
      :close-on-click-modal="false"
    >
      <el-form
        ref="passwordFormRef"
        :model="passwordForm"
        :rules="passwordRules"
        label-width="100px"
      >
        <el-form-item label="原密码" prop="oldPassword">
          <el-input
            v-model="passwordForm.oldPassword"
            type="password"
            show-password
            placeholder="请输入原密码"
          />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input
            v-model="passwordForm.newPassword"
            type="password"
            show-password
            placeholder="请输入新密码"
          />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
            v-model="passwordForm.confirmPassword"
            type="password"
            show-password
            placeholder="请再次输入新密码"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="passwordDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleChangePassword">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </el-container>
</template>

<script setup>
import { useUserStore } from '@/stores/user'
import { useSystemStore } from '@/stores/system'
import { useRouter, useRoute } from 'vue-router'
import { HomeFilled, UserFilled, ArrowDown, SwitchButton, Key, Setting, Menu } from '@element-plus/icons-vue'
import { computed, ref, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { updatePageTitle } from '@/utils/title'

const userStore = useUserStore()
const systemStore = useSystemStore()
const router = useRouter()
const route = useRoute()

// 移动端相关
const isMobile = ref(window.innerWidth <= 768)
const menuVisible = ref(false)

// 监听窗口大小变化
window.addEventListener('resize', () => {
  isMobile.value = window.innerWidth <= 768
  if (!isMobile.value) {
    menuVisible.value = false
  }
})

const toggleMenu = () => {
  menuVisible.value = !menuVisible.value
}

const closeMenu = () => {
  menuVisible.value = false
}

onMounted(() => {
  systemStore.loadSystemSettings()
})

const currentRoute = computed(() => {
  const routeMap = {
    '/': '首页',
    '/users': '用户管理',
    '/settings': '系统设置'
  }
  return routeMap[route.path] || '未知页面'
})

// 监听路由和系统名称变化，更新页面标题
watch(
  [() => route.path, () => systemStore.systemName],
  ([path, systemName]) => {
    updatePageTitle(systemName, currentRoute.value)
  },
  { immediate: true }
)

const passwordDialogVisible = ref(false)
const passwordFormRef = ref(null)
const passwordForm = ref({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const validateConfirmPassword = (rule, value, callback) => {
  if (value !== passwordForm.value.newPassword) {
    callback(new Error('两次输入的密码不一致'))
  } else {
    callback()
  }
}

const passwordRules = {
  oldPassword: [
    { required: true, message: '请输入原密码', trigger: 'blur' }
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 5, message: '密码长度不能小于5位', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入新密码', trigger: 'blur' },
    { validator: validateConfirmPassword, trigger: 'blur' }
  ]
}

const handleCommand = (command) => {
  if (command === 'logout') {
    userStore.logout()
    router.push('/login')
  } else if (command === 'changePassword') {
    passwordDialogVisible.value = true
  }
}

const handleChangePassword = async () => {
  if (!passwordFormRef.value) return
  
  await passwordFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        await userStore.updatePassword(
          passwordForm.value.oldPassword,
          passwordForm.value.newPassword
        )
        ElMessage.success('密码修改成功')
        passwordDialogVisible.value = false
        passwordForm.value = {
          oldPassword: '',
          newPassword: '',
          confirmPassword: ''
        }
      } catch (error) {
        ElMessage.error(error.message || '密码修改失败')
      }
    }
  })
}
</script>

<style scoped>
.layout-container {
  height: 100vh;
  margin: 0;
  padding: 0;
}

.aside {
  background: #fff;
  border-right: 1px solid #e5e6eb;
  box-shadow: none;
  padding: 0;
  min-width: 220px;
}

.logo {
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
  border-bottom: 1px solid #f0f0f0;
  padding: 0 16px;
}

.logo-link {
  text-decoration: none;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo-link span {
  font-size: 24px;
  font-weight: bold;
  color: #1976d2;
  letter-spacing: 1px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.el-menu-vertical {
  background: #fff;
  border-right: none;
  min-height: calc(100vh - 64px);
  padding-top: 8px;
}

.el-menu-vertical :deep(.el-menu-item),
.el-menu-vertical :deep(.el-sub-menu__title) {
  height: 44px;
  line-height: 44px;
  border-radius: 8px;
  margin: 4px 12px;
  color: #222;
  font-size: 15px;
  font-weight: 500;
  transition: background 0.2s, color 0.2s;
  display: flex;
  align-items: center;
  gap: 10px;
}

.el-menu-vertical :deep(.el-menu-item.is-active),
.el-menu-vertical :deep(.el-sub-menu.is-opened > .el-sub-menu__title) {
  background: #e8f3ff !important;
  color: #1976d2 !important;
}

.el-menu-vertical :deep(.el-menu-item:hover),
.el-menu-vertical :deep(.el-sub-menu__title:hover) {
  background: #f4f6fa !important;
  color: #1976d2;
}

.el-menu-vertical :deep(.el-icon) {
  font-size: 18px;
  margin-right: 6px;
  color: #1976d2;
}

.el-menu-vertical :deep(.el-menu-item-group__title) {
  display: none;
}

.aside {
  border-radius: 12px 0 0 12px;
  overflow: hidden;
}

.header {
  background-color: #fff;
  border-bottom: 1px solid #f0f0f0;
  padding: 0 20px;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  height: 60px;
  line-height: 60px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
}

.header-left {
  display: flex;
  align-items: center;
}

.header-right {
  display: flex;
  align-items: center;
}

.user-dropdown {
  cursor: pointer;
  display: flex;
  align-items: center;
  padding: 0 8px;
  height: 48px;
  border-radius: 24px;
  transition: all 0.3s;
}

.user-dropdown:hover {
  background: rgba(0, 0, 0, 0.025);
}

.user-avatar {
  background-color: #1890ff;
  color: #fff;
  font-weight: 500;
}

.username {
  margin: 0 8px;
  font-size: 14px;
  color: #333;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

:deep(.el-dropdown-menu__item) {
  display: flex;
  align-items: center;
  gap: 8px;
}

:deep(.el-breadcrumb__inner) {
  font-weight: 500;
}

:deep(.el-main) {
  padding: 0;
  background-color: #f5f7fa;
}

:deep(.el-menu) {
  border-right: none;
}

:deep(.el-menu-item) {
  margin: 4px 16px;
}

:deep(.el-menu-item.is-active) {
  background-color: #1890ff;
}

/* 移动端菜单按钮样式 */
.menu-btn {
  margin-right: 16px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.menu-btn :deep(.el-button) {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

/* 移除旧的固定菜单按钮样式 */
.mobile-menu-btn {
  display: none;
}

/* 响应式调整 */
@media screen and (max-width: 768px) {
  .header {
    padding: 0 12px;
  }

  .header-content {
    padding-left: 0;
  }
  
  .current-page {
    font-size: 16px;
    font-weight: 500;
    color: #333;
  }
  
  .drawer-content {
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  .user-dropdown {
    padding: 4px;
  }

  :deep(.el-main) {
    padding: 12px;
  }
}

/* 抽屉菜单样式优化 */
:deep(.el-drawer) {
  --el-drawer-bg-color: #fff;
  border-right: 1px solid #e5e6eb;
}

:deep(.el-drawer__body) {
  padding: 0;
}

.logo-img {
  width: 32px;
  height: 32px;
  margin-right: 10px;
  vertical-align: middle;
}
</style> 