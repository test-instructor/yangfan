<template>
  <div class="fs-login">
    <el-dialog
      v-model="dialogVisible"
      title="扫码登录"
      width="30%"
      destroy-on-close
    >
      <div class="scan-ercode">
        <div id="fslogin" />
        <span>使用飞书扫码登录</span>
      </div>

    </el-dialog>
    <el-button
        type="primary"
        size="large"
        style="width: 125px"
        @click="scan"
    >
      扫码登录
    </el-button>
    <el-button
        type="primary"
        size="large"
        style="width: 125px"
        @click="open"
    >
      飞书web登录
    </el-button>
  </div>

</template>

<script setup>
import { nextTick, ref } from 'vue'

const state = window.localStorage.getItem('token')
const client_id = ref(import.meta.env.VITE_FS_APP_ID).value  // 飞书的client_id
const redirect_uri = ref(import.meta.env.VITE_FS_LOGIN).value // 回调地址
const goto = `https://passport.feishu.cn/suite/passport/oauth/authorize?client_id=${client_id}&redirect_uri=${redirect_uri}&response_type=code&state=${state}`

const dialogVisible = ref(false)
const scan = async() => {
  dialogVisible.value = true
  await nextTick()
  var QRLoginObj = QRLogin({
    id: 'fslogin',
    goto: goto,
    width: '500',
    height: '500',
    style: 'width:300px;height:300px'// 可选的，二维码html标签的style属性
  })

  var handleMessage = function(event) {
    var origin = event.origin
    // 使用 matchOrigin 方法来判断 message 来自页面的url是否合法
    if (QRLoginObj.matchOrigin(origin)) {
      var loginTmpCode = event.data
      window.location.href = `${goto}&tmp_code=${loginTmpCode}`
    }
  }
  if (typeof window.addEventListener !== 'undefined') {
    window.addEventListener('message', handleMessage, false)
  } else if (typeof window.attachEvent !== 'undefined') {
    window.attachEvent('onmessage', handleMessage)
  }
}

const open = () => {
  window.open(goto)
}

</script>

<style lang="scss" scoped>
.fs-login {
    .scan-ercode{
        ::v-deep(iframe){
            border: 0;
        }
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }

  display: inline-block;
}
</style>
