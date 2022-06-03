const AutoImport = require('unplugin-auto-import/webpack')
const Components = require('unplugin-vue-components/webpack')
const { ElementPlusResolver } = require('unplugin-vue-components/resolvers')

const HtmlWebpackplugin = require('html-webpack-plugin')
// 防止缓存
const oDate = new Date();
const cacheTime = oDate.getFullYear() +''+ (oDate.getMonth()+1) +''+ oDate.getDate() +''+ ~~(Math.random()*100);
const pageScript = `<script>
        if (window.name != 'bz${cacheTime}') {
            location.href = location.origin + location.pathname + '?vtime=' + Date.now() + location.hash;
            window.name = 'bz${cacheTime}'
        }
    </script>
`

module.exports = {
    publicPath: './',
    lintOnSave: false,
    productionSourceMap: false,
    configureWebpack: config => {
        config.plugins.push(AutoImport({
            resolvers: [ElementPlusResolver({ importStyle: false })],
        }))
        config.plugins.push(Components({
            resolvers: [ElementPlusResolver({ importStyle: false })],
        }))
        const miniCssExtractPlugin = config.plugins.find(
            plugin => plugin.constructor.name === 'MiniCssExtractPlugin'
        );
        if (miniCssExtractPlugin) {
            miniCssExtractPlugin.options.ignoreOrder = true;
        }
        if (process.env.NODE_ENV === 'production') {
            config.mode = 'production';
            // 打包文件大小配置
            config.performance = {
              maxEntrypointSize: 10000000,
              maxAssetSize: 30000000
            }
        }
        plugins: [
            new HtmlWebpackplugin({
                filename: 'index.html',
                template: './public/index.html',
                pageScript
            })
        ]
    },
    devServer: {
        https: false,
        hotOnly: false,
        proxy:{
            '/api': {
                target: 'http://127.0.0.1:8080',
                changeOrigin: true,
                pathRewrjite: {
                    '^/api': '/api'
                }
            }
        }
    },
    css: {
        loaderOptions: {
          sass: {
            // 8版本用prependData:
            prependData:`
              @import "@/assets/styles/variables.scss";  // scss文件地址
              @import "@/assets/styles/mixin.scss";     // scss文件地址
              @import "@/assets/css/uniformStyle.scss";
            `
          }
        }
    }
}
