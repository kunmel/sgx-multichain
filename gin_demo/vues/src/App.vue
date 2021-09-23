<template>
  <el-row>
    <el-col :span="6">
      <el-card class="in-card">
        <el-form ref="inform" :model="inform" label-width="80px" >
          <el-form-item label="属性1" style="width:70%">
            <el-input v-model="inform.arg1"></el-input>
          </el-form-item>
          <el-form-item label="属性2" style="width:70%">
            <el-input v-model="inform.arg2"></el-input>
          </el-form-item>
          <el-form-item label="属性3" style="width:70%">
            <el-input v-model="inform.arg3"></el-input>
          </el-form-item>
          <el-form-item label="属性4" style="width:70%">
            <el-input v-model="inform.arg4"></el-input>
          </el-form-item>
          <el-form-item label="特殊资源">
            <el-radio-group v-model="inform.chainID">
              <el-radio label="ethereum"></el-radio>
              <el-radio label="hyperledger fabric"></el-radio>
              <el-radio label="auto"></el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">立即创建</el-button>
            <el-button @click="onCancel">取消</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
    <el-col :span="6"  >
      <el-card class="out-card" >
        <div slot="header" class="clearfix">
          <span>事务结果</span>
        </div>
        <div class="result-show">事务状态： {{outform.status}}</div>
        <div class="result-show">txID： {{outform.txID}}</div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import axios from 'axios'
export default {
  name: 'App',
  methods: {
    async onSubmit () {
      this.fullscreenLoading = true
      const loading = this.$loading({
        lock: true,
        text: '正在获取结果',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      const res = await axios.get('http://127.0.0.1:8000/v1/onchain', {
        params: {
          arg1: this.inform.arg1,
          arg2: this.inform.arg2,
          arg3: this.inform.arg3,
          arg4: this.inform.arg4,
          chainID: this.inform.chainID
        }
      })
      this.outform.status = res.data.status
      this.outform.txID = res.data.txID
      this.outform.blockID = res.data.blockID
      this.fullscreenLoading = false
      loading.close()
    },
    onCancel () {
      console.log('clearAll')
    }
    // onSubmit () {
    //   console.log('submit')
    // }
  },
  data () {
    return {
      fullscreenLoading: false,
      txID: '!@##',
      blockID: '##@!',
      inform: {
        arg1: '',
        arg2: '',
        arg3: '',
        arg4: '',
        chainID: ''
      },
      outform: {
        status: '',
        txID: '',
        blockID: ''
      }
    }
  }
}
</script>

<style>
.in-card {
  width: 600px;
  margin-top: 100px;
  margin-left: 300px;
}
.out-card {
  width: 600px;
  margin-top: 100px;
  margin-left: 500px;
}
</style>
