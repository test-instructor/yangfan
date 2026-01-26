<template>
  <div class="person-page">
    <a-card class="header-card" :bordered="false">
      <div class="header-content">
        <div class="avatar-section">
          <a-avatar :size="100" shape="circle">
            <img v-if="userInfo.headerImg" :src="userInfo.headerImg" alt="avatar" />
            <IconUser v-else />
          </a-avatar>
        </div>
        <div class="info-section">
          <div class="name-row">
            <span class="name">{{ userInfo.nickName || userInfo.userName || 'User' }}</span>
            <a-button type="text" size="small" @click="openEditNickName">
              <template #icon><IconEdit /></template>
            </a-button>
          </div>
          <div class="detail-row">
            <span class="detail-item"><IconLocation /> 中国·北京市·朝阳区</span>
            <span class="detail-item"><IconCommon /> 北京翻转极光科技有限公司</span>
            <span class="detail-item"><IconUser /> 技术部·前端事业群</span>
          </div>
          <div class="action-row">
            <a-space>
              <a-button type="primary">
                <template #icon><IconMessage /></template>
                发送消息
              </a-button>
              <a-button>
                <template #icon><IconShareInternal /></template>
                分享主页
              </a-button>
            </a-space>
          </div>
        </div>
      </div>
    </a-card>

    <div class="content-grid">
      <div class="left-col">
        <a-card title="基本信息" :bordered="false" class="info-card">
          <a-space direction="vertical" size="large" style="width: 100%">
            <div class="info-item">
              <div class="label"><IconPhone class="icon-blue" /> 手机号码</div>
              <div class="value">{{ userInfo.phone || '未设置' }}</div>
              <a-button type="text" size="small" @click="changePhoneVisible = true">修改</a-button>
            </div>
            <div class="info-item">
              <div class="label"><IconEmail class="icon-green" /> 邮箱地址</div>
              <div class="value">{{ userInfo.email || '未设置' }}</div>
              <a-button type="text" size="small" @click="changeEmailVisible = true">修改</a-button>
            </div>
            <div class="info-item">
              <div class="label"><IconLock class="icon-purple" /> 账号密码</div>
              <div class="value">已设置</div>
              <a-button type="text" size="small" @click="changePasswordVisible = true">修改</a-button>
            </div>
          </a-space>
        </a-card>

        <a-card title="技能特长" :bordered="false" class="skills-card">
          <a-space wrap>
            <a-tag color="green">GoLang</a-tag>
            <a-tag color="orange">JavaScript</a-tag>
            <a-tag color="red">Vue</a-tag>
            <a-tag color="gray">Gorm</a-tag>
            <a-button size="mini" type="text">
              <template #icon><IconPlus /></template>
              添加技能
            </a-button>
          </a-space>
        </a-card>
      </div>

      <div class="right-col">
        <a-card :bordered="false" class="tabs-card">
          <a-tabs default-active-key="stats">
            <a-tab-pane key="stats" title="数据统计">
              <div class="stats-grid">
                <div class="stat-item">
                  <div class="stat-value blue">138</div>
                  <div class="stat-label">项目参与</div>
                </div>
                <div class="stat-item">
                  <div class="stat-value green">2.3k</div>
                  <div class="stat-label">代码提交</div>
                </div>
                <div class="stat-item">
                  <div class="stat-value purple">95%</div>
                  <div class="stat-label">任务完成</div>
                </div>
                <div class="stat-item">
                  <div class="stat-value yellow">12</div>
                  <div class="stat-label">获得勋章</div>
                </div>
              </div>
            </a-tab-pane>
            <a-tab-pane key="activity" title="近期动态">
              <a-timeline>
                <a-timeline-item v-for="(activity, index) in activities" :key="index" :dotColor="activity.color">
                  <div class="timeline-title">{{ activity.title }}</div>
                  <div class="timeline-content">{{ activity.content }}</div>
                  <div class="timeline-time">{{ activity.timestamp }}</div>
                </a-timeline-item>
              </a-timeline>
            </a-tab-pane>
          </a-tabs>
        </a-card>
      </div>
    </div>

    <!-- Modals -->
    <a-modal v-model:visible="editNickNameVisible" title="修改昵称" @ok="saveNickName">
      <a-form :model="nickNameForm">
        <a-form-item field="nickName" label="昵称">
          <a-input v-model="nickNameForm.nickName" placeholder="请输入昵称" />
        </a-form-item>
      </a-form>
    </a-modal>

    <a-modal v-model:visible="changePhoneVisible" title="修改手机号" @ok="savePhone">
      <a-form :model="phoneForm">
        <a-form-item field="phone" label="手机号">
          <a-input v-model="phoneForm.phone" placeholder="请输入手机号" />
        </a-form-item>
        <a-form-item field="code" label="验证码">
          <a-input v-model="phoneForm.code" placeholder="请输入验证码" />
          <a-button class="ml-2" :disabled="timer > 0" @click="sendCode">
            {{ timer > 0 ? `${timer}s` : '获取验证码' }}
          </a-button>
        </a-form-item>
      </a-form>
    </a-modal>

    <a-modal v-model:visible="changeEmailVisible" title="修改邮箱" @ok="saveEmail">
      <a-form :model="emailForm">
        <a-form-item field="email" label="邮箱">
          <a-input v-model="emailForm.email" placeholder="请输入邮箱" />
        </a-form-item>
        <a-form-item field="code" label="验证码">
          <a-input v-model="emailForm.code" placeholder="请输入验证码" />
          <a-button class="ml-2" :disabled="emailTimer > 0" @click="sendEmailCode">
            {{ emailTimer > 0 ? `${emailTimer}s` : '获取验证码' }}
          </a-button>
        </a-form-item>
      </a-form>
    </a-modal>

    <a-modal v-model:visible="changePasswordVisible" title="修改密码" @ok="savePassword">
      <a-form :model="passwordForm">
        <a-form-item field="password" label="原密码">
          <a-input-password v-model="passwordForm.password" placeholder="请输入原密码" />
        </a-form-item>
        <a-form-item field="newPassword" label="新密码">
          <a-input-password v-model="passwordForm.newPassword" placeholder="请输入新密码" />
        </a-form-item>
        <a-form-item field="confirmPassword" label="确认密码">
          <a-input-password v-model="passwordForm.confirmPassword" placeholder="请再次输入新密码" />
        </a-form-item>
      </a-form>
    </a-modal>

  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { Message } from '@arco-design/web-vue'
import { 
  IconUser, IconEdit, IconLocation, IconCommon, IconMessage, IconShareInternal, 
  IconPhone, IconEmail, IconLock, IconPlus 
} from '@arco-design/web-vue/es/icon'
import { getUserInfo, setSelfInfo, changePassword } from '../services/appBridge'

const userInfo = ref({})
const editNickNameVisible = ref(false)
const changePhoneVisible = ref(false)
const changeEmailVisible = ref(false)
const changePasswordVisible = ref(false)

const nickNameForm = reactive({ nickName: '' })
const phoneForm = reactive({ phone: '', code: '' })
const emailForm = reactive({ email: '', code: '' })
const passwordForm = reactive({ password: '', newPassword: '', confirmPassword: '' })

const timer = ref(0)
const emailTimer = ref(0)

const loadUserInfo = async () => {
  try {
    const res = await getUserInfo()
    userInfo.value = res || {}
  } catch (e) {
    console.error(e)
    Message.error('获取用户信息失败')
  }
}

const openEditNickName = () => {
  nickNameForm.nickName = userInfo.value.nickName || ''
  editNickNameVisible.value = true
}

const saveNickName = async () => {
  try {
    await setSelfInfo({ nickName: nickNameForm.nickName })
    Message.success('修改成功')
    await loadUserInfo()
  } catch (e) {
    Message.error(e.message || '修改失败')
  }
}

const sendCode = () => {
  timer.value = 60
  const interval = setInterval(() => {
    timer.value--
    if (timer.value <= 0) clearInterval(interval)
  }, 1000)
}

const savePhone = async () => {
  try {
    // Note: Verification code logic needs backend support, assuming success for now or mock
    await setSelfInfo({ phone: phoneForm.phone })
    Message.success('修改成功')
    await loadUserInfo()
  } catch (e) {
    Message.error(e.message || '修改失败')
  }
}

const sendEmailCode = () => {
  emailTimer.value = 60
  const interval = setInterval(() => {
    emailTimer.value--
    if (emailTimer.value <= 0) clearInterval(interval)
  }, 1000)
}

const saveEmail = async () => {
  try {
    await setSelfInfo({ email: emailForm.email })
    Message.success('修改成功')
    await loadUserInfo()
  } catch (e) {
    Message.error(e.message || '修改失败')
  }
}

const savePassword = async () => {
  if (passwordForm.newPassword !== passwordForm.confirmPassword) {
    Message.error('两次密码不一致')
    return
  }
  try {
    await changePassword({ 
      password: passwordForm.password, 
      newPassword: passwordForm.newPassword 
    })
    Message.success('修改密码成功')
    changePasswordVisible.value = false
  } catch (e) {
    Message.error(e.message || '修改失败')
  }
}

const activities = [
  {
    title: '完成项目里程碑',
    content: '成功完成第三季度主要项目开发任务，获得团队一致好评',
    timestamp: '2024-01-10',
    color: '#165DFF'
  },
  {
    title: '代码审核完成',
    content: '完成核心模块代码审核，提出多项改进建议并获采纳',
    timestamp: '2024-01-11',
    color: '#00B42A'
  },
  {
    title: '技术分享会',
    content: '主持团队技术分享会，分享前端性能优化经验',
    timestamp: '2024-01-12',
    color: '#FF7D00'
  },
  {
    title: '新功能上线',
    content: '成功上线用户反馈的新特性，显著提升用户体验',
    timestamp: '2024-01-13',
    color: '#F53F3F'
  }
]

onMounted(() => {
  loadUserInfo()
})
</script>

<style scoped>
.person-page {
  padding: 0;
  height: 100%;
  overflow-y: auto;
}
.header-card {
  margin-bottom: 24px;
}
.header-content {
  display: flex;
  gap: 24px;
  align-items: center;
}
.info-section {
  flex: 1;
}
.name-row {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 12px;
}
.name {
  font-size: 24px;
  font-weight: bold;
}
.detail-row {
  display: flex;
  gap: 24px;
  color: var(--color-text-3);
  margin-bottom: 16px;
}
.detail-item {
  display: flex;
  align-items: center;
  gap: 4px;
}
.content-grid {
  display: grid;
  grid-template-columns: 350px 1fr;
  gap: 24px;
}
.info-card, .skills-card {
  margin-bottom: 24px;
}
.info-item {
  display: flex;
  align-items: center;
  gap: 8px;
}
.label {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100px;
  color: var(--color-text-2);
}
.value {
  flex: 1;
  color: var(--color-text-1);
}
.icon-blue { color: #165DFF; }
.icon-green { color: #00B42A; }
.icon-purple { color: #722ED1; }

.stats-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  text-align: center;
  padding: 24px 0;
}
.stat-value {
  font-size: 32px;
  font-weight: bold;
  margin-bottom: 8px;
}
.stat-label {
  color: var(--color-text-3);
}
.stat-value.blue { color: #165DFF; }
.stat-value.green { color: #00B42A; }
.stat-value.purple { color: #722ED1; }
.stat-value.yellow { color: #FF7D00; }

.timeline-title {
  font-size: 16px;
  font-weight: 500;
  margin-bottom: 4px;
}
.timeline-content {
  color: var(--color-text-3);
  margin-bottom: 4px;
}
.timeline-time {
  color: var(--color-text-4);
  font-size: 12px;
}
.ml-2 {
  margin-left: 8px;
}
</style>