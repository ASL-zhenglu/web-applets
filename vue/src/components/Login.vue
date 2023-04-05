<template>
  <div class="wrapper">
    <div style="margin: 200px auto; background-color: rgba( 182,193,198, 0.5); width: 350px; height: 300px; padding: 20px; "><!--  (圆角）    border-radius: 10px-->

      <div style="margin: 20px 0; text-align: center; font-size: 24px"><b>登 录</b></div>
      <el-form :model="user" :rules="rules" ref="userForm">

        <el-form-item prop="username">
          <el-input size="medium" style="margin: 10px 0" prefix-icon="el-icon-user" v-model="user.username"></el-input>
        </el-form-item>

        <el-form-item prop="password">
          <el-input size="medium" style="margin: 10px 0" prefix-icon="el-icon-lock" show-password v-model="user.password"></el-input>
        </el-form-item>

        <el-form-item style="margin: 10px 0; text-align: right">
          <el-button type="primary" size="small" autocomplete="off" @click="login">登录</el-button>
<!--          <el-button type="info" size="small" autocomplete="off" @click="$router.push('/register')">重置</el-button>-->
          <el-button type="info" size="small" autocomplete="off" @click="resetLoginForm">重置</el-button>

        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import request from '@/utils/request'
import axios from 'axios'
// import qs from 'qs'
// Vue.prototype.$qs = qs
export default {
  name: "Login",
  data() {
    return {
      user: {
        username: "",
        password: ""
      },
      rules: {
        username: [
          {required: true, message: '请输入用户名', trigger: 'blur'},
          {min: 1, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '请输入密码', trigger: 'blur'},
          {min: 1, max: 10, message: '长度在 1 到 20 个字符', trigger: 'blur'}
        ],
      },
    }
  },
  methods:{
    resetLoginForm () {
      this.$refs.userForm.resetFields() //ref->userForm
    },
    login(){
      this.$refs.userForm.validate(  async (valid)=> {
        // if (valid) {
        //   this.request.post("/login",this.user).then(ans=>{
        //     if(ans.code === '200'){
        //       localStorage.setItem("user",JSON.stringify(ans.data)) //存储用户信息到浏览器上面去
        //       this.$router.push("/")
        //       this.$message.success("登录成功")
        //     } else {
        //       this.$message.error(ans.msg)
        //     }
        //   })
        // }
        if (!valid) {
          return this.$message.error('登录失败！')
        }
         // const { data: res } = await this.$http.post(`/CanteenPC/login/${this.user.username}/${this.user.password}`)
        // axios.post("/CanteenPC/login",{username:this.user.username,password:this.user.password})
        // const{ data : res} = this.$http.post(`CanteenPC/login/${this.user.username}/${this.user.password}`)

        request.post("/CanteenPC/login",{
          username:this.user.username,
          password:this.user.password
        }).then(b=>{
          console.log(b)
          // console.log(typeof b)
          // var r = JSON.parse(b)
          if(b.msg === "1") {
            this.$message.success('登录成功！')
            this.$router.push('/CanteenPC/home')
          } else {
            this.$message.success('登录失败')
          }

        })

      });
    }
  }
}
</script>

<style scoped>
  .wrapper {
    height: 100vh;
    /*100%窗口高度*/
    background-image: url("../assets/bg.jpg");
    overflow: hidden;
    background-size: 100%;
    display: flex;
    align-items: center;
  }
</style>
