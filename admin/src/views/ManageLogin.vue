<template>
    <div class="manageLogin-info">
        <form class="layui-form" action="javascript:void(0);">
            <div class="layui-form-item">
                <img class="logo" src="../assets/image/login.png" />
                <div class="title">博客系统管理后台登录</div>
            </div>
            <div class="layui-form-item">
                <input name="username" v-model="form.username" placeholder="用户名" lay-verify="required" hover class="layui-input"  />
            </div>
            <div class="layui-form-item">
                <input name="password" v-model="form.password" type="password" placeholder="密码" lay-verify="required" hover class="layui-input"  />
            </div>
            <div class="layui-form-item">
                <button type="button" class="pear-btn pear-btn-success login" lay-submit  @click="onSubmit">
                    登 录
                </button>
            </div>
        </form>
    </div>
</template>

<script setup>
import { ref,reactive } from 'vue'
import { login } from "@/api/api"
import { useRouter } from "vue-router";
import { ElMessage } from 'element-plus/lib/components';

    const router = useRouter()

    const form = ref({
        username: 'admin',
        password: '123456'
    })

    const formRef = ref(null)
    const onSubmit = async () => {
        let res = await login({
            username: form.value.username,
            password: form.value.password,
        })

        if(res){
            window.localStorage.setItem('token', res.data);
            router.replace("/")
        }
    }
    window.addEventListener("popstate",function(e) {
        history.pushState(null, null, document.URL);
    }, false);

</script>

<style lang="scss" scoped>
   .manageLogin-info{
        min-height: 100%;
        width: 100%;
        height: 100vh;
        background-image: url("../assets/image/background.svg");
        background-position: 50% 50%;
        background-size: cover;
        background-repeat: no-repeat;
        margin-left: auto;
        margin-right:auto;
        display: flex;
        justify-content: center;
        align-items: center;

        .layui-form {
            width: 320px !important;
            margin: auto !important;
            margin-top: 160px !important;
        }

        .layui-form button {
            width: 100% !important;
            height: 44px !important;
            line-height: 44px !important;
            font-size: 16px !important;
            background-color: #5FB878 !important;
            font-weight: 550 !important;
            border: #36b368;
            color: #fff !important;
        }

        .layui-form-checked[lay-skin=primary] i {
            border-color: #5FB878 !important;
            background-color: #5FB878 !important;
            color: #fff !important;
        }

        .layui-tab-content {
            margin-top: 15px !important;
            padding-left: 0px !important;
            padding-right: 0px !important;
        }

        .layui-form-item {
            margin-top: 20px !important;
        }

        .layui-input {
            height: 44px !important;
            line-height: 44px !important;
            padding-left: 15px !important;
            border-radius: 3px !important;
        }

        .layui-input:focus {
            box-shadow: 0px 0px 2px 1px #5FB878 !important;
        }

        .layui-form-danger:focus{
            box-shadow: 0px 0px 2px 1px #f56c6c !important;
        }

        .logo {
            width: 60px !important;
            margin-top: 10px !important;
            margin-bottom: 10px !important;
            margin-left: 20px !important;
        }

        .title {
            font-size: 30px !important;
            font-weight: 550 !important;
            margin-left: 20px !important;
            color: #5FB878 !important;
            display: inline-block !important;
            height: 60px !important;
            line-height: 60px !important;
            margin-top: 10px !important;
            position: absolute !important;
        }

        .desc {
            width: 100% !important;
            text-align: center !important;
            color: gray !important;
            height: 60px !important;
            line-height: 60px !important;
        }

        body {
            background-repeat:no-repeat;
            background-color: whitesmoke;
            background-size: 100%;
            height: 100%;
        }

        .code {
            float: left;
            margin-right: 13px;
            margin: 0px !important;
            border: #e6e6e6 1px solid;
            display: inline-block!important;
        }

        .codeImage {
            float: right;
            height: 42px;
            border: #e6e6e6 1px solid;
        }

   }

    .item-flex{
        display: flex;
        justify-content: space-between;
    }
   .code-info{
        font-family:Arial;
        font-style:italic;
        color:blue;
        font-size:20px;
        border:0;
        padding: 0 10px;
        letter-spacing:3px;
        font-weight:bolder;
        float:left;
        cursor:pointer;
        width: 80px;
        height:44px;
        line-height:60px;
        text-align:center;
        vertical-align:middle;
        background-color:#D8B7E3;

        display: flex;
        justify-content: center;
        align-items: center;
    }

</style>
