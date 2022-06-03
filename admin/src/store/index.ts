import { createStore } from 'vuex'
import { ElLoading } from 'element-plus';
import createPersistedstate from 'vuex-persistedstate'

export default createStore({
    state: {
         /** 左边导航栏: 当前展开的一级菜单名 */
         openNames: '',
         /** 左边导航栏: 当前选中的菜单名 */
         activeName: '',
    },
    mutations: {
        activeName(state: any, name) {
            state.activeName = name;
        },
        openNames(state: any, name) {
            state.openNames = name;
        },
    },
    actions: {
    },
    modules: {
    },
    plugins: [
        createPersistedstate({
            storage: window.sessionStorage,
            key: "store",
            reducer(state) {
                
                return {
                    menuListAll: state.menuListAll,
                    openNames: state.openNames,
                    activeName: state.activeName
                }
            }
        })
    ]
})
