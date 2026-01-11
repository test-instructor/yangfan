import service from '@/utils/request'
import { useUserStore } from '@/pinia/modules/user'

const userStore = useUserStore()
const projectId = userStore.userInfo.projectId

// @Tags Runner
// @Summary 运行任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body object true "运行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"运行成功"}"
// @Router /runner/api [post]
export const runTask = (data) => {
  data.projectId = projectId
  return service({
    url: '/runner/api',
    method: 'post',
    data
  })
}
