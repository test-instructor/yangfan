<template>
  <div class="container">
    <a-row :gutter="20">
      <a-col :span="8">
        <a-card class="info-card">
          <div class="user-header">
            <a-avatar :size="100" class="avatar-trigger">
              <img v-if="userInfo.headerImg" :src="userInfo.headerImg" />
              <IconUser v-else />
              <div class="avatar-mask">
                <IconCamera />
              </div>
            </a-avatar>
            <div class="user-name">
              {{ userInfo.nickName || userInfo.userName || 'User' }}
              <IconEdit class="edit-icon" @click="openEditNickName" />
            </div>
            <div class="user-desc">UI Automation Engineer</div>
          </div>
          
          <div class="user-detail">
             <div class="detail-item">
               <span class="icon-wrap"><IconUser /></span>
               <span>技术部 · 前端事业群</span>
             </div>
             <div class="detail-item">
               <span class="icon-wrap"><IconCommon /></span>
               <span>北京翻转极光科技有限公司</span>
             </div>
             <div class="detail-item">
               <span class="icon-wrap"><IconLocation /></span>
               <span>中国 · 北京市 · 朝阳区</span>
             </div>
          </div>

          <a-divider />
          
          <div class="section-title">联系方式</div>
          <div class="contact-list">
             <div class="contact-item">
               <IconPhone />
               <span>{{ userInfo.phone || '未设置' }}</span>
               <a-link @click="changePhoneVisible = true">修改</a-link>
             </div>
             <div class="contact-item">
               <IconEmail />
               <span>{{ userInfo.email || '未设置' }}</span>
               <a-link @click="changeEmailVisible = true">修改</a-link>
             </div>
          </div>

          <a-divider />

          <div class="section-title">技能特长</div>
          <div class="tags-list">
             <a-tag v-for="tag in tags" :key="tag" color="arcoblue">{{ tag }}</a-tag>
             <a-tag class="add-tag" @click="addTag"><IconPlus /></a-tag>
          </div>
        </a-card>
      </a-col>

      <a-col :span="16">
        <a-card class="content-card">
          <a-tabs default-active-key="1">
            <a-tab-pane key="1" title="文章">
              <a-empty description="暂无文章" />
            </a-tab-pane>
            <a-tab-pane key="2" title="项目">
               <div class="project-list">
                 <a-card v-for="i in 4" :key="i" class="project-item" hoverable>
                    <template #title>Project {{ i }}</template>
                    UI Automation Framework
                    <div class="project-footer">
                       <span><IconStar /> 120</span>
                       <span><IconThumbUp /> 50</span>
                    </div>
                 </a-card>
               </div>
            </a-tab-pane>
            <a-tab-pane key="3" title="动态">
              <a-timeline>
                <a-timeline-item v-for="(activity, index) in activities" :key="index" :dotColor="activity.color">
                  <div class="timeline-title">{{ activity.title }}</div>
                  <div class="timeline-content">{{ activity.content }}</div>
                  <div class="timeline-time">{{ activity.timestamp }}</div>
                </a-timeline-item>
              </a-timeline>
            </a-tab-pane>
            <a-tab-pane key="4" title="安全设置">
              <a-list>
                <a-list-item>
                  <a-list-item-meta title="登录密码" description="已设置。密码至少6位字符，支持数字、字母和除空格外的特殊字符，且必须同时包含数字和大小写字母。">
                  </a-list-item-meta>
                  <template #actions>
                    <a-button type="text" @click="changePasswordVisible = true">修改</a-button>
                  </template>
                </a-list-item>
                <a-list-item>
                  <a-list-item-meta title="安全手机" :description="userInfo.phone ? `已绑定：${userInfo.phone}` : '未绑定手机'">
                  </a-list-item-meta>
                  <template #actions>
                    <a-button type="text" @click="changePhoneVisible = true">修改</a-button>
                  </template>
                </a-list-item>
                <a-list-item>
                  <a-list-item-meta title="安全邮箱" :description="userInfo.email ? `已绑定：${userInfo.email}` : '未绑定邮箱'">
                  </a-list-item-meta>
                  <template #actions>
                    <a-button type="text" @click="changeEmailVisible = true">修改</a-button>
                  </template>
                </a-list-item>
              </a-list>
            </a-tab-pane>
          </a-tabs>
        </a-card>
      </a-col>
    </a-row>

    <!-- Modals -->
    <a-modal v-model:visible="editNickNameVisible" title="修改昵称" @ok="saveNickName">
      <a-form :model="nickNameForm">
        <a-form-item field="nickName" label="昵称">
          <a-input v-model="nickNameForm.nickName" placeholder="请输入昵称" />
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
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { Message } from '@arco-design/web-vue'
import {
  IconUser,
  IconEdit,
  IconLocation,
  IconCommon,
  IconPhone,
  IconEmail,
  IconPlus,
  IconCamera,
  IconStar,
  IconThumbUp
} from '@arco-design/web-vue/es/icon'
import { getUserInfo, setSelfInfo, changePassword } from '../../services/appBridge'

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

const tags = ref(['GoLang', 'JavaScript', 'Vue', 'Gorm'])

const addTag = () => {
  Message.info('Add Tag Clicked')
}

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
.container {
  padding: 0 20px 20px 20px;
}
.info-card {
  border-radius: 4px;
}
.user-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px 0;
}
.avatar-trigger {
  position: relative;
  cursor: pointer;
  transition: all 0.1s;
}
.avatar-mask {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  opacity: 0;
  transition: opacity 0.1s;
  border-radius: 50%;
}
.avatar-trigger:hover .avatar-mask {
  opacity: 1;
}
.user-name {
  margin-top: 16px;
  font-size: 20px;
  font-weight: 500;
  color: var(--color-text-1);
  display: flex;
  align-items: center;
}
.edit-icon {
  margin-left: 8px;
  font-size: 14px;
  cursor: pointer;
  color: var(--color-text-3);
}
.user-desc {
  margin-top: 4px;
  color: var(--color-text-3);
  font-size: 14px;
}
.user-detail {
  margin-top: 20px;
}
.detail-item {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
  color: var(--color-text-2);
}
.detail-item:last-child {
  margin-bottom: 0;
}
.icon-wrap {
  width: 24px;
  text-align: center;
  margin-right: 8px;
}
.section-title {
  font-weight: 500;
  color: var(--color-text-1);
  margin-bottom: 12px;
}
.contact-item {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
  color: var(--color-text-2);
}
.contact-item span {
  flex: 1;
}
.tags-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.add-tag {
  border-style: dashed;
  cursor: pointer;
  background: transparent;
}

.content-card {
  border-radius: 4px;
  min-height: 600px;
}
.project-list {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}
.project-item {
  cursor: pointer;
}
.project-footer {
  margin-top: 12px;
  display: flex;
  gap: 16px;
  color: var(--color-text-3);
  font-size: 12px;
}
.timeline-title {
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

