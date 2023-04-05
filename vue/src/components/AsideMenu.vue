<template>
  <el-container >


    <!-- 侧边栏菜单区域 -->
    <el-menu background-color="#333744" text-color="#fff" unique-opened router :default-active="activePath">
<!--      router :default-active="activePath"-->
      <!-- 会议室预约 -->
      <el-submenu index="/CanteenPCReserve">
        <template slot="title">
          <i class="el-icon-tickets"></i>
          <span>申请处理</span>
        </template>
        <!-- 会议室预约 -->
        <el-menu-item index="/CanteenPC/Reserve" @click="saveNavState('/CanteenPC/Reserve')">
          <i class="el-icon-house"></i>
          <span slot="title">会议室预约</span>
        </el-menu-item>
        <!-- 管理员申请 -->
        <el-menu-item index="/CanteenPC/AdminApply" @click="saveNavState('/CanteenPC/AdminApply')">
          <i class="el-icon-user-solid"></i>
          <span slot="title">管理员申请</span>
        </el-menu-item>
      </el-submenu>

      <!-- 挂起订单
      <el-menu-item index="/CanteenPC/hangUp" @click="saveNavState('/CanteenPC/hangUp'), drawerVisible=false">
        <i class="el-icon-refresh-left"></i>
        <span slot="title" >
          <span type="primary">挂起订单</span>
        </span>
      </el-menu-item>
      -->

      <!-- 菜系管理 -->
      <el-menu-item index="/CanteenPC/MeetingRoom" @click="saveNavState('/CanteenPC/MeetingRoom')">
        <i class="el-icon-s-management"></i>
        <span slot="title">会议室管理</span>
      </el-menu-item>

      <!-- 订单记录 -->
      <el-menu-item index="/CanteenPC/AppointmentRecord" @click="saveNavState('/CanteenPC/AppointmentRecord')">
        <i class="el-icon-receiving"></i>
        <span slot="title">预约记录</span>
      </el-menu-item>

      <!-- 意见处理 -->
      <el-menu-item index="/CanteenPC/Opinion" @click="saveNavState('/CanteenPC/Opinion')">
        <i class="el-icon-edit"></i>
        <span slot="title">意见处理</span>
      </el-menu-item>

      <!-- 数据统计 -->
      <el-submenu index="/CanteenPCdataAnalysis">
        <template slot="title">
          <i class="el-icon-pie-chart"></i>
          <span>数据统计</span>
        </template>
        <el-menu-item-group>
          <!-- 预约分析 -->
          <el-menu-item index="/CanteenPC/RoomAnalysis" @click="saveNavState('/CanteenPC/RoomAnalysis')">
            <i class="el-icon-data-analysis"></i>
            <span>预约室分析</span>
          </el-menu-item>
          <!-- 菜品分析 -->
          <el-menu-item index="/CanteenPC/PeriodAnalysis" @click="saveNavState('/CanteenPC/PeriodAnalysis')">

            <i class="el-icon-data-line"></i>
            <span>时间段分析</span>
          </el-menu-item>
        </el-menu-item-group>
      </el-submenu>
    </el-menu>

    <!-- 抽屉 -->
    <el-drawer :open="getHangUpOrders" :visible.sync="drawerVisible" size="50%" direction="rtl">
      <el-table :data="hangUpOrders" style="width: 100%">
        <!-- 日期 -->
        <el-table-column label="日期">
          <template slot-scope="scope">
            <span>{{scope.row.time}}</span>
          </template>
        </el-table-column>
        <!-- 订单号 -->
        <el-table-column label="订单号">
          <template slot-scope="scope">
            <div slot="reference" class="name-wrapper">
              <span size="medium">{{scope.row.orderId}}</span>
            </div>
          </template>
        </el-table-column>
        <!-- 菜系 -->
        <el-table-column label="菜系">
          <template slot-scope="scope">
            <div slot="reference" class="name-wrapper">
              <span size="medium">{{scope.row.foodName1}}</span>
            </div>
          </template>
        </el-table-column>
        <!-- 类型 -->
        <el-table-column label="类型">
          <template slot-scope="scope">
            <el-button plain type="info">{{scope.row.type}}</el-button>
          </template>
        </el-table-column>
        <!-- 操作 -->
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="primary" @click="handle(scope.$index, scope.row)">完成</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-drawer>
  </el-container>
</template>

<script>
  export default {
    data () {
      return {
        // 激活的路由
        activePath: '',
        // 控制抽屉的激活状态
        drawerVisible: false,
        // 已挂起的订单
        hangUpOrders: []
      }
    },
    created () {
      // 从缓存中获取已保存的链接地址
      this.activePath = window.sessionStorage.getItem('activePath')
      this.getHangUpOrders()
    },
    methods: {
      // 保存链接的激活状态
      saveNavState (activePath) {
        console.log(activePath)
        window.sessionStorage.setItem('activePath', activePath)
        this.activePath = activePath
      },
      getHangUpOrders () {
        this.$http.get("/hangUp").then(res => { this.hangUpOrders = res.data })
      }
    }
  }
</script>

<style scoped>
  .el-container {
    background-color: #333744;
  }

  .el-menu {
    border-right: none;
  }

  .el-menu-item {
    width: 200px;
  }
</style>
