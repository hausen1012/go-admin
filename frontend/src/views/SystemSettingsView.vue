<template>
  <div class="settings-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>系统设置</span>
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
            @change="value => handleSettingChange('system_name', value)"
          />
          <span class="form-item-tip">设置系统的显示名称（最多6个字符）</span>
        </el-form-item>

        <!-- 允许注册 -->
        <el-form-item label="允许注册">
          <el-switch
            v-model="settings.allow_registration"
            @change="value => handleSettingChange('allow_registration', value.toString())"
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
import axios from 'axios'
import { useSystemStore } from '@/stores/system'

const loading = ref(false)
const systemStore = useSystemStore()
const settingsFormRef = ref(null)
const settings = ref({
  system_name: '',
  allow_registration: false
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
    const response = await axios.get('/api/admin/options')
    const options = response.data
    
    // 将配置数据转换为表单数据
    options.forEach(option => {
      if (option.OptionName === 'system_name') {
        settings.value.system_name = option.OptionValue
        systemStore.systemName = option.OptionValue
      } else if (option.OptionName === 'allow_registration') {
        settings.value.allow_registration = option.OptionValue === 'true'
        systemStore.allowRegistration = option.OptionValue === 'true'
      }
    })
  } catch (error) {
    ElMessage.error('获取系统配置失败')
    console.error('获取系统配置失败:', error)
  } finally {
    loading.value = false
  }
}

// 更新配置
const handleSettingChange = async (optionName, value) => {
  try {
    // 如果是系统名称，先进行表单验证
    if (optionName === 'system_name') {
      try {
        await settingsFormRef.value.validateField('system_name')
      } catch (validationError) {
        // 验证失败，直接返回
        return
      }
    }
    
    loading.value = true
    await axios.put(`/api/admin/options/${optionName}`, {
      option_value: value
    })
    
    // 更新 store 中的配置
    if (optionName === 'system_name') {
      systemStore.systemName = value
    } else if (optionName === 'allow_registration') {
      systemStore.allowRegistration = value === 'true'
    }
    
    ElMessage.success('设置更新成功')
  } catch (error) {
    ElMessage.error('设置更新失败')
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
</style> 