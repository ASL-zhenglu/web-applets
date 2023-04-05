<template>
<div style="line-height:60px;display: flex">
  <div style="flex: 1">
    <span :class="collapseBtnClass" style="cursor: pointer; font-size: 18px" @click="collapse"></span>

    <el-breadcrumb separator="/" style="display: inline-block; margin-left: 10px">
      <el-breadcrumb-item v-for="breadCrumbItem in breadCrumbList" :key="breadCrumbItem.path">{{breadCrumbItem.meta.title}}</el-breadcrumb-item>

    </el-breadcrumb>
  </div>


  <el-dropdown style="width: 100px;cursor: pointer">



    <span>{{user.nickname}}</span> <i class="el-icon-arrow-down" style="margin-left: 5px"></i>

    <el-dropdown-menu slot="dropdown" style="width: 100px;text-align: center">

      <el-dropdown-item style="font-size: 14px;padding: 5px 0">
        <router-link to="/person">个人信息</router-link>
      </el-dropdown-item>

      <el-dropdown-item style="font-size: 14px;padding: 5px 0">
        <span style="text-decoration: none" @click="logout">退出</span>
      </el-dropdown-item>

    </el-dropdown-menu>

  </el-dropdown>
</div>

</template>

<script>
export default {
  name: "Header",
  props:{
    collapseBtnClass:String,
    collapse : Function,

  },
  mounted() {
    console.log(this.$route)
  },
  computed:{
    breadCrumbList(){
      return this.$route.matched;
    }
  },
  data(){
    return {
      user: localStorage.getItem("user") ? JSON.parse(localStorage.getItem("user")) : {}
    }
  },
  methods:{
    logout(){
      this.$router.push("/login")
      localStorage.removeItem("user")
      this.$message.success("退出成功")
    }
  }
}
</script>

<style scoped>

</style>