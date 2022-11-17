<template>
  <div>
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="客户端">
        <el-select v-model="currentClient" placeholder="请选择">
          <el-option
              v-for="item in clients"
              :key="item.ip"
              :label="item.ip"
              :value="item.ip">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form>

  <el-table :data="messages" style="width: 100%" :row-class-name="tableRowClassName">
    <el-table-column label="序号" width="80px">
      <template slot-scope="scope">
        {{ scope.$index + (currentPage - 1) *  pageSize + 1}}
      </template>
    </el-table-column>
    <el-table-column prop="clientIp" label="客户端服务器IP" @row-click="clickRow"></el-table-column>
    <el-table-column prop="sshClientIp" label="SSH客户端IP" ></el-table-column>
    <el-table-column prop="sshUser" label="操作用户" ></el-table-column>
    <el-table-column prop="sshOperateTime" label="操作时间" :formatter="formatTime"></el-table-column>
    <el-table-column prop="serverTime" label="记录时间" :formatter="formatTime"></el-table-column>
    <el-table-column prop="sshCommandLine" label="操作命令" ></el-table-column>
  </el-table>
  <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage"
                 :page-sizes="[5, 10, 15, 20, 50, 100]" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
                 :total="total">
  </el-pagination>
  </div>
</template>

<script>
module.exports =  {
  data() {
    return {
      messages: [],
      currentPage: 1,
      pageSize: 15,
      total: 0,
      clients: [],
      currentClient: '',
      formInline: {}
    }
  },
  props: [],
  mounted() {
    this.listMessage()
    this.listClients()
  },
  methods: {
    clickRow: function (row, column, event) {

    },
    tableRowClassName: function ({row, rowIndex}) {
      if (rowIndex % 2 == 1) {
        return 'success-row';
      } else {
        return '';
      }
    },
    formatTime: function (row, column, cellValue, index) {

      return cellValue != undefined ? this.dateFormat("YYYY-mm-dd HH:MM:SS", new Date(cellValue)): "";
    },
    dateFormat: function(fmt, date) {
      let ret;
      const opt = {
        "Y+": date.getFullYear().toString(),        // 年
        "m+": (date.getMonth() + 1).toString(),     // 月
        "d+": date.getDate().toString(),            // 日
        "H+": date.getHours().toString(),           // 时
        "M+": date.getMinutes().toString(),         // 分
        "S+": date.getSeconds().toString()          // 秒
        // 有其他格式化字符需求可以继续添加，必须转化成字符串
      };
      for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
          fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
        };
      };
      return fmt;
    },
    handleSizeChange: function (size) {
      this.pageSize = size;
      this.listMessage()
    },
    handleCurrentChange: function (currentPage) {
      this.currentPage = currentPage;
      this.listMessage()
    },
    listMessage: function () {
      var that = this;
      var params = new URLSearchParams()
      params.append("currentPage", this.currentPage)
      params.append("pageSize", this.pageSize)
      params.append("currentClient", this.currentClient)
      axios
          .post('/list-client-message', params)
          .then(response => {
            var result = response.data
            that.messages = result.data
            that.total = result.total
            that.pageSize = result.pageSize
            console.log(that.messages)
          })
          .catch(function (error) { // 请求失败处理
            console.log(error);
          });
    },
    listClients: function () {
      axios
          .post('/list-clients')
          .then(response => {
            let arr = []
            for (let client in response.data) {
              let d = response.data[client]
              d.ip = client
              arr.push(d)
            }
            this.clients = arr
            console.log(this.messages)
          })
          .catch(function (error) { // 请求失败处理
            console.log(error);
          });
    },
    onSubmit: function () {
      this.listMessage()
    }
  }
}
</script>