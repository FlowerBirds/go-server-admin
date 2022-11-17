<template>
  <div>
    <el-button type="success" round @click="viewServer">查看服务列表</el-button>
    <el-dialog
        title="提示"
        :visible.sync="dialogVisible"
        width="30%"
        :before-close="handleClose">
      <ul>
        <li v-for="item in servers" :key="item">{{ item }}</li>
      </ul>
      <span slot="footer" class="dialog-footer">
      <el-button @click="dialogVisible = false">取 消</el-button>
      <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
    </span>
    </el-dialog>
  </div>
</template>

<script>
module.exports = {
  data() {
    return {
      dialogVisible: false,
      servers: []
    };
  },
  methods: {
    viewServer: function (){
      axios
          .post('/list-clients')
          .then(response => {
            this.servers = response.data
            this.dialogVisible = true
          })
          .catch(function (error) { // 请求失败处理
            console.log(error);
          });

    },
    handleClose: function () {

    }
  }
}
</script>