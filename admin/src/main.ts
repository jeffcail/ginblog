import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import '@/assets/styles/index.scss'
import Particles from 'particles.vue3'

import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import { vLoading } from 'element-plus/es/components/loading/src/directive'


// 跳转前清除所有通知
router.beforeEach((to, from, next) => {
    // 修改文档标题
    
    next()
})
const app = createApp(App)
app.use(ElementPlus, {locale: zhCn})
app.directive('load', vLoading)
app.use(Particles)
app.use(store)
app.use(router)
app.mount('#app')
