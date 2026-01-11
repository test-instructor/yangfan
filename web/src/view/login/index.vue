<template>
  <div id="userLayout" class="w-full h-full relative flex items-center justify-center overflow-hidden bg-gradient-to-b from-[#f0f6ff] to-[#dbeaff]">
    <!-- Background Tech Elements - Solid Cubes -->
    <div id="bg" class="tech-bg" />

    <!-- Login Container -->
    <div class="login-container z-10 bg-white rounded-2xl shadow-xl px-10 py-12 w-[420px]">
      <div class="logo-section mb-8 text-center">
        <svg class="w-[70px] h-[70px] mx-auto mb-4" viewBox="0 0 100 100" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M50 10L85 30V70L50 90L15 70V30L50 10Z" stroke="#4080ff" stroke-width="3" fill="none" />
          <path d="M35 35V65H60" stroke="#4080ff" stroke-width="5" stroke-linecap="round" stroke-linejoin="round" />
          <circle cx="65" cy="35" r="4" fill="#4080ff" />
        </svg>
        <div class="text-[22px] font-semibold text-[#4080ff] tracking-wide">
          自动化测试平台
        </div>
      </div>

      <el-form
        ref="loginForm"
        :model="loginFormData"
        :rules="rules"
        :validate-on-rule-change="false"
        class="flex flex-col gap-4"
        @keyup.enter="submitForm"
      >
        <el-form-item prop="username" class="mb-0">
          <div class="form-label text-xs text-[#666] font-medium mb-2 uppercase">用户名 / USERNAME</div>
          <el-input
            v-model="loginFormData.username"
            size="large"
            placeholder="请输入用户名"
            class="custom-input"
          />
        </el-form-item>

        <el-form-item prop="password" class="mb-0">
          <div class="form-label text-xs text-[#666] font-medium mb-2 uppercase">密码 / PASSWORD</div>
          <el-input
            v-model="loginFormData.password"
            show-password
            size="large"
            type="password"
            placeholder="请输入密码"
            class="custom-input"
          />
        </el-form-item>

        <el-form-item
          v-if="loginFormData.openCaptcha"
          prop="captcha"
          class="mb-0"
        >
          <div class="form-label text-xs text-[#666] font-medium mb-2 uppercase">验证码 / CAPTCHA</div>
          <div class="flex w-full gap-3">
            <el-input
              v-model="loginFormData.captcha"
              placeholder="请输入验证码"
              size="large"
              class="custom-input flex-1"
            />
            <div class="w-[120px] h-[40px] bg-white border border-[#e1edff] rounded cursor-pointer flex items-center justify-center overflow-hidden" @click="loginVerify()">
              <img
                v-if="picPath"
                class="w-full h-full object-cover"
                :src="picPath"
                alt="验证码"
              >
              <span v-else class="text-gray-400 text-xs">加载中...</span>
            </div>
          </div>
        </el-form-item>

        <div class="flex gap-4 mt-4">
          <el-button
            class="flex-1 !h-[44px] !text-sm !font-semibold shadow-lg shadow-blue-200 !bg-[#4080ff] !border-none hover:!bg-[#3070ef]"
            type="primary"
            size="large"
            @click="submitForm"
          >登录 LOGIN</el-button>
          <el-button
            class="flex-1 !h-[44px] !text-sm !font-semibold !bg-white !text-[#4080ff] !border-[#cfe0fc] hover:!bg-[#f8fbff] hover:!border-[#4080ff]"
            size="large"
            @click="checkInit"
          >初始化 INIT</el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup>
  import { captcha } from '@/api/user'
  import { checkDB } from '@/api/initdb'
  import { reactive, ref, onMounted } from 'vue'
  import { ElMessage } from 'element-plus'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '@/pinia/modules/user'

  defineOptions({
    name: 'Login'
  })

  const router = useRouter()
  // 验证函数
  const checkUsername = (rule, value, callback) => {
    if (value.length < 5) {
      return callback(new Error('请输入正确的用户名'))
    } else {
      callback()
    }
  }
  const checkPassword = (rule, value, callback) => {
    if (value.length < 6) {
      return callback(new Error('请输入正确的密码'))
    } else {
      callback()
    }
  }

  // 获取验证码
  const loginVerify = async() => {
    const ele = await captcha()
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur'
    })
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  }
  loginVerify()

  // 登录相关操作
  const loginForm = ref(null)
  const picPath = ref('')
  const loginFormData = reactive({
    username: 'admin',
    password: '',
    captcha: '',
    captchaId: '',
    openCaptcha: false
  })
  const rules = reactive({
    username: [{ validator: checkUsername, trigger: 'blur' }],
    password: [{ validator: checkPassword, trigger: 'blur' }],
    captcha: [
      {
        message: '验证码格式不正确',
        trigger: 'blur'
      }
    ]
  })

  const userStore = useUserStore()
  const login = async() => {
    return await userStore.LoginIn(loginFormData)
  }
  const submitForm = () => {
    loginForm.value.validate(async(v) => {
      if (!v) {
        // 未通过前端静态验证
        ElMessage({
          type: 'error',
          message: '请正确填写登录信息',
          showClose: true
        })
        await loginVerify()
        return false
      }

      // 通过验证，请求登陆
      const flag = await login()

      // 登陆失败，刷新验证码
      if (!flag) {
        await loginVerify()
        return false
      }

      // 登陆成功
      return true
    })
  }

  // 跳转初始化
  const checkInit = async() => {
    const res = await checkDB()
    if (res.code === 0) {
      if (res.data?.needInit) {
        userStore.NeedInit()
        await router.push({ name: 'Init' })
      } else {
        ElMessage({
          type: 'info',
          message: '已配置数据库信息，无法初始化'
        })
      }
    }
  }

  onMounted(() => {
    const bg = document.getElementById('bg')
    const cubeCount = 20

    for (let i = 0; i < cubeCount; i++) {
      const cube = document.createElement('div')
      cube.classList.add('cube')

      const faces = ['front', 'back', 'right', 'left', 'top', 'bottom']
      faces.forEach(face => {
        const div = document.createElement('div')
        div.classList.add('cube-face', `face-${face}`)
        cube.appendChild(div)
      })

      const sizeVal = Math.random() * 60 + 30
      cube.style.setProperty('--size', `${sizeVal}px`)
      cube.style.setProperty('--left', `${Math.random() * 100}%`)
      cube.style.setProperty('--top', `${Math.random() * 100}%`)

      const duration = Math.random() * 30 + 15
      cube.style.setProperty('--duration', `${duration}s`)
      // 均匀分布延迟，让初始状态更平衡
      cube.style.setProperty('--delay', `-${(i / cubeCount) * duration}s`)

      cube.style.setProperty('--rx', `${Math.random() * 360}deg`)
      cube.style.setProperty('--ry', `${Math.random() * 360}deg`)
      cube.style.setProperty('--rx-end', `${Math.random() * 360 + 360}deg`)
      cube.style.setProperty('--ry-end', `${Math.random() * 360 + 360}deg`)

      bg.appendChild(cube)
    }
  })
</script>

<style lang="scss" scoped>
  .tech-bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 0;
    overflow: hidden;
    perspective: 1000px;
  }

  :deep(.cube) {
    position: absolute;
    width: var(--size);
    height: var(--size);
    transform-style: preserve-3d;
    animation: float linear infinite;
    left: var(--left);
    top: var(--top);
    animation-duration: var(--duration);
    animation-delay: var(--delay);
  }

  :deep(.cube-face) {
    position: absolute;
    width: 100%;
    height: 100%;
    animation: fade linear infinite;
    animation-duration: inherit;
    animation-delay: inherit;
  }

  :deep(.face-front) { transform: rotateY(0deg) translateZ(calc(var(--size) / 2)); background-color: #8ab6f7; }
  :deep(.face-right) { transform: rotateY(90deg) translateZ(calc(var(--size) / 2)); background-color: #5c95e8; }
  :deep(.face-top) { transform: rotateX(90deg) translateZ(calc(var(--size) / 2)); background-color: #b6d4fa; }
  :deep(.face-back) { transform: rotateY(180deg) translateZ(calc(var(--size) / 2)); background-color: #8ab6f7; }
  :deep(.face-left) { transform: rotateY(-90deg) translateZ(calc(var(--size) / 2)); background-color: #5c95e8; }
  :deep(.face-bottom) { transform: rotateX(-90deg) translateZ(calc(var(--size) / 2)); background-color: #5c95e8; }

  @keyframes float {
    0% {
      transform: translateZ(0) translateY(0) rotateX(var(--rx)) rotateY(var(--ry));
    }
    100% {
      transform: translateZ(100px) translateY(-150px) rotateX(var(--rx-end)) rotateY(var(--ry-end));
    }
  }

  @keyframes fade {
    0% {
      opacity: 0;
    }
    10% {
      opacity: 0.9;
    }
    90% {
      opacity: 0.9;
    }
    100% {
      opacity: 0;
    }
  }

  /* Customizing Element Plus Inputs */
  :deep(.custom-input .el-input__wrapper) {
    background-color: #f0f7ff;
    box-shadow: 0 0 0 1px #e1edff inset;
    padding: 4px 15px;
    height: 42px;
  }
  :deep(.custom-input .el-input__wrapper.is-focus) {
    box-shadow: 0 0 0 1px #4080ff inset !important;
    background-color: #ffffff;
  }
  :deep(.custom-input .el-input__inner) {
    color: #333;
    font-size: 14px;
  }
</style>
