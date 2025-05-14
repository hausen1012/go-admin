<template>
  <auth-layout>
    <el-card class="login-card">
      <template #header>
        <h2>用户登录</h2>
      </template>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <div class="button-container">
            <div class="button-group">
              <el-button v-if="allowRegistration" @click="$router.push('/register')">注册</el-button>
              <el-button type="primary" @click="handleLogin" :loading="loading">登录</el-button>
            </div>
          </div>
        </el-form-item>
      </el-form>
    </el-card>
  </auth-layout>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'
import AuthLayout from '@/components/AuthLayout.vue'

const router = useRouter()
const userStore = useUserStore()
const formRef = ref(null)
const loading = ref(false)
const allowRegistration = ref(false)

const form = reactive({
  username: '',
  password: ''
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

onMounted(async () => {
  // 获取注册状态
  allowRegistration.value = await userStore.getRegistrationStatus()
})

const handleLogin = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        await userStore.login(form.username, form.password)
        ElMessage.success('登录成功')
        router.push('/')
      } catch (error) {
        ElMessage.error(error)
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #1976d2 0%, #64b5f6 100%);
  padding: 20px;
}

.login-card {
  width: 100%;
  max-width: 400px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  background: #fff;
}

.login-card :deep(.el-card__header) {
  text-align: center;
  padding: 20px;
  border-bottom: 1px solid #ebeef5;
}

.login-card h2 {
  margin: 0;
  color: #2c3e50;
  font-size: 24px;
  font-weight: 500;
}

:deep(.el-form) {
  padding: 24px 20px;
}

:deep(.el-form-item) {
  margin-bottom: 20px;
}

:deep(.el-input) {
  --el-input-hover-border-color: #409EFF;
  --el-input-focus-border-color: #409EFF;
}

:deep(.el-button) {
  padding: 12px 24px;
  font-weight: 500;
  transition: all 0.3s ease;
  margin-bottom: 0;
}

.button-container {
  width: 100%;
  display: flex;
  justify-content: flex-end;
}

.button-group {
  display: flex;
  gap: 12px;
}

:deep(.el-button--primary) {
  background: #1976d2;
  border: none;
}

:deep(.el-button--primary:hover) {
  background: #1565c0;
  transform: translateY(-2px);
}

:deep(.el-button--default) {
  border: 1px solid #dcdfe6;
}

:deep(.el-button--default:hover) {
  border-color: #1976d2;
  color: #1976d2;
}

/* 移动端适配 */
@media screen and (max-width: 768px) {
  .login-card {
    box-shadow: none;
    max-width: 100%;
  }

  :deep(.el-form) {
    padding: 20px 16px;
  }

  :deep(.el-form-item) {
    margin-bottom: 16px;
  }

  :deep(.el-button) {
    margin-bottom: 8px;
    width: 100%;
  }
}

@media screen and (max-width: 480px) {
  :deep(.el-form-item__label) {
    float: none;
    display: block;
    text-align: left;
    padding: 0 0 8px;
  }
  
  :deep(.el-form-item__content) {
    margin-left: 0 !important;
  }
  
  .button-group {
    flex-direction: column;
    width: 100%;
    gap: 8px;
  }
  
  .button-group .el-button {
    margin: 0;
  }
}
</style> 