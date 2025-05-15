<template>
  <div class="home-container">
    <!-- 系统监控卡片 -->
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="12" :lg="6" v-for="(card, index) in monitorCards" :key="index">
        <el-card class="monitor-card" :body-style="{ padding: '20px' }">
          <div class="card-content">
            <el-icon class="card-icon" :class="card.colorClass">
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

    <!-- 服务器信息 -->
    <el-row :gutter="20" class="mt-20">
      <el-col :span="24">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>服务器信息</span>
            </div>
          </template>
          <el-descriptions :column="descriptionColumns" border>
            <el-descriptions-item label="操作系统">{{ serverInfo.os }}</el-descriptions-item>
            <el-descriptions-item label="运行环境">{{ serverInfo.environment }}</el-descriptions-item>
            <el-descriptions-item label="服务器类型">{{ serverInfo.serverType }}</el-descriptions-item>
            <el-descriptions-item label="数据库">{{ serverInfo.database }}</el-descriptions-item>
            <el-descriptions-item label="服务器内存">{{ serverInfo.memory }}</el-descriptions-item>
            <el-descriptions-item label="系统版本">{{ serverInfo.version }}</el-descriptions-item>
          </el-descriptions>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import {
  User, Timer, Monitor, Cpu
} from '@element-plus/icons-vue'
import { useWindowSize } from '@vueuse/core'

// 窗口尺寸响应式
const { width } = useWindowSize()

// 根据窗口宽度计算描述列数
const descriptionColumns = computed(() => {
  if (width.value < 768) return 1
  if (width.value < 992) return 2
  return 3
})

// 监控卡片数据
const monitorCards = [
  {
    label: '用户数量',
    value: '256',
    icon: User,
    colorClass: 'blue'
  },
  {
    label: '系统运行时间',
    value: '34天12小时',
    icon: Timer,
    colorClass: 'green'
  },
  {
    label: 'CPU利用率',
    value: '32%',
    icon: Cpu,
    colorClass: 'orange'
  },
  {
    label: '内存利用率',
    value: '45%',
    icon: Monitor,
    colorClass: 'purple'
  }
]

// 服务器信息
const serverInfo = {
  os: 'Linux / Windows',
  environment: 'Go 1.20',
  serverType: 'Go HTTP Server',
  database: 'SQLite 3',
  memory: '系统分配',
  version: 'v1.0.0'
}
</script>

<style scoped>
.home-container {
  padding: 20px;
}

.mt-20 {
  margin-top: 20px;
}

.monitor-card {
  height: 120px;
  transition: all 0.3s;
  margin-bottom: 20px;
}

.monitor-card:hover {
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

/* 响应式布局适配 */
@media screen and (max-width: 1200px) {
  .card-icon {
    font-size: 42px;
    margin-right: 15px;
  }
  
  .card-value {
    font-size: 22px;
  }
}

@media screen and (max-width: 992px) {
  .home-container {
    padding: 16px;
  }
  
  .mt-20 {
    margin-top: 16px;
  }
  
  .monitor-card {
    height: 110px;
  }
}

@media screen and (max-width: 768px) {
  .home-container {
    padding: 12px;
  }

  .monitor-card {
    height: 100px;
    margin-bottom: 12px;
  }

  .mt-20 {
    margin-top: 12px;
  }
  
  .card-icon {
    font-size: 36px;
    margin-right: 12px;
    padding: 8px;
  }
  
  .card-value {
    font-size: 20px;
    margin-bottom: 4px;
  }

  .card-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
}

@media screen and (max-width: 576px) {
  .home-container {
    padding: 10px;
  }
  
  /* 移除小屏幕下的垂直布局，保持水平排列 */
  .card-content {
    flex-direction: row;
    text-align: left;
    align-items: center;
  }
  
  .card-icon {
    margin-right: 12px;
    margin-bottom: 0;
    font-size: 32px;
    padding: 6px;
  }
  
  .card-value {
    font-size: 18px;
  }
  
  .card-label {
    font-size: 12px;
  }
}
</style>