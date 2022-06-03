<template>
    <el-container class="app-wrapper">
        <!-- 菜单 -->
        <el-aside :width="asideWidth" class="sidebar-container">
            <el-scrollbar style="height:'100%'">
                <el-menu
                    active-text-color="#ffd04b"
                    background-color="#20222A"
                    class="el-menu-vertical-demo"
                    :default-active="activePath"
                    text-color="#fff"
                    menu-trigger
                    router
                    :collapse="!toggleico"
                >
                    <el-menu-item index="labelList" :class="[activePath == '/labelList'?'is-active':'']"  @click="goUrl('后台管理','','labelList')">
                        <i class="layui-icon layui-icon-engine"></i>
                        <span class="caidan-auth-children">标签管理</span>
                    </el-menu-item>
                    <el-menu-item index="cateList" :class="[activePath == '/cateList'?'is-active':'']"  @click="goUrl('后台管理','','cateList')">
                        <i class="layui-icon layui-icon-engine"></i>
                        <span class="caidan-auth-children">分类管理</span>
                    </el-menu-item>
                    <el-menu-item index="postsList" :class="[activePath == '/postsList'?'is-active':'']"  @click="goUrl('后台管理','','postsList')">
                        <i class="layui-icon layui-icon-engine"></i>
                        <span class="caidan-auth-children">文章管理</span>
                    </el-menu-item>
                    <el-menu-item index="comment" :class="[activePath == '/comment'?'is-active':'']"  @click="goUrl('后台管理','','comment')">
                        <i class="layui-icon layui-icon-engine"></i>
                        <span class="caidan-auth-children">评论管理</span>
                    </el-menu-item>
                    <el-menu-item index="linkList" :class="[activePath == '/linkList'?'is-active':'']"  @click="goUrl('后台管理','','linkList')">
                        <i class="layui-icon layui-icon-engine"></i>
                        <span class="caidan-auth-children">友联管理</span>
                    </el-menu-item>
                    <el-menu-item index="accountList" :class="[activePath == '/accountList'?'is-active':'']"  @click="goUrl('后台管理','','accountList')">
                        <i class="layui-icon layui-icon-engine"></i>
                        <span class="caidan-auth-children">账户管理</span>
                    </el-menu-item>
                </el-menu>
            </el-scrollbar>
        </el-aside>
        
        <!-- 菜单end -->
        <el-container class="container" :class="{hidderContainer: !toggleico}">
            <!-- 头部 -->
            <el-header  style="font-size: 16px;background-color: #ffffff">
                <div class="toolbar">
                    <div class="toolbar-open-close">
                        <div @click="toggleClick">
                            <i class="layui-icon" :class="[!toggleico ? 'layui-icon-shrink-right' : 'layui-icon-spread-left']"></i>
                        </div>
                    </div>
                    <div class="toolbar-select">
                        <el-dropdown>
                            <el-icon class="elIcon"><setting/></el-icon>
                            
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <el-dropdown-item @click="loginOut()">退出</el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
                        <div class="username-sel">
                            <span>{{ username }}</span>
                        </div>
                    </div>
                </div>
            </el-header>
            <!-- 头部end -->
            <!-- 面包屑 -->
            <!-- <breadcrumb></breadcrumb> -->
            <!-- 内容 -->
            <el-main style="background-color: #f2f2f2">
                <router-view />
            </el-main>
            <!-- 内容end -->
        </el-container>
    </el-container>
</template>
<script setup>
import { ref,computed,watch } from "vue"                                                                                                                                                                                                                 
import { useRouter } from "vue-router";
import {Location,Document,Menu as IconMenu,Setting,} from '@element-plus/icons-vue'
import variables from "@/assets/styles/variables.scss"
import store from "@/store/index.ts"

    const router = useRouter()//路由跳转
    const menuList = ref([])//菜单list
    const token = window.localStorage.getItem("token")
    const username = ref(window.localStorage.getItem("username"))//获取登录名称
    const initMenusList = async () =>{//判断是否已经存在token是否已经登录
        if(!window.localStorage.getItem("token")){
            router.push("/manageLogin")                                                                                                                                  
        }else{
            if(sessionStorage.getItem(`url`)){
                router.push(sessionStorage.getItem(`url`))
            }else{
                if(localStorage.getItem("token")){
                    router.push('/labelList')
                }else{
                    router.push('/manageLogin')
                }
            }
        }
    }
    initMenusList()

    const toggleico = ref(true)//后台管理的图片是显示还是隐藏
    const asideWidth = computed(()=>{//菜单栏的宽度
        return toggleico.value ? variables.sideBarWidth : variables.hideSideBarWidth
    })
    const toggleClick = () => {//点击折叠
        toggleico.value = !toggleico.value
    }

    //点击刷新之后
    const activePath = ref(sessionStorage.getItem(`url`) || '/labelList')
    //判断缓存里面是否有url 如果有跳转到对应页面   如果没有就直接跳转到首页看板
    
    
    //保存信息和url到缓存里面
    const goUrl = (authName,authName2,url) => {
        store.commit("openNames", authName);
        store.commit("activeName", authName2);
        sessionStorage.setItem(`url`,`/${url}`)
        //消去订单统计的背景色
        activePath.value = sessionStorage.getItem(`url`)
    }

    //退出登录
    const loginOut = () =>{
        localStorage.clear() //清除缓存
        sessionStorage.removeItem("url");
        router.push("/manageLogin")
    }


</script>
<style lang="scss" scoped>
    @import "../assets/css/manageindex.scss";
</style>
