<template>
  <el-card style="width: 100%;">
    <!-- 预约 -->
    <el-table :data="tableData" style="width: 100%; font-size: 15px">
      <!-- 展开列 -->
        <el-table-column type="expand" width="25px">
          <template slot-scope="scope">
            <!-- 折叠面板 -->
            <el-collapse accordion style="display: flex; flex-direction: column; align-items: center;" >
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
      <!-- 申请时间 -->
      <el-table-column label="申请时间" align="center" width="200" sortable>
        <template slot-scope="scope">
          <span> {{ scope.row.ApplyTime }}</span>
        </template>
      </el-table-column>
      <!-- 日期 -->
      <el-table-column label="预约日期" align="center">
        <template slot-scope="scope">
          <span>{{scope.row.Date}}</span>
<!--          <el-image style="height: 80px; width: 80px; margin: 0;" :src="scope.row.url1"></el-image>-->
<!--          <el-image-->
<!--            style="height: 80px; width: 80px; margin-left: 20px;"-->
<!--            v-if="scope.row.url2 !== 'undefined'"-->
<!--            :src="scope.row.url2"></el-image>-->
        </template>
      </el-table-column>

      <!-- 时间段 -->
      <el-table-column label="时间段" align="center" >
        <template slot-scope="scope">
          <div slot="reference">
            <span>{{scope.row.Period}}</span>
<!--            <span slot="reference" v-if="scope.row.foodName2 !== 'undefined'">-->
<!--              &nbsp;&nbsp;{{scope.row.foodName2}}-->
<!--            </span>-->
          </div>
        </template>
      </el-table-column>

      <!-- 名称 -->
      <el-table-column label="会议室" align="center">
        <template slot-scope="scope">
          <div slot="reference" >
            <span>{{scope.row.RoomName}}</span>
<!--            <el-tag type="danger" size="medium">{{scope.row.allPrice}}</el-tag>-->
          </div>
        </template>
      </el-table-column>

      <!-- 名称 -->
      <el-table-column label="用户" align="center">
        <template slot-scope="scope">
          <div slot="reference" >
            <span>{{scope.row.WxId}}</span>
<!--            <el-tag type="danger" size="medium">{{scope.row.allPrice}}</el-tag>-->
          </div>
        </template>
      </el-table-column>


      <!-- 原因 -->
      <el-table-column label="申请原因" align="center">
        <template slot-scope="scope">
          <div slot="reference" >
            <span>{{scope.row.Reason}}</span>
<!--            <el-tag type="danger" size="medium">{{scope.row.allPrice}}</el-tag>-->
          </div>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button size="medium" type="primary" @click="AgreeReserve(scope.$index, scope.row)">同意</el-button>
          <el-button size="medium" type="danger" @click="DisagreeReserve(scope.$index, scope.row)">拒绝</el-button>
        </template>
      </el-table-column>

      <!-- 操作 -->
<!--      <el-table-column label="操作" align="center">-->
<!--        <template slot-scope="scope">-->
<!--          &lt;!&ndash; 完成按钮 &ndash;&gt;-->
<!--          <el-button-->
<!--            size="medium"-->
<!--            type="primary"-->
<!--            v-show="!scope.row.Finish"-->
<!--            @click="handleEdit(scope.$index, scope.row)">同意</el-button>-->
<!--          &lt;!&ndash; 未挂起 &ndash;&gt;-->
<!--          <el-button-->
<!--            size="medium"-->
<!--            type="primary"-->
<!--            v-show="!scope.row.HangUp"-->
<!--            @click="hangUpOrder(scope.$index, scope.row)">挂起</el-button>-->
<!--          &lt;!&ndash; 已挂起 &ndash;&gt;-->
<!--          <el-button-->
<!--            size="medium"-->
<!--            type="warning"-->
<!--            v-show="scope.row.HangUp"-->
<!--            @click="hangUpOrder(scope.$index, scope.row)">挂起</el-button>-->
<!--        </template>-->
<!--      </el-table-column>-->
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
    this.getReserve()
  },
  methods: {
    // 获取全部预约申请
    getReserve () {
      // this.$http.get("manageOrders").then(res => {
      //   this.tableData = res.data
      // })
      request.get("/CanteenPC/Reserve").then(res => {
        this.tableData = res
      })
    },
    // 同意预约
    AgreeReserve (index, row) {
      console.log(row.ID)
      // this.$http.put(`/finishOrder/${row.orderId}`).then(() => {
      //   this.getFood()
      //   this.$message.success("订单完成")
      // })
      request.put("/CanteenPC/AgreeReserve", {"id": row.ID}).then(() =>{
        this.getReserve()
        this.$notify({
          title: '同意预约',
          type: 'success',
          duration: 1500,
          position: 'bottom-right'
        })
      })
    },
    DisagreeReserve (index, row) {
      request.post("/CanteenPC/DisagreeReserve",{"id":row.ID}).then(()=>{
        this.getReserve()
        this.$notify({
          title: '拒绝预约',
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
