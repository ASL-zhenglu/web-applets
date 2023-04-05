<template>
  <el-card>
    <el-table :data="tableData" style="width: 100%; font-size: 15px" >
      <!-- 展开列 -->
      <el-table-column type="expand">
        <template slot-scope="scope">
          <!-- 折叠面板 -->
          <el-collapse accordion style="display: flex; flex-direction: column; align-items: center;">
            <!-- 地址 -->
            <el-collapse-item name="1" style="width: 90%">
              <template slot="title">
                <i class="el-icon-s-home"></i>
                <span style="margin-left: 10px;">地址</span>
              </template>
              <div>{{scope.row.Address}}</div>
            </el-collapse-item>
            <!-- 联系方式 -->
            <el-collapse-item name="2" style="width: 90%">
              <template slot="title">
                <i class="el-icon-phone"></i>
                <span style="margin-left: 10px;">联系方式</span>
              </template>
              <div>{{scope.row.Phone}}</div>
            </el-collapse-item>
          </el-collapse>
        </template>
      </el-table-column>

      <!-- 索引列 -->
      <el-table-column type="index"></el-table-column>

      <!-- 日期 -->
      <el-table-column label="申请日期" align="center" sortable>
        <template slot-scope="scope">
          <span >{{ scope.row.time }}</span>
        </template>
      </el-table-column>

      <!-- 订单号 -->
      <el-table-column label="微信号" align="center">
        <template slot-scope="scope">
          <div slot="reference" class="name-wrapper">
            <span size="medium">{{ scope.row.WxId }}</span>
          </div>
        </template>
      </el-table-column>

      <!-- 用户名 -->
      <el-table-column label="用户名" align="center">
        <template slot-scope="scope">
          <div slot="reference" class="name-wrapper">
            <span size="medium">{{ scope.row.Username }}</span>
<!--            <span slot="reference" v-if="scope.row.foodName2 !== 'undefined'">-->
<!--              &nbsp;&nbsp;{{scope.row.foodName2}}-->
<!--            </span>-->
          </div>
        </template>
      </el-table-column>

      <!-- 申请原因 -->
      <el-table-column label="申请原因" width="180" align="center">
        <template slot-scope="scope">
          <div slot="reference" class="name-wrapper">
            <span size="medium">{{ scope.row.Reason}}</span>
          </div>
        </template>
      </el-table-column>

      <!-- 操作 -->
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button size="medium" type="primary" @click="AgreeApply(scope.$index, scope.row)">同意</el-button>
          <el-button size="medium" type="danger" @click="DisagreeApply(scope.$index, scope.row)">拒绝</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
import request from '@/utils/request'

export default {
  data () {
    return {
      tableData: []
    }
  },
  created () {
    this.getApply()
  },
  methods: {
    getApply () {
      // this.$http.get("takeOut").then(res => { this.tableData = res.data })
      request.get("/CanteenPC/AdminApply").then(res=>{
        this.tableData = res
      })
    },

    AgreeApply (index, row) {
      request.put("/CanteenPC/AgreeApply",{"WxId":row.WxId}).then(()=>{
      // this.$http.put(`finishTakeOut/${row.orderId}`).then(() => {
        this.getApply()
        this.$notify({
          title: '同意申请',
          type: 'success',
          duration: 1500,
          position: 'bottom-right'
        })
      })
    },
    DisagreeApply(index, row) {
      request.post("/CanteenPC/DisagreeApply",{"WxId":row.WxId}).then(()=>{
        this.getApply()
        this.$notify({
          title: '拒绝申请',
          type: 'warning',
          duration: 1500,
          position: 'bottom-right'
        })
      })
    }
  }
}
</script>

<style scoped>
</style>
