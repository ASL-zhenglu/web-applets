<template>
  <el-container>
    <el-card class="editFood-card" >
      <!-- 添加按钮 -->
      <el-button type="primary" @click="addRoomDialogVisible=true" style="margin-bottom: 15px;">添加会议室</el-button>

      <!-- 展示 -->
      <el-table :data="tableData" border :row-style="{height: '100px'}" style="font-size: 15px">
        <!-- 名称 -->
<!--        <el-table-column label="名称" height="180"   align="center">-->
<!--          <template slot-scope="scope">-->
<!--            <el-image :src="scope.row.Name"></el-image>-->
<!--          </template>-->
<!--        </el-table-column>-->
        <!-- 大小 -->
        <el-table-column label="名称" align="center" width="200">
          <template slot-scope="scope">
            <span> {{ scope.row.RoomName }}</span>
          </template>
        </el-table-column>
        <!-- 大小 -->
        <el-table-column label="大小" align="center" width="200">
          <template slot-scope="scope">
            <span> {{ scope.row.Size }}</span>
          </template>
        </el-table-column>

        <!-- 菜品价格 -->
        <el-table-column label="备注" align="center">
          <template slot-scope="scope">
<!--            <el-tag type="danger" size="medium">{{ scope.row.Remark }}</el-tag>-->
            <span> {{ scope.row.Remark }}</span>
          </template>
        </el-table-column>

        <!-- 操作 -->
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button size="medium" type="primary" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
            <el-button size="medium" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加对话框 -->
    <el-dialog title="添加会议室" :visible.sync="addRoomDialogVisible" width="50%" class="addRoom-dialog">
      <el-form
        :model="addRoomForm"
        ref="addRoomFormRef"
        :rules="addRoomFormRules"
        label-width="80px"
        enctype="multipart/form-data">
        <!-- 名称 -->
        <el-form-item label="名称" label-width="80px" prop="name">
          <el-input v-model="addRoomForm.name" style="width: 60%;" clearable></el-input>
        </el-form-item>

        <!-- 大小 -->
        <el-form-item label="大小" label-width="80px" prop="size">
<!--          <el-input v-model="addFoodForm.price" placeholder="单位：/元" style="width: 60%;" clearable></el-input>-->
        <el-input v-model="addRoomForm.size" placeholder="单位：xx人" style="width: 60%;" clearable></el-input>
        </el-form-item>

        <el-form-item label="备注" label-width="80px" prop="remark">
          <el-input v-model="addRoomForm.remark" style="width: 60%;" clearable></el-input>
        </el-form-item>
        <!-- 封面 -->
<!--        <el-form-item label="图片" prop="image">-->
<!--          <el-upload-->
<!--            ref="upload"-->
<!--            list-type="picture-card"-->
<!--            name="foodPhoto"-->
<!--            multiple-->
<!--            :action=foodImageUrl-->
<!--            :class="{ hide: hideUp }"-->
<!--            :auto-upload="false"-->
<!--            :on-change="imageChange"-->
<!--            :data="addFoodForm"-->
<!--            :limit="1"-->
<!--            :on-preview="handlePictureCardPreview"-->
<!--            :on-remove="handleRemovePicture">-->
<!--            <i class="el-icon-plus"></i>-->
<!--          </el-upload>-->
<!--          &lt;!&ndash; 封面预览 &ndash;&gt;-->
<!--          <el-dialog :visible.sync="dialogVisible">-->
<!--            <img width="100%" :src="dialogImageUrl">-->
<!--          </el-dialog>-->
<!--        </el-form-item>-->
      </el-form>
<!--      <el-button class="btn" plain type="primary" @click="submit()">提 交</el-button>-->
      <el-button class="btn" plain type="primary" @click="handleAdd">提 交</el-button>
    </el-dialog>
  </el-container>
</template>

<script>
import request from '@/utils/request'

export default {
  data () {
    return {
      tableData: [],
      // 增加菜品对话框
      addRoomDialogVisible: false,
      // ---- 下面是添加菜品的数据
      // foodImageUrl: 'http://localhost:8003/CanteenPC/addFood',
      // hideUp: false,
      addRoomForm: {},
      dialogVisible: false,
      // dialogImageUrl: '',
      addRoomFormRules: {
        name: [{ required: true, message: '请输入会议室名称', trigger: 'blur' }],
        size: [{ required: true, message: '请输入会议室大小', trigger: 'blur' }],
        remark: [{ required: true, message: '请输入会议室备注', trigger: 'blur' }]
      }
    }
  },
  created () {
    this.getRoom()
  },
  methods: {
    // 生命周期函数，获取菜单
    getRoom () {
      // this.$http.get("manageFood").then(res => {
      //   this.tableData = res.data
      // })
      request.get("/CanteenPC/MeetingRoom").then(res=>{
        this.tableData = res
      })
    },
    // 编辑菜品按钮
    handleEdit (index, row) {

      this.$prompt(
        '请输入大小'
        , { confirmButtonText: '确定', cancelButtonText: '取消' }).then(({ value }) => {
        // const formData = new FormData()
        // const foodID = row.ID
        // const foodPrice = value

        // formData.append('id', foodID)
        // formData.append('price', foodPrice)

        // this.$http.put('editFood', formData).then(() => { this.getFood() })
        request.put("/CanteenPC/EditRoom",{"id":row.ID,"value":parseInt(value)}).then(()=>{ this.getRoom()})
        this.$message.success('你修改的大小是: ' + value)
      }).catch(() => {
        this.$message.info('取消输入')
      })

    },
    // 删除会议室按钮
    handleDelete (index, row) {
      console.log(row.ID)
      this.$confirm('永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // this.$http.delete(`deleteFood/${row.ID}`).then(() => { this.getFood() })
        request.post("/CanteenPC/DeleteRoom",{"id":row.ID}).then(()=>this.getRoom())
        this.$message.success('删除成功!')
      }).catch(() => {
        this.$message.info('已取消删除')
      })
    },
    handleAdd() {
      this.$refs.addRoomFormRef.validate(  async (valid)=> {

        if (!valid) {
          return this.$message.error('添加失败！')
        }
        // const { data: res } = await this.$http.post(`/CanteenPC/login/${this.user.username}/${this.user.password}`)
        // axios.post("/CanteenPC/login",{username:this.user.username,password:this.user.password})
        // const{ data : res} = this.$http.post(`CanteenPC/login/${this.user.username}/${this.user.password}`)
        request.post("/CanteenPC/AddRoom",{
          name:this.addRoomForm.name.toString(),
          size:this.addRoomForm.size.toString(),
          remark:this.addRoomForm.remark.toString()
        }).then(b=>{
          console.log(b)
          if(b.msg === "1") {
            this.$message.success('添加成功！')
            this.getRoom()
            this.$router.push("/CanteenPC/MeetingRoom")
          } else {
            this.$message.success('添加失败，会议已经存在')
          }

        })

      });
    }
    // ---- 下面是添加会议室的方法
    // imageChange (file, fileList, name) {
    //   this.hideUp = (fileList.length === 1)
    //   console.log(file, fileList)
    // },
    // handlePictureCardPreview (file) {
    //   this.dialogImageUrl = file.url
    //   this.dialogVisible = true
    // },
    // handleRemovePicture (file, fileList) {
    //   this.hideUp = false
    //   console.log(file, fileList)
    // },
    // submit () {
    //   this.addRoomForm = {}
    //   this.$refs.upload.submit()
    //   this.$refs.upload.clearFiles()
    //   // this.hideUp = false
    //   this.$message.success("添加新菜品成功！")
    // }
  }
}
</script>

<style scoped>
.editFood-card {
  width: 100%;
}

.addRoom-dialog {
  height: 100%;
  display: flex;
  flex-direction: column;
  margin-left: 5%;
}

.btn {
  margin-top: 30px;
  margin-left: 70%;
  margin-right: 10%;
  margin-bottom: 8%;
  float: right;
}
</style>
