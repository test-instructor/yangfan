import legacyPlugin from '@vitejs/plugin-legacy'
import { viteLogo } from './src/core/config'
import Banner from 'vite-plugin-banner'
import * as path from 'path'
import * as dotenv from 'dotenv'
import * as fs from 'fs'
import vuePlugin from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import VueFilePathPlugin from './vitePlugin/componentName/index.js'
import { svgBuilder } from 'vite-auto-import-svg'
import { AddSecret } from './vitePlugin/secret'
import UnoCSS from '@unocss/vite'

// @see https://cn.vitejs.dev/config/
export default ({ command, mode }) => {
  AddSecret('0044738584fc14e9a91054d8340c5bff')
  const NODE_ENV = mode || 'development'
  const envFiles = [`.env.${NODE_ENV}`]
  for (const file of envFiles) {
    const envConfig = dotenv.parse(fs.readFileSync(file))
    for (const k in envConfig) {
      process.env[k] = envConfig[k]
    }
  }

  viteLogo(process.env)

  const timestamp = Date.parse(new Date())

  const optimizeDeps = {}

  const alias = {
    '@': path.resolve(__dirname, './src'),
    vue$: 'vue/dist/vue.runtime.esm-bundler.js'
  }

  const isProdBuild = command === 'build' && NODE_ENV === 'production'
  const esbuild = isProdBuild
    ? { drop: ['console', 'debugger'] }
    : {}

  // legacy 构建（额外产出 -legacy 资源）会显著增加构建内存/时间。
  // 约定：VITE_LEGACY=false 时关闭，否则保持开启（兼容现有默认行为）。
  const enableLegacy = process.env.VITE_LEGACY !== 'false'

  const rollupOptions = {
    output: {
      entryFileNames: 'assets/[name].[hash].js',
      chunkFileNames: 'assets/[name].[hash].js',
      assetFileNames: 'assets/[name].[hash].[ext]'
    }
  }

  const base = "/"
  const root = "./"
  const outDir = "dist"

  const config = {
    base: base, // 编译后js导入的资源路径
    root: root, // index.html文件所在位置
    publicDir: 'public', // 静态资源文件夹
    resolve: {
      alias
    },
    define: {
      'process.env': {}
    },
    css: {
      preprocessorOptions: {
        scss: {
          api: 'modern-compiler' // or "modern"
        }
      }
    },
    server: {
      // 如果使用docker-compose开发模式，设置为false
      open: true,
      port: process.env.VITE_CLI_PORT,
      proxy: {
        // 把key的路径代理到target位置
        // detail: https://cli.vuejs.org/config/#devserver-proxy
        [process.env.VITE_BASE_API]: {
          // 需要代理的路径   例如 '/api'
          target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`, // 代理到 目标路径
          changeOrigin: true,
          rewrite: (path) =>
            path.replace(new RegExp('^' + process.env.VITE_BASE_API), '')
        }
      }
    },
    build: {
      // 默认用 esbuild，内存占用显著低于 terser（尤其是大 chunk/legacy 构建场景）
      minify: 'esbuild',
      manifest: false, // 是否产出manifest.json
      sourcemap: false, // 是否产出sourcemap.json
      outDir: outDir, // 产出目录
      // 仅影响构建日志统计；关闭可减少额外 CPU/内存开销
      reportCompressedSize: false,
      rollupOptions
    },
    esbuild,
    optimizeDeps,
    plugins: [
      process.env.VITE_POSITION === 'open' &&
        vueDevTools({ launchEditor: process.env.VITE_EDITOR }),
      enableLegacy &&
        legacyPlugin({
          targets: [
            'Android > 39',
            'Chrome >= 60',
            'Safari >= 10.1',
            'iOS >= 10.3',
            'Firefox >= 54',
            'Edge >= 15'
          ]
        }),
      vuePlugin(),
      svgBuilder(['./src/plugin/','./src/assets/icons/'],base, outDir,'assets', NODE_ENV),
      [Banner(`\n Build based on gin-vue-admin \n Time : ${timestamp}`)],
      VueFilePathPlugin('./src/pathInfo.json'),
      UnoCSS()
    ]
  }
  return config
}
