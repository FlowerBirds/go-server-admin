<template>
  <div>
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="客户端">

      </el-form-item>

    </el-form>

    <el-table :data="clients" style="width: 100%" :row-class-name="tableRowClassName">
      <el-table-column label="序号" width="80px">
        <template slot-scope="scope">
          {{ scope.$index + (currentPage - 1) *  pageSize + 1}}
        </template>
      </el-table-column>
      <el-table-column prop="n" label="主机名" @row-click="clickRow"></el-table-column>
      <el-table-column prop="ip" label="主机IP" ></el-table-column>
      <el-table-column prop="m.t" label="总内存" :formatter="formatMemory"></el-table-column>
      <el-table-column prop="m.f" label="剩余内存" :formatter="formatMemory"></el-table-column>
      <el-table-column prop="c.c" label="CPU核数" ></el-table-column>
      <el-table-column prop="d" label="磁盘总容量" :formatter="formatDiskTotal"></el-table-column>
      <el-table-column prop="d" label="磁盘空闲容量" :formatter="formatDiskFree"></el-table-column>
      <el-table-column prop="t" label="心跳时间" :formatter="formatTime"></el-table-column>
    </el-table>
    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage"
                   :page-sizes="[5, 10, 20, 50, 100]" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
                   :total="total">
    </el-pagination>
  </div>
</template>

<script>
const bits = ['B', 'KB', 'MB', 'GB', 'TB', 'PB']
module.exports =  {
  data() {
    return {
      messages: [],
      currentPage: 1,
      pageSize: 10,
      total: 0,
      clients: [],
      currentClient: '',
      formInline: {}
    }
  },
  props: [],
  mounted() {
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
    formatMemory: function (row, column, cellValue, index) {
      let arr = this.numberFormat(cellValue, 0)
      return Number(arr[0]).toFixed(4) + ' ' + bits[arr[1]] + ''
    },
    formatDiskTotal: function (row, column, cellValue, index) {
      let total = 0
      for (let i in cellValue) {
        total += cellValue[i].t
      }
      let arr = this.numberFormat(total, 0)
      return Number(arr[0]).toFixed(4) + ' ' + bits[arr[1]] + ''
    },
    formatDiskFree: function (row, column, cellValue, index) {
      let free = 0
      for (let i in cellValue) {
        free += cellValue[i].f
      }
      let arr = this.numberFormat(free, 0)
      return Number(arr[0]).toFixed(4) + ' ' + bits[arr[1]] + ''
    },
    numberFormat: function (value, index) {
      if (value >= 1024) {
        return this.numberFormat(value / 1024.0, index + 1)
      } else {
        return [value, index]
      }
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
            this.total = this.clients.length
            console.log(this.messages)
          })
          .catch(function (error) { // 请求失败处理
            console.log(error);
          });
    },
    onSubmit: function () {

    }
  }
}
</script>

<style>

</style>