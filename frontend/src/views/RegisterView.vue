<template>
  <auth-layout>
    <el-card class="register-card">
      <template #header>
        <h2>用户注册</h2>
      </template>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input v-model="form.confirmPassword" type="password" placeholder="请再次输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <div class="button-container">
            <div class="button-group">
              <el-button @click="$router.push('/login')">返回</el-button>
              <el-button type="primary" @click="handleRegister" :loading="loading">注册</el-button>
            </div>
          </div>
        </el-form-item>
      </el-form>
    </el-card>
  </auth-layout>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { useSystemStore } from '@/stores/system'
import { ElMessage } from 'element-plus'
import AuthLayout from '@/components/AuthLayout.vue'

const router = useRouter()
const userStore = useUserStore()
const systemStore = useSystemStore()
const formRef = ref(null)
const loading = ref(false)

// 使用计算属性获取注册状态
const allowRegistration = computed(() => systemStore.allowRegistration)

onMounted(() => {
  // 如果不允许注册，重定向到登录页面
  if (!allowRegistration.value) {
    ElMessage.error('注册功能已禁用')
    router.push('/login')
  }
})

const form = reactive({
  username: '',
  password: '',
  confirmPassword: ''
})

const validatePass = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请输入密码'))
  } else if (value.length < 5) {
    callback(new Error('密码长度不能小于5位'))
  } else {
    if (form.confirmPassword !== '') {
      formRef.value?.validateField('confirmPassword')
    }
    callback()
  }
}

const validatePass2 = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请再次输入密码'))
  } else if (value !== form.password) {
    callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, validator: validatePass, trigger: 'blur' }],
  confirmPassword: [{ required: true, validator: validatePass2, trigger: 'blur' }]
}

const handleRegister = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        await userStore.register(form.username, form.password)
        ElMessage.success('注册成功')
        router.push('/login')
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
.register-container {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #1976d2 0%, #64b5f6 100%);
  padding: 20px;
}

.register-card {
  width: 100%;
  max-width: 400px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  background: #fff;
}

.register-card :deep(.el-card__header) {
  text-align: center;
  padding: 20px;
  border-bottom: 1px solid #ebeef5;
}

.register-card h2 {
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

.button-group .el-button {
  flex: 0 0 auto;
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
  .register-card {
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