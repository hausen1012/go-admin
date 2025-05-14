<template>
  <div class="home-container">
    <!-- 统计卡片 -->
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="6" v-for="(card, index) in statisticsCards" :key="index">
        <el-card class="statistics-card" :body-style="{ padding: '20px' }">
          <div class="card-content">
            <el-icon class="card-icon" :class="card.color">
              <component :is="card.icon" />
            </el-icon>
            <div class="card-info">
              <div class="card-value">{{ card.value }}</div>
              <div class="card-label">{{ card.label }}</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 主要内容区域 -->
    <el-row :gutter="20" class="mt-20">
      <!-- 左侧区域 -->
      <el-col :xs="24" :sm="24" :md="16">
        <!-- 系统公告 -->
        <el-card class="mb-20">
          <template #header>
            <div class="card-header">
              <span>系统公告</span>
              <el-button type="primary" link>查看全部</el-button>
            </div>
          </template>
          <div class="announcement-list">
            <div v-for="(notice, index) in announcements" :key="index" class="announcement-item">
              <el-tag :type="notice.type" size="small" class="mr-10">{{ notice.tag }}</el-tag>
              <span class="announcement-title">{{ notice.title }}</span>
              <span class="announcement-time">{{ notice.time }}</span>
            </div>
          </div>
        </el-card>

        <!-- 图表区域 -->
        <el-card>
          <template #header>
            <div class="card-header">
              <span>数据统计</span>
              <el-radio-group v-model="activityTimeRange" size="small">
                <el-radio-button label="week">本周</el-radio-button>
                <el-radio-button label="month">本月</el-radio-button>
                <el-radio-button label="year">全年</el-radio-button>
              </el-radio-group>
            </div>
          </template>
          <!-- 图表内容 -->
        </el-card>
      </el-col>

      <!-- 右侧区域 -->
      <el-col :xs="24" :sm="24" :md="8">
        <!-- 快捷操作 -->
        <el-card class="mb-20">
          <template #header>
            <div class="card-header">
              <span>快捷操作</span>
            </div>
          </template>
          <div class="quick-actions">
            <el-button
              v-for="action in quickActions"
              :key="action.name"
              :type="action.type"
              class="quick-action-btn"
              @click="handleQuickAction(action.name)"
            >
              <el-icon><component :is="action.icon" /></el-icon>
              <span>{{ action.label }}</span>
            </el-button>
          </div>
        </el-card>

        <!-- 待办事项 -->
        <el-card>
          <template #header>
            <div class="card-header">
              <span>待办事项</span>
              <el-button type="primary" link @click="addTodo">
                <el-icon><Plus /></el-icon>添加
              </el-button>
            </div>
          </template>
          <el-checkbox-group v-model="checkedTodos">
            <div v-for="todo in todos" :key="todo.id" class="todo-item">
              <el-checkbox :label="todo.id">
                <span :class="{ 'todo-done': checkedTodos.includes(todo.id) }">
                  {{ todo.content }}
                </span>
              </el-checkbox>
              <el-tag size="small" :type="todo.priority">{{ todo.priorityLabel }}</el-tag>
            </div>
          </el-checkbox-group>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  User, View, Bell, Calendar,
  Plus, Setting, Upload, Download,
  Edit, Delete
} from '@element-plus/icons-vue'

// 统计卡片数据
const statisticsCards = [
  {
    label: '待处理任务',
    value: '8',
    icon: 'Bell',
    color: 'orange'
  },
  {
    label: '本月活跃用户',
    value: '128',
    icon: 'User',
    color: 'blue'
  },
  {
    label: '今日访问量',
    value: '1,234',
    icon: 'View',
    color: 'green'
  },
  {
    label: '待审核内容',
    value: '6',
    icon: 'Calendar',
    color: 'purple'
  }
]

// 系统公告数据
const announcements = [
  {
    title: '系统将于本周日进行例行维护，请提前做好准备',
    time: '2024-03-24',
    type: 'warning',
    tag: '维护'
  },
  {
    title: '新功能上线：用户现可以自定义仪表盘展示内容',
    time: '2024-03-22',
    type: 'success',
    tag: '更新'
  },
  {
    title: '关于加强系统安全性的重要通知',
    time: '2024-03-20',
    type: 'info',
    tag: '通知'
  }
]

// 最近活动数据
const activityTimeRange = ref('week')
const activities = [
  {
    content: '张三创建了新的用户账号',
    time: '2024-03-24 10:30',
    type: 'primary',
    color: '#409EFF'
  },
  {
    content: '系统自动备份完成',
    time: '2024-03-24 03:00',
    type: 'success',
    color: '#67C23A'
  },
  {
    content: '李四更新了系统配置',
    time: '2024-03-23 15:20',
    type: 'warning',
    color: '#E6A23C'
  },
  {
    content: '王五删除了过期数据',
    time: '2024-03-23 11:45',
    type: 'info',
    color: '#909399'
  }
]

// 快捷操作
const quickActions = [
  {
    name: 'createUser',
    label: '创建用户',
    icon: Plus,
    type: 'primary'
  },
  {
    name: 'systemSettings',
    label: '系统设置',
    icon: Setting,
    type: 'info'
  },
  {
    name: 'importData',
    label: '导入数据',
    icon: Upload,
    type: 'success'
  },
  {
    name: 'exportData',
    label: '导出数据',
    icon: Download,
    type: 'warning'
  }
]

// 待办事项数据
const todos = [
  {
    id: 1,
    content: '审核新用户申请',
    priority: 'danger',
    priorityLabel: '紧急'
  },
  {
    id: 2,
    content: '更新系统文档',
    priority: 'warning',
    priorityLabel: '重要'
  },
  {
    id: 3,
    content: '检查系统日志',
    priority: 'info',
    priorityLabel: '普通'
  }
]
const checkedTodos = ref([])

// 方法
const handleQuickAction = (actionName) => {
  // 处理快捷操作点击事件
  console.log('Quick action clicked:', actionName)
}

const addTodo = () => {
  // 添加待办事项
  console.log('Add todo clicked')
}
</script>

<style scoped>
.home-container {
  padding: 20px;
}

.mt-20 {
  margin-top: 20px;
}

.mb-20 {
  margin-bottom: 20px;
}

.mr-10 {
  margin-right: 10px;
}

.statistics-card {
  height: 120px;
  transition: all 0.3s;
}

.statistics-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);
}

.card-content {
  display: flex;
  align-items: center;
  height: 100%;
}

.card-icon {
  font-size: 48px;
  margin-right: 20px;
  padding: 10px;
  border-radius: 8px;
}

.blue {
  background-color: #e6f7ff;
  color: #1890ff;
}

.green {
  background-color: #f6ffed;
  color: #52c41a;
}

.orange {
  background-color: #fff7e6;
  color: #fa8c16;
}

.purple {
  background-color: #f9f0ff;
  color: #722ed1;
}

.card-info {
  flex: 1;
}

.card-value {
  font-size: 24px;
  font-weight: bold;
  color: #333;
  margin-bottom: 8px;
}

.card-label {
  font-size: 14px;
  color: #666;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.announcement-list {
  .announcement-item {
    display: flex;
    align-items: center;
    padding: 12px 0;
    border-bottom: 1px solid #f0f0f0;

    &:last-child {
      border-bottom: none;
    }
  }

  .announcement-title {
    flex: 1;
    margin: 0 12px;
    color: #333;
  }

  .announcement-time {
    color: #999;
    font-size: 13px;
  }
}

.activity-list {
  padding: 0 20px;
}

.quick-actions {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 12px;
}

.quick-action-btn {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.todo-list {
  .todo-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 0;
    border-bottom: 1px solid #f0f0f0;

    &:last-child {
      border-bottom: none;
    }
  }
}

/* 移动端适配 */
@media screen and (max-width: 768px) {
  .home-container {
    padding: 12px;
  }

  .statistics-card {
    margin-bottom: 12px;
  }

  .mt-20 {
    margin-top: 12px;
  }

  .mb-20 {
    margin-bottom: 12px;
  }

  .card-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .announcement-item {
    flex-direction: column;
    gap: 8px;
  }

  .announcement-title {
    width: 100%;
  }

  .announcement-time {
    width: 100%;
    text-align: left;
  }

  .quick-actions {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 8px;
  }

  .quick-action-btn {
    width: 100%;
  }

  .todo-item {
    padding: 8px 0;
  }
}

/* 基础样式优化 */
.quick-actions {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 12px;
}

.quick-action-btn {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.announcement-item {
  display: flex;
  align-items: center;
  padding: 12px 0;
  border-bottom: 1px solid #f0f0f0;
}

.announcement-item:last-child {
  border-bottom: none;
}

.announcement-title {
  flex: 1;
  margin: 0 12px;
}

.announcement-time {
  color: #909399;
  font-size: 13px;
}

.todo-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 0;
  border-bottom: 1px solid #f0f0f0;
}

.todo-item:last-child {
  border-bottom: none;
}

.todo-done {
  text-decoration: line-through;
  color: #909399;
}
</style>