<template>
  <el-card>
    <el-table :data = tableData style="width: 100% ;font-size: 15px" max-height="600">

      <el-table-column fixed prop="Time" label="发送日期" width="200" align="center" sortable></el-table-column>
<!--      <el-table-column prop="FoodName" label="菜名" width="120" align="center"></el-table-column>-->
      <el-table-column prop="Advise" label="建议" width="550" align="center"></el-table-column>
      <el-table-column fixed="right" label="操作" width="240" align="center">

        <template slot-scope="scope">
          <el-button @click.native.prevent="deleteRow(scope.$index, tableData)"
            @click="deleteOpinion(scope.row)"
            type="text"
            size="medium">收到</el-button>
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
      this.getOpinions()
    },
    methods: {
      getOpinions () {
        // this.$http.get("/CanteenPC/Opinion").then(res => { this.tableData = res.data })
        request.get("/CanteenPC/Opinion",).then(res=>{
          this.tableData= res
          // console.log(res)
          console.log(res.length)
        })

      },

      deleteRow (index, rows) {
        rows.splice(index, 1)
      },
      deleteOpinion (val) {
        // const data = new FormData()
        const a = val.ID
        // data.append('Id', a)
        console.log(a)
        // this.$http.post('deleteOpinion', data)

        request.post("/CanteenPC/DeleteOpinion",{"Id":a})
        this.$notify({
          title: '收到意见',
          type: 'success',
          duration: 1500,
          position: 'bottom-right'
        })
      }
    }
  }
</script>

<style scoped>

</style>
