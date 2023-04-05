<template>
  <el-card>
    <!-- 搜索区域 -->
    <el-row :gutter="20" style="margin-bottom: 20px; font-size: 15px" >
      <el-col :span="8">
        <el-input placeholder="请输入内容" v-model="queryInfo.query" clearable @clear="getRecord">
          <el-button slot="append" icon="el-icon-search" @click="getRecordSearch"></el-button>
        </el-input>
      </el-col>
    </el-row>

    <!-- 表格区域 -->
    <el-table ref="filterTable"
              :data="tableData.slice(queryInfo.pagesize*(queryInfo.pagenum-1), queryInfo.pagesize*queryInfo.pagenum)"
              border stripe
              style="font-size: 15px">
      <!-- 日期 -->
      <el-table-column prop="ApplyTime" label="申请日期" sortable align="center"></el-table-column>
      <!-- 时间段 -->
      <el-table-column prop="Date" label="预约日期" align="center"></el-table-column>
      <!-- 会议室 -->
      <el-table-column label="时间段" align="center">
        <template slot-scope="scope">
          <div slot="reference">
            <span>{{scope.row.Period}}</span>
<!--            <span slot="reference" v-if="scope.row.foodName2 !== 'undefined'">-->
<!--              &nbsp;&nbsp;{{scope.row.foodName2}}-->
<!--            </span>-->
          </div>
        </template>
      </el-table-column>
      <!-- 价格 -->
      <el-table-column prop="RoomName" label="会议室" align="center"></el-table-column>
      <!-- 是否已收货 -->
      <el-table-column prop="WxId" label="用户" align="center"></el-table-column>
      <!-- 类别 -->
      <el-table-column prop="Status" label="审核" align="center"
                       :filters="[{ text: '通过', value: 2 }, { text: '未通过', value: 0 }]"
                       :filter-method="filterTag"
                       filter-placement="bottom-end">
        <template slot-scope="scope">
<!--          <el-tag :type="scope.row.Status === '0' ? 'primary' : 'success'" disable-transitions>-->
<!--            {{scope.row.Status}}-->
<!--          </el-tag>-->
          <div v-if="scope.row.Status === 0">
            未通过
          </div>
          <div v-else>
            通过
          </div>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页区域 -->
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="queryInfo.pagenum"
      :page-sizes="[1, 2, 3, 5]"
      :page-size="queryInfo.pagesize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"></el-pagination>
  </el-card>
</template>

<script>
import request from '@/utils/request'

export default {
  data () {
    return {
      // 获取用户列表的参数对象
      queryInfo: {
        query: '逸夫楼101',
        // 当前的页码
        pagenum: 1,
        // 当前每页显示多少条数据
        pagesize: 5
      },
      // 表格信息
      tableData: [],
      total: 0
    }
  },
  created () {
    this.getRecord()
  },
  methods: {
    getRecord () {
      // this.$http.get("ordersRecord").then(res => {
      //   this.tableData = res.data
      //   this.total = res.data.length
      // })
      request.get("/CanteenPC/AppointmentRecord").then(res=>{
        this.tableData = res.msg
        this.total = res.length
        console.log(res.length)
      })
    },
    getRecordSearch () {
      console.log(this.queryInfo.query)
      request.post("/CanteenPC/RecordSearch",{"name":this.queryInfo.query}).then(res=>{
        this.tableData = res.msg
        this.total = res.length
        console.log(res.length)
      })
      // this.$http.get(`ordersRecordSearch/${this.queryInfo.query}`).then(res => {
      //   this.tableData = res.data
      //   this.total = res.data.length
      // })
    },
    filterTag (value, row) {
      return row.Status === value
    },
    // 监听 pagesize 改变的事件
    handleSizeChange (newSize) {
      this.queryInfo.pagesize = newSize
      this.getRecord()
    },
    // 监听页码值改变的事件
    handleCurrentChange (newPage) {
      this.queryInfo.pagenum = newPage
      this.getRecord()
    }
  }
}
</script>

<style scoped>
.el-pagination {
  margin-top: 20px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
</style>
