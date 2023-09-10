<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="16">
        <el-card>
          <template #header>
            <el-divider>扬帆测试平台</el-divider>
          </template>
          <p>
            扬帆测试平台是一款高效、可靠的自动化测试平台，旨在帮助团队提升测试效率、降低测试成本。该平台包括用例管理、定时任务、执行记录等功能模块，支持多种类型的测试用例，目前支持API(http和grpc协议)、性能，并且可定制化，灵活满足不同场景的需求。其中，用例管理模块支持上传、编辑、复制、删除等基础操作，同时支持批量执行、并发执行等高级功能。通过用例设置，可以设置用例的基本信息、运行配置、环境变量等，灵活控制用例的执行。而定时任务模块支持引用一个或多个用例，实现用例的自动执行，支持并发运行、任务标签等功能，后续支持CI/CD集成，实现全自动化的测试流程。扬帆测试平台还提供执行记录模块，记录测试用例的执行结果，支持查看测试报告、执行日志等详细信息，方便对测试结果进行分析和优化。
          </p>
        </el-card>
        <el-card style="margin-top: 20px">
          <template #header>
            <div>联系我们</div>
          </template>
          <div style="display: flex">
            <img
              :src="`http://qiniu.yangfan.gd.cn/about/author.jpeg?time=${timestamp}`"
              alt="Author Image"
            />
            <img
              :src="`http://qiniu.yangfan.gd.cn/about/group.jpeg?time=${timestamp}`"
              alt="Group Image"
            />
            <img
              :src="`http://qiniu.yangfan.gd.cn/about/mp.jpg?time=${timestamp}`"
              alt="mp Image"
            />
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <div>提交记录</div>
          </template>
          <div>
            <el-timeline>
              <el-timeline-item
                v-for="(item, index) in dataTimeline"
                :key="index"
                :timestamp="item.from"
                placement="top"
              >
                <el-card>
                  <h4>{{ item.title }}</h4>
                  <p>{{ item.message }}</p>
                </el-card>
              </el-timeline-item>
            </el-timeline>
          </div>
          <el-button class="load-more" type="primary" link @click="loadMore"
            >Load more</el-button
          >
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "About",
};
</script>

<script setup>
import { ref } from "vue";
import { Commits, Members } from "@/api/github";
import { formatTimeToStr } from "@/utils/date";
const page = ref(0);

const loadMore = () => {
  page.value++;
  loadCommits();
};

const timestamp = Date.now();

const dataTimeline = ref([]);
const loadCommits = () => {
  Commits(page.value).then(({ data }) => {
    data.forEach((element) => {
      if (element.commit.message) {
        dataTimeline.value.push({
          from: formatTimeToStr(element.commit.author.date, "yyyy-MM-dd"),
          title: element.commit.author.name,
          showDayAndMonth: true,
          message: element.commit.message,
        });
      }
    });
  });
};

const members = ref([]);
const loadMembers = () => {
  Members().then(({ data }) => {
    members.value = data;
    members.value.sort();
  });
};

loadCommits();
loadMembers();
</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}

img {
  width: 360px;
  height: 520px;
}
</style>
