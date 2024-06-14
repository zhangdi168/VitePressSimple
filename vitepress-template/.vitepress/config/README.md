- index.ts 入口文件,也客自行修改成index.mts 然后使用export default defineConfig({ // ... })
- vpsimple.ts 存放来自vpsimple的配置,请勿动！！！！
- custom.ts 存放一些自定义的配置，或者需要覆盖vpsimple.ts的配置，比如第三方插件的配置等
> custom config优先级更大,如果两个对象有相同的属性，customer的属性值会覆盖前面的。