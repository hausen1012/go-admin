<template>
  <div class="users-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span class="title">用户管理</span>
          <div class="header-right">
            <el-input
              v-model="searchQuery"
              placeholder="搜索用户名"
              style="width: 200px; margin-right: 16px;"
              clearable
              @input="handleSearch"
            >
              <template #prefix>
                <el-icon><Search /></el-icon>
              </template>
            </el-input>
            <el-button type="primary" @click="handleAdd">添加用户</el-button>
          </div>
        </div>
      </template>

      <el-table :data="paginatedUsers" style="width: 100%" v-loading="loading">
        <el-table-column prop="id" label="ID" width="60" />
        <el-table-column prop="username" label="用户名" min-width="80" />
        <el-table-column prop="is_admin" label="管理员" width="80">
          <template #default="{ row }">
            <el-tag :type="row.is_admin ? 'danger' : 'info'">
              {{ row.is_admin ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" min-width="120" class-name="hide-on-mobile">
          <template #default="{ row }">
            {{ new Date(row.created_at).toLocaleString() }}
          </template>
        </el-table-column>
        <el-table-column label="操作" min-width="160" class-name="action-column">
          <template #default="{ row }">
            <div class="action-buttons">
              <el-link 
                type="primary" 
                @click="handleEdit(row)" 
                style="margin-right: 8px;"
                :disabled="row.username === 'admin'"
                v-if="row.username !== 'admin'"
              >编辑</el-link>
              <el-link 
                type="primary" 
                @click="handleResetPassword(row)" 
                style="margin-right: 8px;"
                v-if="row.username !== 'admin' || (row.username === 'admin' && isCurrentUserAdmin)"
              >重置密码</el-link>
              <el-link 
                type="danger" 
                @click="handleDelete(row)"
                :disabled="row.username === 'admin'"
                v-if="row.username !== 'admin'"
              >删除</el-link>
            </div>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="totalUsers"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 添加/编辑用户对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogType === 'add' ? '添加用户' : '编辑用户'"
      width="500px"
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item v-if="dialogType === 'add'" label="密码" prop="password">
          <el-input v-model="form.password" type="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item label="管理员">
          <el-switch v-model="form.is_admin"></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit" :loading="submitting">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useUserStore } from '@/stores/user'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRouter } from 'vue-router'
import { Search } from '@element-plus/icons-vue'

const userStore = useUserStore()
const router = useRouter()
const users = ref([])
const loading = ref(false)
const dialogVisible = ref(false)
const dialogType = ref('add')
const submitting = ref(false)
const formRef = ref(null)
const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref(10)
const totalUsers = ref(0)

const form = reactive({
  id: null,
  username: '',
  password: '',
  is_admin: false
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const filteredUsers = computed(() => {
  if (!searchQuery.value) {
    return users.value
  }
  return users.value.filter(user => 
    user.username.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

const paginatedUsers = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredUsers.value.slice(start, end)
})

const handleSizeChange = (val) => {
  pageSize.value = val
  currentPage.value = 1
}

const handleCurrentChange = (val) => {
  currentPage.value = val
}

const fetchUsers = async () => {
  loading.value = true
  try {
    const allUsers = await userStore.fetchUsers()
    users.value = allUsers
    totalUsers.value = allUsers.length
  } catch (error) {
    ElMessage.error(error)
    if (error.includes('未提供认证信息') || error.includes('无效的认证信息')) {
      userStore.logout()
      router.push('/login')
    }
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  dialogType.value = 'add'
  form.id = null
  form.username = ''
  form.password = ''
  form.is_admin = false
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogType.value = 'edit'
  form.id = row.id
  form.username = row.username
  form.is_admin = row.is_admin
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      submitting.value = true
      try {
        if (dialogType.value === 'add') {
          await userStore.createUser(form.username, form.password)
          ElMessage.success('用户创建成功')
        } else {
          await userStore.updateUser(form.id, {
            username: form.username,
            is_admin: form.is_admin
          })
          ElMessage.success('用户更新成功')
        }
        dialogVisible.value = false
        fetchUsers()
      } catch (error) {
        ElMessage.error(error)
        if (error.includes('未提供认证信息') || error.includes('无效的认证信息')) {
          userStore.logout()
          router.push('/login')
        }
      } finally {
        submitting.value = false
      }
    }
  })
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该用户吗？', '警告', {
      type: 'warning'
    })
    await userStore.deleteUser(row.id)
    ElMessage.success('用户删除成功')
    fetchUsers()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(error)
      if (error.includes('未提供认证信息') || error.includes('无效的认证信息')) {
        userStore.logout()
        router.push('/login')
      }
    }
  }
}

const handleResetPassword = async (row) => {
  try {
    await ElMessageBox.confirm('确定要重置该用户的密码吗？', '警告', {
      type: 'warning'
    })
    const result = await userStore.resetUserPassword(row.id)
    ElMessage.success(`密码重置成功，新密码为：${result.password}`)
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(error)
      if (error.includes('未提供认证信息') || error.includes('无效的认证信息')) {
        userStore.logout()
        router.push('/login')
      }
    }
  }
}

const handleSearch = () => {
  // 搜索逻辑已通过计算属性实现
}

// 添加计算属性来判断当前用户是否为admin
const isCurrentUserAdmin = computed(() => {
  return userStore.user?.username === 'admin'
})

onMounted(() => {
  fetchUsers()
})
</script>

<style scoped>
.users-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-right {
  display: flex;
  align-items: center;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

@media screen and (max-width: 768px) {
  .users-container {
    padding: 10px;
  }
  
  .header-right .el-input {
    width: 180px !important;
    margin-right: 10px !important;
  }
  
  .pagination-container {
    justify-content: center;
  }
  
  :deep(.hide-on-mobile) {
    display: none;
  }
  
  .title {
    display: none;
  }
  
  .card-header {
    justify-content: flex-end;
  }
  
  :deep(.el-table) {
    width: 100% !important;
  }

  :deep(.action-column) {
    width: auto !important;
    min-width: 120px !important;
  }

  :deep(.el-table__header-wrapper),
  :deep(.el-table__body-wrapper) {
    width: 100% !important;
  }

  :deep(.el-table__cell) {
    padding: 8px 4px !important;
  }
}

@media screen and (max-width: 576px) {
  .card-header {
    flex-direction: column;
    align-items: stretch;
    gap: 10px;
  }
  
  .header-right {
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
  }
  
  .header-right .el-input {
    flex: 1;
    margin-right: 10px !important;
  }
  
  .header-right .el-button {
    width: auto;
  }
  
  .action-buttons {
    display: flex;
    flex-direction: column;
    gap: 5px;
    align-items: flex-start;
  }
  
  .action-buttons .el-link {
    margin-right: 0 !important;
  }
  
  :deep(.el-table) {
    font-size: 12px;
  }

  :deep(.action-buttons) {
    display: flex;
    flex-wrap: wrap;
    gap: 4px;
  }

  :deep(.action-buttons .el-link) {
    font-size: 12px;
    margin-right: 4px !important;
  }
}

@media screen and (max-width: 480px) {
  :deep(.el-table) {
    font-size: 12px;
  }
  
  :deep(.el-table .cell) {
    padding: 0 5px;
  }
  
  :deep(.el-pagination) {
    flex-wrap: wrap;
    justify-content: center;
  }
}
</style> 