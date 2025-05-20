<template>
  <div class="settings-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span class="title">系统设置</span>
          <el-button 
            type="primary" 
            @click="handleSaveSettings" 
            :loading="loading"
          >
            保存
          </el-button>
        </div>
      </template>

      <el-form
        v-loading="loading"
        :model="settings"
        :rules="rules"
        ref="settingsFormRef"
        label-width="auto"
        :label-position="'right'"
        class="settings-form"
      >
        <!-- 系统名称 -->
        <el-form-item label="系统名称" prop="system_name">
          <el-input
            v-model="settings.system_name"
            placeholder="请输入系统名称"
            @change="handleFormChange"
          />
          <span class="form-item-tip">设置系统的显示名称（最多6个字符）</span>
        </el-form-item>

        <!-- 允许注册 -->
        <el-form-item label="允许注册">
          <el-switch
            v-model="settings.allow_registration"
            @change="handleFormChange"
          />
          <span class="form-item-tip">开启后允许新用户自行注册账号</span>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useSystemStore } from '@/stores/system'

const loading = ref(false)
const systemStore = useSystemStore()
const settingsFormRef = ref(null)
const settings = ref({
  system_name: '',
  allow_registration: false,
  hasChanges: false // 用于追踪是否有更改
})

// 表单验证规则
const rules = {
  system_name: [
    { required: true, message: '系统名称不能为空', trigger: 'blur' },
    { max: 6, message: '系统名称最多6个字符', trigger: 'blur' }
  ]
}

// 获取系统配置
const fetchSettings = async () => {
  try {
    loading.value = true
    const options = await systemStore.loadSystemSettings()
    
    // 将配置数据转换为表单数据
    options.forEach(option => {
      if (option.OptionName === 'system_name') {
        settings.value.system_name = option.OptionValue
      } else if (option.OptionName === 'allow_registration') {
        settings.value.allow_registration = option.OptionValue === 'true'
      }
    })
  } catch (error) {
    ElMessage.error('获取系统配置失败')
    console.error('获取系统配置失败:', error)
  } finally {
    loading.value = false
  }
}

// 表单变更处理
const handleFormChange = () => {
  settings.value.hasChanges = true
}

// 保存所有设置
const handleSaveSettings = async () => {
  try {
    // 表单验证
    await settingsFormRef.value.validate()
    
    loading.value = true
    const settingsData = [
      {
        option_name: 'system_name',
        option_value: settings.value.system_name
      },
      {
        option_name: 'allow_registration',
        option_value: settings.value.allow_registration.toString()
      }
    ]

    await systemStore.updateSystemSettings(settingsData)
    settings.value.hasChanges = false
    ElMessage.success('设置更新成功')
  } catch (error) {
    if (error?.message) {
      ElMessage.error(error.message)
    } else {
      ElMessage.error('设置更新失败')
    }
    console.error('设置更新失败:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchSettings()
})
</script>

<style scoped>
.settings-container {
  padding: 20px;
}

.settings-form {
  max-width: 800px;
}

.form-item-tip {
  display: block;
  margin-top: 4px;
  color: #909399;
  font-size: 13px;
  width: 100%;
  clear: both;
}

:deep(.el-form-item) {
  margin-bottom: 24px;
}

:deep(.el-form-item__label) {
  padding-right: 12px;
  align-self: flex-start;
  margin-top: 0;
  line-height: 28px;
}

:deep(.el-form-item__content) {
  flex-wrap: wrap;
}

:deep(.el-input__wrapper) {
  margin-top: 0;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

@media screen and (max-width: 768px) {
  .title {
    display: none;
  }
  
  .card-header {
    justify-content: flex-end;
  }
}
</style> 