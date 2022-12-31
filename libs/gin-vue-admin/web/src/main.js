import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import './style/element_visiable.scss'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
// 引入gin-vue-admin前端初始化相关内容
// import './core/gin-vue-admin'
// by-xm 由于是default匿名函数，所以可以用别名来指代该函数，如本次使用的run
import run from '@/core/gin-vue-admin.js'
import { store } from '@/pinia'
import auth from '@/directive/auth'
import App from './App.vue'
import { initDom } from './utils/positionToCode'

/*
 * @description 导入加载进度条，防止首屏加载时间过长，用户等待
 */
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'
Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 500 })
Nprogress.start()

/*
 * 无需在这块结束，会在路由中间件中结束此块内容
 */

const app = createApp(App)
// 当为开发模式时，请求了一个拼接的url
// initDom(app)

app.config.errorHandler = (err, instance, info) => {
  // 向追踪服务报告错误
  console.log("errorHandler", err, instance, info)
}
app.config.warnHandler = (msg, instance, trace) => {
  // `trace` is the component hierarchy trace
  // 向追踪服务报告警告
  console.log("warnHandler", msg, instance, trace)
}
// app.config.performance = true

app
  .use(run)
  .use(store)
  .use(auth)
  .use(router)
  .use(ElementPlus, { locale: zhCn })
  .mount('#app')

export default app
