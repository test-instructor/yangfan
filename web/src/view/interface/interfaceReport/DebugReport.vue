<template>
  <div>
    <el-table
        :data="[summary]"
        size="medium"
        style="width: 100%"
        border
        stripe
        :header-cell-style="{textAlign:'center', background: '#F8F8FA'}"
        :cell-style="{textAlign:'center'}"
    >
      <el-table-column label="测试时间" width="160">
        <template #default="scope">
          <span>{{ scope.row.time.start_at|timestampToTime }}</span>
        </template>
      </el-table-column>

      <el-table-column label="持续时间" width="100">
        <template #default="scope">
          <span v-text="scope.row.time.duration.toFixed(3)+' 秒'"></span>
        </template>
      </el-table-column>

      <el-table-column label="Total" width="100">
        <template #default="scope">
          <el-tag>{{ scope.row.stat.testsRun }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="Success" width="100">
        <template #default="scope">
          <el-tag type="success">{{ scope.row.stat.successes }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="Failed" width="100">
        <template #default="scope">
          <el-tag type="danger">{{ scope.row.stat.failures }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="Error" width="100">
        <template #default="scope">
          <el-tag type="warning">{{ scope.row.stat.errors }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="Skipped" width="100">
        <template #default="scope">
          <el-tag type="info">{{ scope.row.stat.skipped }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="Platform">
        <template #default="scope">
          <el-popover trigger="hover" placement="top">
            <p>HttpRunner: {{ scope.row.platform.httprunner_version }}</p>
            <p>Python: {{ scope.row.platform.python_version }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">{{ scope.row.platform.platform }}</el-tag>
            </div>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <br/>
    <br/>

    <slot v-for="item in summary.details">
      <div>
        <span style="font-weight: bold; font-size: medium">{{ item.name }}</span>
        <el-popover placement="top-start" width="400" trigger="hover">
          <pre class="code-block">{{ item.in_out }}</pre>
          <el-button slot="reference" round type="text">parameters & output</el-button>
        </el-popover>
      </div>
      <el-table
          :data="item.records"
          style="width: 100%"
          border
          :header-cell-style="{textAlign:'center', background: '#F8F8FA'}"
          :cell-style="{textAlign:'center'}"
      >
        <el-table-column type="expand" fixed>
          <template slot-scope="props">
            <el-tabs>
              <el-tab-pane label="Request">
                <pre class="code-block" v-html="handleRequest(props.row.meta_data.request)"></pre>
              </el-tab-pane>
              <el-tab-pane label="Content" v-if="props.row.meta_data.response.content !== null">
                <pre
                    class="code-block"
                    v-text="handleContent(props.row.meta_data.response.content)"
                ></pre>
              </el-tab-pane>
              <el-tab-pane label="Response">
                <pre class="code-block" v-text="handleResponse(props.row.meta_data.response)"></pre>
              </el-tab-pane>
              <el-tab-pane label="Validators" v-if="props.row.meta_data.validators.length !== 0">
                <!--                                <pre class="code-block" v-html="props.row.meta_data.validators"></pre>-->
                <el-table
                    :data="props.row.meta_data.validators"
                    stripe
                    border
                    style="width: 100%">
                  <el-table-column
                      prop="check_result"
                      label="是否通过"
                      width="180">
                  </el-table-column>
                  <el-table-column

                      prop="check"
                      label="取值表达式"
                      width="180">
                  </el-table-column>
                  <el-table-column
                      prop="check_value"
                      label="实际值">
                  </el-table-column>
                  <el-table-column
                      prop="comparator"
                      label="比较器">
                  </el-table-column>
                  <el-table-column
                      prop="expect"
                      label="期望值">
                  </el-table-column>
                </el-table>
              </el-tab-pane>
              <el-tab-pane label="Exception" v-if="props.row.attachment !== ''">
                <pre class="code-block" v-html="props.row.attachment"></pre>
              </el-tab-pane>
              <el-tab-pane label="Extract" v-if="props.row.meta_data.response.content !== null">
                <ResContent :data="props.row.meta_data.response.content"></ResContent>
              </el-tab-pane>
            </el-tabs>
          </template>
        </el-table-column>

        <el-table-column label="名 称">
          <template #default="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="请求地址">
          <template #default="scope">
            <span>{{ scope.row.meta_data.request.url }}</span>
          </template>
        </el-table-column>

        <el-table-column label="请求方法">
          <template #default="scope">
            <span
                :class="scope.row.meta_data.request.method"
            >{{ scope.row.meta_data.request.method }}</span>
          </template>
        </el-table-column>

        <el-table-column label="响应时间 (ms)">
          <template #default="scope">
            <span>{{ scope.row.meta_data.response.elapsed_ms }}</span>
          </template>
        </el-table-column>

        <el-table-column label="测试结果">
          <template #default="scope">
            <div :class="scope.row.status">{{ scope.row.status }}</div>
          </template>
        </el-table-column>

      </el-table>
    </slot>
  </div>
</template>

<script>
import ResContent from "./ResContent";

export default {
  name: "DebugReport",
  props: {
    summary: {
      require: true,
    },
  },
  components: {
    ResContent
  },
  methods: {
    handleRequest(request) {
      const keys = ["start_timestamp"];

      keys.forEach(function (item) {
        delete request[item];
      });
      try {
        request["body"] = JSON.parse(request["body"]);
      } catch (e) {
      }

      return request;
    },

    handleContent(content) {
      try {
        content = JSON.parse(content);
      } catch (e) {
      }

      return content;
    },

    handleResponse(response) {
      const keys = [
        "response_time_ms",
        "encoding",
        "ok",
        "reason",
        "url",
        "text",
        "json",
        "content_size",
        "content_type",
      ];

      keys.forEach(function (item) {
        delete response[item];
      });

      return response;
    },
  },
};
</script>

<style scoped>
</style>
