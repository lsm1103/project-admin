{{ .Html }}

<script>
export default {
  name: "List2"
}
</script>

<script setup>
import { ref } from 'vue'
import axios from "axios";
import {ElMessage} from "element-plus";
import { formatDate } from '@/utils/format'
import { Search,CloseBold } from '@element-plus/icons-vue'

const tableData = ref([])
const searchKW = ref('')
const selectd = ref([])
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const nameList = ref([])
let searchHistoryKW = ""
let rowHistory = {}

const baseUrl = "http://127.0.0.1:810"
const service = axios.create({
  baseURL: baseUrl,
  timeout: 99999
})
// http response 拦截器
service.interceptors.response.use(
    response => {
      if (response.data.code === 200 || response.headers.success === 'true') {
        if (response.headers.msg) {
          response.data.msg = decodeURI(response.headers.msg)
        }
        return response.data
      } else {
        ElMessage({
          showClose: true,
          message: response.data.msg || decodeURI(response.headers.msg),
          type: 'error'
        })
        return response.data.msg ? response.data : response
      }
    },
    error => {
      if (!error.response) {
        ElMessageBox.confirm(`
        <p>检测到请求错误</p>
        <p>${error}</p>
        `, '请求报错', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '稍后重试',
          cancelButtonText: '取消'
        })
        return
      }
      switch (error.response.status) {
        case 500:
          ElMessageBox.confirm(`
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>
        `, '接口报错', {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '清理缓存',
            cancelButtonText: '取消'
          })
              .then(() => {
                const userStore = useUserStore()
                userStore.token = ''
                localStorage.clear()
                router.push({ name: 'Login', replace: true })
              })
          break
        case 404:
          ElMessageBox.confirm(`
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
          `, '接口报错', {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '我知道了',
            cancelButtonText: '取消'
          })
          break
      }
      return error
    }
)

// 获取列表信息
const getList = async (val) => {
  console.log("getList",val)

  let res = await service({
    url: '/admin/Application/v1/gets',
    method: 'post',
    data:val
  })
  if (res.code === 200) {
    if (!res.data.list){
      tableData.value = []
      ElMessage({
        showClose: true,
        message: "找不到数据",
        type: 'warning'
      })
      return
    }

    const tmp = []
    const tmp_ = []
    res.data.list.forEach((item) => {
      item.create_time = formatDate(item.create_time)
      //人名筛选功能
      if ( !tmp_.includes(item.create_user )){
        tmp.push({ text: item.create_user, value: item.create_user })
        tmp_.push(item.create_user)
      }
    });
    nameList.value = tmp

    tableData.value = res.data.list
    currentPage.value = res.data.current
    const total_ = res.data.current===1? res.data.list.length : (res.data.pageSize*(res.data.current-1))+res.data.list.length
    total.value = res.data.isNext? total_+1 : total_
    pageSize.value = res.data.pageSize
    console.log("total", total_, total.value, currentPage.value, pageSize.value)
  }
}
getList({
  "current": currentPage.value,
  "orderBy": "create_time",
  "pageSize": pageSize.value,
  "query": [],
  "sort": "desc"
})

//搜索
const toSearch = () => {
  console.log("toSearch",searchKW.value, tableData.value.length == 0, tableData.value)
  if(searchKW.value === "") {
    if (tableData.value.length == 0 || searchHistoryKW != ""){
      getList({
        "current": 1,
        "orderBy": "create_time",
        "pageSize": pageSize.value,
        "query": [],
        "sort": "desc"
      })
    }
    searchHistoryKW = searchKW.value
    return;
  } else if (searchHistoryKW == searchKW.value){
    return;
  }else {
    // 请求查询接口，将列表展现出来
    getList({
      "current": 1,
      "orderBy": "create_time",
      "pageSize": pageSize.value,
      "query": [
        {
          "handle": "like",
          "key": "zh_name",
          "nextHandle": "and",
          "val": searchKW.value
        }
      ],
      "sort": "desc"
    })
    searchHistoryKW = searchKW.value
  }
}

//列表多选事件触发处理
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  console.log("multipleSelection",multipleSelection.value, val)
  multipleSelection.value = val
}
//人名筛选功能
const filterHandler = (value, row, column) => {
  const property = column['property']
  return row[property] === value
}

//add
const add = (val) => {
  if (val.hasOwnProperty("create_user")){
    val.create_user = parseInt(val.create_user)
  }
  return service({
    url: '/admin/Application/v1/',
    method: 'post',
    data:val
  })
}
//del
const del = (val) =>{
  return service({
    url: '/admin/Application/v1/',
    method: 'delete',
    data:val
  })
}
//update
const update = (val) =>{
  if (val.hasOwnProperty("create_user")){
    val.create_user = parseInt(val.create_user)
  }
  return service({
    url: '/admin/Application/v1/',
    method: 'put',
    data:val
  })
}

//状态更改
const handleStateChange = (row) => {
  console.log("handleStateChange", row)
  update({
    "id": row.id,
    "state": row.state,
  })
  ElMessage({
    showClose: true,
    message: row.id+"操作成功",
    type: 'success'
  })
}
//状态更改
const handleBatchStateChange = () => {
  console.log("handleBatchStateChange",multipleSelection.value)
  multipleSelection.value.forEach((item) => {
    if (item.state != -1){
      update({
        "id": item.id,
        "state": -1,
      })
      ElMessage({
        showClose: true,
        message: item.id+"操作成功",
        type: 'success'
      })
      item.state = -1
    }
  })
}
//新增
const handleAdd = () => {
  dialogType.value = "add"
  formData.value = {}
  dialogFormVisible.value = true
}
//修改
let editIndex = 0
const handleEdit = (index, row) => {
  console.log("handleEdit", index, row)
  editIndex = index
  dialogType.value = "edit"
  formData.value = JSON.parse(JSON.stringify(row))
  // 历史记录
  rowHistory = JSON.parse(JSON.stringify(row))
  dialogFormVisible.value = true
}
//删除
const handleDelete = async(index, row) => {
  console.log("handleDelete", index, row)
  let res = await del({
        "id": row.id
  })
  console.log("handleDelete-res", res)
  if (res.code === 200) {
    ElMessage({
      showClose: true,
      message: row.id+"操作成功",
      type: 'success'
    })
    getList({
      "current": 1,
      "orderBy": "create_time",
      "pageSize": pageSize.value,
      "query": [],
      "sort": "desc"
    })
  }
  row.visible = false
}
//批量删除
const handleBatchDelete = async() => {
  console.log("handleBatchDelete", multipleSelection)
  for (let item of multipleSelection.value) {
    let res = await del({
      "id": item.id
    })
    if (res.code === 200) {
      ElMessage({
        showClose: true,
        message: item.id+"操作成功",
        type: 'success'
      })
      console.log("handleBatchDelete-res", res)
    }
  }

  getList({
    "current": 1,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": [],
    "sort": "desc"
  })
}

//分页
const handleSizeChange = (val) => {
  pageSize.value = val
  let query = []
  if(searchKW.value != ""){
    query = [
      {
        "handle": "like",
        "key": "zh_name",
        "nextHandle": "and",
        "val": searchKW.value
      }
    ]
  }
  getList({
    "current": currentPage.value,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": query,
    "sort": "desc"
  })
}
const handleCurrentChange = (val) => {
  currentPage.value = val
  let query = []
  if(searchKW.value != ""){
    query = [
      {
        "handle": "like",
        "key": "zh_name",
        "nextHandle": "and",
        "val": searchKW.value
      }
    ]
  }
  getList({
    "current": currentPage.value,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": query,
    "sort": "desc"
  })
}

// 弹窗
const dialogType = ref('add')
const dialogTitle = ref('新增角色')
const dialogFormVisible = ref(false)
const formData = ref({})
const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('请输入正整数'))
  }
  return callback()
}
const rules = ref({
  create_user: [
    { required: true, message: '请输入创建人id', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  ]
})
const enterDialog = async() => {
  console.log("enterDialog", formData.value)
  if (dialogType.value == "add"){
    let res = await add(formData.value)
    console.log("enterDialog-add-res", res)
    if (res.code === 200) {
      ElMessage({
        showClose: true,
        message: "新增操作成功",
        type: 'success'
      })
      getList({
        "current": 1,
        "orderBy": "create_time",
        "pageSize": pageSize.value,
        "query": [],
        "sort": "desc"
      })
    }
  } else if (dialogType.value == "edit"){
    let data = {}
    formData.value.forEach(function(value,key){
      if ( value != rowHistory[key] ){
        data[key] = value
      }
    });
    console.log("data", data)
    let res = await update(data)
    console.log("enterDialog-update-res", res)
    if (res.code === 200) {
      ElMessage({
        showClose: true,
        message: "编辑操作成功",
        type: 'success'
      })
      getList({
        "current": 1,
        "orderBy": "create_time",
        "pageSize": pageSize.value,
        "query": [],
        "sort": "desc"
      })
    }
  }
  // 历史记录
  rowHistory = JSON.parse(JSON.stringify(formData.value))
  tableData.value[editIndex] = JSON.parse(JSON.stringify(formData.value))
  dialogFormVisible.value = false
}
const closeDialog = () => {
  dialogFormVisible.value = false
}
const beforeReferenc = () => {
  formData.value = rowHistory
}

</script>

{{ .Style }}