import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

import ManageIndex from '@/views/ManageIndex.vue';
import ManageLogin from '@/views/ManageLogin.vue';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'manageIndex',
        component: ManageIndex,
        redirect:'labelList',
        children: [
            {
                path: '/labelList', name: 'labelList',
                component: () => import('@/views/labelList.vue'),
                meta: {
                    title: '标签',
                }
            },
            {
                path: '/cateList', name: 'cateList',
                component: () => import('@/views/cateList.vue'),
                meta: {
                    title: '分类',
                }
            },
            {
                path: '/postsList', name: 'postsList',
                component: () => import('@/views/postsList.vue'),
                meta: {
                    title: '文章',
                }
            },
            {
                path: '/comment', name: 'comment',
                component: () => import('@/views/comment.vue'),
                meta: {
                    title: '评论',
                }
            },
            {
                path: '/linkList', name: 'linkList',
                component: () => import('@/views/linkList.vue'),
                meta: {
                    title: '友联',
                }
            },
            {
                path: '/accountList', name: 'accountList',
                component: () => import('@/views/accountList.vue'),
                meta: {
                    title: '账号',
                }
            }
        ]

    },
    {
        path: '/manageLogin',//登录页面
        name: 'manageLogin',
        component: ManageLogin
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
