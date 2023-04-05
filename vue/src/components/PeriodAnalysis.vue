<template>
  <el-card>
    <div id="main" style=""></div>
  </el-card>
</template>

<script>
// const echarts = require('echarts')
import * as echarts from "echarts"
import request from "@/utils/request";

export default {
  mounted () {
    this.initChart()
  },
  created () {
    this.getRoomAnalysis()
  },
  data () {
    return {
      list: {}
    }
  },
  methods: {
    // 获得数据
    getRoomAnalysis () {
      // this.$http.get("foodAnalysis").then(res => {
      //   this.list = res.data
      // })
      request.get("/CanteenPC/PeriodAnalysis").then(res=>{ this.list = res
        console.log(res)
      })
    },
    initChart () {
      this.char = echarts.init(document.getElementById('main'))
      this.char.setOption({
        color: ['#3398DB'],
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow'
          }
        },
        title: {
          text: '时间段折线图',
          textStyle: {
            color: '#333744'
          },
          padding: [10, 0, 0, 100]
        },
        xAxis: [{ type: 'category' }],
        yAxis: [{ type: 'value' }],
        series: [
          {
            name: '预约次数',
            type: 'line',
            barWidth: '60%',
            showBackground: true,
            backgroundStyle: {
              color: 'rgba(180, 180, 180, 0.2)'
            }
          }
        ]
      })
      // this.$http.get('http://localhost:8003/CanteenPC/foodAnalysis')
      request.get("/CanteenPC/PeriodAnalysis").then((res) => {
        console.log('访问后台')
        this.list = res
        this.char.setOption({
          series: [
            {
              name: '预约次数',
              type: 'line',
              barWidth: '60%',
              data: this.list.cnt
            }
          ],
          xAxis: [
            {
              type: 'category',
              data: this.list.name
            }
          ]
        })
      })
    }
  }
}
</script>

<style scoped>
#main {
  height: 600px;
  display: flex;
  align-items: center;
  justify-content: space-around;
}
</style>
